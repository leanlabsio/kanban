package models

import (
	"github.com/spf13/viper"
	"gitlab.com/kanban/kanban/modules/gitlab"
	"golang.org/x/oauth2"
	"gopkg.in/redis.v3"
	"strings"
	"log"
)

var (
	c *redis.Client
)

// NewEngine creates new services for data from config settings
func NewEngine() error {

	gh := strings.TrimSuffix(viper.GetString("gitlab.host"), "/")
	d := strings.TrimSuffix(viper.GetString("server.domain"), "/")

	gitlab.NewEngine(&gitlab.Config{
		BasePath: gh + "/api/v3",
		Domain:   d,
		Oauth2: &oauth2.Config{
			ClientID:     viper.GetString("gitlab.oauth_client_id"),
			ClientSecret: viper.GetString("gitlab.oauth_client_secret"),
			Endpoint: oauth2.Endpoint{
				AuthURL:  gh + "/oauth/authorize",
				TokenURL: gh + "/oauth/token",
			},
			RedirectURL: d + "/assets/html/user/views/oauth.html",
		},
	})

	c = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host"),
		Password: viper.GetString("redis.passwd"),
		DB:       int64(viper.GetInt("redis.db")),
	})

	_, err := c.Ping().Result()

	if err != nil {
		log.Fatalf("Error connection to redis %s", err.Error())
	}

	return nil
}
