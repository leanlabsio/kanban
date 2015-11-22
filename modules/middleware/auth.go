package middleware

import (
	"github.com/Unknwon/macaron"
	"gitlab.com/leanlabsio/kanban/models"
	"gitlab.com/leanlabsio/kanban/modules/auth"
)

// Auther checks jwt authentication
func Auther() macaron.Handler {
	return func(ctx *macaron.Context, c *Context) {
		u, err := auth.SignedInUser(ctx)

		if err != nil {
			unauthorized(ctx)
		}

		c.User = u
		c.IsSigned = true
	}
}

// unauthorized is a helper method to respond with HTTP 401
func unauthorized(ctx *macaron.Context) {
	ctx.JSON(401, models.ResponseError{Success: false, Message: "Unauthorized"})
}
