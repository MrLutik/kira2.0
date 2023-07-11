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

var log = logging.Log

func main() {
	// TODO: change level by flag
	log.SetLevel(logrus.DebugLevel)

	systemd.DockerServiceManagement()

	dockerManager, err := docker.NewTestDockerManager()
	errors.HandleFatalErr("Can't create instance of docker manager", err)
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
		"sekai-linux-amd64.deb",
		"interx-linux-amd64.deb",
		10500,
	)

	docker.VerifyingDockerImage(ctx, dockerManager, cfg.DockerImageName+":"+cfg.DockerImageVersion)

	adapters.DownloadBinaries(ctx, cfg, cfg.SekaiDebFileName, cfg.InterxDebFileName)

	sekaiManager, err := manager.NewSekaidManager(dockerManager, cfg)
	errors.HandleFatalErr("Error creating new 'sekai' manager instance", err)
	sekaiManager.InitAndRun(ctx)

	interxManager, err := manager.NewInterxManager(dockerManager, cfg)
	errors.HandleFatalErr("Error creating new 'interx' manager instance:", err)
	interxManager.InitAndRun(ctx)
}
