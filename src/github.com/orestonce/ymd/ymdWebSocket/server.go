package ymdWebSocket

import (
	"net/http"
	"github.com/gorilla/websocket"
)

func Upgrade(w http.ResponseWriter, r *http.Request) (conn *WebSocket, err error) {
	up := websocket.Upgrader{}
	inConn, err := up.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	return &WebSocket{
		inConn: inConn,
	}, nil
}
