package blacklist

import (
	"errors"
	"fmt"

	"github.com/docker/docker/client"
	"github.com/mrlutik/kira2.0/internal/config/controller"
	"github.com/mrlutik/kira2.0/internal/config/handler"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/firewall/manager"
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

var ErrConflictingFlags = errors.New("conflicting flags")

func Blacklist(log *logging.Logger) *cobra.Command {
	blacklistCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, _ []string) {
			if err := validateFlags(cmd, log); err != nil {
				log.Errorf("Some flag is not valid: %s", err)
				if err := cmd.Help(); err != nil {
					log.Fatalf("Error displaying help: %s", err)
				}
				return
			}
			mainBlacklist(cmd, log)
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

func validateFlags(cmd *cobra.Command, log *logging.Logger) error {
	ipValue, err := cmd.Flags().GetString(ipFlag)
	if err != nil {
		return fmt.Errorf("error retrieving '%s' flag: %w", ipFlag, err)
	}

	utilsOS := osutils.NewOSUtils(log)

	check, err := utilsOS.CheckIfIPIsValid(ipValue)
	if err != nil || !check {
		return fmt.Errorf("cannot parse ip '%s': %w", ipValue, err)
	}

	add, err := cmd.Flags().GetBool(addingFlag)
	if err != nil {
		return fmt.Errorf("retrieving '%s' flag failed: %w", addingFlag, err)
	}
	remove, err := cmd.Flags().GetBool(removingFlag)
	if err != nil {
		return fmt.Errorf("retrieving '%s' flag failed: %w", removingFlag, err)
	}

	if add && remove {
		return fmt.Errorf("%w: --%s and --%s flags cannot be both 'true'", ErrConflictingFlags, addingFlag, removingFlag)
	}
	if !add && !remove {
		return fmt.Errorf("%w: --%s and --%s flags cannot be both 'false'", ErrConflictingFlags, addingFlag, removingFlag)
	}

	return nil
}

func mainBlacklist(cmd *cobra.Command, log *logging.Logger) {
	ip, err := cmd.Flags().GetString(ipFlag)
	if err != nil {
		log.Fatalf("Cannot get blacklist-ip: %s", err)
	}

	utilsOS := osutils.NewOSUtils(log)

	isIpValid, err := utilsOS.CheckIfIPIsValid(ip)
	if err != nil {
		log.Fatalf("Cannot check if ip is valid: %s", err)
	}

	add, err := cmd.Flags().GetBool(addingFlag)
	if err != nil {
		log.Fatalf("Retrieving '%s' flag failed: %s", addingFlag, err)
	}

	remove, err := cmd.Flags().GetBool(removingFlag)
	if err != nil {
		log.Fatalf("Retrieving '%s' flag failed: %s", removingFlag, err)
	}

	configController := controller.NewConfigController(handler.NewHandler(utilsOS, log), utilsOS, log)

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

	if add {
		if ip != "" {
			if isIpValid {
				err = firewallManager.BlackListIP(ip)
				if err != nil {
					log.Fatalf("Can't blacklist IP: %s", err)
				}
			}
		}
	}
	if remove {
		if ip != "" {
			if isIpValid {
				err = firewallManager.RemoveFromBlackListIP(ip)
				if err != nil {
					log.Fatalf("Can't remove IP from blacklist firewall: %s", err)
				}
			}
		}
	}
}
