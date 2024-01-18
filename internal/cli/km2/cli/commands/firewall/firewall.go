package firewall

import (
	"context"
	"errors"

	"github.com/mrlutik/kira2.0/internal/config/configFileController"
	"github.com/mrlutik/kira2.0/internal/docker"
	errUtils "github.com/mrlutik/kira2.0/internal/errors"
	"github.com/mrlutik/kira2.0/internal/firewall/firewallManager"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/manager/cli/commands/firewall/blacklist"
	"github.com/mrlutik/kira2.0/internal/manager/cli/commands/firewall/closePort"
	"github.com/mrlutik/kira2.0/internal/manager/cli/commands/firewall/openPort"
	"github.com/mrlutik/kira2.0/internal/manager/cli/commands/firewall/whitelist"
	"github.com/spf13/cobra"
)

const (
	use   = "firewall"
	short = "Setting up firewalld"
	long  = "Setting up firewalld"
)

var (
	log = logging.Log

	ErrOnlyOneFlagAllowed = errors.New("only one flag at a time is allowed")
)

func Firewall() *cobra.Command {
	log.Info("Adding `firewall` command...")
	firewallCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, _ []string) {
			if err := validateFlags(cmd); err != nil {
				log.Errorf("Some flag are not valid: %s", err)
				if err := cmd.Help(); err != nil {
					log.Fatalf("Error displaying help: %s", err)
				}
				return
			}
			mainFirewall(cmd)
		},
	}

	firewallCmd.AddCommand(openPort.OpenPort())
	firewallCmd.AddCommand(closePort.ClosePort())
	firewallCmd.AddCommand(blacklist.Blacklist())
	firewallCmd.AddCommand(whitelist.Whitelist())

	firewallCmd.Flags().Bool("close-ports", false, "Set this flag to block all ports except ssh")
	firewallCmd.Flags().Bool("open-ports", false, "Set this flag to open all km2 default ports")

	firewallCmd.Flags().Bool("default", false, "Set this flag to restore default setting for firewall (what km2 is set after node installation)")

	return firewallCmd
}

func validateFlags(cmd *cobra.Command) error {
	return nil
}

func mainFirewall(cmd *cobra.Command) {
	kiraCfg, err := configFileController.ReadOrCreateConfig()
	errUtils.HandleFatalErr("Error while reading cfg file", err)

	log.Info("Validating flags")

	openPorts, err := cmd.Flags().GetBool("open-ports")
	errUtils.HandleFatalErr("cannot parse flag", err)

	closePorts, err := cmd.Flags().GetBool("close-ports")
	errUtils.HandleFatalErr("cannot parse flag", err)

	defaultB, err := cmd.Flags().GetBool("default")
	errUtils.HandleFatalErr("cannot parse flag", err)

	err = validateBoolFlags(openPorts, closePorts, defaultB)
	errUtils.HandleFatalErr("only 1 flag can be accepted ", err)

	dockerManager, err := docker.NewTestDockerManager()
	errUtils.HandleFatalErr("Can't create instance of docker manager", err)
	defer dockerManager.Cli.Close()

	fm := firewallManager.NewFirewallManager(dockerManager, kiraCfg)
	ctx := context.Background()

	switch {
	case closePorts:
		err = fm.ClostAllOpenedPorts(ctx)
		errUtils.HandleFatalErr("Error while closing ports", err)
	case openPorts:
		err = fm.OpenConfigPorts(ctx)
		errUtils.HandleFatalErr("Error while opening ports", err)
	case defaultB:
		err = fm.SetUpFirewall(ctx)
		errUtils.HandleFatalErr("Error while closing ports", err)
	}
}

// checks if only 1 flag is set to true
func validateBoolFlags(flags ...bool) error {
	sum := 0
	for _, val := range flags {
		if val {
			sum++
		}
	}

	if sum > 1 {
		return ErrOnlyOneFlagAllowed
	}
	return nil
}
