package impl

import (
	"net/http"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func Handler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err    error
		data   []byte
		conn   *Connection
	)

	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}
	if conn, err = InConnection(wsConn); err != nil {
		goto ERROR
	}
	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERROR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERROR
		}
	}
ERROR:
	conn.Close()
}
