package init

import (
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/manager/cli/commands/init/join"
	"github.com/mrlutik/kira2.0/internal/manager/cli/commands/init/new"
	"github.com/spf13/cobra"
)

var log = logging.Log

const (
	use   = "init"
	short = "init your node"
	long  = "init your node with creating new network or joining to existing one"
)

func Init() *cobra.Command {
	log.Info("Adding `firewall` command...")
	initCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, _ []string) {
			if err := validateFlags(cmd); err != nil {
				log.Errorf("Some flag are not valid: %s", err)
				cmd.Help()
				return
			}
			mainInit(cmd)
		},
	}

	initCmd.AddCommand(join.Join())
	initCmd.AddCommand(new.New())

	return initCmd
}

func validateFlags(cmd *cobra.Command) error {
	return nil
}

func mainInit(cmd *cobra.Command) {
	cmd.Help()
}
