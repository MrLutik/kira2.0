package manager

import (
	"context"
	"errors"
	"fmt"
	"os/exec"

	dockerTypes "github.com/docker/docker/api/types"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/firewall/controller"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/osutils"
	"github.com/mrlutik/kira2.0/internal/types"
)

type (
	FirewallManager struct {
		controllerKiraZone *controller.FirewallDController
		dockerMaintenance  DockerMaintenance
		firewallConfig     *FirewallConfig
		kiraConfig         *config.KiraConfig
		utils              *osutils.OSUtils

		log *logging.Logger
	}

	DockerMaintenance interface {
		RestartDockerService() error
		VerifyDockerInstallation(context.Context) error
		NetworkInspect(context.Context, string) (*dockerTypes.NetworkResource, error)
	}

	FirewallConfig struct {
		zoneName    string
		portsToOpen []types.Port
	}
)

var ErrFirewallDNotInstalled = errors.New("firewalld is not installed on the system")

// Port range: 0 - 65535
// Type: udp or tcp
//
// Example: 39090 tcp or 53 udp

func NewFirewallConfig(kiraCfg *config.KiraConfig) *FirewallConfig {
	return &FirewallConfig{
		zoneName: "validator",
		portsToOpen: []types.Port{
			{Port: kiraCfg.InterxPort, Type: "tcp"},
			{Port: kiraCfg.GrpcPort, Type: "tcp"},
			{Port: kiraCfg.P2PPort, Type: "tcp"},
			{Port: kiraCfg.PrometheusPort, Type: "tcp"},
			{Port: kiraCfg.RpcPort, Type: "tcp"},
		},
	}
}

func NewFirewallManager(dockerMaintenance DockerMaintenance, osUtils *osutils.OSUtils, kiraConfig *config.KiraConfig, logger *logging.Logger) (*FirewallManager, error) {
	firewallConfig := NewFirewallConfig(kiraConfig)
	controller, err := controller.NewFirewallDController(osUtils, firewallConfig.zoneName)
	if err != nil {
		return nil, fmt.Errorf("initialization of the firewall manager error: %w", err)
	}

	return &FirewallManager{
		controllerKiraZone: controller,
		dockerMaintenance:  dockerMaintenance,
		firewallConfig:     firewallConfig,
		kiraConfig:         kiraConfig,
		utils:              osUtils,
		log:                logger,
	}, nil
}

func (f *FirewallManager) CheckFirewallSetUp(ctx context.Context) (bool, error) {
	_, err := exec.LookPath("firewall-cmd")
	if err != nil {
		return false, ErrFirewallDNotInstalled
	}

	// Checking if validator zone exist
	check, err := f.checkFirewallZone(f.firewallConfig.zoneName)
	if err != nil {
		return false, fmt.Errorf("error while checking validator zone: %w", err)
	}
	if !check {
		return false, nil
	}
	return true, nil
}

func (f *FirewallManager) SetUpFirewall(ctx context.Context) error {
	f.log.Infof("Firewall setup:")

	f.log.Infof("Restarting docker service")
	err := f.dockerMaintenance.RestartDockerService()
	if err != nil {
		return fmt.Errorf("cannot restart docker service: %w", err)
	}

	f.log.Infof("Checking if docker is running")
	err = f.dockerMaintenance.VerifyDockerInstallation(ctx)
	if err != nil {
		return fmt.Errorf("verifying docker installation failed: %w", err)
	}

	// TODO DOCKER ZONE
	f.log.Infof("Checking and deleting default docker zone")
	dockerZoneName := "docker"
	check, err := f.checkFirewallZone(dockerZoneName)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if check {
		f.log.Infof("docker zone exist: %s, deleting", dockerZoneName)

		var output string
		// TODO Docker firewall!
		output, err = f.controllerKiraZone.DeleteFirewallZone()
		if err != nil {
			return fmt.Errorf("%s\n%w", output, err)
		}
		output, err = f.controllerKiraZone.ReloadFirewall()
		if err != nil {
			return fmt.Errorf("%s\n%w", output, err)
		}
	} else {
		f.log.Infof("zone '%s' does not exist", dockerZoneName)
	}
	// TODO DOCKER ZONE

	f.log.Infof("Checking if '%s' zone exists", f.firewallConfig.zoneName)
	check, err = f.checkFirewallZone(f.firewallConfig.zoneName)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if !check {
		f.log.Infof("Creating new firewall zone '%s'", f.firewallConfig.zoneName)

		var output string
		output, err = f.controllerKiraZone.CreateNewZone()
		if err != nil {
			return fmt.Errorf("%s\n%w", output, err)
		}
		output, err = f.controllerKiraZone.ReloadFirewall()
		if err != nil {
			return fmt.Errorf("%s\n%w", output, err)
		}
	}

	f.log.Infof("Switching into '%s' firewall zone", f.firewallConfig.zoneName)
	o, err := f.controllerKiraZone.ChangeDefaultZone()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	o, err = f.controllerKiraZone.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	f.log.Infof("Checking ports %+v ", f.firewallConfig.portsToOpen)
	err = f.checkPorts(f.firewallConfig.portsToOpen)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	f.log.Infof("Opening system ports")
	systemPorts := []types.Port{
		{Port: "22", Type: "tcp"},
		{Port: "53", Type: "udp"},
	}
	err = f.OpenPorts(systemPorts)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	f.log.Infof("Opening kira ports")
	err = f.OpenPorts(f.firewallConfig.portsToOpen)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	// Adding interface that has internet access
	f.log.Infof("Adding interface that has internet access")
	internetInterface := f.utils.GetInternetInterface()
	o, err = f.controllerKiraZone.AddInterfaceToTheZone(internetInterface)
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	dockerInterface, err := f.getDockerNetworkInterface(ctx, f.kiraConfig.DockerNetworkName)
	interfaceName := "br-" + dockerInterface.ID[0:11]
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	f.log.Infof("Adding %s interface to the zone and enabling routing", interfaceName)
	o, err = f.controllerKiraZone.AddInterfaceToTheZone(interfaceName)
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	f.log.Debugf("issuing docker interface subnet")

	dockerInterfaceConfig := dockerInterface.IPAM.Config
	f.log.Debugf("docker interace subnet: %s", dockerInterfaceConfig[0].Subnet)
	o, err = f.controllerKiraZone.AddRichRule(fmt.Sprintf("rule family=ipv4 source address=%s accept", dockerInterfaceConfig[0].Subnet))
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	o, err = f.controllerKiraZone.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	o, err = f.controllerKiraZone.TurnOnMasquerade()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	// Adding docker to the zone and enabling routing
	f.log.Infof("Adding docker0 interface to the zone and enabling routing")
	o, err = f.controllerKiraZone.AddInterfaceToTheZone("docker0")
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	o, err = f.controllerKiraZone.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	f.log.Infof("Reloading firewall")
	o, err = f.controllerKiraZone.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	f.log.Infof("Restarting docker service")
	err = f.dockerMaintenance.RestartDockerService()
	if err != nil {
		return fmt.Errorf("cannot restart docker service: %w", err)
	}
	return nil
}

// closing all opened ports in default km2 firewalld's zone except 22 and 53 (ssh and dns port)
func (f *FirewallManager) CloseAllOpenedPorts(ctx context.Context) error {
	_, ports, err := f.controllerKiraZone.GetOpenedPorts()
	if err != nil {
		return fmt.Errorf("cannot get opened ports: %w", err)
	}

	var portsToClose []types.Port

	for _, p := range ports {
		var port types.Port
		port, err = f.convertFirewallDPortToKM2Port(p)
		if err != nil {
			return fmt.Errorf("convert port: %w", err)
		}
		portsToClose = append(portsToClose, port)
	}

	err = f.checkPorts(portsToClose)
	if err != nil {
		return fmt.Errorf("cannot while checking ports are valid: %w", err)
	}

	err = f.ClosePorts(portsToClose)
	if err != nil {
		return fmt.Errorf("cannot close ports: %w", err)
	}

	o, err := f.controllerKiraZone.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	return nil
}

func (f *FirewallManager) OpenConfigPorts(ctx context.Context) error {
	err := f.checkPorts(f.firewallConfig.portsToOpen)
	if err != nil {
		return fmt.Errorf("cannot while checking ports are valid: %w", err)
	}
	err = f.OpenPorts(f.firewallConfig.portsToOpen)
	if err != nil {
		return fmt.Errorf("cannot open ports: %w", err)
	}
	o, err := f.controllerKiraZone.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	return nil
}

func (f *FirewallManager) RestoreDefaultFirewalldSettingsForKM2(ctx context.Context) error {
	check, err := f.checkFirewallZone(f.firewallConfig.zoneName)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if check {
		var output string
		f.log.Infof("docker zone exist: %s, deleting", f.firewallConfig.zoneName)
		output, err = f.controllerKiraZone.DeleteFirewallZone()
		if err != nil {
			return fmt.Errorf("%s\n%w", output, err)
		}
		output, err = f.controllerKiraZone.ReloadFirewall()
		if err != nil {
			return fmt.Errorf("%s\n%w", output, err)
		}
	}

	err = f.SetUpFirewall(ctx)
	if err != nil {
		return fmt.Errorf("setting up firewall failed: %w", err)
	}
	return nil
}
