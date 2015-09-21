package middleware

import (
	"github.com/Unknwon/macaron"
	"gitlab.com/kanban/kanban/models"
	"gitlab.com/kanban/kanban/modules/auth"
	"net/http"
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
		var err error

		ctx.User, err = auth.SignedInUser(ctx.Context)

		// Hardcore default data provider
		ctx.Provider = "gitlab"

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, `{"success": false}`)
		}

		ctx.IsSigned = true

		c.Map(ctx)
	}
}
