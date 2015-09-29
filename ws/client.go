package ws

import (
	"encoding/json"
	"log"
)

//Client struct represents user connected via websocket
type client struct {
	ReceivingChan  <-chan string
	SendingChan    chan<- string
	DoneChan       <-chan bool
	DisconnectChan chan<- int
	ErrChan        <-chan error
	subscriptions  []string
}

//Handle starts client message handling loop
func (c *client) Handle() {
	for {
		select {
		case msg := <-c.ReceivingChan:
			var m message
			err := json.Unmarshal([]byte(msg), &m)
			if nil != err {
				log.Printf("%s", err)
			}
			p, err := m.resolve()
			c.process(p)
			log.Printf("%s, %+v", err, p)
			log.Printf("%s: %+v", "Received message", msg)
		case msg := <-c.DoneChan:
			for _, r := range c.subscriptions {
				s := Server(r)
				s.unsubscribe(c)
			}
			log.Printf("%+v", msg)
			return
		}
	}
}

// process handles different message types with appropriate method
func (c *client) process(m interface{}) {
	switch m.(type) {
	case *subscribe:
		m := m.(*subscribe)
		c.subscribe(m.RoutingKey)
	}
}

// subscribe binds client to listen for messages sent with routing key r
func (c *client) subscribe(r string) {
	h := Server(r)
	c.subscriptions = append(c.subscriptions, r)
	h.subscribe(c)
}
