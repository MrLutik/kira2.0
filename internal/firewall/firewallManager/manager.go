package firewallManager

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/firewall/firewallController"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/osutils"
)

type firewallManager struct {
	ZoneName           string
	PortsToOpen        []Port
	firewallController *firewallController.FirewalldController
	dockerManager      *docker.DockerManager
}

// port range 0-65535
// type udp or tcp
//
// for example 39090 tcp or 53 udp
type Port struct {
	Port string
	Type string
}

func NewFirewallmanager(dockerManager *docker.DockerManager, zoneName string, portsToOpen []Port) *firewallManager {
	c := firewallController.NewFireWalldController(zoneName)

	return &firewallManager{ZoneName: zoneName, PortsToOpen: portsToOpen, firewallController: c, dockerManager: dockerManager}
}

var log = logging.Log

func (fm *firewallManager) SetUpFirewall(ctx context.Context, kiraConfig *config.KiraConfig) error {
	log.Infof("***FIREWALL SETUP***\n")

	log.Infof("checking and deleting default docker zone\n")
	dockerZoneName := "docker"
	check, err := fm.checkFirewallZone(dockerZoneName)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	if check {
		log.Infof("docker zone exist: %s, deleting\n", dockerZoneName)
		o, err := fm.firewallController.DeleteFirewallZone(dockerZoneName)
		if err != nil {
			return fmt.Errorf("%s\n%w", o, err)
		}
		o, err = fm.firewallController.ReloadFirewall()
		if err != nil {
			return fmt.Errorf("%s\n%w", o, err)
		}
	} else {
		log.Infof("docker zone: %s not exist\n", dockerZoneName)
	}

	log.Infof("checking if %s zone exist\n", fm.ZoneName)
	check, err = fm.checkFirewallZone(fm.ZoneName)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	if !check {
		log.Infof("Creating new firewalldZone %s, check = %v\n ", fm.ZoneName, check)
		o, err := fm.firewallController.CreateNewFirewalldZone()
		if err != nil {
			return fmt.Errorf("%s\n%w", o, err)
		}
		o, err = fm.firewallController.ReloadFirewall()
		if err != nil {
			return fmt.Errorf("%s\n%w", o, err)
		}
	}

	log.Infof("Switching into %s firewalldZone\n", fm.ZoneName)
	o, err := fm.firewallController.ChangeDefaultZone()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	o, err = fm.firewallController.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	log.Infof("Closing all ports\n")
	o, err = fm.firewallController.CloseAllPorts()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	log.Infof("Checking ports %+v \n", fm.PortsToOpen)
	err = fm.checkPorts()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	log.Infof("Opening ports\n")
	err = fm.openPorts()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	//adding interface that has internet acces
	log.Infof("Adding interface that has internet acces\n")
	internetInterface := osutils.GetInternetInterface()
	o, err = fm.firewallController.AddInterfaceToTheZone(internetInterface)
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	interfaceName, err := fm.getDockerNetworkInterfaceName(ctx, kiraConfig.DockerNetworkName)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	log.Infof("Adding %s interface to the zone and enabling routing\n", interfaceName)
	o, err = fm.firewallController.AddInterfaceToTheZone(interfaceName)
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	//adding docket to the zone and enabling routing
	log.Infof("Adding docker0 interface to the zone and enabling routing\n")
	o, err = fm.firewallController.AddInterfaceToTheZone("docker0")
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	o, err = fm.firewallController.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	o, err = fm.firewallController.EnableDockerRouting("docker0")
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	//save
	o, err = fm.firewallController.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	return nil
}

func (fm *firewallManager) editDockerService() error {
	const dockerOverrideDir = "/etc/systemd/system/docker.service.d"
	const dockerOverrideFile = "override.conf"

	// Create directory if it doesn't exist
	if err := os.MkdirAll(dockerOverrideDir, 0755); err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}

	// Configuration content
	content := `[Service]
ExecStart=
ExecStart=/usr/bin/dockerd --iptables=false
`

	// Write content to file
	filePath := filepath.Join(dockerOverrideDir, dockerOverrideFile)
	if err := ioutil.WriteFile(filePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

func (fm *firewallManager) openPorts() error {
	for _, port := range fm.PortsToOpen {
		log.Debugf("Opening %s/%s port\n", port.Port, port.Type)
		o, err := fm.firewallController.OpenPort(fmt.Sprintf("%s/%s", port.Port, port.Type))
		if err != nil {
			return fmt.Errorf("%s\n%w", o, err)
		}
	}
	return nil
}

func (fm *firewallManager) checkPorts() error {
	var ok bool
	for _, port := range fm.PortsToOpen {
		ok = osutils.CheckIfPortIsValid(port.Port)
		if !ok {
			return fmt.Errorf("port <%s> is not valid", port)
		}
		if port.Type != "tcp" && port.Type != "udp" {
			return fmt.Errorf("port type <%s> is not valid", port.Type)
		}
	}
	return nil
}

func (fm *firewallManager) checkFirewallZone(zoneName string) (bool, error) {
	out, zones, err := fm.firewallController.GetAllFirewallZones()
	log.Debugf("%s\n%+v\n%s\n", string(out), zones, err)
	if err != nil {
		return false, fmt.Errorf("%s\n%w", out, err)
	}
	for _, zone := range zones {
		log.Debugf("%s %s", zone, zoneName)
		if zone == zoneName {
			return true, nil
		}
	}
	return false, nil
}

func (fm *firewallManager) getDockerNetworkInterfaceName(ctx context.Context, dockerNetworkName string) (interfaceName string, err error) {
	networks, err := fm.dockerManager.GetNetworksInfo(ctx)
	if err != nil {
		return interfaceName, fmt.Errorf("cannot get docker network info: %w", err)
	}

	for _, network := range networks {
		if network.Name == dockerNetworkName {
			interfaceName = "br-" + network.ID[0:11]
		}
	}
	return interfaceName, nil
}
