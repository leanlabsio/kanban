package cmd

import (
	"fmt"
	"github.com/Unknwon/macaron"
	"github.com/codegangsta/cli"
	"log"
	"net/http"
)

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
			Value: "https://gitlab.org",
			Usage: "GitLab host",
		},
	},
	Action: daemon,
}

func daemon(c *cli.Context) {
	m := macaron.New()

	m.Use(macaron.Recovery())
	m.Use(macaron.Logger())
	m.Use(macaron.Renderer(macaron.RenderOptions{Directory: "templates"}))
	m.Use(macaron.Static("web",
		macaron.StaticOptions{
			Prefix: c.App.Version,
		}))

	m.Get("/*", func(ctx *macaron.Context) {
		ctx.Data["Version"] = c.App.Version
		ctx.HTML(200, "index")
	})

	listenAddr := fmt.Sprintf("%s:%s", c.String("ip"), c.String("port"))
	log.Printf("Listen: %s", listenAddr)

	http.ListenAndServe(listenAddr, m)
}
