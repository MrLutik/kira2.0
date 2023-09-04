package firewallHandler

import (
	"context"
	"fmt"
	"net"

	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/firewall/firewallController"
	"github.com/mrlutik/kira2.0/internal/types"

	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/osutils"
)

var log = logging.Log

func NewFirewallHandler(dockerManager *docker.DockerManager, firewalldController *firewallController.FirewalldController) *FirewallHandler {
	return &FirewallHandler{firewalldController: firewalldController, dockerManager: dockerManager}
}

type FirewallHandler struct {
	firewalldController *firewallController.FirewalldController
	dockerManager       *docker.DockerManager
}

func (fh *FirewallHandler) OpenPorts(portsToOpen []types.Port) error {
	for _, port := range portsToOpen {
		log.Debugf("Opening %s/%s port\n", port.Port, port.Type)
		o, err := fh.firewalldController.OpenPort(fmt.Sprintf("%s/%s", port.Port, port.Type))
		if err != nil {
			return fmt.Errorf("%s\n%w", o, err)
		}
	}
	return nil
}

func (fh *FirewallHandler) CheckPorts(portsToOpen []types.Port) error {
	var ok bool
	for _, port := range portsToOpen {
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

func (fh *FirewallHandler) CheckFirewallZone(zoneName string) (bool, error) {
	out, zones, err := fh.firewalldController.GetAllFirewallZones()
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

// geting docker's custom interface name
func (fh *FirewallHandler) GetDockerNetworkInterfaceName(ctx context.Context, dockerNetworkName string) (interfaceName string, err error) {
	networks, err := fh.dockerManager.GetNetworksInfo(ctx)
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

// blacklisting ip, still thinking if i shoud do realoading in this func or latter seperate, because reloading taking abit time
func (fh *FirewallHandler) BlackListIP(ip string) error {
	ipCheck := net.ParseIP(ip)
	if ipCheck != nil {
		fh.firewalldController.RejectIp(ip)
	} else {
		return fmt.Errorf("%s is not a valid ip", ip)
	}
	_, err := fh.firewalldController.ReloadFirewall()
	if err != nil {
		return err
	}
	return nil
}

func (fh *FirewallHandler) WhiteListIp(ip string) error {
	ipCheck := net.ParseIP(ip)
	if ipCheck != nil {
		fh.firewalldController.AcceptIp(ip)
	} else {
		return fmt.Errorf("%s is not a valid ip", ip)
	}
	_, err := fh.firewalldController.ReloadFirewall()
	if err != nil {
		return err
	}
	return nil
}
