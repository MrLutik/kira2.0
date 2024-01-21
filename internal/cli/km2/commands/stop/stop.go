package stop

import (
	"context"
	"time"

	"github.com/docker/docker/client"
	"github.com/mrlutik/kira2.0/internal/config/controller"
	"github.com/mrlutik/kira2.0/internal/config/handler"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/manager"
	"github.com/mrlutik/kira2.0/internal/osutils"
	"github.com/mrlutik/kira2.0/internal/systemd"
	"github.com/mrlutik/kira2.0/internal/utils"
	"github.com/spf13/cobra"
)

const (
	use   = "stop"
	short = "Stop kira node"
	long  = "Stop node if running"
)

func Stop(log *logging.Logger) *cobra.Command {
	log.Info("Adding `start` command...")
	startCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(cmd *cobra.Command, _ []string) {
			mainStop(cmd, log)
		},
	}

	return startCmd
}

func mainStop(_ *cobra.Command, log *logging.Logger) {
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

	helper := utils.NewHelperManager(containerManager, containerManager, utilsOS, cfg, log)

	sekaiManager, err := manager.NewSekaidManager(containerManager, helper, dockerManager, cfg, log)
	if err != nil {
		log.Fatalf("Can't create new 'sekai' manager instance: %s", err)
	}
	err = sekaiManager.StopSekaid(ctx)
	if err != nil {
		log.Fatalf("Stopping of joiner node failed: %s", err)
	}

	interxManager, err := manager.NewInterxManager(containerManager, cfg, log)
	if err != nil {
		log.Fatalf("Can't create new 'interx' manager instance: %s", err)
	}
	err = interxManager.StopInterx(ctx)
	if err != nil {
		log.Fatalf("Stopping interx failed: %s", err)
	}
}
