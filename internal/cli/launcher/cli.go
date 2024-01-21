package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/mrlutik/kira2.0/internal/cli/launcher/deploy"
	"github.com/mrlutik/kira2.0/internal/cli/launcher/keys"
	"github.com/mrlutik/kira2.0/internal/cli/version"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/spf13/cobra"
)

const (
	// Command information
	use   = "kira2_launcher"
	short = "short description"
	long  = "long description"

	// Flags
	loggingLevelFlag = "log-level"
)

func NewCLI(commands []*cobra.Command, log *logging.Logger) *cobra.Command {
	log.Debug("Creating new CLI...")
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
	rootCmd.PersistentFlags().Bool("verbose", false, "Verbosity level. Default: `false` ")
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
		version.Version(log),
		deploy.Node(log),
		keys.Generate(log),
	}

	c := NewCLI(commands, log)
	if err := c.Execute(); err != nil {
		log.Fatalf("Failed to execute command %v", err)
	}
}
