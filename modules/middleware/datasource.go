package middleware

import (
	"gitlab.com/leanlabsio/kanban/datasource/gitlab"
	"gitlab.com/leanlabsio/kanban/models"
	"gopkg.in/macaron.v1"
	"gopkg.in/redis.v3"
)

func Datasource() macaron.Handler {
	return func(ctx *Context, u *models.User, r *redis.Client) {
		gds := gitlab.New(u.Credential["gitlab"].Token,
			u.Credential["gitlab"].PrivateToken,
			r)
		ctx.DataSource = gds
	}
}
