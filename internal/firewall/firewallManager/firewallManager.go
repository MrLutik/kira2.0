package firewallManager

import (
	"context"
	"errors"
	"fmt"
	"os/exec"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/firewall/firewallController"
	"github.com/mrlutik/kira2.0/internal/firewall/firewallManager/firewallHandler"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/osutils"
	"github.com/mrlutik/kira2.0/internal/types"
)

type (
	FirewallManager struct {
		FirewalldController *firewallController.FirewalldController
		DockerManager       *docker.DockerManager
		FirewallHandler     *firewallHandler.FirewallHandler
		FirewallConfig      *FirewallConfig
		KiraConfig          *config.KiraConfig
	}

	FirewallConfig struct {
		ZoneName    string
		PortsToOpen []types.Port
	}
)

var (
	log = logging.Log

	ErrFirewallDNotInstalled = errors.New("firewalld is not installed on the system")
)

// Port range: 0 - 65535
// Type: udp or tcp
//
// Example: 39090 tcp or 53 udp

func NewFirewallConfig(kiraCfg *config.KiraConfig) *FirewallConfig {
	return &FirewallConfig{
		ZoneName: "validator",
		PortsToOpen: []types.Port{
			{Port: kiraCfg.InterxPort, Type: "tcp"},
			{Port: kiraCfg.GrpcPort, Type: "tcp"},
			{Port: kiraCfg.P2PPort, Type: "tcp"},
			{Port: kiraCfg.PrometheusPort, Type: "tcp"},
			{Port: kiraCfg.RpcPort, Type: "tcp"},
		},
	}
}

func NewFirewallManager(dockerManager *docker.DockerManager, kiraCfg *config.KiraConfig) *FirewallManager {
	cfg := NewFirewallConfig(kiraCfg)
	c := firewallController.NewFireWalldController(cfg.ZoneName)
	h := firewallHandler.NewFirewallHandler(c)

	return &FirewallManager{FirewalldController: c, DockerManager: dockerManager, FirewallHandler: h, FirewallConfig: cfg, KiraConfig: kiraCfg}
}

func (fm *FirewallManager) CheckFirewallSetUp(ctx context.Context) (bool, error) {
	_, err := exec.LookPath("firewall-cmd")
	if err != nil {
		return false, ErrFirewallDNotInstalled
	}

	// Checking if validator zone exist
	check, err := fm.FirewallHandler.CheckFirewallZone(fm.FirewallConfig.ZoneName)
	if err != nil {
		return false, fmt.Errorf("error while checking validator zone %w", err)
	}
	if !check {
		return false, nil
	}
	return true, nil
}

func (fm *FirewallManager) SetUpFirewall(ctx context.Context) error {
	log.Infof("***FIREWALL SETUP***\n")

	log.Infof("Restarting docker service\n")
	err := fm.DockerManager.RestartDockerService()
	if err != nil {
		return fmt.Errorf("cannot restart docker service: %w", err)
	}
	log.Infof("Checking if docker is running\n")
	err = fm.DockerManager.VerifyDockerInstallation(ctx)
	if err != nil {
		return err
	}

	log.Infof("checking and deleting default docker zone\n")
	dockerZoneName := "docker"
	check, err := fm.FirewallHandler.CheckFirewallZone(dockerZoneName)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	if check {
		log.Infof("docker zone exist: %s, deleting\n", dockerZoneName)
		o, err := fm.FirewalldController.DeleteFirewallZone(dockerZoneName)
		if err != nil {
			return fmt.Errorf("%s\n%w", o, err)
		}
		o, err = fm.FirewalldController.ReloadFirewall()
		if err != nil {
			return fmt.Errorf("%s\n%w", o, err)
		}
	} else {
		log.Infof("docker zone: %s not exist\n", dockerZoneName)
	}

	log.Infof("checking if %s zone exist\n", fm.FirewallConfig.ZoneName)
	check, err = fm.FirewallHandler.CheckFirewallZone(fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	if !check {
		log.Infof("Creating new firewalldZone %s, check = %v\n ", fm.FirewallConfig.ZoneName, check)
		o, err := fm.FirewalldController.CreateNewFirewalldZone(fm.FirewallConfig.ZoneName)
		if err != nil {
			return fmt.Errorf("%s\n%w", o, err)
		}
		o, err = fm.FirewalldController.ReloadFirewall()
		if err != nil {
			return fmt.Errorf("%s\n%w", o, err)
		}
	}

	log.Infof("Switching into %s firewalldZone\n", fm.FirewallConfig.ZoneName)
	o, err := fm.FirewalldController.ChangeDefaultZone(fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	o, err = fm.FirewalldController.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	log.Infof("Checking ports %+v \n", fm.FirewallConfig.PortsToOpen)
	err = fm.FirewallHandler.CheckPorts(fm.FirewallConfig.PortsToOpen)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	log.Infof("Opening system ports\n")
	sysports := []types.Port{
		{Port: "22", Type: "tcp"},
		{Port: "53", Type: "udp"},
	}
	err = fm.FirewallHandler.OpenPorts(sysports, fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	log.Infof("Opening kira ports\n")
	err = fm.FirewallHandler.OpenPorts(fm.FirewallConfig.PortsToOpen, fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	// Adding interface that has internet acces
	log.Infof("Adding interface that has internet acces\n")
	internetInterface := osutils.GetInternetInterface()
	o, err = fm.FirewalldController.AddInterfaceToTheZone(internetInterface, fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	dockerInterface, err := fm.FirewallHandler.GetDockerNetworkInterface(ctx, fm.KiraConfig.DockerNetworkName, fm.DockerManager)
	interfaceName := "br-" + dockerInterface.ID[0:11]
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	log.Infof("Adding %s interface to the zone and enabling routing\n", interfaceName)
	o, err = fm.FirewalldController.AddInterfaceToTheZone(interfaceName, fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	log.Debugf("issuing docker interface subnet\n")

	dockerInterfaceConfig := dockerInterface.IPAM.Config
	log.Debugf("docker interace subnet: %s\n", dockerInterfaceConfig[0].Subnet)
	o, err = fm.FirewalldController.AddRichRule(fmt.Sprintf("rule family=ipv4 source address=%s accept", dockerInterfaceConfig[0].Subnet), fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	o, err = fm.FirewalldController.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	o, err = fm.FirewalldController.TurnOnMasquarade(fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	// Adding docker to the zone and enabling routing
	log.Infof("Adding docker0 interface to the zone and enabling routing\n")
	o, err = fm.FirewalldController.AddInterfaceToTheZone("docker0", fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	o, err = fm.FirewalldController.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	log.Infof("Reloading firewall\n")
	o, err = fm.FirewalldController.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	log.Infof("Restarting docker service\n")
	err = fm.DockerManager.RestartDockerService()
	if err != nil {
		return fmt.Errorf("cannot restart docker service: %w", err)
	}
	return nil
}

// closing all opened ports in default km2 firewalld's zone except 22 and 53 (ssh and dns port)
func (fm *FirewallManager) ClostAllOpenedPorts(ctx context.Context) error {
	_, ports, err := fm.FirewalldController.GetOpenedPorts(fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("cannot get opened ports: %w", err)
	}

	var portsToClose []types.Port

	for _, p := range ports {
		port, err := fm.FirewallHandler.ConvertFirewallDPortToKM2Port(p)
		if err != nil {
			return fmt.Errorf("convert port: %w", err)
		}
		portsToClose = append(portsToClose, port)
	}

	err = fm.FirewallHandler.CheckPorts(portsToClose)
	if err != nil {
		return fmt.Errorf("cannot while checking ports are valid: %w", err)
	}

	err = fm.FirewallHandler.ClosePorts(portsToClose, fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("cannot close ports: %w", err)
	}

	o, err := fm.FirewalldController.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	return nil
}

func (fm *FirewallManager) OpenConfigPorts(ctx context.Context) error {
	portsToOpen := fm.FirewallConfig.PortsToOpen
	err := fm.FirewallHandler.CheckPorts(portsToOpen)
	if err != nil {
		return fmt.Errorf("cannot while checking ports are valid: %w", err)
	}
	err = fm.FirewallHandler.OpenPorts(portsToOpen, fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("cannot open ports: %w", err)
	}
	o, err := fm.FirewalldController.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	return nil
}

func (fm *FirewallManager) RestoreDefaultFirewalldSetingsForKM2(ctx context.Context) error {
	check, err := fm.FirewallHandler.CheckFirewallZone(fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	if check {
		log.Infof("docker zone exist: %s, deleting\n", fm.FirewallConfig.ZoneName)
		o, err := fm.FirewalldController.DeleteFirewallZone(fm.FirewallConfig.ZoneName)
		if err != nil {
			return fmt.Errorf("%s\n%w", o, err)
		}
		o, err = fm.FirewalldController.ReloadFirewall()
		if err != nil {
			return fmt.Errorf("%s\n%w", o, err)
		}
	}

	err = fm.SetUpFirewall(ctx)
	if err != nil {
		return fmt.Errorf("errow while setuping firewalld: %w", err)
	}
	return nil
}
