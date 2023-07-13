package main

import (
	"context"
	"time"

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

	containerManager, err := docker.NewTestContainerManager()
	errors.HandleFatalErr("Can't create instance of container docker manager", err)
	defer containerManager.Cli.Close()

	ctx := context.Background()

	// TODO: Instead of HARDCODE - reading config file
	// Note: we do not need the constructor for config, it is not readable right now
	// Using initialization of structure on the way reads better
	cfg := &config.KiraConfig{
		NetworkName:         "testnet-1",
		SekaidHome:          "/data/.sekai",
		InterxHome:          "/data/.interx",
		KeyringBackend:      "test",
		DockerImageName:     "ghcr.io/kiracore/docker/kira-base",
		DockerImageVersion:  "v0.13.11",
		DockerNetworkName:   "kira_network",
		SekaiVersion:        "latest", // or v0.3.16
		InterxVersion:       "latest", // or v0.4.33
		SekaidContainerName: "sekaid",
		InterxContainerName: "interx",
		VolumeName:          "kira_volume:/data",
		MnemonicDir:         "~/mnemonics",
		RpcPort:             "26657",
		GrpcPort:            "9090",
		InterxPort:          "11000",
		Moniker:             "VALIDATOR",
		SekaiDebFileName:    "sekai-linux-amd64.deb",
		InterxDebFileName:   "interx-linux-amd64.deb",
		TimeBetweenBlocks:   time.Second * 10,
	}

	docker.VerifyingDockerImage(ctx, dockerManager, cfg)

	// TODO Do we need to safe deb packages in temporary directory?
	// Right now the files are downloaded in current directory, where the program starts
	adapters.DownloadBinaries(ctx, cfg)

	sekaiManager, err := manager.NewSekaidManager(containerManager, cfg)
	errors.HandleFatalErr("Error creating new 'sekai' manager instance", err)
	sekaiManager.InitAndRun(ctx)

	interxManager, err := manager.NewInterxManager(containerManager, cfg)
	errors.HandleFatalErr("Error creating new 'interx' manager instance:", err)
	interxManager.InitAndRun(ctx)
}
