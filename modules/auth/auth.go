package auth

import (
	"errors"
	"github.com/Unknwon/macaron"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"gitlab.com/kanban/kanban/models"
)

// SignedInUser returns models.User instance if user exists
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

	if err != nil {
		return nil, err
	}

	user.Token = jwtToken

	return user, nil
}
