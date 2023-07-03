package main

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/logging"
)

var log = logging.Log

func main() {
	log.SetLevel(logrus.InfoLevel)

	dockerManager, err := docker.NewTestDockerManager()
	if err != nil {
		log.Fatalln("Can't create instance of docker manager", err)
	}
	defer dockerManager.Cli.Close()

	ctx := context.Background()

	running, _ := dockerManager.IsContainerRunning(ctx, "test")
	log.Infof("Container 'test' is running: %t", running)

	// HEALTHCHECK is needed
	healthy, _ := dockerManager.IsContainerHealthy(ctx, "test")
	log.Infof("Container 'test' healthy: %t", healthy)

	_ = dockerManager.PauseContainer(ctx, "test")
	log.Infof("Container 'test' is paused")

	_ = dockerManager.UnpauseContainer(ctx, "test")
	log.Infof("Container 'test' is unpaused")

	_ = dockerManager.RestartContainer(ctx, "test", 2)
	log.Infof("Container 'test' is restarted")

	_ = dockerManager.KillAndDeleteContainer(ctx, "test")
	log.Infof("Container 'test' is killed")
}
