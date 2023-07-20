package join

import (
	"context"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/errors"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/manager"
	"github.com/mrlutik/kira2.0/internal/systemd"
	"github.com/spf13/cobra"
)

const (
	use   = "join"
	short = "Join to sekaid network"
	long  = "Joining a running sekaid network"
)

var log = logging.Log

func Join() *cobra.Command {
	log.Info("Adding `join` command...")
	joinCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(_ *cobra.Command, _ []string) {
			mainJoin()
		},
	}

	return joinCmd
}

func mainJoin() {
	systemd.DockerServiceManagement()

	dockerManager, err := docker.NewTestDockerManager()
	errors.HandleFatalErr("Can't create instance of docker manager", err)
	defer dockerManager.Cli.Close()

	containerManager, err := docker.NewTestContainerManager()
	errors.HandleFatalErr("Can't create instance of container docker manager", err)
	defer containerManager.Cli.Close()

	ctx := context.Background()

	kiraCfg := &config.KiraConfig{
		DockerImageName:    "ghcr.io/kiracore/docker/kira-base",
		DockerImageVersion: "v0.13.11",
		DockerNetworkName:  "test",
	}

	docker.VerifyingDockerEnvironment(ctx, dockerManager, kiraCfg)

	// Information about validator we need to join
	joinerCfg := &manager.JoinerKiraConfig{
		IpAddress:     "172.18.0.1",
		InterxPort:    "11000",
		SekaidRPCPort: "26657",
	}
	joinerManager := manager.NewJoinerManager(joinerCfg)
	genesis, err := joinerManager.GetGenesis(ctx)
	errors.HandleFatalErr("Can't get genesis file", err)

	log.Info(string(genesis))

	// TODO generate config based on genesis
	// for sekaid and interx manager
	// After - start sekaid and interx containers
	// using generated config in isolated docker network
}
