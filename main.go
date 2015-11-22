package main // import "gitlab.com/leanlabsio/kanban"

import (
	"github.com/spf13/cobra"
	"gitlab.com/leanlabsio/kanban/cmd"
)

// AppVer defines application version
const AppVer = "1.2.4"

func main() {
	kbCmd := &cobra.Command{
		Use: "kanban",
		Long: `Free OpenSource self hosted Kanban board for GitLab issues.

Full documentation is available on http://kanban.leanlabs.io/.

Report issues to <support@leanlabs.io> or https://gitter.im/leanlabsio/kanban.
                `,
	}

	kbCmd.AddCommand(&cmd.DaemonCmd)
	kbCmd.Execute()
}
