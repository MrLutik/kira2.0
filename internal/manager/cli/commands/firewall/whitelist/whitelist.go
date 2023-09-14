package whitelist

import (
	"fmt"

	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/errors"
	"github.com/mrlutik/kira2.0/internal/firewall/firewallManager"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/osutils"
	"github.com/spf13/cobra"
)

var log = logging.Log

const (
	use   = "whitelist"
	short = "subcommand for whitelisting ips"
	long  = `subcommand for adding to whitelist or removing from whitelist specific ips
example: 
	km2 firewall whitelist --ip 8.8.8.8 --add --log-level debug
	
	km2 firewall whitelist --ip 8.8.8.8 --remove --log-level debug`
)

func Whitelist() *cobra.Command {
	closePortCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, _ []string) {
			if err := validateFlags(cmd); err != nil {
				log.Errorf("Some flag is not valid: %s", err)
				cmd.Help()
				return
			}
			mainWhitelist(cmd)
		},
	}

	closePortCmd.Flags().String("ip", "", "target ip")
	closePortCmd.MarkFlagRequired("ip")

	closePortCmd.Flags().Bool("add", false, "if TRUE adding ip to whitelist")
	closePortCmd.Flags().Bool("remove", false, "if TRUE removing ip from whitelist")

	return closePortCmd
}
func validateFlags(cmd *cobra.Command) error {
	ip, err := cmd.Flags().GetString("ip")
	if err != nil {
		return fmt.Errorf("error retrieving 'ip' flag: %s", err)
	}
	check, err := osutils.CheckIfIPIsValid(ip)
	if err != nil || !check {
		return fmt.Errorf("cannot parse ip <%v>: %s", ip, err)
	}

	add, err := cmd.Flags().GetBool("add")
	if err != nil {
		return fmt.Errorf("error retrieving 'add' flag: %s", err)
	}
	remove, err := cmd.Flags().GetBool("remove")
	if err != nil {
		return fmt.Errorf("error retrieving 'remove' flag: %s", err)
	}

	if add && remove {
		return fmt.Errorf("--add and --remove flags cannot be both true")
	}
	if !add && !remove {
		return fmt.Errorf("--add and --remove flags cannot be both false")
	}

	return nil
}

func mainWhitelist(cmd *cobra.Command) {
	log.Println("mainWhitelist")
	ip, err := cmd.Flags().GetString("ip")
	errors.HandleFatalErr("cannot Get whitelist-ip", err)
	ok, err := osutils.CheckIfIPIsValid(ip)
	errors.HandleFatalErr("cannot check if ip is valid", err)

	add, err := cmd.Flags().GetBool("add")
	if err != nil {
		errors.HandleFatalErr("error retrieving 'add' flag", err)
	}
	remove, err := cmd.Flags().GetBool("remove")
	if err != nil {
		errors.HandleFatalErr("error retrieving 'remove' flag", err)
	}

	kiraCfg := firewallManager.GenerateKiraConfigForFirewallManager()

	if add {
		if ip != "" {
			ok, err := osutils.CheckIfIPIsValid(ip)
			errors.HandleFatalErr("cannot check if ip is valid", err)
			if ok {
				dockerManager, err := docker.NewTestDockerManager()
				errors.HandleFatalErr("Can't create instance of docker manager", err)
				defer dockerManager.Cli.Close()
				fm := firewallManager.NewFirewallManager(dockerManager, kiraCfg)
				fm.FirewallHandler.WhiteListIp(ip, fm.FirewallConfig.ZoneName)
			}
		}
	}
	if remove {
		if ip != "" {
			if ok {
				dockerManager, err := docker.NewTestDockerManager()
				errors.HandleFatalErr("Can't create instance of docker manager", err)
				defer dockerManager.Cli.Close()
				fm := firewallManager.NewFirewallManager(dockerManager, kiraCfg)
				fm.FirewallHandler.RemoveFromWhitelistIP(ip, fm.FirewallConfig.ZoneName)
			}
		}
	}

}
