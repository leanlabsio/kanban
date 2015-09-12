package ws

import (
	"encoding/json"
	"log"
	"sync"
)

var _ = json.NewDecoder

// Server is
type Server struct {
	sync.Mutex
	clients []*Client
}

// Message is
type Message struct {
	Data string `json:"data"`
}

//append add client to current server instance
func (serv *Server) append(c *Client) {
	serv.clients = append(serv.clients, c)
}

//Broadcast send message to all subscribed clients
func (serv *Server) Broadcast(m *Message) {
	for _, c := range serv.clients {
		c.SendingChan <- m
	}
}

// ListenAndServe start ws
func (serv *Server) ListenAndServe(r <-chan *Message, s chan<- *Message, d <-chan bool, disc chan<- int, e <-chan error) {
	c := &Client{
		ReceivingChan:  r,
		SendingChan:    s,
		DoneChan:       d,
		DisconnectChan: disc,
		ErrChan:        e,
	}
	serv.append(c)
	log.Printf("%+v", serv.clients)
	for {
		select {
		case msg := <-c.ReceivingChan:
			log.Printf("%s: %+v", "Recieved message", msg)
			serv.Broadcast(msg)
		}
	}
}
