package ws

//Hub struct represents websocket Client's group
type hub struct {
	server
	Name string
}

//Broadcast sends message to all Client's on Hub
func (h *hub) broadcast(m string) {
	for _, c := range h.clients {
		c.SendingChan <- m
	}
}
