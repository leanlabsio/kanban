package middleware

import (
	"gitlab.com/leanlabsio/kanban/models"
	"gitlab.com/leanlabsio/kanban/modules/auth"
	"gopkg.in/macaron.v1"
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
		ctx.Map(u)
	}
}

// unauthorized is a helper method to respond with HTTP 401
func unauthorized(ctx *macaron.Context) {
	ctx.JSON(401, models.ResponseError{Success: false, Message: "Unauthorized"})
}
