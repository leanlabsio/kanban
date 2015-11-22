package models

import (
	"github.com/spf13/viper"
	"gitlab.com/leanlabsio/kanban/modules/gitlab"
	"golang.org/x/oauth2"
	"gopkg.in/redis.v3"
	"log"
	"strings"
)

var (
	c *redis.Client
)

// NewEngine creates new services for data from config settings
func NewEngine() error {
	gh := strings.TrimSuffix(viper.GetString("gitlab.url"), "/")
	d := strings.TrimSuffix(viper.GetString("server.hostname"), "/")

	gitlab.NewEngine(&gitlab.Config{
		BasePath: gh + "/api/v3",
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

	c = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       int64(viper.GetInt("redis.db")),
	})

	_, err := c.Ping().Result()

	if err != nil {
		log.Fatalf("Error connection to redis %s", err.Error())
	}

	return nil
}
