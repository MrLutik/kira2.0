package join

import (
	"context"
	"fmt"
	"regexp"
	"time"

	"github.com/spf13/cobra"

	"github.com/mrlutik/kira2.0/internal/adapters"
	"github.com/mrlutik/kira2.0/internal/config/configFileController"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/errors"
	"github.com/mrlutik/kira2.0/internal/firewall/firewallManager"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/manager"
	"github.com/mrlutik/kira2.0/internal/systemd"
)

const (
	// Command information
	use   = "join"
	short = "Join to network"
	long  = "Join to existing network"

	// Flags naming
	sekaiVersionFlag  = "sekai-version"
	interxVersionFlag = "interx-version"

	// Regular expression to match IPv4 and IPv6 addresses.
	ipRegex = `^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$|^(?:[0-9A-Fa-f]{1,4}:){7}[0-9A-Fa-f]{1,4}$`

	// Regular expression to match valid port numbers (from 1 to 65535).
	portRegex = `^([1-9]\d{0,4}|[1-5]\d{4}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])$`
)

var (
	log     = logging.Log
	recover bool
)

func Join() *cobra.Command {
	log.Info("Adding `join` command...")
	joinCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, _ []string) {
			if err := validateFlags(cmd); err != nil {
				log.Errorf("Some flag is not valid: %s", err)
				cmd.Help()
				return
			}
			mainJoin(cmd)
		},
	}

	joinCmd.Flags().String("ip", "", "IP address of the validator to join")
	joinCmd.MarkFlagRequired("ip")

	joinCmd.Flags().String("interx-port", "11000", "Interx port of the validator")
	joinCmd.Flags().String("rpc-port", "26657", "Sekaid RPC port of the validator")
	joinCmd.Flags().String("p2p-port", "26656", "Sekaid P2P port of the validator")
	joinCmd.PersistentFlags().Bool("recover", false, "If true recover keys and mnemonic from master mnemonic, otherwise generate random one")

	joinCmd.Flags().String(sekaiVersionFlag, "latest", "Set this flag to choose what sekai version will be initialized")
	joinCmd.Flags().String(interxVersionFlag, "latest", "Set this flag to choose what interx version will be initialized")
	return joinCmd
}

func validateFlags(cmd *cobra.Command) error {
	ip, err := cmd.Flags().GetString("ip")
	if err != nil {
		return fmt.Errorf("error retrieving 'ip' flag: %s", err)
	}
	if !isValidIP(ip) {
		return fmt.Errorf("'%s' is not a valid IP address", ip)
	}

	interxPort, err := cmd.Flags().GetString("interx-port")
	if err != nil {
		return fmt.Errorf("error retrieving 'interx-port' flag: %s", err)
	}
	if !isValidPort(interxPort) {
		return fmt.Errorf("'%s' is not a valid Interx port", interxPort)
	}

	rpcPort, err := cmd.Flags().GetString("rpc-port")
	if err != nil {
		return fmt.Errorf("error retrieving 'rpc-port' flag: %s", err)
	}
	if !isValidPort(rpcPort) {
		return fmt.Errorf("'%s' is not a valid Sekaid RPC port", rpcPort)
	}

	p2pPort, err := cmd.Flags().GetString("p2p-port")
	if err != nil {
		return fmt.Errorf("error retrieving 'p2p-port' flag: %s", err)
	}
	if !isValidPort(p2pPort) {
		return fmt.Errorf("'%s' is not a valid Sekaid P2P port", p2pPort)
	}

	_, err = cmd.Flags().GetString(sekaiVersionFlag)
	if err != nil {
		return fmt.Errorf("error retrieving %s flag: %s", sekaiVersionFlag, err)
	}
	_, err = cmd.Flags().GetString(interxVersionFlag)
	if err != nil {
		return fmt.Errorf("error retrieving %s flag: %s", interxVersionFlag, err)
	}
	return nil
}

func isValidIP(ip string) bool {
	match, err := regexp.MatchString(ipRegex, ip)
	if err != nil {
		log.Errorf("Can't match string, error: %s", err)
		return false
	}

	return match
}

func isValidPort(port string) bool {
	match, err := regexp.MatchString(portRegex, port)
	if err != nil {
		log.Errorf("Can't match string, error: %s", err)
		return false
	}

	return match
}

func mainJoin(cmd *cobra.Command) {
	systemd.DockerServiceManagement()

	dockerManager, err := docker.NewTestDockerManager()
	errors.HandleFatalErr("Can't create instance of docker manager", err)
	defer dockerManager.Cli.Close()

	containerManager, err := docker.NewTestContainerManager()
	errors.HandleFatalErr("Can't create instance of container docker manager", err)
	defer containerManager.Cli.Close()

	// TODO make flexible setting timeout
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancelFunc()

	// Skip errors here due to validateFlags method
	ip, _ := cmd.Flags().GetString("ip")
	interxPort, _ := cmd.Flags().GetString("interx-port")
	sekaidRPCPort, _ := cmd.Flags().GetString("rpc-port")
	sekaidP2PPort, _ := cmd.Flags().GetString("p2p-port")

	// Information about validator we need to join
	joinerCfg := &manager.TargetSeedKiraConfig{
		IpAddress:     ip,
		InterxPort:    interxPort,
		SekaidRPCPort: sekaidRPCPort,
		SekaidP2PPort: sekaidP2PPort,
	}
	joinerManager := manager.NewJoinerManager(joinerCfg)
	recover, _ = cmd.Flags().GetBool("recover")
	cfg, err := joinerManager.GenerateKiraConfig(ctx, recover)
	errors.HandleFatalErr("Can't get kira config", err)

	sekaiVersion, _ := cmd.Flags().GetString(sekaiVersionFlag)
	interxVersion, _ := cmd.Flags().GetString(interxVersionFlag)
	if sekaiVersion != cfg.SekaiVersion || interxVersion != cfg.InterxVersion {
		cfg.SekaiVersion = sekaiVersion
		cfg.InterxVersion = interxVersion
		err = configFileController.ChangeConfigFile(cfg)
		errors.HandleFatalErr("Can't change config file", err)
	}
	// TODO method called twice
	genesis, err := joinerManager.GetVerifiedGenesisFile(ctx)
	errors.HandleFatalErr("Can't get genesis", err)

	// todo this docker service restart has to be after docker and firewalld instalation, im doin it here because laucnher is not ready
	// temp remove docker restarting, only need once after firewalld instalation
	// err = dockerManager.RestartDockerService()
	errors.HandleFatalErr("Restarting docker service", err)
	docker.VerifyingDockerEnvironment(ctx, dockerManager, cfg)
	err = containerManager.CleanupContainersAndVolumes(ctx, cfg)
	errors.HandleFatalErr("Cleaning docker volume and containers", err)
	// TODO Do we need to safe deb packages in temporary directory?
	// Right now the files are downloaded in current directory, where the program starts
	adapters.MustDownloadBinaries(ctx, cfg)

	firewallManager := firewallManager.NewFirewallManager(dockerManager, cfg)
	check, err := firewallManager.CheckFirewallSetUp(ctx)
	errors.HandleFatalErr("Error while checking valid firewalld setup", err)
	if !check {
		err = firewallManager.SetUpFirewall(ctx)
		errors.HandleFatalErr("Error while setuping firewall", err)
	}

	sekaiManager, err := manager.NewSekaidManager(containerManager, dockerManager, cfg)
	errors.HandleFatalErr("Can't create new 'sekai' manager instance", err)
	sekaiManager.MustInitJoiner(ctx, genesis)
	sekaiManager.MustRunSekaid(ctx)
	log.Printf("Waiting for %v\n", cfg.TimeBetweenBlocks)
	time.Sleep(cfg.TimeBetweenBlocks + time.Second)
	interxManager, err := manager.NewInterxManager(containerManager, cfg)
	errors.HandleFatalErr("Can't create new 'interx' manager instance", err)
	interxManager.MustInitInterx(ctx)
	interxManager.MustRunInterx(ctx)
}
