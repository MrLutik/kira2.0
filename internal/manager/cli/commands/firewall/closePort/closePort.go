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
	// Command information
	use   = "close-port"
	short = "subcommand for port closing"
	long  = "long description"

	// Flags
	portFlag     = "port"
	typeOfIPFlag = "type"
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
				if err := cmd.Help(); err != nil {
					log.Fatalf("Error displaying help: %s", err)
				}
				return
			}
			mainClosePort(cmd)
		},
	}

	closePortCmd.Flags().String(portFlag, "", "Port to close (between 0 and 65535)")
	closePortCmd.Flags().String(typeOfIPFlag, "", "<tcp> or <udp>")

	if err := closePortCmd.MarkFlagRequired(portFlag); err != nil {
		log.Fatalf("Failed to mark '%s' flag as required: %s", portFlag, err)
	}
	if err := closePortCmd.MarkFlagRequired(typeOfIPFlag); err != nil {
		log.Fatalf("Failed to mark '%s' flag as required: %s", typeOfIPFlag, err)
	}

	return closePortCmd
}

func validateFlags(cmd *cobra.Command) error {
	portToClose, err := cmd.Flags().GetString(portFlag)
	if err != nil {
		return fmt.Errorf("error retrieving '%s' flag: %w", portFlag, err)
	}
	check, err := osutils.CheckIfPortIsValid(portToClose)
	if err != nil || !check {
		return fmt.Errorf("cannot parse port <%v>: %w", portToClose, err)
	}

	portType, err := cmd.Flags().GetString(typeOfIPFlag)
	if err != nil {
		return fmt.Errorf("error retrieving '%s' flag: %w", typeOfIPFlag, err)
	}
	if portType != "tcp" && portType != "udp" {
		return fmt.Errorf("wrong port type: <%s>, can only be <tcp> or <udp>", portType)
	}

	return nil
}

func mainClosePort(cmd *cobra.Command) {
	var port types.Port
	var err error

	port.Port, err = cmd.Flags().GetString(portFlag)
	errors.HandleFatalErr(fmt.Sprintf("cannot get '%s' flag", portFlag), err)
	port.Type, err = cmd.Flags().GetString(typeOfIPFlag)
	errors.HandleFatalErr(fmt.Sprintf("cannot get '%s' flag", typeOfIPFlag), err)

	fc := firewallController.NewFireWalldController("validator")
	log.Infof("Adding %s port with %s type\n", port.Port, port.Type)
	_, err = fc.ClosePort(port, fc.ZoneName)
	errors.HandleFatalErr(fmt.Sprintf("error while closing port %v", port), err)

	_, err = fc.ReloadFirewall()
	errors.HandleFatalErr("error while reloading firewall", err)
}
