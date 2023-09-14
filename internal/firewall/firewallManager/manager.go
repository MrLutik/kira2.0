package firewallManager

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/firewall/firewallController"
	"github.com/mrlutik/kira2.0/internal/firewall/firewallManager/firewallHandler"
	"github.com/mrlutik/kira2.0/internal/types"

	// D:\Coding\Go\KIRA\kira2.0\kira2.0\internal\firewall\firewallManager\firewallHandlers\firewallHandlers.go
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/osutils"
)

type FirewallManager struct {
	FirewalldController *firewallController.FirewalldController
	DockerManager       *docker.DockerManager
	FirewallHandler     *firewallHandler.FirewallHandler
	FirewallConfig      *FirewallConfig
	KiraConfig          *config.KiraConfig
}

type FirewallConfig struct {
	ZoneName    string
	PortsToOpen []types.Port
}

// port range 0-65535
// type udp or tcp
//
// for example 39090 tcp or 53 udp

func GenerateKiraConfigForFirewallManager() *config.KiraConfig {
	return &config.KiraConfig{
		SekaidHome:          "/data/.sekai",
		InterxHome:          "/data/.interx",
		KeyringBackend:      "test",
		DockerImageName:     "ghcr.io/kiracore/docker/kira-base",
		DockerImageVersion:  "v0.13.11",
		DockerNetworkName:   "kira_network",
		SekaiVersion:        "latest", // or v0.3.16
		InterxVersion:       "latest", // or v0.4.33
		SekaidContainerName: "sekaid",
		InterxContainerName: "interx",
		VolumeName:          "kira_volume:/data",
		MnemonicDir:         "~/mnemonics",
		RpcPort:             "26657",
		P2PPort:             "26656",
		GrpcPort:            "9090",
		InterxPort:          "11000",
		PrometheusPort:      "26660",
		Moniker:             "VALIDATOR",
		SekaiDebFileName:    "sekai-linux-amd64.deb",
		InterxDebFileName:   "interx-linux-amd64.deb",
		TimeBetweenBlocks:   time.Second * 10,
	}
}
func NewFirewallConfig(kiraCfg *config.KiraConfig) *FirewallConfig {
	return &FirewallConfig{
		ZoneName: "validator",
		PortsToOpen: []types.Port{
			{Port: kiraCfg.InterxPort, Type: "tcp"},
			{Port: kiraCfg.GrpcPort, Type: "tcp"},
			{Port: kiraCfg.P2PPort, Type: "tcp"},
			{Port: kiraCfg.PrometheusPort, Type: "tcp"},
			{Port: kiraCfg.RpcPort, Type: "tcp"},
		},
	}
}
func NewFirewallManager(dockerManager *docker.DockerManager, kiraCfg *config.KiraConfig) *FirewallManager {
	cfg := NewFirewallConfig(kiraCfg)
	c := firewallController.NewFireWalldController(cfg.ZoneName)
	h := firewallHandler.NewFirewallHandler(c)

	return &FirewallManager{FirewalldController: c, DockerManager: dockerManager, FirewallHandler: h, FirewallConfig: cfg, KiraConfig: kiraCfg}
}

var log = logging.Log

func (fm *FirewallManager) CheckFirewallSetUp(ctx context.Context) (bool, error) {
	_, err := exec.LookPath("firewall-cmd")
	if err != nil {
		return false, fmt.Errorf("firewalld is not installed on the system")
	}

	//check if validator zone exist
	check, err := fm.FirewallHandler.CheckFirewallZone(fm.FirewallConfig.ZoneName)
	if err != nil {
		return false, fmt.Errorf("error while checking validator zone %w", err)
	}
	if !check {
		return false, nil
	}
	return true, nil
}

func (fm *FirewallManager) SetUpFirewall(ctx context.Context) error {
	log.Infof("***FIREWALL SETUP***\n")

	log.Infof("Disabling docker iptables\n")
	fm.FirewallHandler.DisableIpTablesForDocker()
	log.Infof("Restarting docker service\n")
	err := fm.FirewallHandler.RestartDockerService()
	if err != nil {
		return fmt.Errorf("cannot restart docker service: %s", err)
	}
	log.Infof("Checking if docker is running\n")
	err = fm.DockerManager.VerifyDockerInstallation(ctx)
	if err != nil {
		return err
	}

	log.Infof("checking and deleting default docker zone\n")
	dockerZoneName := "docker"
	check, err := fm.FirewallHandler.CheckFirewallZone(dockerZoneName)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	if check {
		log.Infof("docker zone exist: %s, deleting\n", dockerZoneName)
		o, err := fm.FirewalldController.DeleteFirewallZone(dockerZoneName)
		if err != nil {
			return fmt.Errorf("%s\n%w", o, err)
		}
		o, err = fm.FirewalldController.ReloadFirewall()
		if err != nil {
			return fmt.Errorf("%s\n%w", o, err)
		}
	} else {
		log.Infof("docker zone: %s not exist\n", dockerZoneName)
	}

	log.Infof("checking if %s zone exist\n", fm.FirewallConfig.ZoneName)
	check, err = fm.FirewallHandler.CheckFirewallZone(fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	if !check {
		log.Infof("Creating new firewalldZone %s, check = %v\n ", fm.FirewallConfig.ZoneName, check)
		o, err := fm.FirewalldController.CreateNewFirewalldZone(fm.FirewallConfig.ZoneName)
		if err != nil {
			return fmt.Errorf("%s\n%w", o, err)
		}
		o, err = fm.FirewalldController.ReloadFirewall()
		if err != nil {
			return fmt.Errorf("%s\n%w", o, err)
		}
	}

	log.Infof("Switching into %s firewalldZone\n", fm.FirewallConfig.ZoneName)
	o, err := fm.FirewalldController.ChangeDefaultZone(fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	o, err = fm.FirewalldController.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	// log.Infof("Closing all ports\n")
	// o, err = fm.FirewalldController.DropAllPorts()
	// if err != nil {
	// 	return fmt.Errorf("%s\n%w", o, err)
	// }

	log.Infof("Checking ports %+v \n", fm.FirewallConfig.PortsToOpen)
	err = fm.FirewallHandler.CheckPorts(fm.FirewallConfig.PortsToOpen)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	log.Infof("Opening system ports\n")
	sysports := []types.Port{
		{Port: "22", Type: "tcp"},
		{Port: "53", Type: "udp"},
	}
	err = fm.FirewallHandler.OpenPorts(sysports, fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	log.Infof("Opening kira ports\n")
	err = fm.FirewallHandler.OpenPorts(fm.FirewallConfig.PortsToOpen, fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	//adding interface that has internet acces
	log.Infof("Adding interface that has internet acces\n")
	internetInterface := osutils.GetInternetInterface()
	o, err = fm.FirewalldController.AddInterfaceToTheZone(internetInterface, fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	dockerInterface, err := fm.FirewallHandler.GetDockerNetworkInterface(ctx, fm.KiraConfig.DockerNetworkName, fm.DockerManager)
	interfaceName := "br-" + dockerInterface.ID[0:11]
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	fm.DockerManager.GetNetworksInfo(ctx)
	log.Infof("Adding %s interface to the zone and enabling routing\n", interfaceName)
	o, err = fm.FirewalldController.AddInterfaceToTheZone(interfaceName, fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	log.Debugf("issuing docker interface subnet\n")

	dockerInterfaceConfig := dockerInterface.IPAM.Config
	log.Debugf("docker interace subnet: %s\n", dockerInterfaceConfig[0].Subnet)
	o, err = fm.FirewalldController.AddRichRule(fmt.Sprintf("rule family=ipv4 source address=%s accept", dockerInterfaceConfig[0].Subnet), fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	o, err = fm.FirewalldController.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	o, err = fm.FirewalldController.TurnOnMasquarade(fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	// os.Exit(1)
	//adding docker to the zone and enabling routing
	log.Infof("Adding docker0 interface to the zone and enabling routing\n")
	o, err = fm.FirewalldController.AddInterfaceToTheZone("docker0", fm.FirewallConfig.ZoneName)
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	o, err = fm.FirewalldController.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	// o, err = fm.FirewalldController.EnableDockerRouting("docker0")
	// if err != nil {
	// 	return fmt.Errorf("%s\n%w", o, err)
	// }
	log.Infof("Reloading firewall\n")
	o, err = fm.FirewalldController.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	log.Infof("Restarting docker service\n")
	err = fm.FirewallHandler.RestartDockerService()
	if err != nil {
		return fmt.Errorf("cannot restart docker service: %s", err)
	}
	return nil
}
