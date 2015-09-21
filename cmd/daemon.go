package cmd

import (
	"fmt"
	"github.com/Unknwon/macaron"
	"github.com/codegangsta/cli"
	"github.com/macaron-contrib/bindata"
	"github.com/macaron-contrib/binding"
	"gitlab.com/kanban/kanban/templates"
	"gitlab.com/kanban/kanban/web"
	"log"
	"net/http"

	"gitlab.com/kanban/kanban/modules/auth"
	"gitlab.com/kanban/kanban/modules/setting"

	"gitlab.com/kanban/kanban/models"
	"gitlab.com/kanban/kanban/modules/middleware"
	"gitlab.com/kanban/kanban/routers"
	"gitlab.com/kanban/kanban/routers/board"
	"gitlab.com/kanban/kanban/routers/user"
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
			Name:  "config",
			Value: "",
			Usage: "Custom config file",
		},
	},
	Action: daemon,
}

func daemon(c *cli.Context) {
	m := macaron.New()

	if c.String("config") != "" {
		setting.CustomPath = c.String("config")
	}

	setting.NewContext()
	models.NewEngine()
	setting.App_Version = c.App.Version
	m.Use(middleware.Contexter())
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

	m.Get("/assets/html/user/views/oauth.html", user.OauthHandler)

	m.Combo("/api/oauth").
		Get(user.OauthUrl).
		Post(binding.Json(auth.Oauth2{}), user.OauthLogin)

	m.Post("/api/login", binding.Json(auth.SignIn{}), user.SignIn)
	m.Post("/api/register", binding.Json(auth.SignUp{}), user.SignUp)

	m.Get("/api/boards", board.ListBoards)
	m.Get("/api/board", board.ItemBoard)
	m.Get("/api/labels", board.ListLabels)
	m.Get("/api/cards", board.ListCards)
	m.Get("/api/milestones", board.ListMilestones)
	m.Get("/api/users", board.ListMembers)
	m.Get("/api/comments", board.ListComments)

	m.Get("/*", routers.Home)

	listenAddr := fmt.Sprintf("%s:%s", c.String("ip"), c.String("port"))
	log.Printf("Starting listening on %s", listenAddr)
	http.ListenAndServe(listenAddr, m)
}
