package openPort

import (
	"errors"
	"fmt"

	errUtils "github.com/mrlutik/kira2.0/internal/errors"
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

var (
	log = logging.Log

	ErrWrongPortType = errors.New("wrong port type, can only be <tcp> or <udp>")
)

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
		return fmt.Errorf("error retrieving '%s' flag: %w", portFlag, err)
	}

	if osutils.CheckIfPortIsValid(portToOpen) {
		return fmt.Errorf("cannot parse port <%v>: %w", portToOpen, err)
	}

	portType, err := cmd.Flags().GetString(typeOfIPFlag)
	if err != nil {
		return fmt.Errorf("error retrieving '%s' flag: %w", typeOfIPFlag, err)
	}
	if portType != "tcp" && portType != "udp" {
		return fmt.Errorf("%w: current value is '%s'", ErrWrongPortType, portType)
	}

	return nil
}

func mainOpenPort(cmd *cobra.Command) {
	var port types.Port
	var err error

	port.Port, err = cmd.Flags().GetString("port")
	errUtils.HandleFatalErr("cannot get port flag", err)
	port.Type, err = cmd.Flags().GetString("type")
	errUtils.HandleFatalErr("cannot get type flag", err)

	fc := firewallController.NewFireWalldController("validator")
	log.Infof("Adding %s port with %s type\n", port.Port, port.Type)
	_, err = fc.OpenPort(port, fc.ZoneName)
	errUtils.HandleFatalErr(fmt.Sprintf("error while opening port %v", port), err)

	_, err = fc.ReloadFirewall()
	errUtils.HandleFatalErr("error while reloading firewall", err)
}
