package cmd

import (
	"fmt"
	"github.com/Unknwon/macaron"
	"github.com/macaron-contrib/binding"
	"github.com/codegangsta/cli"
	"github.com/macaron-contrib/bindata"
	"gitlab.com/kanban/kanban/templates"
	"gitlab.com/kanban/kanban/web"
	"gitlab.com/kanban/kanban/modules/models"
	"log"
	"net/http"
	"gitlab.com/kanban/kanban/modules/gitlab"
	"golang.org/x/oauth2"
	"strings"
)

// DaemonCmd is implementation of command to run application in daemon mode
var DaemonCmd = cli.Command{
	Name:  "daemon",
	Usage: "Start serving web traffic",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "ip",
			Value: "0.0.0.0",
			Usage: "IP address to listen on",
		},
		cli.StringFlag{
			Name:  "port",
			Value: "9000",
			Usage: "port to bind",
		},
		cli.StringFlag{
			Name:  "gh",
			Value: "https://gitlab.com",
			Usage: "GitLab host",
		},
		cli.StringFlag{
			Name:  "domain, d",
			Value: "http://localhost:9000",
			Usage: "Domain for using kanban",
		},
		cli.StringFlag{
			Name:  "gitlab_oauth_client_id, gc",
			Value: "qwerty",
			Usage: "GitLab Oauth client id",
		},
		cli.StringFlag{
			Name:  "gitlab_oauth_client_secret, gs",
			Value: "qwerty",
			Usage: "GitLab host",
		},
		cli.StringFlag{
			Name:  "secret_key, s",
			Value: "secret",
			Usage: "Kanban secret key for encript password",
		},
		cli.StringFlag{
			Name:  "redis, rh",
			Value: "localhost:6379",
			Usage: "Host for redis server",
		},
	},
	Action: daemon,
}

func daemon(c *cli.Context) {
	m := macaron.New()

	d := strings.TrimSuffix(c.String("domain"), "/")
	gh := strings.TrimSuffix(c.String("gh"), "/")
	gitlabApi := gitlab.New(&gitlab.Config{
		BasePath: gh + "/api/v3",
		Domain: d + "/assets/html/user/views/oauth.html",
		Oauth2: &oauth2.Config{
			ClientID:     c.String("gitlab_oauth_client_id"),
			ClientSecret: c.String("gitlab_oauth_client_secret"),
			Endpoint: oauth2.Endpoint{
				AuthURL:  gh+"/oauth/authorize",
				TokenURL: gh+"/oauth/token",
			},
			RedirectURL: d + "/assets/html/user/views/oauth.html",
		},
	}, c.String("secret_key"))

	m.Use(macaron.Recovery())
	m.Use(macaron.Logger())
	m.Use(macaron.Renderer(
		macaron.RenderOptions{
			Directory: "templates",
			TemplateFileSystem: bindata.Templates(bindata.Options{
				Asset:      templates.Asset,
				AssetDir:   templates.AssetDir,
				AssetNames: templates.AssetNames,
				Prefix:     "",
			}),
		},
	))
	m.Use(macaron.Static("web/images",
		macaron.StaticOptions{
			Prefix: "images",
			FileSystem: bindata.Static(bindata.Options{
				Asset:      web.Asset,
				AssetDir:   web.AssetDir,
				AssetNames: web.AssetNames,
				Prefix:     "web/images",
			}),
		},
	))
	m.Use(macaron.Static("web/template",
		macaron.StaticOptions{
			Prefix: "template",
			FileSystem: bindata.Static(bindata.Options{
				Asset:      web.Asset,
				AssetDir:   web.AssetDir,
				AssetNames: web.AssetNames,
				Prefix:     "web/template",
			}),
		},
	))
	m.Use(macaron.Static("web",
		macaron.StaticOptions{
			FileSystem: bindata.Static(bindata.Options{
				Asset:      web.Asset,
				AssetDir:   web.AssetDir,
				AssetNames: web.AssetNames,
				Prefix:     "web",
			}),
			Prefix: c.App.Version,
		},
	))

	m.Get("/assets/html/user/views/oauth.html", func(ctx *macaron.Context) {
		ctx.HTML(200, "templates/oauth")
	})

	m.Combo("/api/oauth").
		Get(gitlabApi.OauthUrl).
		Post(binding.Json(model.Oauth2{}), gitlabApi.OauthLogin)

	m.Get("/api/boards", gitlabApi.ListProjects)
	m.Get("/api/board", gitlabApi.SingleProjects)
	m.Get("/api/labels", gitlabApi.ListLabels)
	m.Get("/api/cards", gitlabApi.ListIssues)
	m.Get("/api/milestones", gitlabApi.ListMilestones)
	m.Get("/api/users", gitlabApi.ListProjectMembers)
	m.Get("/api/comments", gitlabApi.ListComments)

	m.Get("/*", func(ctx *macaron.Context) {
		ctx.Data["Version"] = c.App.Version
		ctx.Data["GitlabHost"] = c.String("gh")
		ctx.HTML(200, "templates/index")
	})

	listenAddr := fmt.Sprintf("%s:%s", c.String("ip"), c.String("port"))
	log.Printf("Starting listening on %s", listenAddr)
	http.ListenAndServe(listenAddr, m)
}
