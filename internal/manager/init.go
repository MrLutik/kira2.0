package manager

import (
	"context"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/errors"
)

const debFileDestInContainer = "/tmp/"

func InitAndRunSekaid(ctx context.Context, dockerManager *docker.DockerManager, cfg *config.KiraConfig, sekaiDebFileName string) {
	check, err := dockerManager.CheckForContainersName(ctx, cfg.SekaidContainerName)
	errors.HandleErr("Checking container names", err)
	if check {
		err = dockerManager.StopAndDeleteContainer(ctx, cfg.SekaidContainerName)
		errors.HandleErr("Deleting container", err)
	}

	err = dockerManager.CheckOrCreateNetwork(ctx, cfg.DockerNetworkName)
	errors.HandleErr("Docker networking", err)

	sekaidManager, err := NewSekaidManager(dockerManager, cfg)
	errors.HandleErr("Sekaid managing", err)
	err = dockerManager.InitAndCreateContainer(ctx, sekaidManager.ContainerConfig, sekaidManager.SekaidNetworkingConfig, sekaidManager.SekaiHostConfig, cfg.SekaidContainerName)
	errors.HandleErr("Sekaid initialization", err)

	err = dockerManager.SendFileToContainer(ctx, sekaiDebFileName, debFileDestInContainer, cfg.SekaidContainerName)
	errors.HandleErr("Sending file to container", err)

	err = dockerManager.InstallDebPackage(ctx, cfg.SekaidContainerName, debFileDestInContainer+sekaiDebFileName)
	errors.HandleErr("Installing dep package in container", err)

	err = sekaidManager.RunSekaidContainer(ctx)
	errors.HandleErr("Setup container", err)
}

func InitAndRunInterxd(ctx context.Context, dockerManager *docker.DockerManager, cfg *config.KiraConfig, interxDebFileName string) {
	check, err := dockerManager.CheckForContainersName(ctx, cfg.InterxContainerName)
	errors.HandleErr("Checking container names", err)
	if check {
		dockerManager.StopAndDeleteContainer(ctx, cfg.InterxContainerName)
		errors.HandleErr("Deleting container", err)
	}

	interxManager, err := NewInterxManager(dockerManager, cfg)
	errors.HandleErr("Interx managing", err)
	err = dockerManager.InitAndCreateContainer(ctx, interxManager.ContainerConfig, interxManager.SekaidNetworkingConfig, interxManager.SekaiHostConfig, cfg.InterxContainerName)
	errors.HandleErr("Interx initialization", err)

	err = dockerManager.SendFileToContainer(ctx, interxDebFileName, debFileDestInContainer, cfg.InterxContainerName)
	errors.HandleErr("Sending file to container", err)

	err = dockerManager.InstallDebPackage(ctx, cfg.InterxContainerName, debFileDestInContainer+interxDebFileName)
	errors.HandleErr("Installing dep package in container", err)

	err = interxManager.RunInterxContainer(ctx)
	errors.HandleErr("Setup container", err)
}
