package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/Unknwon/macaron"
	"gitlab.com/kanban/kanban/models"
	"gitlab.com/kanban/kanban/modules/setting"
	"errors"
)

func SignedInUser(ctx *macaron.Context) (*models.User, error) {
	if 0 == len(ctx.Req.Header["X-Kb-Access-Token"]) {
		return &models.User{
			Id: 0,
		}, nil
	}

	jwtToken, err := jwt.Parse(ctx.Req.Header["X-Kb-Access-Token"][0], func(token *jwt.Token) (interface{}, error) {
		return []byte(setting.Cfg.Section("security").Key("SECRET_KEY").String()), nil
	})

	if err != nil {
		return nil, err
	}

	if !jwtToken.Valid {
		return nil, errors.New("Invalid jwt token")
	}

	uname, _ := jwtToken.Claims["name"].(string)
	user, err := models.LoadUserByUsername(uname)
	user.Token = jwtToken

	return user, nil
}
