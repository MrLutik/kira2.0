package closePort

import (
	"fmt"

	"github.com/mrlutik/kira2.0/internal/errors"
	"github.com/mrlutik/kira2.0/internal/firewall/firewallController"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/osutils"
	"github.com/mrlutik/kira2.0/internal/types"
	"github.com/spf13/cobra"
)

const (
	use   = "close-port"
	short = "subcommand for port closing"
	long  = "long description"
)

var log = logging.Log

func ClosePort() *cobra.Command {
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
			mainClosePort(cmd)
		},
	}

	closePortCmd.Flags().String("port", "", "Port to close (between 0 and 65535)")
	closePortCmd.Flags().String("type", "", "<tcp> or <udp>")

	closePortCmd.MarkFlagRequired("port")
	closePortCmd.MarkFlagRequired("type")

	return closePortCmd
}

func validateFlags(cmd *cobra.Command) error {
	portToClose, err := cmd.Flags().GetString("port")
	if err != nil {
		return fmt.Errorf("error retrieving 'port' flag: %s", err)
	}
	check, err := osutils.CheckIfPortIsValid(portToClose)
	if err != nil || !check {
		return fmt.Errorf("cannot parse port <%v>: %s", portToClose, err)
	}

	portType, err := cmd.Flags().GetString("type")
	if err != nil {
		return fmt.Errorf("error retrieving 'type' flag: %s", err)
	}
	if portType != "tcp" && portType != "udp" {
		return fmt.Errorf("wrong port type: <%s>, can only be <tcp> or <udp>", portType)
	}

	return nil
}

func mainClosePort(cmd *cobra.Command) {
	var port types.Port
	var err error

	port.Port, err = cmd.Flags().GetString("port")
	errors.HandleFatalErr("cannot get port flag", err)
	port.Type, err = cmd.Flags().GetString("type")
	errors.HandleFatalErr("cannot get type flag", err)

	fc := firewallController.NewFireWalldController("validator")
	log.Printf("Adding %s port with %s type\n", port.Port, port.Type)
	_, err = fc.ClosePort(port, fc.ZoneName)
	errors.HandleFatalErr(fmt.Sprintf("error while closing port %v", port), err)

	_, err = fc.ReloadFirewall()
	errors.HandleFatalErr("error while reloading firewall", err)
}
