package manager

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"

	dockerTypes "github.com/docker/docker/api/types"
	"github.com/mrlutik/kira2.0/internal/types"
)

type (
	NetworkInspector interface {
		NetworkInspect(ctx context.Context, networkID string) (dockerTypes.NetworkResource, error)
	}
)

var (
	ErrInvalidIPAddress = errors.New("invalid IP address")
	ErrInvalidPortType  = errors.New("port type is not valid")
	ErrInvalidPort      = errors.New("invalid port")
)

func (f *FirewallManager) OpenPorts(portsToOpen []types.Port) error {
	for _, port := range portsToOpen {
		f.log.Debugf("Opening '%s/%s' port", port.Port, port.Type)
		_, err := f.controllerKiraZone.OpenPort(port)
		if err != nil {
			return fmt.Errorf("opening port '%s/%s' %w", port.Port, port.Type, err)
		}
	}

	output, err := f.controllerKiraZone.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", output, err)
	}

	return nil
}

func (f *FirewallManager) ClosePorts(portsToClose []types.Port) error {
	for _, port := range portsToClose {
		if port.Port != "53" && port.Port != "22" {
			f.log.Debugf("Closing %s/%s port\n", port.Port, port.Type)
			o, err := f.controllerKiraZone.ClosePort(port)
			if err != nil {
				return fmt.Errorf("%s\n%w", o, err)
			}
		} else {
			f.log.Debugf("skipping %s port (system port)", port.Port)
		}
	}

	output, err := f.controllerKiraZone.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", output, err)
	}

	return nil
}

func (f *FirewallManager) convertFirewallDPortToKM2Port(firewallDPort string) (types.Port, error) {
	parts := strings.Split(firewallDPort, "/")
	if len(parts) != 2 {
		return types.Port{}, fmt.Errorf("invalid port format '%s': %w", firewallDPort, ErrInvalidPort)
	}

	port, portType := parts[0], parts[1]

	if portType != "tcp" && portType != "udp" {
		return types.Port{}, fmt.Errorf("invalid port type '%s': %w", portType, ErrInvalidPortType)
	}

	if !f.utils.CheckIfPortIsValid(port) {
		return types.Port{}, fmt.Errorf("%w: '%s'", ErrInvalidPort, port)
	}

	return types.Port{Port: port, Type: portType}, nil
}

func (f *FirewallManager) checkPorts(portsToOpen []types.Port) error {
	for _, port := range portsToOpen {
		if !f.utils.CheckIfPortIsValid(port.Port) {
			return fmt.Errorf("%w: '%s'", ErrInvalidPort, port)
		}
		if port.Type != "tcp" && port.Type != "udp" {
			return fmt.Errorf("%w: '%s' is not valid", ErrInvalidPortType, port.Type)
		}
	}
	return nil
}

func (f *FirewallManager) checkFirewallZone(zoneName string) (bool, error) {
	out, zones, err := f.controllerKiraZone.GetAllFirewallZones()
	if err != nil {
		return false, fmt.Errorf("%s\n%w", out, err)
	}

	f.log.Debugf("Output:%s\nZones: %+v", string(out), zones)

	for _, zone := range zones {
		f.log.Debugf("Current zone: %s, expected zone: %s", zone, zoneName)
		if zone == zoneName {
			return true, nil
		}
	}

	return false, nil
}

// getDockerNetworkInterface gets docker's custom interface name
func (f *FirewallManager) getDockerNetworkInterface(ctx context.Context, dockerNetworkName string) (dockerInterface *dockerTypes.NetworkResource, err error) {
	network, err := f.dockerMaintenance.NetworkInspect(ctx, dockerNetworkName)
	if err != nil {
		return dockerInterface, fmt.Errorf("cannot get docker network info: %w", err)
	}
	return network, nil
}

// BlackListIP makes ip blacklisted
// TODO: still thinking if I should do reloading in this func or latter separately, because reloading taking a bit time
func (f *FirewallManager) BlackListIP(ip string) error {
	ipCheck := net.ParseIP(ip)
	if ipCheck == nil {
		return fmt.Errorf("%w: %s is not valid", ErrInvalidIPAddress, ip)
	}

	output, err := f.controllerKiraZone.RejectIp(ip)
	if err != nil {
		return fmt.Errorf("rejecting IP error: %w", err)
	}
	f.log.Infof("Output: %s", output)

	out, err := f.controllerKiraZone.ReloadFirewall()
	if err != nil {
		return err
	}
	f.log.Debugf("%s", out)

	return nil
}

func (f *FirewallManager) RemoveFromBlackListIP(ip string) error {
	ipCheck := net.ParseIP(ip)
	if ipCheck == nil {
		return fmt.Errorf("%w: %s is not valid", ErrInvalidIPAddress, ip)
	}

	output, err := f.controllerKiraZone.RemoveRejectRuleIp(ip)
	if err != nil {
		return fmt.Errorf("removing rejecting rule error: %w", err)
	}
	f.log.Infof("Output: %s", output)

	output, err = f.controllerKiraZone.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", output, err)
	}

	f.log.Debugf("Output: %s", output)
	return nil
}

func (f *FirewallManager) WhiteListIp(ip string) error {
	ipCheck := net.ParseIP(ip)
	if ipCheck == nil {
		return fmt.Errorf("%w: %s is not valid", ErrInvalidIPAddress, ip)
	}

	output, err := f.controllerKiraZone.AcceptIp(ip)
	if err != nil {
		return fmt.Errorf("accepting IP error: %w", err)
	}
	f.log.Infof("Output: %s", output)

	output, err = f.controllerKiraZone.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", output, err)
	}

	f.log.Debugf("Output: %s", output)
	return nil
}

func (f *FirewallManager) RemoveFromWhitelistIP(ip string) error {
	ipCheck := net.ParseIP(ip)
	if ipCheck == nil {
		return fmt.Errorf("%w: %s is not valid", ErrInvalidIPAddress, ip)
	}

	output, err := f.controllerKiraZone.RemoveAllowRuleIp(ip)
	if err != nil {
		return fmt.Errorf("removing allowing rule error: %w", err)
	}
	f.log.Infof("Output: %s", output)

	output, err = f.controllerKiraZone.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", output, err)
	}

	f.log.Debugf("Output: %s", output)
	return nil
}
