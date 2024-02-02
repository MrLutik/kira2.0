package firewall

import (
	"context"
	"errors"
	"time"

	"github.com/docker/docker/client"
	"github.com/mrlutik/kira2.0/internal/cli/km2/commands/firewall/blacklist"
	"github.com/mrlutik/kira2.0/internal/cli/km2/commands/firewall/closeport"
	"github.com/mrlutik/kira2.0/internal/cli/km2/commands/firewall/openport"
	"github.com/mrlutik/kira2.0/internal/cli/km2/commands/firewall/whitelist"
	"github.com/mrlutik/kira2.0/internal/config/controller"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/firewall/manager"
	"github.com/mrlutik/kira2.0/internal/osutils"

	"github.com/mrlutik/kira2.0/internal/logging"

	"github.com/spf13/cobra"
)

const (
	// Command information
	use   = "firewall"
	short = "Setting up firewalld"
	long  = "Setting up firewalld"

	// Flags
	closingPortFlag = "close-ports"
	openingPortFlag = "open-ports"
	defaultFlag     = "default"
)

var ErrOnlyOneFlagAllowed = errors.New("only one flag at a time is allowed")

func Firewall(log *logging.Logger) *cobra.Command {
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
			mainFirewall(cmd, log)
		},
	}

	firewallCmd.AddCommand(openport.OpenPort(log))
	firewallCmd.AddCommand(closeport.ClosePort(log))
	firewallCmd.AddCommand(blacklist.Blacklist(log))
	firewallCmd.AddCommand(whitelist.Whitelist(log))

	firewallCmd.Flags().Bool(closingPortFlag, false, "Set this flag to block all ports except ssh")
	firewallCmd.Flags().Bool(openingPortFlag, false, "Set this flag to open all km2 default ports")

	firewallCmd.Flags().Bool(defaultFlag, false, "Set this flag to restore default setting for firewall (what km2 is set after node installation)")

	return firewallCmd
}

func validateFlags(cmd *cobra.Command) error {
	return nil
}

func mainFirewall(cmd *cobra.Command, log *logging.Logger) {
	log.Info("Validating flags")

	openPorts, err := cmd.Flags().GetBool(openingPortFlag)
	if err != nil {
		log.Fatalf("Cannot get '%s' flag: %s", openingPortFlag, err)
	}

	closePorts, err := cmd.Flags().GetBool(closingPortFlag)
	if err != nil {
		log.Fatalf("Cannot get '%s' flag: %s", closingPortFlag, err)
	}

	defaultB, err := cmd.Flags().GetBool(defaultFlag)
	if err != nil {
		log.Fatalf("Cannot get '%s' flag: %s", defaultFlag, err)
	}

	err = validateBoolFlags(openPorts, closePorts, defaultB)
	if err != nil {
		log.Fatalf("Only 1 flag can be accepted: %s", err)
	}

	utilsOS := osutils.NewOSUtils(log)

	configController := controller.NewConfigController(utilsOS, log)

	kiraCfg, err := configController.ReadOrCreateConfig()
	if err != nil {
		log.Fatalf("Reading config file failed: %s", err)
	}

	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Can't initialize the Docker client: %s", err)
	}

	dockerManager := docker.NewTestDockerManager(client, utilsOS, log)
	if err != nil {
		log.Fatalf("Can't create instance of docker manager: %s", err)
	}
	defer dockerManager.CloseClient()

	firewallManager, err := manager.NewFirewallManager(dockerManager, utilsOS, kiraCfg, log)
	if err != nil {
		log.Fatalf("Initialization of firewall manager failed: %s", err)
	}

	// TODO make flexible setting timeout
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute*1)
	defer cancelFunc()

	switch {
	case closePorts:
		err = firewallManager.CloseAllOpenedPorts(ctx)
		if err != nil {
			log.Fatalf("Closing ports failed: %s", err)
		}
	case openPorts:
		err = firewallManager.OpenConfigPorts(ctx)
		if err != nil {
			log.Fatalf("Opening ports failed: %s", err)
		}
	case defaultB:
		err = firewallManager.SetUpFirewall(ctx)
		if err != nil {
			log.Fatalf("Setup for firewall failed: %s", err)
		}
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
