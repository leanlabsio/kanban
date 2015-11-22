package middleware

import (
	"encoding/json"
	"github.com/Unknwon/macaron"
	"gitlab.com/leanlabsio/kanban/models"
	"gitlab.com/leanlabsio/kanban/ws"
)

type Context struct {
	*macaron.Context
	User        *models.User
	IsAdmin     bool
	IsSigned    bool
	IsBasicAuth bool

	Provider string
}

// Contexter initializes a classic context for a request.
func Contexter() macaron.Handler {
	return func(c *macaron.Context) {
		ctx := &Context{
			Context: c,
		}

		ctx.Provider = "gitlab"

		c.Map(ctx)
	}
}

// Broadcast sends message via WebSocket to all subscribed to r users
func (*Context) Broadcast(r string, d interface{}) {
	res, _ := json.Marshal(d)
	go ws.Server(r).Broadcast(string(res))
}
