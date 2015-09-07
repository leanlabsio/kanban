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
}

// Client is
type Client struct {
	ReceivingChan  <-chan *Message
	SendingChan    chan<- *Message
	DoneChan       <-chan bool
	DisconnectChan chan<- int
	ErrChan        <-chan error
}

// Message is
type Message struct {
	Data string `json:"data"`
}

func (me *Server) ListenAndServe(r <-chan *Message, s chan<- *Message, d <-chan bool, disc chan<- int, e <-chan error) {
	c := &Client{
		ReceivingChan:  r,
		SendingChan:    s,
		DoneChan:       d,
		DisconnectChan: disc,
		ErrChan:        e,
	}
	log.Printf("%+v", c)
	for {
		select {
		case msg := <-c.ReceivingChan:
			log.Printf("%+v", msg)
		}
	}
}
