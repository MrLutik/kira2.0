package initialization

import (
	"github.com/mrlutik/kira2.0/internal/cli/km2/commands/initialization/join"
	"github.com/mrlutik/kira2.0/internal/cli/km2/commands/initialization/new"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/spf13/cobra"
)

const (
	use   = "init"
	short = "init your node"
	long  = "init your node with creating new network or joining to existing one"
)

func Init(log *logging.Logger) *cobra.Command {
	log.Info("Adding `firewall` command...")
	initCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, _ []string) {
			mainInit(cmd, log)
		},
	}

	initCmd.AddCommand(join.Join(log))
	initCmd.AddCommand(new.New(log))

	return initCmd
}

func mainInit(cmd *cobra.Command, log *logging.Logger) {
	if err := cmd.Help(); err != nil {
		log.Fatalf("Error displaying help: %s", err)
	}
}
