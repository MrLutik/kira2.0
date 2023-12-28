package new

import (
	"context"
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

var log = logging.Log
var recover bool

const (
	use   = "new"
	short = "Create new blockchain network"
	long  = "Create new blockchain network from genesis file"
)

func New() *cobra.Command {
	log.Info("Adding `join` command...")
	newCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, _ []string) {
			mainNew(cmd)
		},
	}

	newCmd.MarkFlagRequired("ip")

	return newCmd
}

func mainNew(*cobra.Command) {
	systemd.DockerServiceManagement()

	dockerManager, err := docker.NewTestDockerManager()
	errors.HandleFatalErr("Can't create instance of docker manager", err)
	defer dockerManager.Cli.Close()

	containerManager, err := docker.NewTestContainerManager()
	errors.HandleFatalErr("Can't create instance of container docker manager", err)
	defer containerManager.Cli.Close()

	ctx := context.Background()

	cfg, err := configFileController.ReadOrCreateConfig()
	errors.HandleFatalErr("Error while reading cfg file", err)
	cfg.Recover = recover
	log.Traceln(recover)
	//todo this docker service restart has to be after docker and firewalld instalation, im doin it here because im laucnher is not ready
	err = dockerManager.RestartDockerService()
	errors.HandleFatalErr("Restarting docker service", err)
	docker.VerifyingDockerEnvironment(ctx, dockerManager, cfg)
	containerManager.CleanupContainersAndVolumes(ctx, cfg)
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
	time.Sleep(cfg.TimeBetweenBlocks + time.Second)
	interxManager, err := manager.NewInterxManager(containerManager, cfg)
	errors.HandleFatalErr("Error creating new 'interx' manager instance:", err)
	interxManager.MustInitInterx(ctx)
	interxManager.MustRunInterx(ctx)
}
