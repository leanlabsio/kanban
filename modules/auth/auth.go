package auth

import (
	"errors"
	"github.com/Unknwon/macaron"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"gitlab.com/leanlabsio/kanban/models"
)

// SignedInUser returns models.User instance if user exists
func SignedInUser(ctx *macaron.Context) (*models.User, error) {
	h := ctx.Req.Header.Get("X-KB-Access-Token")
	if len(h) == 0 {
		return nil, errors.New("X-KB-Access-Token header missed")
	}

	jwtToken, err := jwt.Parse(h, func(token *jwt.Token) (interface{}, error) {
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
