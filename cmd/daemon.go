package cmd

import (
	"fmt"
	"github.com/Unknwon/macaron"
	"github.com/codegangsta/cli"
	"github.com/macaron-contrib/bindata"
	"github.com/macaron-contrib/binding"
	"github.com/macaron-contrib/sockets"
	"gitlab.com/kanban/kanban/templates"
	"gitlab.com/kanban/kanban/web"
	"gitlab.com/kanban/kanban/ws"
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
		cli.StringFlag{
			Name: "cache-host",
			Value: "",
			Usage: "Cache host 127.0.0.1:6379",
		},
		cli.StringFlag{
			Name: "gitlab-client-id",
			Value: "",
			Usage: "Gitlab oauth2 client id",
		},
		cli.StringFlag{
			Name: "gitlab-client-secret",
			Value: "",
			Usage: "Gitlab oauth2 client secret",
		},
	},
	Action: daemon,
}

func daemon(c *cli.Context) {
	m := macaron.New()

	setting.NewContext(c)
	models.NewEngine()
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

	m.Group("/api", func() {
		m.Combo("/oauth").
			Get(user.OauthUrl).
			Post(binding.Json(auth.Oauth2{}), user.OauthLogin)

		m.Post("/login", binding.Json(auth.SignIn{}), user.SignIn)
		m.Post("/register", binding.Json(auth.SignUp{}), user.SignUp)

		m.Get("/boards", board.ListBoards)
		m.Get("/board", board.ItemBoard)
		m.Get("/labels", board.ListLabels)
		m.Get("/cards", board.ListCards)
		m.Get("/milestones", board.ListMilestones)
		m.Get("/users", board.ListMembers)
		m.Combo("/comments").
			Get(board.ListComments).
			Post(binding.Json(models.CommentRequest{}), board.CreateComment)

		m.Combo("/card").
			Post(binding.Json(models.CardRequest{}), board.CreateCard).
			Put(binding.Json(models.CardRequest{}), board.UpdateCard).
			Delete(binding.Json(models.CardRequest{}), board.DeleteCard)

		m.Put("/card/move", binding.Json(models.CardRequest{}), board.MoveToCard)

	})
	m.Get("/*", routers.Home)
	m.Get("/ws/", sockets.Messages(), ws.ListenAndServe)
	listenAddr := fmt.Sprintf("%s:%s", c.String("ip"), c.String("port"))
	log.Printf("Starting listening on %s", listenAddr)
	http.ListenAndServe(listenAddr, m)
}
