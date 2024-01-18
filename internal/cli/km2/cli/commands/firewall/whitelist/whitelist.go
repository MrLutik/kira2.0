package whitelist

import (
	"errors"
	"fmt"

	"github.com/mrlutik/kira2.0/internal/config/configFileController"
	"github.com/mrlutik/kira2.0/internal/docker"
	errUtils "github.com/mrlutik/kira2.0/internal/errors"
	"github.com/mrlutik/kira2.0/internal/firewall/firewallManager"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/osutils"
	"github.com/spf13/cobra"
)

const (
	use   = "whitelist"
	short = "subcommand for whitelisting ips"
	long  = `subcommand for adding to whitelist or removing from whitelist specific ips
	example: 
	km2 firewall whitelist --ip 8.8.8.8 --add --log-level debug
	
	km2 firewall whitelist --ip 8.8.8.8 --remove --log-level debug`

	// Flags
	ipFlag       = "ip"
	addingFlag   = "add"
	removingFlag = "remove"
)

var (
	log = logging.Log

	ErrConflictingFlags = errors.New("conflicting flags")
)

func Whitelist() *cobra.Command {
	closePortCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, _ []string) {
			if err := validateFlags(cmd); err != nil {
				log.Errorf("Some flag is not valid: %s", err)
				if err := cmd.Help(); err != nil {
					log.Fatalf("Error displaying help: %s", err)
				}
				return
			}
			mainWhitelist(cmd)
		},
	}

	closePortCmd.Flags().String(ipFlag, "", "target ip")
	if err := closePortCmd.MarkFlagRequired(ipFlag); err != nil {
		log.Fatalf("Failed to mark '%s' flag as required: %s", ipFlag, err)
	}

	closePortCmd.Flags().Bool(addingFlag, false, "if TRUE adding ip to whitelist")
	closePortCmd.Flags().Bool(removingFlag, false, "if TRUE removing ip from whitelist")

	return closePortCmd
}

func validateFlags(cmd *cobra.Command) error {
	ip, err := cmd.Flags().GetString(ipFlag)
	if err != nil {
		return fmt.Errorf("error retrieving '%s' flag: %w", ipFlag, err)
	}
	check, err := osutils.CheckIfIPIsValid(ip)
	if err != nil || !check {
		return fmt.Errorf("cannot parse ip <%v>: %w", ip, err)
	}

	add, err := cmd.Flags().GetBool(addingFlag)
	if err != nil {
		return fmt.Errorf("error retrieving '%s' flag: %w", addingFlag, err)
	}
	remove, err := cmd.Flags().GetBool(removingFlag)
	if err != nil {
		return fmt.Errorf("error retrieving '%s' flag: %w", removingFlag, err)
	}

	if add && remove {
		return fmt.Errorf("%w: --%s and --%s flags cannot be both 'true'", ErrConflictingFlags, addingFlag, removingFlag)
	}
	if !add && !remove {
		return fmt.Errorf("%w: --%s and --%s flags cannot be both 'false'", ErrConflictingFlags, addingFlag, removingFlag)
	}

	return nil
}

func mainWhitelist(cmd *cobra.Command) {
	ip, err := cmd.Flags().GetString(ipFlag)
	errUtils.HandleFatalErr("cannot Get whitelist-ip", err)
	ok, err := osutils.CheckIfIPIsValid(ip)
	errUtils.HandleFatalErr("cannot check if ip is valid", err)

	add, err := cmd.Flags().GetBool(addingFlag)
	if err != nil {
		errUtils.HandleFatalErr(fmt.Sprintf("error retrieving '%s' flag", addingFlag), err)
	}
	remove, err := cmd.Flags().GetBool(removingFlag)
	if err != nil {
		errUtils.HandleFatalErr(fmt.Sprintf("error retrieving '%s' flag", removingFlag), err)
	}

	kiraCfg, err := configFileController.ReadOrCreateConfig()
	errUtils.HandleFatalErr("Error while reading cfg file", err)

	if add {
		if ip != "" {
			isValidIP, err := osutils.CheckIfIPIsValid(ip)
			errUtils.HandleFatalErr("cannot check if ip is valid", err)
			if isValidIP {
				dockerManager, err := docker.NewTestDockerManager()
				errUtils.HandleFatalErr("Can't create instance of docker manager", err)
				defer dockerManager.Cli.Close()

				fm := firewallManager.NewFirewallManager(dockerManager, kiraCfg)
				err = fm.FirewallHandler.WhiteListIp(ip, fm.FirewallConfig.ZoneName)
				errUtils.HandleFatalErr("Can't whitelist IP", err)
			}
		}
	}
	if remove {
		if ip != "" {
			if ok {
				dockerManager, err := docker.NewTestDockerManager()
				errUtils.HandleFatalErr("Can't create instance of docker manager", err)
				defer dockerManager.Cli.Close()

				fm := firewallManager.NewFirewallManager(dockerManager, kiraCfg)
				err = fm.FirewallHandler.RemoveFromWhitelistIP(ip, fm.FirewallConfig.ZoneName)
				errUtils.HandleFatalErr("Can't remove IP from whitelist firewall", err)
			}
		}
	}
}