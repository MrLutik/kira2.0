package blacklist

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
	use   = "blacklist"
	short = "subcommand for blacklisting ips"
	long  = `subcommand for adding to blacklist or removing from blacklist specific ips
	example: 
	km2 firewall blacklist --ip 8.8.8.8 --add --log-level debug
	
	km2 firewall blacklist --ip 8.8.8.8 --remove --log-level debug`

	// Flags
	ipFlag       = "ip"
	addingFlag   = "add"
	removingFlag = "remove"
)

var (
	log = logging.Log

	ErrConflictingFlags = errors.New("conflicting flags")
)

func Blacklist() *cobra.Command {
	blacklistCmd := &cobra.Command{
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
			mainBlacklist(cmd)
		},
	}

	blacklistCmd.Flags().String(ipFlag, "", "target ip")
	if err := blacklistCmd.MarkFlagRequired(ipFlag); err != nil {
		log.Fatalf("Failed to mark '%s' flag as required: %s", ipFlag, err)
	}

	blacklistCmd.Flags().Bool(addingFlag, false, "if TRUE adding ip to blacklist")
	blacklistCmd.Flags().Bool(removingFlag, false, "if TRUE removing ip from blacklist")

	return blacklistCmd
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

	add, err := cmd.Flags().GetBool("add")
	if err != nil {
		return fmt.Errorf("error retrieving '%s' flag: %w", addingFlag, err)
	}
	remove, err := cmd.Flags().GetBool("remove")
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

func mainBlacklist(cmd *cobra.Command) {
	ip, err := cmd.Flags().GetString("ip")
	errUtils.HandleFatalErr("cannot Get blacklist-ip", err)
	ok, err := osutils.CheckIfIPIsValid(ip)
	errUtils.HandleFatalErr("cannot check if ip is valid", err)

	add, err := cmd.Flags().GetBool("add")
	if err != nil {
		errUtils.HandleFatalErr("error retrieving 'add' flag", err)
	}
	remove, err := cmd.Flags().GetBool("remove")
	if err != nil {
		errUtils.HandleFatalErr("error retrieving 'remove' flag", err)
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
				err = fm.FirewallHandler.BlackListIP(ip, fm.FirewallConfig.ZoneName)
				errUtils.HandleFatalErr("Can't blacklist IP", err)
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
				err = fm.FirewallHandler.RemoveFromBlackListIP(ip, fm.FirewallConfig.ZoneName)
				errUtils.HandleFatalErr("Can't remove IP from blacklist firewall", err)
			}
		}
	}
}
