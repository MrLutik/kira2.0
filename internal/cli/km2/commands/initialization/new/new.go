package new

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/docker/docker/client"
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
	"github.com/spf13/cobra"
)

const (
	// Command information
	use   = "new"
	short = "Create new blockchain network"
	long  = "Create new blockchain network from genesis file"

	// Flags naming
	sekaiVersionFlag  = "sekai-version"
	interxVersionFlag = "interx-version"
	recoveringFlag    = "recover"

	envGithubTokenVariableName = "GITHUB_TOKEN"
)

func New(log *logging.Logger) *cobra.Command {
	log.Info("Adding `join` command...")
	newCmd := &cobra.Command{
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
			mainNew(cmd, log)
		},
	}

	newCmd.Flags().String(sekaiVersionFlag, "latest", "Set this flag to choose what sekai version will be initialized")
	newCmd.Flags().String(interxVersionFlag, "latest", "Set this flag to choose what interx version will be initialized")
	newCmd.PersistentFlags().Bool(recoveringFlag, false, "If true recover keys and mnemonic from master mnemonic, otherwise generate random one")

	return newCmd
}

func validateFlags(cmd *cobra.Command) error {
	_, err := cmd.Flags().GetString(sekaiVersionFlag)
	if err != nil {
		return fmt.Errorf("error retrieving '%s' flag: %w", sekaiVersionFlag, err)
	}
	_, err = cmd.Flags().GetString(interxVersionFlag)
	if err != nil {
		return fmt.Errorf("error retrieving '%s flag: %w", interxVersionFlag, err)
	}
	_, err = cmd.Flags().GetBool(recoveringFlag)
	if err != nil {
		return fmt.Errorf("error retrieving '%s flag: %w", recoveringFlag, err)
	}
	return nil
}

func mainNew(cmd *cobra.Command, log *logging.Logger) {
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

	configController := controller.NewConfigController(handler.NewHandler(utilsOS, log), utilsOS, log)
	cfg, err := configController.ReadOrCreateConfig()
	if err != nil {
		log.Fatalf("Can't read or create config file: %s", err)
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

	recover, err := cmd.Flags().GetBool(recoveringFlag)
	if err != nil {
		log.Fatalf("Error retrieving flag '%s': %s", recoveringFlag, err)
	}

	// TODO Dmytro task - add `recovery` usage!
	log.Tracef("Recover flag is: %t", recover)

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

	// TODO Do we need to safe deb packages in temporary directory?
	// Right now the files are downloaded in current directory, where the program starts
	token, exists := os.LookupEnv(envGithubTokenVariableName)
	if !exists {
		log.Fatalf("'%s' variable is not set", envGithubTokenVariableName)
	}
	adapterGitHub := adapters.NewGitHubAdapter(ctx, token, log)
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

	err = sekaiManager.InitNew(ctx)
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
