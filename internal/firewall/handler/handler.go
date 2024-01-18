package handler

import (
	"context"
	"fmt"
	"net"
	"strings"

	dockerTypes "github.com/docker/docker/api/types"
	"github.com/mrlutik/kira2.0/internal/firewall/controller"
	"github.com/mrlutik/kira2.0/internal/types"

	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/osutils"
)

type (
	FirewallHandler struct {
		controller       *controller.FirewallDController
		networkInspector NetworkInspector
		utils            *osutils.OSUtils

		log *logging.Logger
	}
	NetworkInspector interface {
		NetworkInspect(ctx context.Context, networkID string) (dockerTypes.NetworkResource, error)
	}
)

func NewFirewallHandler(firewallDController *controller.FirewallDController, utils *osutils.OSUtils, networkInspector NetworkInspector, logger *logging.Logger) *FirewallHandler {
	return &FirewallHandler{
		controller:       firewallDController,
		networkInspector: networkInspector,
		utils:            utils,
		log:              logger,
	}
}

func (f *FirewallHandler) OpenPorts(portsToOpen []types.Port) error {
	for _, port := range portsToOpen {
		f.log.Debugf("Opening '%s/%s' port", port.Port, port.Type)
		_, err := f.controller.OpenPort(port)
		if err != nil {
			return fmt.Errorf("opening port '%s/%s' %w", port.Port, port.Type, err)
		}
	}
	return nil
}

func (f *FirewallHandler) ClosePorts(portsToClose []types.Port) error {
	for _, port := range portsToClose {
		if port.Port != "53" && port.Port != "22" {
			f.log.Debugf("Closing %s/%s port\n", port.Port, port.Type)
			o, err := f.controller.ClosePort(port)
			if err != nil {
				return fmt.Errorf("%s\n%w", o, err)
			}
		} else {
			f.log.Debugf("skipping %s port (sys port)", port.Port)
		}
	}
	return nil
}

func (FirewallHandler) ConvertFirewallDPortToKM2Port(firewallDPort string) (types.Port, error) {
	parts := strings.Split(firewallDPort, "/")
	if len(parts) != 2 {
		return types.Port{}, fmt.Errorf("invalid port format '%s': %w", firewallDPort, ErrInvalidPort)
	}

	port, portType := parts[0], parts[1]

	if portType != "tcp" && portType != "udp" {
		return types.Port{}, fmt.Errorf("invalid port type '%s': %w", portType, ErrInvalidPortType)
	}

	return types.Port{Port: port, Type: portType}, nil
}

func (f *FirewallHandler) CheckPorts(portsToOpen []types.Port) error {
	for _, port := range portsToOpen {
		if f.utils.CheckIfPortIsValid(port.Port) {
			return fmt.Errorf("%w: '%s'", ErrInvalidPort, port)
		}
		if port.Type != "tcp" && port.Type != "udp" {
			return fmt.Errorf("%w: '%s' is not valid", ErrInvalidPortType, port.Type)
		}
	}
	return nil
}

func (f *FirewallHandler) CheckFirewallZone(zoneName string) (bool, error) {
	out, zones, err := f.controller.GetAllFirewallZones()
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

// GetDockerNetworkInterface gets docker's custom interface name
func (f *FirewallHandler) GetDockerNetworkInterface(ctx context.Context, dockerNetworkName string) (dockerInterface dockerTypes.NetworkResource, err error) {
	network, err := f.networkInspector.NetworkInspect(ctx, dockerNetworkName)
	if err != nil {
		return dockerInterface, fmt.Errorf("cannot get docker network info: %w", err)
	}
	return network, nil
}

// BlackListIP makes ip blacklisted
// TODO: still thinking if I should do reloading in this func or latter separately, because reloading taking a bit time
func (f *FirewallHandler) BlackListIP(ip string) error {
	ipCheck := net.ParseIP(ip)
	if ipCheck == nil {
		return fmt.Errorf("%w: %s is not valid", ErrInvalidIPAddress, ip)
	}

	output, err := f.controller.RejectIp(ip)
	if err != nil {
		return fmt.Errorf("rejecting IP error: %w", err)
	}
	f.log.Infof("Output: %s", output)

	out, err := f.controller.ReloadFirewall()
	if err != nil {
		return err
	}
	f.log.Debugf("%s", out)

	return nil
}

func (f *FirewallHandler) RemoveFromBlackListIP(ip string) error {
	ipCheck := net.ParseIP(ip)
	if ipCheck == nil {
		return fmt.Errorf("%w: %s is not valid", ErrInvalidIPAddress, ip)
	}

	output, err := f.controller.RemoveRejectRuleIp(ip)
	if err != nil {
		return fmt.Errorf("removing rejecting rule error: %w", err)
	}
	f.log.Infof("Output: %s", output)

	out, err := f.controller.ReloadFirewall()
	if err != nil {
		return err
	}

	f.log.Debugf("Output: %s", out)
	return nil
}

func (f *FirewallHandler) WhiteListIp(ip string) error {
	ipCheck := net.ParseIP(ip)
	if ipCheck == nil {
		return fmt.Errorf("%w: %s is not valid", ErrInvalidIPAddress, ip)
	}

	output, err := f.controller.AcceptIp(ip)
	if err != nil {
		return fmt.Errorf("accepting IP error: %w", err)
	}
	f.log.Infof("Output: %s", output)

	out, err := f.controller.ReloadFirewall()
	if err != nil {
		return err
	}

	f.log.Debugf("Output: %s", out)
	return nil
}

func (f *FirewallHandler) RemoveFromWhitelistIP(ip string) error {
	ipCheck := net.ParseIP(ip)
	if ipCheck == nil {
		return fmt.Errorf("%w: %s is not valid", ErrInvalidIPAddress, ip)
	}

	output, err := f.controller.RemoveAllowRuleIp(ip)
	if err != nil {
		return fmt.Errorf("removing allowing rule error: %w", err)
	}
	f.log.Infof("Output: %s", output)

	out, err := f.controller.ReloadFirewall()
	if err != nil {
		return err
	}

	f.log.Debugf("Output: %s", out)
	return nil
}
