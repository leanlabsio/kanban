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
			Usage: "ip to bind",
		},
		cli.StringFlag{
			Name:  "port",
			Value: "9000",
			Usage: "port to bind",
		},
	},
	Action: daemon,
}

func daemon(c *cli.Context) {
	m := macaron.New()
	m.Use(macaron.Static("cmd"))
	m.Use(macaron.Recovery())
	listenAddr := fmt.Sprintf("%s:%s", c.String("ip"), c.String("port"))
	log.Printf("Listen: %s", listenAddr)
	http.ListenAndServe(listenAddr, m)
}
