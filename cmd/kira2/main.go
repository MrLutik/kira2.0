package main

import (
	"context"

	"github.com/mrlutik/kira2.0/internal/adapters"
	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/errors"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/manager"
	"github.com/mrlutik/kira2.0/internal/systemd"
	"github.com/sirupsen/logrus"
)

const (
	sekaiDebFileName  = "sekai-linux-amd64.deb"
	interxDebFileName = "interx-linux-amd64.deb"
)

var log = logging.Log

func main() {
	// TODO: change level by flag
	log.SetLevel(logrus.DebugLevel)

	systemd.DockerServiceManagement()

	dockerManager, err := docker.NewTestDockerManager()
	errors.HandleErr("Can't create instance of docker manager", err)
	defer dockerManager.Cli.Close()

	ctx := context.Background()
	// TODO: Instead of HARDCODE - using config file
	cfg := config.NewKiraConfig()

	docker.VerifyingDockerImage(ctx, dockerManager, cfg)

	adapters.DownloadBinaries(ctx, cfg, sekaiDebFileName, interxDebFileName)

	manager.InitAndRunSekaid(ctx, dockerManager, cfg, sekaiDebFileName)
	manager.InitAndRunInterxd(ctx, dockerManager, cfg, interxDebFileName)
}
