package closeport

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

var ErrWrongPortType = errors.New("wrong port type, can only be 'tcp' or 'udp'")

func ClosePort(log *logging.Logger) *cobra.Command {
	closePortCmd := &cobra.Command{
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
			mainClosePort(cmd, log)
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

func validateFlags(cmd *cobra.Command, log *logging.Logger) error {
	portToClose, err := cmd.Flags().GetString(portFlag)
	if err != nil {
		return fmt.Errorf("error retrieving '%s' flag: %w", portFlag, err)
	}

	utilsOS := osutils.NewOSUtils(log)

	if utilsOS.CheckIfPortIsValid(portToClose) {
		return fmt.Errorf("cannot parse port '%s': %w", portToClose, err)
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

func mainClosePort(cmd *cobra.Command, log *logging.Logger) {
	var port types.Port
	var err error

	port.Port, err = cmd.Flags().GetString(portFlag)
	if err != nil {
		log.Fatalf("Cannot get '%s' flag: %s", portFlag, err)
	}

	port.Type, err = cmd.Flags().GetString(typeOfIPFlag)
	if err != nil {
		log.Fatalf("Cannot get '%s' flag: %s", typeOfIPFlag, err)
	}

	utilsOS := osutils.NewOSUtils(log)

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

	log.Infof("Adding '%s' port with '%s' type to closed list", port.Port, port.Type)
	err = firewallManager.ClosePorts([]types.Port{port})
	if err != nil {
		log.Fatalf("Closing port '%s' failed: %s", port, err)
	}
}
