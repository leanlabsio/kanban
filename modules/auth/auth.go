package auth

import (
	"errors"
	"github.com/Unknwon/macaron"
	"github.com/dgrijalva/jwt-go"
	"gitlab.com/kanban/kanban/models"
	"github.com/spf13/viper"
)

func SignedInUser(ctx *macaron.Context) (*models.User, error) {
	if 0 == len(ctx.Req.Header["X-Kb-Access-Token"]) {
		return &models.User{
			Id: 0,
		}, nil
	}

	jwtToken, err := jwt.Parse(ctx.Req.Header["X-Kb-Access-Token"][0], func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("security.secret_key")), nil
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
