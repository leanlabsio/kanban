package cmd

import (
	"fmt"
	"github.com/Unknwon/macaron"
	"github.com/codegangsta/cli"
	"github.com/macaron-contrib/bindata"
	"github.com/macaron-contrib/sockets"
	"gitlab.com/kanban/kanban/templates"
	"gitlab.com/kanban/kanban/web"
	"gitlab.com/kanban/kanban/ws"
	"log"
	"net/http"
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
	},
	Action: daemon,
}

func daemon(c *cli.Context) {
	m := macaron.New()

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

	m.Get("/*", func(ctx *macaron.Context) {
		ctx.Data["Version"] = c.App.Version
		ctx.Data["GitlabHost"] = c.String("gh")
		ctx.HTML(200, "templates/index")
	})
	m.Get("/ws/", sockets.Messages(), ws.S.ListenAndServe)
	m.Get("/ws/plugin", sockets.JSON(ws.Message{}), ws.S.ListenAndServePlugin)
	listenAddr := fmt.Sprintf("%s:%s", c.String("ip"), c.String("port"))
	log.Printf("Starting listening on %s", listenAddr)
	http.ListenAndServe(listenAddr, m)
}
