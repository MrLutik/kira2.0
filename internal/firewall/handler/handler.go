package handler

import (
	"context"
	"fmt"
	"net"
	"regexp"
	"strings"

	dockerTypes "github.com/docker/docker/api/types"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/firewall/firewallController"
	"github.com/mrlutik/kira2.0/internal/types"

	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/osutils"
)

type FirewallHandler struct {
	firewalldController *firewallController.FirewalldController
}

var log = logging.Log

func NewFirewallHandler(firewalldController *firewallController.FirewalldController) *FirewallHandler {
	return &FirewallHandler{firewalldController: firewalldController}
}

func (fh *FirewallHandler) OpenPorts(portsToOpen []types.Port, zoneName string) error {
	for _, port := range portsToOpen {
		log.Debugf("Opening %s/%s port\n", port.Port, port.Type)
		o, err := fh.firewalldController.OpenPort(port, zoneName)
		if err != nil {
			return fmt.Errorf("%s\n%w", o, err)
		}
	}
	return nil
}

func (fh *FirewallHandler) ClosePorts(portsToClose []types.Port, zoneName string) error {
	for _, port := range portsToClose {
		if port.Port != "53" && port.Port != "22" {
			log.Debugf("Closing %s/%s port\n", port.Port, port.Type)
			o, err := fh.firewalldController.ClosePort(port, zoneName)
			if err != nil {
				return fmt.Errorf("%s\n%w", o, err)
			}
		} else {
			log.Debugf("skiping %s port (sys port)", port.Port)
		}
	}
	return nil
}

func (fh *FirewallHandler) ConvertFirewallDPortToKM2Port(firewallDPort string) (types.Port, error) {
	re := regexp.MustCompile(`(?P<Port>\d+)/(?P<Type>tcp|udp)`)
	matches := re.FindStringSubmatch(firewallDPort)

	if matches == nil {
		return types.Port{}, fmt.Errorf("cannot convert '%s' port: %w", firewallDPort, ErrNoPortMatch)
	}

	// Extract matches based on named groups
	portIndex := re.SubexpIndex("Port")
	typeIndex := re.SubexpIndex("Type")

	port := matches[portIndex]
	portType := strings.TrimPrefix(matches[typeIndex], "/")
	if osutils.CheckIfPortIsValid(port) {
		return types.Port{}, fmt.Errorf("%w: '%s'", ErrInvalidPort, port)
	}
	return types.Port{
		Port: port,
		Type: portType,
	}, nil
}

func (fh *FirewallHandler) CheckPorts(portsToOpen []types.Port) error {
	for _, port := range portsToOpen {
		if osutils.CheckIfPortIsValid(port.Port) {
			return fmt.Errorf("%w: '%s'", ErrInvalidPort, port)
		}
		if port.Type != "tcp" && port.Type != "udp" {
			return fmt.Errorf("%w: '%s' is not valid", ErrInvalidPortType, port.Type)
		}
	}
	return nil
}

func (fh *FirewallHandler) CheckFirewallZone(zoneName string) (bool, error) {
	out, zones, err := fh.firewalldController.GetAllFirewallZones()
	log.Debugf("Output:%s\nZones: %+v\nError: %s\n", string(out), zones, err)
	if err != nil {
		return false, fmt.Errorf("%s\n%w", out, err)
	}
	for _, zone := range zones {
		log.Debugf("Current zone: %s, expected zone: %s", zone, zoneName)
		if zone == zoneName {
			return true, nil
		}
	}
	return false, nil
}

// GetDockerNetworkInterface gets docker's custom interface name
func (fh *FirewallHandler) GetDockerNetworkInterface(ctx context.Context, dockerNetworkName string, dockerManager *docker.DockerManager) (dockerInterface dockerTypes.NetworkResource, err error) {
	network, err := dockerManager.Cli.NetworkInspect(ctx, dockerNetworkName, dockerTypes.NetworkInspectOptions{})
	if err != nil {
		return dockerInterface, fmt.Errorf("cannot get docker network info: %w", err)
	}
	return network, nil
}

// BlackListIP makes ip blacklisted
// TODO: still thinking if I should do reloading in this func or latter separately, because reloading taking a bit time
func (fh *FirewallHandler) BlackListIP(ip, zoneName string) error {
	ipCheck := net.ParseIP(ip)
	if ipCheck == nil {
		return fmt.Errorf("%w: %s is not valid", ErrInvalidIPAddress, ip)
	}

	output, err := fh.firewalldController.RejectIp(ip, zoneName)
	if err != nil {
		return fmt.Errorf("rejecting IP error: %w", err)
	}
	log.Infof("Output: %s", output)

	out, err := fh.firewalldController.ReloadFirewall()
	log.Debugf("%s", out)
	if err != nil {
		return err
	}
	return nil
}

func (fh *FirewallHandler) RemoveFromBlackListIP(ip, zoneName string) error {
	ipCheck := net.ParseIP(ip)
	if ipCheck == nil {
		return fmt.Errorf("%w: %s is not valid", ErrInvalidIPAddress, ip)
	}

	output, err := fh.firewalldController.RemoveRejectRuleIp(ip, zoneName)
	if err != nil {
		return fmt.Errorf("removing rejecting rule error: %w", err)
	}
	log.Infof("Output: %s", output)

	out, err := fh.firewalldController.ReloadFirewall()
	log.Debugf("%s", out)
	if err != nil {
		return err
	}
	return nil
}

func (fh *FirewallHandler) WhiteListIp(ip, zoneName string) error {
	ipCheck := net.ParseIP(ip)
	if ipCheck == nil {
		return fmt.Errorf("%w: %s is not valid", ErrInvalidIPAddress, ip)
	}

	output, err := fh.firewalldController.AcceptIp(ip, zoneName)
	if err != nil {
		return fmt.Errorf("accepting IP error: %w", err)
	}
	log.Infof("Output: %s", output)

	out, err := fh.firewalldController.ReloadFirewall()
	log.Debugf("%s", out)
	if err != nil {
		return err
	}
	return nil
}

func (fh *FirewallHandler) RemoveFromWhitelistIP(ip, zoneName string) error {
	ipCheck := net.ParseIP(ip)
	if ipCheck == nil {
		return fmt.Errorf("%w: %s is not valid", ErrInvalidIPAddress, ip)
	}

	output, err := fh.firewalldController.RemoveAllowRuleIp(ip, zoneName)
	if err != nil {
		return fmt.Errorf("removing allowing rule error: %w", err)
	}
	log.Infof("Output: %s", output)

	out, err := fh.firewalldController.ReloadFirewall()
	log.Debugf("%s", out)
	if err != nil {
		return err
	}
	return nil
}
