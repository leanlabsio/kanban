package ws

import (
	"sync"
)

// Server is generic type describing WebSocket server
type server struct {
	name string
	sync.Mutex
	clients map[*client]int
}

var servers = make(map[string]*server)

// Server returns existing named Server or creates new one
func Server(name string) *server {
	_, ok := servers[name]
	if !ok {
		servers[name] = &server{
			name:    name,
			clients: make(map[*client]int),
		}
	}

	return servers[name]
}

// subscribe adds client to current server instance
func (serv *server) subscribe(c *client) {
	serv.clients[c] = 0
}

// unsubscribe removes client from specified server instance
func (serv *server) unsubscribe(c *client) {
	delete(serv.clients, c)
}

//Broadcast sends message to all subscribed clients
func (serv *server) Broadcast(m string) {
	for c := range serv.clients {
		c.SendingChan <- m
	}
}

// ListenAndServe start ws
func ListenAndServe(r <-chan string, s chan<- string, d <-chan bool, disc chan<- int, e <-chan error) {
	c := &client{
		ReceivingChan:  r,
		SendingChan:    s,
		DoneChan:       d,
		DisconnectChan: disc,
		ErrChan:        e,
	}
	go c.Handle()
}
