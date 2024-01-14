package start

import (
	"context"

	"github.com/mrlutik/kira2.0/internal/config/configFileController"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/errors"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/manager"
	"github.com/mrlutik/kira2.0/internal/systemd"
	"github.com/spf13/cobra"
)

const (
	use   = "start"
	short = "Start new sekaid network"
	long  = "Starting new genesis validator network"
)

var log = logging.Log

func Start() *cobra.Command {
	log.Info("Adding `start` command...")
	startCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, _ []string) {
			mainStart(cmd)
		},
	}

	return startCmd
}

func mainStart(cmd *cobra.Command) {
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

	// todo this docker service restart has to be after docker and firewalld instalation, im doin it here because im laucnher is not ready
	docker.VerifyingDockerEnvironment(ctx, dockerManager, cfg)
	// TODO Do we need to safe deb packages in temporary directory?
	// Right now the files are downloaded in current directory, where the program starts

	sekaiManager, err := manager.NewSekaidManager(containerManager, dockerManager, cfg)
	errors.HandleFatalErr("Error creating new 'sekai' manager instance", err)
	sekaiManager.MustRunSekaid(ctx)

	interxManager, err := manager.NewInterxManager(containerManager, cfg)
	errors.HandleFatalErr("Error creating new 'interx' manager instance:", err)
	interxManager.MustRunInterx(ctx)
}
