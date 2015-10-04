package middleware

import (
	"github.com/Unknwon/macaron"
	"gitlab.com/kanban/kanban/models"
	"gitlab.com/kanban/kanban/modules/auth"
	"gitlab.com/kanban/kanban/ws"
	"encoding/json"
)

type Context struct {
	*macaron.Context
	User        *models.User
	IsAdmin     bool
	IsSigned    bool
	IsBasicAuth bool

	Provider    string
}

// Contexter initializes a classic context for a request.
func Contexter() macaron.Handler {
	return func(c *macaron.Context) {
		ctx := &Context{
			Context: c,
		}
		var err error

		ctx.User, err = auth.SignedInUser(ctx.Context)
		// Hardcore default data provider
		ctx.Provider = "gitlab"

		if err != nil {

		}

		ctx.IsSigned = true

		c.Map(ctx)
	}
}

func (*Context) Broadcast(r string, d interface{}) {
	res, _ := json.Marshal(d)
	go ws.Server(r).Broadcast(string(res))
}
