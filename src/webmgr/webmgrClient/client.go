package main

import (
	"time"
	"github.com/orestonce/ymd/ymdWebSocket"
	"webmgr/webmgrProtocol"
	"sync"
	"encoding/json"
	"log"
	"os/exec"
	"bytes"
	"runtime"
	"github.com/orestonce/ymd/ymdOs"
	"strings"
	"context"
	"github.com/orestonce/ymd/ymdEncoding"
)

func main() {
	RunClient(`wss://ip.linux9.org/admin/WsConnect`)
}

func RunClient(serverAddress string) {
	for {
		err := runMainLogic(serverAddress)
		log.Println(`RunClient`, err)
		time.Sleep(time.Second * 3)
	}
}

func runMainLogic(serverAddress string) (err error) {
	conn, err := ymdWebSocket.Dial(serverAddress)
	if err != nil {
		return
	}
	defer conn.Close()

	wg := sync.WaitGroup{}
	exitCtx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func() {
		defer wg.Done()
		writePing := func() error {
			return conn.WritePacket(webmgrProtocol.PTPingRequest, webmgrProtocol.PingRequest{
				GoOs: runtime.GOOS,
				Pwd:  ymdOs.MustGetWd(),
			})
		}
		var err1 error
		err1 = writePing()
		for err1 == nil {
			select {
			case <-time.After(time.Second * 3):
				err1 = writePing()
			case <-exitCtx.Done():
				err1 = exitCtx.Err()
			}
		}
		log.Println(`runMainLogic `, err1)
		conn.Close()
	}()
	var err2 error
	for err2 == nil {
		var mt string
		var data []byte
		mt, data, err2 = conn.ReadPacket()
		switch mt {
		case webmgrProtocol.PTPingResponse:
		case webmgrProtocol.PTShellRequest:
			go runCmd(conn, data, exitCtx)
		}
	}
	conn.Close()
	cancel()
	wg.Wait()
	return
}

func runCmd(conn *ymdWebSocket.WebSocket, data []byte, exitCtx context.Context) {
	var req webmgrProtocol.ShellRequest
	err2 := json.Unmarshal(data, &req)
	if err2 != nil {
		log.Println(`runCmd`, err2)
		return
	}
	buf := bytes.NewBuffer(nil)
	if cmdSlice := strings.Split(req.Cmd, ` `); len(cmdSlice) > 0 {
		cmd := exec.CommandContext(exitCtx, cmdSlice[0], cmdSlice[1:]...)
		var content []byte
		content, err2 = cmd.CombinedOutput()
		if runtime.GOOS == `windows` {
			content, _ = ymdEncoding.GbkToUtf8(content)
		}
		buf.Write(content)
		if err2 != nil {
			buf.WriteString(err2.Error())
		}
	} else {
		buf.WriteString(`Unkown command ` + req.Cmd)
	}
	var resp webmgrProtocol.ShellResponse
	resp.Result = buf.String()
	resp.Id = req.Id
	select {
	case <-exitCtx.Done():
		err2 = exitCtx.Err()
	default:
		err2 = conn.WritePacket(webmgrProtocol.PTShellResponse, resp)
	}
	if err2 != nil {
		log.Println(`runCmd.WritePacket`, req.Id, err2)
	}
}
