package main // import "gitlab.com/kanban/kanban"

import (
	"github.com/codegangsta/cli"
	"gitlab.com/kanban/kanban/cmd"
	"os"
)

// AppVer defines application version
const AppVer = "1.2.6"

func main() {
	app := cli.NewApp()
	app.Name = "kanban"
	app.Email = "support@leanlabs.io"
	app.Usage = "Leanlab.io kanban board"
	app.Version = AppVer
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
