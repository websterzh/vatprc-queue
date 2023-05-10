package sockets

import (
	"github.com/goccy/go-json"
	"nhooyr.io/websocket"
	"sync"
	"vatprc-queue/gin/services"
)

type ClientPool struct {
	sync.RWMutex
	Clients map[*Client]interface{}
}

var pool *ClientPool

func AddClient(conn *websocket.Conn, airport string) {
	newClient := NewClient(conn, airport)
	defer DropClient(newClient)
	pool.Lock()
	pool.Clients[newClient] = nil
	pool.Unlock()
	newClient.Start()
}

func DropClient(conn *Client) {
	pool.Lock()
	delete(pool.Clients, conn)
	pool.Unlock()
	conn.Close()
}

func Tick() {
	for client := range pool.Clients {
		msg, err := json.Marshal(services.GetQueueResult(client.Airport, true))
		if err != nil {
			DropClient(client)
			continue
		}
		select {
		case client.ToSend <- msg:
			continue
		default:
			DropClient(client)
		}
	}
}

func Broadcast(msg string) {
	for client := range pool.Clients {
		select {
		case client.ToSend <- []byte(msg):
			continue
		default:
			DropClient(client)
		}
	}
}

func init() {
	pool = &ClientPool{
		Clients: make(map[*Client]interface{}),
	}
}
