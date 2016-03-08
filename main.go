package main // import "gitlab.com/leanlabsio/kanban"

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/leanlabsio/kanban/cmd"
)

// AppVer defines application version
var AppVer string = "dev"

func main() {
	kbCmd := &cobra.Command{
		Use: "kanban",
		Long: `Free OpenSource self hosted Kanban board for GitLab issues.

Full documentation is available on http://kanban.leanlabs.io/.

Report issues to <support@leanlabs.io> or https://gitter.im/leanlabsio/kanban.
                `,
	}
	viper.SetDefault("version", AppVer)

	kbCmd.AddCommand(&cmd.DaemonCmd)
	kbCmd.Execute()
}
