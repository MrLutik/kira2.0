package firewallController

import (
	"fmt"
	"strings"

	"github.com/mrlutik/kira2.0/internal/osutils"
	"github.com/mrlutik/kira2.0/internal/types"
)

type FirewalldController struct {
	ZoneName string
}

func NewFireWalldController(zoneName string) *FirewalldController {
	return &FirewalldController{ZoneName: zoneName}
}

func (f *FirewalldController) CreateNewFirewalldZone(zoneName string) (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--new-zone="+zoneName, "--permanent")
	return string(out), err
}

func (f *FirewalldController) ChangeDefaultZone(zoneName string) (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--set-default-zone="+zoneName)
	return string(out), err
}

func (f *FirewalldController) DropAllConnections() (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--set-target=DROP", "--permanent")
	return string(out), err
}

func (f *FirewalldController) AllowAllConnections() (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--set-target=ALLOW", "--permanent")
	return string(out), err
}

func (f *FirewalldController) OpenPort(port types.Port, zoneName string) (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--zone="+zoneName, "--add-port="+port.Port+"/"+port.Type, "--permanent")
	return string(out), err
}

func (f *FirewalldController) ClosePort(port types.Port, zoneName string) (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--zone="+zoneName, "--remove-port="+port.Port+"/"+port.Type, "--permanent")
	return string(out), err
}

func (f *FirewalldController) ReloadFirewall() (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--reload")
	return string(out), err
}

func (f *FirewalldController) GetAllFirewallZones() (string, []string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--get-zones")
	zones := strings.Fields(string(out))
	return string(out), zones, err
}

func (f *FirewalldController) AddInterfaceToTheZone(interfaceName, zoneName string) (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--zone="+zoneName, "--add-interface="+interfaceName, "--permanent")
	return string(out), err
}

func (f *FirewalldController) EnableDockerRouting(interfaceName string) (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--direct", "--permanent", "--add-rule", "ipv4", "filter", "FORWARD", "0", "-i", interfaceName, "-o", interfaceName, "-j", "ACCEPT")
	return string(out), err
}

func (f *FirewalldController) DeleteFirewallZone(zonename string) (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--delete-zone="+zonename, "--permanent")
	return string(out), err
}

// adding rule for rejecting ip
func (f *FirewalldController) RejectIp(ip, zoneName string) (string, error) {
	// out, err := osutils.RunCommand("sudo", "firewall-cmd", "--permanent", "--zone="+zoneName, "--add-rich-rule="+fmt.Sprintf(`"rule family='ipv4' source address='%s' reject"`, ip))
	out, err := osutils.RunCommandV2(fmt.Sprintf(`firewall-cmd --permanent --zone=%s  --add-rich-rule="rule family='ipv4' source address='%s' reject"`, zoneName, ip))
	return string(out), err
}

// removing rule for rejecting ip
func (f *FirewalldController) RemoveRejectRuleIp(ip, zoneName string) (string, error) {
	// out, err := osutils.RunCommand("sudo", "firewall-cmd", "--permanent", "--zone="+zoneName, "--add-rich-rule="+fmt.Sprintf(`"rule family='ipv4' source address='%s' reject"`, ip))
	out, err := osutils.RunCommandV2(fmt.Sprintf(`firewall-cmd --permanent --zone=%s  --remove-rich-rule="rule family='ipv4' source address='%s' reject"`, zoneName, ip))
	return string(out), err
}

// adding rule for accepting ip
func (f *FirewalldController) AcceptIp(ip, zoneName string) (string, error) {
	out, err := osutils.RunCommandV2(fmt.Sprintf(`firewall-cmd --permanent --zone=%s  --add-rich-rule="rule family='ipv4' source address='%s' accept"`, zoneName, ip))
	return string(out), err
}

// removing rule for accepting ip
func (f *FirewalldController) RemoveAllowRuleIp(ip, zoneName string) (string, error) {
	// out, err := osutils.RunCommand("sudo", "firewall-cmd", "--permanent", "--zone="+zoneName, "--add-rich-rule="+fmt.Sprintf(`"rule family='ipv4' source address='%s' reject"`, ip))
	out, err := osutils.RunCommandV2(fmt.Sprintf(`firewall-cmd --permanent --zone=%s  --remove-rich-rule="rule family='ipv4' source address='%s' accept"`, zoneName, ip))
	return string(out), err
}

func (f *FirewalldController) AddRichRule(rule string, zoneName string) (string, error) {
	out, err := osutils.RunCommandV2(fmt.Sprintf(`firewall-cmd --permanent --zone=%s  --add-rich-rule="%s"`, zoneName, rule))
	return string(out), err
}

func (f *FirewalldController) TurnOnMasquarade(zoneName string) (string, error) {
	out, err := osutils.RunCommandV2(fmt.Sprintf("firewall-cmd --zone=%s --add-masquerade --permanent", zoneName))
	return string(out), err
}

func (f *FirewalldController) GetOpenedPorts(zoneName string) (string, []string, error) {
	out, err := osutils.RunCommandV2(fmt.Sprintf("firewall-cmd --zone=%s --list-ports", zoneName))
	ports := strings.Fields(string(out))
	return string(out), ports, err
}
