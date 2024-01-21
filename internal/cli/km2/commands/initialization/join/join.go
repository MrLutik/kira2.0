package join

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/docker/docker/client"
	"github.com/spf13/cobra"

	"github.com/mrlutik/kira2.0/internal/adapters"
	"github.com/mrlutik/kira2.0/internal/config/controller"
	"github.com/mrlutik/kira2.0/internal/config/handler"
	"github.com/mrlutik/kira2.0/internal/docker"
	firewallManager "github.com/mrlutik/kira2.0/internal/firewall/manager"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/manager"
	"github.com/mrlutik/kira2.0/internal/osutils"
	"github.com/mrlutik/kira2.0/internal/systemd"
	"github.com/mrlutik/kira2.0/internal/utils"
)

const (
	// Command information
	use   = "join"
	short = "Join to network"
	long  = "Join to existing network"

	// Flags naming
	sekaiVersionFlag  = "sekai-version"
	interxVersionFlag = "interx-version"
	interxPortFlag    = "interx-port"
	recoveringFlag    = "recover"
	ipFlag            = "ip"
	rpcPortFlag       = "rpc-port"
	p2pPortFlag       = "p2p-port"

	// Regular expression to match IPv4 and IPv6 addresses.
	ipRegex = `^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$|^(?:[0-9A-Fa-f]{1,4}:){7}[0-9A-Fa-f]{1,4}$`

	// Regular expression to match valid port numbers (from 1 to 65535).
	portRegex = `^([1-9]\d{0,4}|[1-5]\d{4}|6[0-4]\d{3}|65[0-4]\d{2}|655[0-2]\d|6553[0-5])$`

	envGithubTokenVariableName = "GITHUB_TOKEN"
)

func Join(log *logging.Logger) *cobra.Command {
	log.Info("Adding `join` command...")
	joinCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, _ []string) {
			if err := validateFlags(cmd); err != nil {
				log.Errorf("Some flag is not valid: %s", err)
				if err := cmd.Help(); err != nil {
					log.Fatalf("Error displaying help: %s", err)
				}
				return
			}
			mainJoin(cmd, log)
		},
	}

	joinCmd.Flags().String(ipFlag, "", "IP address of the validator to join")
	if err := joinCmd.MarkFlagRequired(ipFlag); err != nil {
		log.Fatalf("Failed to mark '%s' flag as required: %s", ipFlag, err)
	}

	joinCmd.Flags().String(interxPortFlag, "11000", "Interx port of the validator")
	joinCmd.Flags().String(rpcPortFlag, "26657", "Sekaid RPC port of the validator")
	joinCmd.Flags().String(p2pPortFlag, "26656", "Sekaid P2P port of the validator")
	joinCmd.PersistentFlags().Bool(recoveringFlag, false, "If true recover keys and mnemonic from master mnemonic, otherwise generate random one")
	joinCmd.Flags().String(sekaiVersionFlag, "latest", "Set this flag to choose what sekai version will be initialized")
	joinCmd.Flags().String(interxVersionFlag, "latest", "Set this flag to choose what interx version will be initialized")

	return joinCmd
}

func validateFlags(cmd *cobra.Command) error {
	ip, err := cmd.Flags().GetString("ip")
	if err != nil {
		return fmt.Errorf("error retrieving 'ip' flag: %w", err)
	}
	if !isValidIP(ip) {
		return fmt.Errorf("%w: '%s' is not valid", ErrInvalidIPAddress, ip)
	}

	interxPort, err := cmd.Flags().GetString("interx-port")
	if err != nil {
		return fmt.Errorf("error retrieving 'interx-port' flag: %w", err)
	}
	if !isValidPort(interxPort) {
		return fmt.Errorf("%w: '%s' is not valid", ErrInvalidInterxPort, interxPort)
	}

	rpcPort, err := cmd.Flags().GetString("rpc-port")
	if err != nil {
		return fmt.Errorf("error retrieving 'rpc-port' flag: %w", err)
	}
	if !isValidPort(rpcPort) {
		return fmt.Errorf("%w: '%s' is not valid", ErrInvalidSekaidRPCPort, rpcPort)
	}

	p2pPort, err := cmd.Flags().GetString("p2p-port")
	if err != nil {
		return fmt.Errorf("error retrieving 'p2p-port' flag: %w", err)
	}
	if !isValidPort(p2pPort) {
		return fmt.Errorf("%w: '%s' is not valid", ErrInvalidSekaidP2PPort, p2pPort)
	}

	_, err = cmd.Flags().GetString(sekaiVersionFlag)
	if err != nil {
		return fmt.Errorf("error retrieving %s flag: %w", sekaiVersionFlag, err)
	}
	_, err = cmd.Flags().GetString(interxVersionFlag)
	if err != nil {
		return fmt.Errorf("error retrieving %s flag: %w", interxVersionFlag, err)
	}
	return nil
}

func isValidIP(ip string) bool {
	match, err := regexp.MatchString(ipRegex, ip)
	if err != nil {
		return false
	}

	return match
}

func isValidPort(port string) bool {
	match, err := regexp.MatchString(portRegex, port)
	if err != nil {
		return false
	}

	return match
}

func mainJoin(cmd *cobra.Command, log *logging.Logger) {
	err := systemd.DockerServiceManagement(log)
	if err != nil {
		log.Fatalf("Docker service management failed: %s", err)
	}

	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Can't initialize the Docker client: %s", err)
	}

	utilsOS := osutils.NewOSUtils(log)

	dockerManager := docker.NewTestDockerManager(client, utilsOS, log)
	if err != nil {
		log.Fatalf("Can't create instance of docker manager: %s", err)
	}
	defer dockerManager.CloseClient()

	containerManager := docker.NewTestContainerManager(client, log)
	if err != nil {
		log.Fatalf("Can't create instance of container manager: %s", err)
	}
	defer containerManager.CloseClient()

	// TODO make flexible setting timeout
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancelFunc()

	ip, err := cmd.Flags().GetString(ipFlag)
	if err != nil {
		log.Fatalf("Error retrieving flag '%s': %s", ipFlag, err)
	}
	interxPort, err := cmd.Flags().GetString(interxPortFlag)
	if err != nil {
		log.Fatalf("Error retrieving flag '%s': %s", interxPortFlag, err)
	}
	sekaidRPCPort, err := cmd.Flags().GetString(rpcPortFlag)
	if err != nil {
		log.Fatalf("Error retrieving flag '%s': %s", rpcPortFlag, err)
	}
	sekaidP2PPort, err := cmd.Flags().GetString(p2pPortFlag)
	if err != nil {
		log.Fatalf("Error retrieving flag '%s': %s", p2pPortFlag, err)
	}

	// Information about validator we need to join
	joinerCfg := &manager.TargetSeedKiraConfig{
		IpAddress:     ip,
		InterxPort:    interxPort,
		SekaidRPCPort: sekaidRPCPort,
		SekaidP2PPort: sekaidP2PPort,
	}
	joinerManager := manager.NewJoinerManager(joinerCfg)

	recover, err := cmd.Flags().GetBool(recoveringFlag)
	if err != nil {
		log.Fatalf("Error retrieving flag '%s': %s", recoveringFlag, err)
	}

	cfg, err := joinerManager.GenerateKiraConfig(ctx, recover)
	if err != nil {
		log.Fatalf("Can't get kira config: %s", err)
	}

	sekaiVersion, err := cmd.Flags().GetString(sekaiVersionFlag)
	if err != nil {
		log.Fatalf("Error retrieving flag '%s': %s", sekaiVersionFlag, err)
	}
	interxVersion, err := cmd.Flags().GetString(interxVersionFlag)
	if err != nil {
		log.Fatalf("Error retrieving flag '%s': %s", interxVersionFlag, err)
	}

	if sekaiVersion != cfg.SekaiVersion || interxVersion != cfg.InterxVersion {
		cfg.SekaiVersion = sekaiVersion
		cfg.InterxVersion = interxVersion

		configController := controller.NewConfigController(handler.NewHandler(utilsOS, log), utilsOS, log)
		err = configController.ChangeConfigFile(cfg)
		if err != nil {
			log.Fatalf("Can't change config file: %s", err)
		}
	}

	// TODO method called twice
	genesis, err := joinerManager.GetVerifiedGenesisFile(ctx)
	if err != nil {
		log.Fatalf("Can't get genesis: %s", err)
	}

	// TODO this docker service restart has to be after docker and firewalld installation, i'm doing it here because launcher is not ready
	// temp remove docker restarting, only need once after firewalld installation
	// err = dockerManager.RestartDockerService()
	// if err != nil {
	//     log.Fatalf("Restarting docker service: %s", err)
	// }

	err = docker.VerifyingDockerEnvironment(ctx, dockerManager, cfg)
	if err != nil {
		log.Fatalf("Verifying docker environment failed: %s", err)
	}

	err = containerManager.CleanupContainersAndVolumes(ctx, cfg)
	if err != nil {
		log.Fatalf("Cleaning docker volume and containers: %s", err)
	}

	token, exists := os.LookupEnv(envGithubTokenVariableName)
	if !exists {
		log.Fatalf("'%s' variable is not set", envGithubTokenVariableName)
	}
	adapterGitHub := adapters.NewGitHubAdapter(ctx, token)
	// TODO Do we need to safe deb packages in temporary directory?
	// Right now the files are downloaded in current directory, where the program starts
	adapterGitHub.MustDownloadBinaries(ctx, cfg)

	firewallManager, err := firewallManager.NewFirewallManager(dockerManager, utilsOS, cfg, log)
	if err != nil {
		log.Fatalf("Initialization of firewall manager failed: %s", err)
	}

	check, err := firewallManager.CheckFirewallSetUp(ctx)
	if err != nil {
		log.Fatalf("Checking valid firewalld setup failed: %s", err)
	}

	if !check {
		err = firewallManager.SetUpFirewall(ctx)
		if err != nil {
			log.Fatalf("Setup firewall failed: %s", err)
		}
	}

	helper := utils.NewHelperManager(containerManager, containerManager, utilsOS, cfg, log)

	sekaiManager, err := manager.NewSekaidManager(containerManager, helper, dockerManager, cfg, log)
	if err != nil {
		log.Fatalf("Can't create new 'sekai' manager instance: %s", err)
	}

	err = sekaiManager.InitJoiner(ctx, genesis)
	if err != nil {
		log.Fatalf("Initialization of joiner node failed: %s", err)
	}
	err = sekaiManager.RunSekaid(ctx)
	if err != nil {
		log.Fatalf("Running of joiner node failed: %s", err)
	}

	log.Infof("Waiting for %v\n", cfg.TimeBetweenBlocks)
	time.Sleep(cfg.TimeBetweenBlocks + time.Second)

	interxManager, err := manager.NewInterxManager(containerManager, cfg, log)
	if err != nil {
		log.Fatalf("Can't create new 'interx' manager instance: %s", err)
	}

	err = interxManager.InitInterx(ctx)
	if err != nil {
		log.Fatalf("Initialization of interx failed: %s", err)
	}
	err = interxManager.RunInterx(ctx)
	if err != nil {
		log.Fatalf("Running interx failed: %s", err)
	}
}
