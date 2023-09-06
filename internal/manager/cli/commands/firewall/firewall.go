package firewall

import (
	"fmt"

	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/manager/cli/commands/firewall/closePort"
	openport "github.com/mrlutik/kira2.0/internal/manager/cli/commands/firewall/openPort"
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
				log.Errorf("Some flag is not valid: %s", err)
				cmd.Help()
				return
			}
			mainFirewall(cmd)
		},
	}

	firewallCmd.AddCommand(openport.OpenPort())
	firewallCmd.AddCommand(closePort.ClosePort())

	firewallCmd.Flags().String("blacklist-ip", "", "IP address to block")
	firewallCmd.Flags().String("whitelist-ip", "", "IP address to allow")
	firewallCmd.Flags().Bool("drop-all", false, "Set this flag to block all ports except those needed for node functioning")
	firewallCmd.Flags().Bool("allow-all", false, "Set this flag to open all ports, after installation is opened by default")

	firewallCmd.Flags().Bool("default", false, "Set this flag to restore default setting for firewall (what km2 is set after node installation)")

	return firewallCmd
}
func validateFlags(cmd *cobra.Command) error {
	blacklistIP, err := cmd.Flags().GetString("blacklist-ip")
	if err != nil {
		return fmt.Errorf("error retrieving 'ip' flag: %s", err)
	}
	_, err = osutils.CheckIfIPIsValid(blacklistIP)
	if err != nil {
		return fmt.Errorf("cannot parse ip: %s", err)
	}

	whitelistIP, err := cmd.Flags().GetString("whitelist-ip")
	if err != nil {
		return fmt.Errorf("error retrieving 'ip' flag: %s", err)
	}
	_, err = osutils.CheckIfIPIsValid(whitelistIP)
	if err != nil {
		return fmt.Errorf("cannot parse ip: %s", err)
	}

	portToClose, err := cmd.Flags().GetString("close-port")
	if err != nil {
		return fmt.Errorf("error retrieving 'ip' flag: %s", err)
	}
	_, err = osutils.CheckIfPortIsValid(portToClose)
	if err != nil {
		return fmt.Errorf("cannot parse port <%s>: %s", portToClose, err)
	}

	portToOpen, err := cmd.Flags().GetString("open-port")
	if err != nil {
		return fmt.Errorf("error retrieving 'ip' flag: %s", err)
	}
	_, err = osutils.CheckIfPortIsValid(portToOpen)
	if err != nil {
		return fmt.Errorf("cannot parse port <%s>: %s", portToOpen, err)
	}
	return nil
}

func mainFirewall(cmd *cobra.Command) {

}
