package ws

//Client struct represents user connected via websocket
type Client struct {
	ReceivingChan  <-chan *Message
	SendingChan    chan<- *Message
	DoneChan       <-chan bool
	DisconnectChan chan<- int
	ErrChan        <-chan error
}
