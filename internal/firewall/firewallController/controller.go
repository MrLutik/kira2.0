package firewallController

import (
	"fmt"
	"strings"

	"github.com/mrlutik/kira2.0/internal/osutils"
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

func (f *FirewalldController) CloseAllPorts() (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--set-target=DROP", "--permanent")
	return string(out), err
}

func (f *FirewalldController) OpenPort(port string) (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--zone=validator", "--add-port="+port, "--permanent")
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

func (f *FirewalldController) RejectIp(ip string) (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--permanent", "--zone="+f.zoneName, fmt.Sprintf(`--add-rich-rule="rule family='ipv4' source address=%s reject"`, ip))
	return string(out), err
}

func (f *FirewalldController) AcceptIp(ip string) (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--permanent", "--zone="+f.zoneName, fmt.Sprintf(`--add-rich-rule="rule family='ipv4' source address=%s accept"`, ip))
	return string(out), err
}
