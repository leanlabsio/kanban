package ws

import (
	"sync"
)

// server is generic type describing WebSocket server
type server struct {
	sync.Mutex
	clients []*client
	hubs    map[string]*hub
}

// Server is an websocket server instance
var Server = &server{
	hubs: make(map[string]*hub),
}

//GetHub returns existing named Hub or creates new one
func (serv *server) hub(id string) *hub {
	h, ok := serv.hubs[id]
	if !ok {
		h = &hub{}
		serv.hubs[id] = h
	}

	return h
}

//append adds client to current server instance
func (serv *server) append(c *client) {
	serv.clients = append(serv.clients, c)
}

//Broadcast sends message to all subscribed clients
func (serv *server) Broadcast(r string, m string) {
	h := serv.hub(r)
	h.broadcast(m)
}

// ListenAndServe start ws
func (serv *server) ListenAndServe(r <-chan string, s chan<- string, d <-chan bool, disc chan<- int, e <-chan error) {
	c := &client{
		ReceivingChan:  r,
		SendingChan:    s,
		DoneChan:       d,
		DisconnectChan: disc,
		ErrChan:        e,
	}
	go c.Handle()
}
