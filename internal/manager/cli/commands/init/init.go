package init

import (
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/manager/cli/commands/init/join"
	"github.com/mrlutik/kira2.0/internal/manager/cli/commands/init/new"
	"github.com/spf13/cobra"
)

const (
	use   = "init"
	short = "init your node"
	long  = "init your node with creating new network or joining to existing one"
)

var log = logging.Log

func Init() *cobra.Command {
	log.Info("Adding `firewall` command...")
	initCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, _ []string) {
			mainInit(cmd)
		},
	}

	initCmd.AddCommand(join.Join())
	initCmd.AddCommand(new.New())

	return initCmd
}

func mainInit(cmd *cobra.Command) {
	cmd.Help()
}
