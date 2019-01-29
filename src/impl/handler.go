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
		//设置读的缓存的大小
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func Handler(hub *Hub, w http.ResponseWriter, r *http.Request) {
	var (
		data   []byte
		err    error
		wsCoon *websocket.Conn
	)

	if wsCoon, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}
	client := &Client{hub: hub, wsConn: wsCoon,tick:make(chan int,1)}
	client.hub.register <- client
	if client.coon, err = InConnection(wsCoon); err != nil {
		goto ERROR
	}
	//心跳
	go client.TickTime()
	for {
		if data, err = client.coon.ReadMessage(); err != nil {
			goto ERROR
		}
		switch string(data) {
		case "broadcast":
			client.hub.broadcast <- data
			continue
		case "close":
			client.hub.unregister <- client
			continue
		case "ping":
			client.tick <- 1
			continue
		}
		if err = client.coon.WriteMessage(data); err != nil {
			goto ERROR
		}
	}
ERROR:
	client.coon.Close()
}
