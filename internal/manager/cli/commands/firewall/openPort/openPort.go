package openPort

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
	use   = "open-port"
	short = "subcommand for port opening"
	long  = "long description"

	// Flags
	portFlag     = "port"
	typeOfIPFlag = "type"
)

var log = logging.Log

func OpenPort() *cobra.Command {
	openPortCmd := &cobra.Command{
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
			mainOpenPort(cmd)
		},
	}

	openPortCmd.Flags().String(portFlag, "", "Port to open (between 0 and 65535)")
	openPortCmd.Flags().String(typeOfIPFlag, "", "<tcp> or <udp>")

	if err := openPortCmd.MarkFlagRequired(portFlag); err != nil {
		log.Fatalf("Failed to mark '%s' flag as required: %s", portFlag, err)
	}
	if err := openPortCmd.MarkFlagRequired(typeOfIPFlag); err != nil {
		log.Fatalf("Failed to mark '%s' flag as required: %s", typeOfIPFlag, err)
	}

	return openPortCmd
}

func validateFlags(cmd *cobra.Command) error {
	portToOpen, err := cmd.Flags().GetString(portFlag)
	if err != nil {
		return fmt.Errorf("error retrieving '%s' flag: %s", portFlag, err)
	}
	check, err := osutils.CheckIfPortIsValid(portToOpen)
	if err != nil || !check {
		return fmt.Errorf("cannot parse port <%v>: %s", portToOpen, err)
	}

	portType, err := cmd.Flags().GetString(typeOfIPFlag)
	if err != nil {
		return fmt.Errorf("error retrieving '%s' flag: %s", typeOfIPFlag, err)
	}
	if portType != "tcp" && portType != "udp" {
		return fmt.Errorf("wrong port type: <%s>, can only be <tcp> or <udp>", portType)
	}

	return nil
}

func mainOpenPort(cmd *cobra.Command) {
	var port types.Port
	var err error

	port.Port, err = cmd.Flags().GetString("port")
	errors.HandleFatalErr("cannot get port flag", err)
	port.Type, err = cmd.Flags().GetString("type")
	errors.HandleFatalErr("cannot get type flag", err)

	fc := firewallController.NewFireWalldController("validator")
	log.Printf("Adding %s port with %s type\n", port.Port, port.Type)
	_, err = fc.OpenPort(port, fc.ZoneName)
	errors.HandleFatalErr(fmt.Sprintf("error while opening port %v", port), err)

	_, err = fc.ReloadFirewall()
	errors.HandleFatalErr("error while reloading firewall", err)
}
