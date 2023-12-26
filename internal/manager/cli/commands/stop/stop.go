package stop

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
	use   = "stop"
	short = "Stop kira node"
	long  = "Stop node if running"
)

var log = logging.Log

func Stop() *cobra.Command {
	log.Info("Adding `start` command...")
	startCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, _ []string) {
			mainStop(cmd)
		},
	}

	return startCmd
}

func mainStop(cmd *cobra.Command) {
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
	docker.VerifyingDockerEnvironment(ctx, dockerManager, cfg)
	sekaiManager, err := manager.NewSekaidManager(containerManager, dockerManager, cfg)
	errors.HandleFatalErr("Error creating new 'sekai' manager instance", err)
	sekaiManager.MustStopSekaid(ctx)

	interxManager, err := manager.NewInterxManager(containerManager, cfg)
	errors.HandleFatalErr("Error creating new 'interx' manager instance:", err)
	interxManager.MustStopInterxd(ctx)
}
