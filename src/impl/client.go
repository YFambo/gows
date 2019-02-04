package impl

import (
	"github.com/gorilla/websocket"
	"time"
)

type Client struct {
	hub *Hub

	// The websocket connection.
	wsConn *websocket.Conn

	//the objcoon
	coon *Connection

	tick chan int
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = time.Second * 10

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

func (c *Client) TickTime() {
	for {
		select {
		case <-c.tick:
			c.coon.outChan <- []byte("pong")
		case <-time.After(pingPeriod):
			c.coon.outChan <- []byte("time out !")
			time.Sleep(time.Second * 1)
			c.hub.unregister <- c
			return
		}
	}
}
