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

	// TODO: Instead of HARDCODE - reading config file
	cfg := config.NewConfig(
		"testnet-1",
		"/data/.sekai",
		"/data/.interx",
		"test",
		"ghcr.io/kiracore/docker/kira-base",
		"v0.13.11",
		"kira_network",
		"latest", // or v0.3.16
		"latest", // or v0.4.33
		"sekaid",
		"interx",
		"kira_volume:/data",
		"~/mnemonics",
		"26657",
		"9090",
		"11000",
		"VALIDATOR",
	)

	docker.VerifyingDockerImage(ctx, dockerManager, cfg.DockerImageName+":"+cfg.DockerImageVersion)

	adapters.DownloadBinaries(ctx, cfg, sekaiDebFileName, interxDebFileName)

	manager.InitAndRunSekaid(ctx, dockerManager, cfg, sekaiDebFileName)
	manager.InitAndRunInterxd(ctx, dockerManager, cfg, interxDebFileName)
}
