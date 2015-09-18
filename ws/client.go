package ws

import (
	"log"
)

//Client struct represents user connected via websocket
type Client struct {
	ReceivingChan  <-chan *Message
	SendingChan    chan<- *Message
	DoneChan       <-chan bool
	DisconnectChan chan<- int
	ErrChan        <-chan error
}

//Handle starts client message handling loop
func (c *Client) Handle() {
	for {
		select {
		case msg := <-c.ReceivingChan:
			b := msg.Data["board"]
			str, ok := b.(string)
			if !ok {
				log.Printf("BoardId is not s string %s", ok)
				panic("Could not resolve hub")
			}
			h := S.GetHub(str)
			h.append(c)
			log.Printf("%+v", h.clients)
			log.Printf("%s: %+v", "Received message", msg)
		}
	}
}
