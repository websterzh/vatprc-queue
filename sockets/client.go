package sockets

import (
	"context"
	"nhooyr.io/websocket"
	"vatprc-queue/config"
)

type Client struct {
	ctx     context.Context
	cancel  context.CancelFunc
	Airport string
	Socket  *websocket.Conn
	ToSend  chan []byte
}

func (c *Client) Close() {
	c.cancel()
	c.Socket.Close(websocket.StatusNormalClosure, "bye :)")
	close(c.ToSend)
}

func (c *Client) Start() {
	c.ctx = c.Socket.CloseRead(c.ctx)
	for {
		select {
		case message := <-c.ToSend:
			c.Socket.Write(c.ctx, websocket.MessageText, message)
		case <-c.ctx.Done():
			return
		}
	}
}

func NewClient(Socket *websocket.Conn, airport string) *Client {
	client := &Client{
		Airport: airport,
		Socket:  Socket,
		ToSend:  make(chan []byte, config.File.Section("app").Key("socket_message_buffer").MustInt(0)),
	}
	client.ctx, client.cancel = context.WithCancel(context.Background())
	return client
}
