package models

import (
	"strings"

	"github.com/spf13/viper"
	"gitlab.com/leanlabsio/kanban/modules/gitlab"
	"golang.org/x/oauth2"
	"gopkg.in/redis.v3"
)

var (
	c *redis.Client
)

// NewEngine creates new services for data from config settings
func NewEngine(r *redis.Client) error {
	gh := strings.TrimSuffix(viper.GetString("gitlab.url"), "/")
	d := strings.TrimSuffix(viper.GetString("server.hostname"), "/")
	c = r

	gitlab.NewEngine(&gitlab.Config{
		BasePath: gh + "/api/v4",
		Domain:   d,
		Oauth2: &oauth2.Config{
			ClientID:     viper.GetString("gitlab.client"),
			ClientSecret: viper.GetString("gitlab.secret"),
			Endpoint: oauth2.Endpoint{
				AuthURL:  gh + "/oauth/authorize",
				TokenURL: gh + "/oauth/token",
			},
			RedirectURL: d + "/assets/html/user/views/oauth.html",
		},
	})

	return nil
}
