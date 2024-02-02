package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/mrlutik/kira2.0/internal/cli/km2/commands/firewall"
	"github.com/mrlutik/kira2.0/internal/cli/km2/commands/initialization"
	"github.com/mrlutik/kira2.0/internal/cli/km2/commands/maintenance"
	"github.com/mrlutik/kira2.0/internal/cli/km2/commands/monitoring"
	"github.com/mrlutik/kira2.0/internal/cli/km2/commands/start"
	"github.com/mrlutik/kira2.0/internal/cli/km2/commands/stop"
	"github.com/mrlutik/kira2.0/internal/logging"
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

func NewKiraCLI(commands []*cobra.Command, log *logging.Logger) *cobra.Command {
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
				err := log.SetLoggingLevel(logLevel)
				if err != nil {
					log.Fatalf("Setting logging level error: %s", err)
				}
			}
		},
	}

	for _, cmd := range commands {
		rootCmd.AddCommand(cmd)
	}

	rootCmd.PersistentFlags().String(
		"log-level",
		"panic",
		fmt.Sprintf(
			"Messages with this level and above will be logged. Valid levels are: %s",
			strings.Join(logging.GetValidLogLevels(), ", "),
		),
	)

	return rootCmd
}

func Start() {
	log, err := logging.InitLogger(logging.GetHooks(), "debug")
	if err != nil {
		fmt.Fprintf(os.Stdout, "Logging initialization error: %s\n", err)
		os.Exit(1)
	}

	commands := []*cobra.Command{
		initialization.Init(log),
		start.Start(log),
		stop.Stop(log),
		maintenance.Maintenance(log),
		firewall.Firewall(log),
		monitoring.Monitoring(log),
	}
	c := NewKiraCLI(commands, log)
	if err := c.Execute(); err != nil {
		log.Fatalf("Failed to execute command, error: %s", err)
	}
}
