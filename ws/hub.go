package ws

//Hub struct represents websocket Client's group
type Hub struct {
	Server
	Name string
}

//Broadcast sends message to all Client's on Hub
func (h *Hub) Broadcast(m *Message) {
	for _, c := range h.clients {
		c.SendingChan <- m
	}
}
