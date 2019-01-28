package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"impl"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {
	http.HandleFunc("/ws", handler)
	http.ListenAndServe(":7777", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		err    error
		data   []byte
		conn   *impl.Connection
	)

	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}
	if conn, err = impl.InConnection(wsConn); err != nil {
		goto ERROR
	}

	for {
		if err = conn.WriteMessage([]byte("go , go go !")); err != nil {
			return
		}
		time.Sleep(time.Second * 2)
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
