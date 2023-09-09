package firewallController

import (
	"fmt"
	"strings"

	"github.com/mrlutik/kira2.0/internal/osutils"
	"github.com/mrlutik/kira2.0/internal/types"
)

type FirewalldController struct {
	zoneName string
}

func NewFireWalldController(zoneName string) *FirewalldController {
	return &FirewalldController{zoneName: zoneName}
}

func (f *FirewalldController) CreateNewFirewalldZone() (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--new-zone="+f.zoneName, "--permanent")
	return string(out), err
}

func (f *FirewalldController) ChangeDefaultZone() (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--set-default-zone="+f.zoneName)
	return string(out), err
}

func (f *FirewalldController) DropAllConnections() (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--set-target=DROP", "--permanent")
	return string(out), err
}

func (f *FirewalldController) AllowAllConnections() (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--set-target=DROP", "--permanent")
	return string(out), err
}

func (f *FirewalldController) OpenPort(port types.Port) (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--zone="+f.zoneName, "--add-port="+port.Port+"/"+port.Type, "--permanent")
	return string(out), err
}

func (f *FirewalldController) ClosePort(port types.Port) (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--zone="+f.zoneName, "--remove-port="+port.Port+"/"+port.Type, "--permanent")
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

func (f *FirewalldController) AddInterfaceToTheZone(interfaceName string) (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--zone="+f.zoneName, "--add-interface="+interfaceName, "--permanent")
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
func (f *FirewalldController) RejectIp(ip string) (string, error) {
	// out, err := osutils.RunCommand("sudo", "firewall-cmd", "--permanent", "--zone="+f.zoneName, "--add-rich-rule="+fmt.Sprintf(`"rule family='ipv4' source address='%s' reject"`, ip))
	out, err := osutils.RunCommandV2(fmt.Sprintf(`firewall-cmd --permanent --zone=%s  --add-rich-rule="rule family='ipv4' source address='%s' reject"`, f.zoneName, ip))
	return string(out), err
}

// removing rule for rejecting ip
func (f *FirewalldController) RemoveRejectRuleIp(ip string) (string, error) {
	// out, err := osutils.RunCommand("sudo", "firewall-cmd", "--permanent", "--zone="+f.zoneName, "--add-rich-rule="+fmt.Sprintf(`"rule family='ipv4' source address='%s' reject"`, ip))
	out, err := osutils.RunCommandV2(fmt.Sprintf(`firewall-cmd --permanent --zone=%s  --remove-rich-rule="rule family='ipv4' source address='%s' reject"`, f.zoneName, ip))
	return string(out), err
}

// adding rule for accepting ip
func (f *FirewalldController) AcceptIp(ip string) (string, error) {
	out, err := osutils.RunCommandV2(fmt.Sprintf(`firewall-cmd --permanent --zone=%s  --add-rich-rule="rule family='ipv4' source address='%s' accept"`, f.zoneName, ip))
	return string(out), err
}

// removing rule for accepting ip
func (f *FirewalldController) RemoveAllowRuleIp(ip string) (string, error) {
	// out, err := osutils.RunCommand("sudo", "firewall-cmd", "--permanent", "--zone="+f.zoneName, "--add-rich-rule="+fmt.Sprintf(`"rule family='ipv4' source address='%s' reject"`, ip))
	out, err := osutils.RunCommandV2(fmt.Sprintf(`firewall-cmd --permanent --zone=%s  --remove-rich-rule="rule family='ipv4' source address='%s' accept"`, f.zoneName, ip))
	return string(out), err
}

func (f *FirewalldController) AddRichRule(rule string, zoneName string) (string, error) {
	out, err := osutils.RunCommandV2(fmt.Sprintf(`firewall-cmd --permanent --zone=%s  --add-rich-rule="%s"`, zoneName, rule))
	return string(out), err
}

func (f *FirewalldController) TurnOnMasquarade() (string, error) {
	out, err := osutils.RunCommandV2(fmt.Sprintf("firewall-cmd --zone=%s --add-masquerade --permanent", f.zoneName))
	return string(out), err
}
