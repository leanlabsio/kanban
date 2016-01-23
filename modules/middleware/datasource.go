package middleware

import (
	"gitlab.com/leanlabsio/kanban/datasource/gitlab"
	"gitlab.com/leanlabsio/kanban/models"
	"gopkg.in/macaron.v1"
)

func Datasource() macaron.Handler {
	return func(ctx *Context, u *models.User) {
		gds := gitlab.New(u.Credential["gitlab"].Token, u.Credential["gitlab"].PrivateToken)
		ctx.DataSource = gds
	}
}
