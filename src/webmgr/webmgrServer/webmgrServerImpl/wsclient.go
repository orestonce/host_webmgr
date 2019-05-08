package webmgrServerImpl

import (
	"github.com/orestonce/ymd/ymdGin"
	"github.com/orestonce/ymd/ymdWebSocket"
	"log"
	"sync"
	"github.com/orestonce/ymd/ymdUuid"
	"time"
	"webmgr/webmgrProtocol"
	"encoding/json"
	"github.com/orestonce/ymd/ymdView"
	"github.com/orestonce/ymd/ymdView/ymdLyear"
	"sort"
	"strconv"
	"github.com/orestonce/ymd/ymdIpToRegion"
)

var gWsClientMap = map[string]*WsClient{}
var gWsClientMapLocker sync.Mutex

type WsClient struct {
	ClientId        string
	Conn            *ymdWebSocket.WebSocket
	ClientIp        string
	ConnectTime     time.Time
	UpdateTime      time.Time
	GoOs            string
	Pwd             string
	RequestIdAlloc  int64
	Region          string
	ResponseWaitMap map[int64]chan webmgrProtocol.ShellResponse
}

func (adminWeb) WsClientListPage(ctx *ymdGin.Context) {
	var list ymdView.HtmlRendererList
	table := ymdLyear.Table{
		TitleRow: []string{
			`#`, `客户端IP`, `首次链接时间`, `最后心跳时间`, `地点`, `GoOs`, `Pwd`, `RpcWaitMapSize`, `操作`,
		},
	}
	gWsClientMapLocker.Lock()
	defer gWsClientMapLocker.Unlock()
	clientList := []*WsClient{}
	for _, one := range gWsClientMap {
		clientList = append(clientList, one)
	}
	sort.Slice(clientList, func(i, j int) bool {
		return clientList[i].ConnectTime.Before(clientList[j].ConnectTime)
	})
	for idx, one := range clientList {
		table.AddOneRow(ymdLyear.AddOneRowRequest{
			PreList: []string{
				strconv.Itoa(idx + 1),
				one.ClientIp,
				getTimeStr(one.ConnectTime),
				getTimeStr(one.UpdateTime),
				one.Region,
				one.GoOs,
				one.Pwd,
				strconv.Itoa(len(one.ResponseWaitMap)),
			},
			PostList: []ymdView.HtmlRenderer{
				ymdView.HtmlRendererList{
					ymdLyear.Button{
						Href: adminUrlArgs(`WsAction`, map[string]string{
							`Action`:   `Delete`,
							`ClientId`: one.ClientId,
						}),
						IsAjaxSendAndReload: true,
						IsDanger:            true,
						ShowContent:         `删除`,
					},
					ymdLyear.Blank{
						Count: 4,
					},
					ymdLyear.Button{
						Href: adminUrlArgs(`WsShellPage`, map[string]string{
							`ClientId`: one.ClientId,
						}),
						ShowContent: `执行shell`,
					},
				},
			},
		})
	}
	list.Add(table)
	adminUi(ctx.Writer, `WsClientListPage`, list)
}

func (adminWeb) WsAction(ctx *ymdGin.Context) {
	clientId := ctx.InStr(`ClientId`)
	switch ctx.InStr(`Action`) {
	case `Delete`:
		editClient(clientId, func(client *WsClient) {
			client.Conn.Close()
		})
	}
}

func (adminWeb) WsConnect(ctx *ymdGin.Context) {
	conn, err := ymdWebSocket.Upgrade(ctx.Writer, ctx.Request)
	if err != nil {
		log.Println(`WsConnect`, err)
		return
	}
	defer conn.Close()

	info, err := ymdIpToRegion.GetIpInfo(ctx.ClientIP())
	if err != nil {
		log.Println(`WsConnect getIpInfo`, info)
		return
	}
	clientId := ymdUuid.MustNewUUID()
	client := &WsClient{
		ClientId:        clientId,
		Conn:            conn,
		ClientIp:        ctx.ClientIP(),
		ConnectTime:     time.Now(),
		UpdateTime:      time.Now(),
		ResponseWaitMap: map[int64]chan webmgrProtocol.ShellResponse{},
		Region:          info.String(),
	}
	gWsClientMapLocker.Lock()
	gWsClientMap[client.ClientId] = client
	gWsClientMapLocker.Unlock()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err error
		for err == nil {
			var mt string
			var data []byte
			mt, data, err = conn.ReadPacket()
			if err != nil {
				break
			}
			switch mt {
			case webmgrProtocol.PTPingRequest:
				var p webmgrProtocol.PingRequest
				json.Unmarshal(data, &p)
				now := time.Now()
				editClient(clientId, func(client *WsClient) {
					client.UpdateTime = now
					client.GoOs = p.GoOs
					client.Pwd = p.Pwd
				})
				err = conn.WritePacket(webmgrProtocol.PTPingResponse, struct{}{})
			case webmgrProtocol.PTShellResponse:
				var resp webmgrProtocol.ShellResponse
				err = json.Unmarshal(data, &resp)
				if err != nil {
					break
				}
				editClient(clientId, func(client *WsClient) {
					ch := client.ResponseWaitMap[resp.Id]
					if ch == nil {
						return
					}
					select {
					case ch <- resp:
					default:
					}
					delete(client.ResponseWaitMap, resp.Id)
				})
			default:
				log.Println(`WsConnect unkown packet`, clientId)
			}
		}
		log.Println(`WsConnect end`, clientId, err)
		conn.Close()
	}()
	log.Println(`WsConnect success`, clientId)
	wg.Wait()
	gWsClientMapLocker.Lock()
	for id, one := range client.ResponseWaitMap {
		select {
		case one <- webmgrProtocol.ShellResponse{
			Id:     id,
			Result: `ugfke0lr server kick client.`,
		}:
		default:
		}
	}
	delete(gWsClientMap, clientId)
	gWsClientMapLocker.Unlock()
}

func (adminWeb) WsShellPage(ctx *ymdGin.Context) {
	var list ymdView.HtmlRendererList
	cmd := ctx.InStr(`Cmd`)
	clientId := ctx.InStr(`ClientId`)
	form := ymdLyear.Form{
		TopTitle: `执行shell`,
		ActionUrl: adminUrlArgs(`WsShellPage`, map[string]string{
			`ClientId`: clientId,
		}),
		IsPost: true,
	}
	form.InputList = append(form.InputList, ymdLyear.InputString{
		Name:     `ClientId`,
		Value:    clientId,
		ReadOnly: true,
	})
	editClient(clientId, func(client *WsClient) {
		form.InputList = append(form.InputList, ymdLyear.InputString{
			Name:     `Ip`,
			Value:    client.ClientIp,
			ReadOnly: true,
		})
		form.InputList = append(form.InputList, ymdLyear.InputString{
			Name:     `Region`,
			Value:    client.Region,
			ReadOnly: true,
		})
		form.InputList = append(form.InputList, ymdLyear.InputString{
			Name:     `GoOs`,
			Value:    client.GoOs,
			ReadOnly: true,
		})
		form.InputList = append(form.InputList, ymdLyear.InputString{
			Name:     `Pwd`,
			Value:    client.Pwd,
			ReadOnly: true,
		})
	})
	form.InputList = append(form.InputList, ymdLyear.InputString{
		Name:  `Cmd`,
		Value: cmd,
	})
	list.Add(form)
	card := ymdLyear.Card{
		Title: `输出`,
		Body:  ymdLyear.String(``),
	}
	if cmd != `` {
		card.Body = ymdLyear.Pre(runCmdInWsClient(cmd, clientId))
	}
	list.Add(card)
	adminUi(ctx.Writer, `WsShellPage`, list)
}

func runCmdInWsClient(cmd string, clientId string) (outstr string) {
	var conn *ymdWebSocket.WebSocket
	var id int64
	ch := make(chan webmgrProtocol.ShellResponse, 10)
	editClient(clientId, func(client *WsClient) {
		conn = client.Conn
		client.RequestIdAlloc++
		id = client.RequestIdAlloc
		client.ResponseWaitMap[id] = ch
	})
	if conn == nil {
		outstr = `SendShell conn not found .`
		return
	}
	err := conn.WritePacket(webmgrProtocol.PTShellRequest, webmgrProtocol.ShellRequest{
		Cmd: cmd,
		Id:  id,
	})
	if err != nil {
		outstr = `SendShell ` + err.Error()
		return
	}
	resp := <-ch
	outstr = resp.Result
	return
}

func editClient(clientId string, cb func(client *WsClient)) {
	gWsClientMapLocker.Lock()
	defer gWsClientMapLocker.Unlock()

	client := gWsClientMap[clientId]
	if client != nil {
		cb(client)
	}
}
