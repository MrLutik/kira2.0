package firewallController

import (
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

func (f *FirewalldController) ChangeZone() (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--change-interface=eth0", "--zone="+f.zoneName, "--permanent")
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

func (f *FirewalldController) SaveChanges() (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--reload")
	return string(out), err
}

func (f *FirewalldController) GetAllFirewallZones() (string, []string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--get-zones")
	zones := strings.Fields(string(out))
	return string(out), zones, err
}

func (f *FirewalldController) AddDockerToTheZone() (string, error) {
	out, err := osutils.RunCommand("sudo", "firewall-cmd", "--zone="+f.zoneName, "--permanent", "--add-interface=docker0")
	return string(out), err
}
