package cli

import (
	"fmt"
	"strings"

	"github.com/mrlutik/kira2.0/internal/errors"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/manager/cli/commands/firewall"
	initnode "github.com/mrlutik/kira2.0/internal/manager/cli/commands/init"
	"github.com/mrlutik/kira2.0/internal/manager/cli/commands/maintenance"
	"github.com/mrlutik/kira2.0/internal/manager/cli/commands/monitoring"
	"github.com/mrlutik/kira2.0/internal/manager/cli/commands/start"
	"github.com/mrlutik/kira2.0/internal/manager/cli/commands/stop"
	"github.com/spf13/cobra"
)

const (
	// Command info
	use   = "kira2"
	short = "kira2 manager for Kira network"
	long  = "kira2 manager for Kira network"

	// Flags
	loggingLevelFlag = "log-level"
)

var log = logging.Log

func NewKiraCLI(commands []*cobra.Command) *cobra.Command {
	log.Info("Creating new Kira manager CLI...")
	rootCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		PersistentPreRun: func(cmd *cobra.Command, _ []string) {
			logLevel, err := cmd.Flags().GetString(loggingLevelFlag)
			if err != nil {
				log.Fatalf("Retrieving '%s' flag error: %s", loggingLevelFlag, err)
			}

			if logLevel != "" {
				logging.SetLevel(logLevel)
			}
		},
	}

	for _, cmd := range commands {
		rootCmd.AddCommand(cmd)
	}

	rootCmd.PersistentFlags().String(loggingLevelFlag, "panic", fmt.Sprintf("Messages with this level and above will be logged. Valid levels are: %s", strings.Join(logging.ValidLogLevels, ", ")))
	return rootCmd
}

func Start() {
	commands := []*cobra.Command{start.Start(), monitoring.Monitoring(), firewall.Firewall(), maintenance.Maintenance(), initnode.Init(), stop.Stop()}
	c := NewKiraCLI(commands)
	if err := c.Execute(); err != nil {
		errors.HandleFatalErr("Failed to execute command", err)
	}
}
