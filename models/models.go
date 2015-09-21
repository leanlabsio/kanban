package models

import (
	"fmt"
	"gitlab.com/kanban/kanban/modules/gitlab"
	"gitlab.com/kanban/kanban/modules/setting"
	"golang.org/x/oauth2"
	"gopkg.in/redis.v3"
	"strings"
)

var (
	gc *gitlab.GitlabClient
	c  *redis.Client
)

// NewEngine creates new services for data from config settings
func NewEngine() error {

	gh := strings.TrimSuffix(setting.Cfg.Section("gitlab").Key("ROOT_URL").String(), "/")
	d := strings.TrimSuffix(setting.Cfg.Section("server").Key("ROOT_URL").String(), "/")

	gitlab.NewEngine(&gitlab.Config{
		BasePath: gh + "/api/v3",
		Domain:   d,
		Oauth2: &oauth2.Config{
			ClientID:     setting.Cfg.Section("gitlab").Key("OAUTH_CLIENT_ID").String(),
			ClientSecret: setting.Cfg.Section("gitlab").Key("OAUTH_SECRET_KEY").String(),
			Endpoint: oauth2.Endpoint{
				AuthURL:  gh + "/oauth/authorize",
				TokenURL: gh + "/oauth/token",
			},
			RedirectURL: d + "/assets/html/user/views/oauth.html",
		},
	})

	db, _ := setting.Cfg.Section("cache").Key("DB").Int64()

	c = redis.NewClient(&redis.Options{
		Addr:     setting.Cfg.Section("cache").Key("HOST").String(),
		Password: setting.Cfg.Section("cache").Key("PASS").String(),
		DB:       db,
	})

	_, err := c.Ping().Result()

	if err != nil {
		fmt.Println("%s", err.Error())
	}

	return nil
}
