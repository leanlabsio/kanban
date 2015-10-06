package middleware

import (
	"github.com/Unknwon/macaron"
	"gitlab.com/kanban/kanban/modules/auth"
	"net/http"
)

// Auther checks jwt authentication
func Auther() macaron.Handler {
	return func(ctx *macaron.Context, c *Context) {
		u, err := auth.SignedInUser(ctx)
		if err != nil {
			unauthorized(ctx.Resp)
		}

		c.User = u
		c.IsSigned = true
	}
}

// unauthorized is a helper method to respond with HTTP 401
func unauthorized(resp http.ResponseWriter) {
	http.Error(resp, "401 Unauthorized", http.StatusUnauthorized)
}
