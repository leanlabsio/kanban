package middleware

import (
	"gopkg.in/macaron.v1"
	"gitlab.com/leanlabsio/kanban/datasource/gitlab"
	"gitlab.com/leanlabsio/kanban/models"
)

func Datasource() macaron.Handler {
	return func(ctx *Context, u *models.User) {
		gds := gitlab.New(u.Credential["gitlab"].Token, u.Credential["gitlab"].PrivateToken)
		ctx.DataSource = gds
	}
}
