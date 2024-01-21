package controller

import (
	"fmt"
	"strings"

	"github.com/mrlutik/kira2.0/internal/types"
)

type (
	FirewallDController struct {
		commandExecutor CommandExecutor

		zoneName string
	}

	CommandExecutor interface {
		RunCommand(command string, args ...string) ([]byte, error)
		RunCommandV2(commandStr string) ([]byte, error)
	}
)

func NewFirewallDController(commandExecutor CommandExecutor, zoneName string) (*FirewallDController, error) {
	if err := validateZoneName(zoneName); err != nil {
		return nil, err
	}

	return &FirewallDController{
		commandExecutor: commandExecutor,
		zoneName:        zoneName,
	}, nil
}

func (f *FirewallDController) CreateNewZone() (string, error) {
	out, err := f.commandExecutor.RunCommand(
		"sudo",
		"firewall-cmd",
		fmt.Sprintf("--new-zone=%s", f.zoneName),
		"--permanent",
	)
	if err != nil {
		return "", fmt.Errorf("creating new firewall zone '%s' failed, error: %w", f.zoneName, err)
	}
	return string(out), nil
}

func (f *FirewallDController) ChangeDefaultZone() (string, error) {
	out, err := f.commandExecutor.RunCommand(
		"sudo",
		"firewall-cmd",
		fmt.Sprintf("--set-default-zone=%s", f.zoneName),
	)
	if err != nil {
		return "", fmt.Errorf("changing default zone to '%s' failed, error: %w", f.zoneName, err)
	}
	return string(out), nil
}

func (f *FirewallDController) DropAllConnections() (string, error) {
	out, err := f.commandExecutor.RunCommand(
		"sudo",
		"firewall-cmd",
		"--set-target=DROP",
		"--permanent",
	)
	if err != nil {
		return "", fmt.Errorf("dropping all connections failed, error: %w", err)
	}
	return string(out), nil
}

func (f *FirewallDController) AllowAllConnections() (string, error) {
	out, err := f.commandExecutor.RunCommand(
		"sudo",
		"firewall-cmd",
		"--set-target=ALLOW",
		"--permanent",
	)
	if err != nil {
		return "", fmt.Errorf("allowing all connections failed, error: %w", err)
	}
	return string(out), nil
}

func (f *FirewallDController) OpenPort(port types.Port) (string, error) {
	if err := validatePort(port); err != nil {
		return "", err
	}

	out, err := f.commandExecutor.RunCommand(
		"sudo",
		"firewall-cmd",
		fmt.Sprintf("--zone=%s", f.zoneName),
		fmt.Sprintf("--add-port=%s/%s", port.Port, port.Type),
		"--permanent",
	)
	if err != nil {
		return "", fmt.Errorf("opening port '%s/%s' in zone '%s' failed, error: %w", port.Port, port.Type, f.zoneName, err)
	}
	return string(out), nil
}

func (f *FirewallDController) ClosePort(port types.Port) (string, error) {
	if err := validatePort(port); err != nil {
		return "", err
	}

	out, err := f.commandExecutor.RunCommand(
		"sudo",
		"firewall-cmd",
		fmt.Sprintf("--zone=%s", f.zoneName),
		fmt.Sprintf("--remove-port=%s/%s", port.Port, port.Type),
		"--permanent",
	)
	if err != nil {
		return "", fmt.Errorf("closing port '%s/%s' in zone '%s' failed, error: %w", port.Port, port.Type, f.zoneName, err)
	}
	return string(out), nil
}

func (f *FirewallDController) ReloadFirewall() (string, error) {
	out, err := f.commandExecutor.RunCommand("sudo", "firewall-cmd", "--reload")
	if err != nil {
		return "", fmt.Errorf("reloading firewall failed, error: %w", err)
	}
	return string(out), nil
}

func (f *FirewallDController) GetAllFirewallZones() (string, []string, error) {
	out, err := f.commandExecutor.RunCommand("sudo", "firewall-cmd", "--get-zones")
	if err != nil {
		return "", nil, fmt.Errorf("retrieving all firewall zones failed: %w", err)
	}
	zones := strings.Fields(string(out))
	return string(out), zones, nil
}

func (f *FirewallDController) AddInterfaceToTheZone(interfaceName string) (string, error) {
	if err := validateInterfaceName(interfaceName); err != nil {
		return "", err
	}

	out, err := f.commandExecutor.RunCommand(
		"sudo",
		"firewall-cmd",
		fmt.Sprintf("--zone=%s", f.zoneName),
		fmt.Sprintf("--add-interface=%s", interfaceName),
		"--permanent",
	)
	if err != nil {
		return "", fmt.Errorf("adding interface '%s' to zone '%s' failed, error: %w", interfaceName, f.zoneName, err)
	}
	return string(out), nil
}

// TODO Dmytro: How DOCKER is connected with controller?
func (f *FirewallDController) EnableDockerRouting(interfaceName string) (string, error) {
	out, err := f.commandExecutor.RunCommand(
		"sudo",
		"firewall-cmd",
		"--direct",
		"--permanent",
		"--add-rule",
		"ipv4", "filter", "FORWARD", "0",
		fmt.Sprintf("-i %s", interfaceName),
		fmt.Sprintf("-o %s", interfaceName),
		"-j", "ACCEPT",
	)
	if err != nil {
		return "", fmt.Errorf("enabling Docker routing on interface '%s' failed, error: %w", interfaceName, err)
	}
	return string(out), nil
}

func (f *FirewallDController) DeleteFirewallZone() (string, error) {
	out, err := f.commandExecutor.RunCommand(
		"sudo",
		"firewall-cmd",
		fmt.Sprintf("--delete-zone=%s", f.zoneName),
		"--permanent",
	)
	if err != nil {
		return "", fmt.Errorf("deleting firewall zone '%s' failed, error: %w", f.zoneName, err)
	}
	return string(out), nil
}

// RejectIp adds a rule to reject an IP address in the specified zone.
func (f *FirewallDController) RejectIp(ip string) (string, error) {
	if err := validateIP(ip); err != nil {
		return "", err
	}

	cmd := fmt.Sprintf(`sudo firewall-cmd --permanent --zone=%s --add-rich-rule="rule family='ipv4' source address='%s' reject"`, f.zoneName, ip)
	out, err := f.commandExecutor.RunCommandV2(cmd)
	if err != nil {
		return "", fmt.Errorf("failed to reject IP '%s' in zone '%s', error: %w", ip, f.zoneName, err)
	}
	return string(out), nil
}

// RemoveRejectRuleIp removes a rule that rejects an IP address from the specified zone.
func (f *FirewallDController) RemoveRejectRuleIp(ip string) (string, error) {
	if err := validateIP(ip); err != nil {
		return "", err
	}

	cmd := fmt.Sprintf(`sudo firewall-cmd --permanent --zone=%s --remove-rich-rule="rule family='ipv4' source address='%s' reject"`, f.zoneName, ip)
	out, err := f.commandExecutor.RunCommandV2(cmd)
	if err != nil {
		return "", fmt.Errorf("failed to remove reject rule for IP '%s' from zone '%s', error: %w", ip, f.zoneName, err)
	}
	return string(out), nil
}

// AcceptIp adds a rule to accept an IP address in the specified zone.
func (f *FirewallDController) AcceptIp(ip string) (string, error) {
	if err := validateIP(ip); err != nil {
		return "", err
	}

	cmd := fmt.Sprintf(`sudo firewall-cmd --permanent --zone=%s --add-rich-rule="rule family='ipv4' source address='%s' accept"`, f.zoneName, ip)
	out, err := f.commandExecutor.RunCommandV2(cmd)
	if err != nil {
		return "", fmt.Errorf("failed to accept IP '%s' in zone '%s', error: %w", ip, f.zoneName, err)
	}
	return string(out), nil
}

// RemoveAllowRuleIp removes a rule that accepts an IP address from the specified zone.
func (f *FirewallDController) RemoveAllowRuleIp(ip string) (string, error) {
	if err := validateIP(ip); err != nil {
		return "", err
	}

	cmd := fmt.Sprintf(`sudo firewall-cmd --permanent --zone=%s --remove-rich-rule="rule family='ipv4' source address='%s' accept"`, f.zoneName, ip)
	out, err := f.commandExecutor.RunCommandV2(cmd)
	if err != nil {
		return "", fmt.Errorf("failed to remove allow rule for IP '%s' from zone '%s', error: %w", ip, f.zoneName, err)
	}
	return string(out), nil
}

// AddRichRule adds a rich rule to the specified zone.
func (f *FirewallDController) AddRichRule(rule string) (string, error) {
	cmd := fmt.Sprintf(`sudo firewall-cmd --permanent --zone=%s --add-rich-rule="%s"`, f.zoneName, rule)
	out, err := f.commandExecutor.RunCommandV2(cmd)
	if err != nil {
		return "", fmt.Errorf("failed to add rich rule '%s' to zone '%s', error: %w", rule, f.zoneName, err)
	}
	return string(out), nil
}

// TurnOnMasquerade enables masquerading in the specified zone.
func (f *FirewallDController) TurnOnMasquerade() (string, error) {
	cmd := fmt.Sprintf("sudo firewall-cmd --zone=%s --add-masquerade --permanent", f.zoneName)
	out, err := f.commandExecutor.RunCommandV2(cmd)
	if err != nil {
		return "", fmt.Errorf("failed to turn on masquerading in zone '%s': %w", f.zoneName, err)
	}
	return string(out), nil
}

// GetOpenedPorts retrieves a list of open ports for the specified zone.
func (f *FirewallDController) GetOpenedPorts() (string, []string, error) {
	cmd := fmt.Sprintf("sudo firewall-cmd --zone=%s --list-ports", f.zoneName)
	out, err := f.commandExecutor.RunCommandV2(cmd)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get opened ports for zone '%s', error: %w", f.zoneName, err)
	}
	ports := strings.Fields(string(out))
	return string(out), ports, nil
}
