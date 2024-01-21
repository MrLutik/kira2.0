package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/mrlutik/kira2.0/internal/cli/deploy"
	"github.com/mrlutik/kira2.0/internal/cli/keys"
	"github.com/mrlutik/kira2.0/internal/cli/version"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/spf13/cobra"
)

const (
	// Command information
	use   = "kira2_launcher"
	short = "short descritpion"
	long  = "long description"

	// Flags
	loggingLevelFlag = "log-level"
)

var log = logging.Log

func NewCLI(cmds []*cobra.Command) *cobra.Command {
	log.Debug("Creating new CLI...")
	rootCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logLevel, err := cmd.Flags().GetString(loggingLevelFlag)
			if err != nil {
				log.Fatalf("Retrieving '%s' flag error: %s", loggingLevelFlag, err)
			}

			if logLevel != "" {
				logging.SetLevel(logLevel)
			}
		},
	}
	for _, cmd := range cmds {
		rootCmd.AddCommand(cmd)
	}
	rootCmd.PersistentFlags().Bool("verbose", false, "Verbosity level. Default: `false` ")
	rootCmd.PersistentFlags().String("log-level", "panic", fmt.Sprintf("Messages with this level and above will be logged. Valid levels are: %s", strings.Join(logging.ValidLogLevels, ", ")))
	return rootCmd
}

func Start() {
	cmds := []*cobra.Command{version.Version(), deploy.Node(), keys.Generate()}
	c := NewCLI(cmds)
	if err := c.Execute(); err != nil {
		log.Errorf("Failed to execute command %v\n", err)
		os.Exit(1)
	}
}
