package main // import "gitlab.com/kanban/kanban"

import (
	"github.com/spf13/cobra"
	"gitlab.com/kanban/kanban/cmd"
)

// AppVer defines application version
const AppVer = "1.2.4"

func main() {
	kbCmd := &cobra.Command{
		Use:  "kanban",
		Long: "Here should be brief desc http://kanban.leanlabs.io",
	}

	kbCmd.AddCommand(&cmd.DaemonCmd)
	kbCmd.Execute()
}
