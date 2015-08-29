package main

import (
	"github.com/codegangsta/cli"
	"gitlab.com/kanban/kanban/cmd"
	"os"
)

const APP_VER = "1.2.4"

func main() {
	app := cli.NewApp()
	app.Name = "kanban"
	app.Email = "support@leanlabs.io"
	app.Usage = "Leanlab.io kanban board"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.DaemonCmd,
	}
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "V",
			Email: "support@leanlabs.io",
		},
		cli.Author{
			Name:  "cnam",
			Email: "support@leanlabs.io",
		},
	}
	app.Run(os.Args)
}
