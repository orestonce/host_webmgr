package ymdWebSocket

import (
	"github.com/gorilla/websocket"
	"time"
	"encoding/json"
	"sync"
)

type WebSocket struct {
	inConn *websocket.Conn
	mtxR   sync.Mutex
	mtxW   sync.Mutex
}

func Dial(wsUrl string) (conn *WebSocket, err error) {
	dialer := websocket.Dialer{}
	inConn, _, err := dialer.Dial(wsUrl, nil)
	if err != nil {
		return
	}
	return &WebSocket{
		inConn: inConn,
	}, nil
}

type exchangePacket struct {
	Mt   string          `json:",omitempty"`
	Data json.RawMessage `json:",omitempty"`
}

func (this *WebSocket) WritePacket(mt string, a interface{}) (err error) {
	this.mtxW.Lock()
	defer this.mtxW.Unlock()

	this.inConn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	data, err := json.Marshal(a)
	if err != nil {
		return
	}
	err = this.inConn.WriteJSON(exchangePacket{
		Mt:   mt,
		Data: data,
	})
	if err != nil {
		return
	}
	err = this.inConn.SetWriteDeadline(time.Time{})
	return
}

func (this *WebSocket) ReadPacket() (mt string, data []byte, err error) {
	this.mtxR.Lock()
	defer this.mtxR.Unlock()

	this.inConn.SetReadLimit(2 * 1024 * 1024)
	var tmp exchangePacket
	err = this.inConn.ReadJSON(&tmp)
	if err != nil {
		return
	}
	return tmp.Mt, tmp.Data, nil
}

func (this *WebSocket) Close() {
	this.inConn.Close()
}
