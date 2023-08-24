package firewallManager

import (
	"fmt"

	"github.com/mrlutik/kira2.0/internal/firewall/firewallController"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/osutils"
)

type firewallManager struct {
	ZoneName           string
	PortsToOpen        []Port
	firewallController *firewallController.FirewalldController
}

// port range 0-65535
// type udp or tcp
//
// for example 39090 tcp or 53 udp
type Port struct {
	Port string
	Type string
}

func NewFirewallmanager(zoneName string, portsToOpen []Port) *firewallManager {
	c := firewallController.NewFireWalldController(zoneName)
	return &firewallManager{ZoneName: zoneName, PortsToOpen: portsToOpen, firewallController: c}
}

var log = logging.Log

func (cfg *firewallManager) SetUpFirewall() error {
	log.Debugf("HEHEHEHHEHEHEHEHEHERRRRRRREEEEEEEEEEEE\n")

	log.Debugf("Configuring new firewalld controller\n")

	check, err := cfg.checkFirewallZone()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	if !check {
		log.Debugf("Creating new firewalldZone %s\n ", cfg.ZoneName)
		o, err := cfg.firewallController.CreateNewFirewalldZone()
		if err != nil {
			return fmt.Errorf("%s\n%w", o, err)
		}
	}

	log.Debugf("Switching into %s firewalldZone\n", cfg.ZoneName)
	o, err := cfg.firewallController.ChangeZone()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	log.Debugf("Closing all ports\n")
	o, err = cfg.firewallController.CloseAllPorts()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	log.Debugf("Checking ports %+v \n", cfg.PortsToOpen)
	err = cfg.checkPorts()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	log.Debugf("Checking checkFirewallZone\n")

	err = cfg.openPorts()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	o, err = cfg.firewallController.AddDockerToTheZone()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	//save
	o, err = cfg.firewallController.SaveChanges()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	return nil
}

func (cfg *firewallManager) openPorts() error {
	for _, port := range cfg.PortsToOpen {
		log.Debugf("Opening %s/%s port\n", port.Port, port.Type)
		o, err := cfg.firewallController.OpenPort(fmt.Sprintf("%s/%s", port.Port, port.Type))
		if err != nil {
			return fmt.Errorf("%s\n%w", o, err)
		}
	}
	return nil
}

func (cfg *firewallManager) checkPorts() error {
	var ok bool
	for _, port := range cfg.PortsToOpen {
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
func (cfg *firewallManager) checkFirewallZone() (bool, error) {
	out, zones, err := cfg.firewallController.GetAllFirewallZones()
	log.Debugf("%s\n%+v\n%w\n", string(out), zones, err)
	if err != nil {
		return false, fmt.Errorf("%s\n%w", out, err)
	}
	for _, zone := range zones {
		if zone == cfg.ZoneName {
			return true, nil
		}
	}
	return false, nil
}
