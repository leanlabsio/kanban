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
	hubs    map[string]*Hub
}

// Message is
type Message struct {
	Event string                 `json:"event"`
	Data  map[string]interface{} `json:"data"`
}

//New creates new Server instance
func New() *Server {
	return &Server{hubs: make(map[string]*Hub)}
}

//GetHub returns existing named Hub or creates new one
func (serv *Server) GetHub(id string) *Hub {
	h, ok := serv.hubs[id]
	if !ok {
		h = &Hub{}
		serv.hubs[id] = h
	}
	return h
}

//append add client to current server instance
func (serv *Server) append(c *Client) {
	serv.clients = append(serv.clients, c)
}

//Broadcast sends message to all subscribed clients
func (serv *Server) Broadcast(m *Message) {
	for _, h := range serv.hubs {
		h.Broadcast(m)
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
	for {
		select {
		case msg := <-c.ReceivingChan:
			b := msg.Data["board"]
			str, ok := b.(string)
			if !ok {
				log.Printf("BoardId if not a string %s", ok)
				panic("Could not resolve Hub")
			}
			h := serv.GetHub(str)
			h.append(c)
			log.Printf("%+v", h.clients)
			log.Printf("%s: %+v", "Recieved message", msg)
		}
	}
}

//ListenAndServePlugin start ws endpoint for plugins
func (serv *Server) ListenAndServePlugin(r <-chan *Message, s chan<- *Message, d <-chan bool, disc chan<- int, e <-chan error) {
	c := &Client{
		ReceivingChan:  r,
		SendingChan:    s,
		DoneChan:       d,
		DisconnectChan: disc,
		ErrChan:        e,
	}
	serv.append(c)
	for {
		select {
		case msg := <-c.ReceivingChan:
			serv.Broadcast(msg)
		}
	}
}
