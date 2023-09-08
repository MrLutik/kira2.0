package firewallManager

import (
	"context"
	"fmt"
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

	return &FirewallManager{FirewalldController: c, DockerManager: dockerManager, FirewallHandler: h, FirewallConfig: cfg}
}

var log = logging.Log

func (fm *FirewallManager) SetUpFirewall(ctx context.Context) error {
	log.Infof("***FIREWALL SETUP***\n")

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
		o, err := fm.FirewalldController.CreateNewFirewalldZone()
		if err != nil {
			return fmt.Errorf("%s\n%w", o, err)
		}
		o, err = fm.FirewalldController.ReloadFirewall()
		if err != nil {
			return fmt.Errorf("%s\n%w", o, err)
		}
	}

	log.Infof("Switching into %s firewalldZone\n", fm.FirewallConfig.ZoneName)
	o, err := fm.FirewalldController.ChangeDefaultZone()
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
	err = fm.FirewallHandler.OpenPorts(sysports)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	log.Infof("Opening kira ports\n")
	err = fm.FirewallHandler.OpenPorts(fm.FirewallConfig.PortsToOpen)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	//adding interface that has internet acces
	log.Infof("Adding interface that has internet acces\n")
	internetInterface := osutils.GetInternetInterface()
	o, err = fm.FirewalldController.AddInterfaceToTheZone(internetInterface)
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	interfaceName, err := fm.FirewallHandler.GetDockerNetworkInterfaceName(ctx, fm.KiraConfig.DockerNetworkName, fm.DockerManager)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	log.Infof("Adding %s interface to the zone and enabling routing\n", interfaceName)
	o, err = fm.FirewalldController.AddInterfaceToTheZone(interfaceName)
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	//adding docker to the zone and enabling routing
	log.Infof("Adding docker0 interface to the zone and enabling routing\n")
	o, err = fm.FirewalldController.AddInterfaceToTheZone("docker0")
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	o, err = fm.FirewalldController.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}
	o, err = fm.FirewalldController.EnableDockerRouting("docker0")
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	//save
	o, err = fm.FirewalldController.ReloadFirewall()
	if err != nil {
		return fmt.Errorf("%s\n%w", o, err)
	}

	return nil
}
