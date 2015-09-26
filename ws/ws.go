package ws

import (
	"sync"
)

// Server is generic type describing WebSocket server
type server struct {
	name string
	sync.Mutex
	clients []*client
	//	servers map[string]*Server
}

var servers = make(map[string]*server)

// Server returns existing named Server or creates new one
func Server(name string) *server {
	_, ok := servers[name]
	if !ok {
		servers[name] = &server{
			name: name,
		}
	}

	return servers[name]
}

//append adds client to current server instance
func (serv *server) append(c *client) {
	serv.clients = append(serv.clients, c)
}

//Broadcast sends message to all subscribed clients
func (serv *server) Broadcast(r string, m string) {
	s := Server(r)
	for _, c := range s.clients {
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
