package ws

import (
	"log"
)

//Client struct represents user connected via websocket
type client struct {
	ReceivingChan  <-chan string
	SendingChan    chan<- string
	DoneChan       <-chan bool
	DisconnectChan chan<- int
	ErrChan        <-chan error
}

//Handle starts client message handling loop
func (c *client) Handle() {
	for {
		select {
		case msg := <-c.ReceivingChan:
			h := Server.GetHub("hub1")
			h.append(c)
			log.Printf("%+v", h.clients)
			log.Printf("%s: %+v", "Received message", msg)
		}
	}
}
