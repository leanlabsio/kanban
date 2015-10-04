package ws

import (
	"encoding/json"
)

// message represents general client-server communication unit
type message struct {
	Event string `json:"event"`
	Data  json.RawMessage
}

// subscribe represents subscribe message payload
type subscribe struct {
	RoutingKey string `json:"routing_key"`
}

// resolve unserializes message.Data to an appropriate event stuct
func (m *message) resolve() (interface{}, error) {
	var d interface{}
	switch m.Event {
	case "subscribe":
		d = new(subscribe)
	}

	err := json.Unmarshal([]byte(m.Data), &d)
	if nil != err {
		return nil, err
	}

	return d, nil
}
