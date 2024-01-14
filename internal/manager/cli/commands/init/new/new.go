package new

import (
	"context"
	"fmt"
	"time"

	"github.com/mrlutik/kira2.0/internal/adapters"
	"github.com/mrlutik/kira2.0/internal/config/configFileController"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/errors"
	"github.com/mrlutik/kira2.0/internal/firewall/firewallManager"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/manager"
	"github.com/mrlutik/kira2.0/internal/systemd"
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
)

var (
	log     = logging.Log
	recover bool
)

func New() *cobra.Command {
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
			mainNew(cmd)
		},
	}

	newCmd.Flags().String(sekaiVersionFlag, "latest", "Set this flag to choose what sekai version will be initialized")
	newCmd.Flags().String(interxVersionFlag, "latest", "Set this flag to choose what interx version will be initialized")

	return newCmd
}

func validateFlags(cmd *cobra.Command) error {
	sekaiVersion, err := cmd.Flags().GetString(sekaiVersionFlag)
	if err != nil {
		return fmt.Errorf("error retrieving <%s> flag: %s", sekaiVersion, err)
	}
	interxVersion, err := cmd.Flags().GetString(interxVersionFlag)
	if err != nil {
		return fmt.Errorf("error retrieving <%s> flag: %s", interxVersion, err)
	}
	return nil
}

func mainNew(cmd *cobra.Command) {
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

	cfg, err := configFileController.ReadOrCreateConfig()
	errors.HandleFatalErr("Error while reading cfg file", err)

	sekaiVersion, err := cmd.Flags().GetString(sekaiVersionFlag)
	if err != nil {
		errors.HandleFatalErr(fmt.Sprintf("Error retrieving flag '%s'", sekaiVersionFlag), err)
	}

	interxVersion, err := cmd.Flags().GetString(interxVersionFlag)
	if err != nil {
		errors.HandleFatalErr(fmt.Sprintf("Error retrieving flag '%s'", interxVersionFlag), err)
	}

	if sekaiVersion != cfg.SekaiVersion || interxVersion != cfg.InterxVersion {
		cfg.SekaiVersion = sekaiVersion
		cfg.InterxVersion = interxVersion
		err = configFileController.ChangeConfigFile(cfg)
		errors.HandleFatalErr("Can't change config file", err)
	}

	cfg.Recover = recover
	log.Traceln(recover)
	// todo this docker service restart has to be after docker and firewalld instalation, im doin it here because im laucnher is not ready
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
	errors.HandleFatalErr("Error creating new 'sekai' manager instance", err)
	sekaiManager.MustInitNew(ctx)
	sekaiManager.MustRunSekaid(ctx)
	log.Printf("Waiting for %v\n", cfg.TimeBetweenBlocks)
	time.Sleep(cfg.TimeBetweenBlocks + time.Second)
	interxManager, err := manager.NewInterxManager(containerManager, cfg)
	errors.HandleFatalErr("Error creating new 'interx' manager instance:", err)
	interxManager.MustInitInterx(ctx)
	interxManager.MustRunInterx(ctx)
}
