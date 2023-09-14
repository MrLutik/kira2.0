package firewallHandler

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"

	dockerTypes "github.com/docker/docker/api/types"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/firewall/firewallController"
	"github.com/mrlutik/kira2.0/internal/types"

	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/osutils"
)

var log = logging.Log

//	func NewFirewallHandler(dockerManager *docker.DockerManager, firewalldController *firewallController.FirewalldController) *FirewallHandler {
//		return &FirewallHandler{firewalldController: firewalldController, dockerManager: dockerManager}
//	}
func NewFirewallHandler(firewalldController *firewallController.FirewalldController) *FirewallHandler {
	return &FirewallHandler{firewalldController: firewalldController}
}

type FirewallHandler struct {
	firewalldController *firewallController.FirewalldController
	// dockerManager       *docker.DockerManager
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

func (fh *FirewallHandler) ConvertFirewalldPortToKM2Port(firewalldPort string) (types.Port, error) {
	re := regexp.MustCompile(`(?P<Port>\d+)/(?P<Type>tcp|udp)`)
	matches := re.FindStringSubmatch(firewalldPort)

	if matches == nil {
		return types.Port{}, fmt.Errorf("cannot convert %s port, no matches", firewalldPort)
	}

	// Extract matches based on named groups
	portIndex := re.SubexpIndex("Port")
	typeIndex := re.SubexpIndex("Type")

	port := matches[portIndex]
	portType := strings.TrimPrefix(matches[typeIndex], "/")
	check, err := osutils.CheckIfPortIsValid(port)
	if err != nil {
		return types.Port{}, fmt.Errorf("cannot check if <%s> is a valid port", port)
	}
	if !check {
		return types.Port{}, fmt.Errorf("<%s> is not a valid port", port)
	}
	return types.Port{
		Port: port,
		Type: portType,
	}, nil
}

func (fh *FirewallHandler) CheckPorts(portsToOpen []types.Port) (err error) {
	var ok bool
	for _, port := range portsToOpen {
		ok, err = osutils.CheckIfPortIsValid(port.Port)
		if err != nil {
			return fmt.Errorf("error when parsinh <%s>: %w", port, err)
		}
		if !ok {
			return fmt.Errorf("port <%s> is not valid: %w", port, err)
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
		// log.Debugf("%s %s", zone, zoneName)
		if zone == zoneName {
			return true, nil
		}
	}
	return false, nil
}

// geting docker's custom interface name
func (fh *FirewallHandler) GetDockerNetworkInterface(ctx context.Context, dockerNetworkName string, dockerManager *docker.DockerManager) (dockerInterface dockerTypes.NetworkResource, err error) {
	network, err := dockerManager.Cli.NetworkInspect(ctx, dockerNetworkName, dockerTypes.NetworkInspectOptions{})
	if err != nil {
		return dockerInterface, fmt.Errorf("cannot get docker network info: %w", err)
	}
	return network, nil
}

// blacklisting ip, still thinking if i shoud do realoading in this func or latter seperate, because reloading taking abit time
func (fh *FirewallHandler) BlackListIP(ip string, zoneName string) error {
	ipCheck := net.ParseIP(ip)
	if ipCheck != nil {
		fh.firewalldController.RejectIp(ip, zoneName)
	} else {
		return fmt.Errorf("%s is not a valid ip", ip)
	}
	out, err := fh.firewalldController.ReloadFirewall()
	log.Debugf("%s", out)
	if err != nil {
		return err
	}
	return nil
}

func (fh *FirewallHandler) RemoveFromBlackListIP(ip, zoneName string) error {
	ipCheck := net.ParseIP(ip)
	if ipCheck != nil {
		fh.firewalldController.RemoveRejectRuleIp(ip, zoneName)
	} else {
		return fmt.Errorf("%s is not a valid ip", ip)
	}
	out, err := fh.firewalldController.ReloadFirewall()
	log.Debugf("%s", out)
	if err != nil {
		return err
	}
	return nil
}

func (fh *FirewallHandler) WhiteListIp(ip, zoneName string) error {
	ipCheck := net.ParseIP(ip)
	if ipCheck != nil {
		fh.firewalldController.AcceptIp(ip, zoneName)
	} else {
		return fmt.Errorf("%s is not a valid ip", ip)
	}
	out, err := fh.firewalldController.ReloadFirewall()
	log.Debugf("%s", out)
	if err != nil {
		return err
	}
	return nil
}

func (fh *FirewallHandler) RemoveFromWhitelistIP(ip, zoneName string) error {
	ipCheck := net.ParseIP(ip)
	if ipCheck != nil {
		fh.firewalldController.RemoveAllowRuleIp(ip, zoneName)
	} else {
		return fmt.Errorf("%s is not a valid ip", ip)
	}
	out, err := fh.firewalldController.ReloadFirewall()
	log.Debugf("%s", out)
	if err != nil {
		return err
	}
	return nil
}

func (fh *FirewallHandler) RestartDockerService() error {
	out, err := osutils.RunCommandV2("sudo systemctl restart docker")
	if err != nil {
		return fmt.Errorf("failed to restart:\n %s\n%w", string(out), err)
	}
	return nil
}

func (fh *FirewallHandler) DisableIpTablesForDocker() error {
	filepath := "/etc/docker/daemon.json"
	type dockerServiceConfig struct {
		Iptables bool `json:"iptables"`
	}
	var config dockerServiceConfig
	file, err := os.Open(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			config.Iptables = false
		} else {
			return err
		}
	} else {
		defer file.Close()
		if err := json.NewDecoder(file).Decode(&config); err != nil {
			return err
		}
		config.Iptables = false
	}
	outFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer outFile.Close()
	encoder := json.NewEncoder(outFile)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(config); err != nil {
		return err
	}
	return nil
}
