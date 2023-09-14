package firewall

import (
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/errors"
	"github.com/mrlutik/kira2.0/internal/firewall/firewallManager"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/manager/cli/commands/firewall/blacklist"
	"github.com/mrlutik/kira2.0/internal/manager/cli/commands/firewall/closePort"
	openport "github.com/mrlutik/kira2.0/internal/manager/cli/commands/firewall/openPort"
	"github.com/mrlutik/kira2.0/internal/manager/cli/commands/firewall/whitelist"
	"github.com/mrlutik/kira2.0/internal/osutils"
	"github.com/spf13/cobra"
)

var log = logging.Log

const (
	use   = "firewall"
	short = "Seting up firewalld"
	long  = "Seting up firewalld"
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
				cmd.Help()
				return
			}
			mainFirewall(cmd)
		},
	}

	firewallCmd.AddCommand(openport.OpenPort())
	firewallCmd.AddCommand(closePort.ClosePort())
	firewallCmd.AddCommand(blacklist.Blacklist())
	firewallCmd.AddCommand(whitelist.Whitelist())

	firewallCmd.Flags().Bool("drop-all", false, "Set this flag to block all ports except ssh")
	firewallCmd.Flags().Bool("allow-all", false, "Set this flag to open all ports")

	firewallCmd.Flags().Bool("default", false, "Set this flag to restore default setting for firewall (what km2 is set after node installation)")

	return firewallCmd
}

func validateFlags(cmd *cobra.Command) error {
	log.Println("validateFlags")
	// blacklistIP, err := cmd.Flags().GetString("blacklist-ip")
	// if err != nil {
	// 	return fmt.Errorf("error retrieving 'ip' flag: %s", err)
	// }
	// if blacklistIP != "" {
	// 	_, err = osutils.CheckIfIPIsValid(blacklistIP)
	// 	if err != nil {
	// 		return fmt.Errorf("cannot parse ip: %s", err)
	// 	}
	// }

	// whitelistIP, err := cmd.Flags().GetString("whitelist-ip")
	// if err != nil {
	// 	return fmt.Errorf("error retrieving 'ip' flag: %s", err)
	// }

	// if whitelistIP != "" {
	// 	_, err = osutils.CheckIfIPIsValid(whitelistIP)
	// 	if err != nil {
	// 		return fmt.Errorf("cannot parse ip: %s", err)
	// 	}
	// }

	return nil
}

func mainFirewall(cmd *cobra.Command) {
	blacklistIP, err := cmd.Flags().GetString("blacklist-ip")
	errors.HandleFatalErr("cannot Get blacklist-ip", err)

	kiraCfg := firewallManager.GenerateKiraConfigForFirewallManager()

	if blacklistIP != "" {
		ok, err := osutils.CheckIfIPIsValid(blacklistIP)
		errors.HandleFatalErr("cannot check if ip is valid", err)

		if ok {
			dockerManager, err := docker.NewTestDockerManager()
			errors.HandleFatalErr("Can't create instance of docker manager", err)
			defer dockerManager.Cli.Close()
			fm := firewallManager.NewFirewallManager(dockerManager, kiraCfg)
			fm.FirewallHandler.BlackListIP(blacklistIP, fm.FirewallConfig.ZoneName)
		}
	}

	whitelistIP, err := cmd.Flags().GetString("whitelist-ip")
	errors.HandleFatalErr("cannot Get whitelist-ip", err)

	if whitelistIP != "" {
		log.Printf("Checking %s ip is its valid\n", whitelistIP)
		ok, err := osutils.CheckIfIPIsValid(whitelistIP)
		errors.HandleFatalErr("cannot check if ip is valid", err)

		if ok {
			dockerManager, err := docker.NewTestDockerManager()
			errors.HandleFatalErr("Can't create instance of docker manager", err)
			defer dockerManager.Cli.Close()
			fm := firewallManager.NewFirewallManager(dockerManager, kiraCfg)
			fm.FirewallHandler.WhiteListIp(blacklistIP, fm.FirewallConfig.ZoneName)
		}

	}

}
