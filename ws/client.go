package ws

//Client is
type Client struct {
	ReceivingChan  <-chan *Message
	SendingChan    chan<- *Message
	DoneChan       <-chan bool
	DisconnectChan chan<- int
	ErrChan        <-chan error
}
