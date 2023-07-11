package manager

import (
	"context"

	"github.com/mrlutik/kira2.0/internal/errors"
)

const debFileDestInContainer = "/tmp/"

type Repository interface {
	InitAndRun(context.Context)
}

func (s *SekaidManager) InitAndRun(ctx context.Context) {
	check, err := s.dockerManager.CheckForContainersName(ctx, s.config.SekaidContainerName)
	errors.HandleErr("Checking container names", err)
	if check {
		err = s.dockerManager.StopAndDeleteContainer(ctx, s.config.SekaidContainerName)
		errors.HandleErr("Deleting container", err)
	}

	err = s.dockerManager.CheckOrCreateNetwork(ctx, s.config.DockerNetworkName)
	errors.HandleErr("Docker networking", err)

	errors.HandleErr("Sekaid managing", err)
	err = s.dockerManager.InitAndCreateContainer(ctx, s.ContainerConfig, s.SekaidNetworkingConfig, s.SekaiHostConfig, s.config.SekaidContainerName)
	errors.HandleErr("Sekaid initialization", err)

	err = s.dockerManager.SendFileToContainer(ctx, s.config.SekaiDebFileName, debFileDestInContainer, s.config.SekaidContainerName)
	errors.HandleErr("Sending file to container", err)

	err = s.dockerManager.InstallDebPackage(ctx, s.config.SekaidContainerName, debFileDestInContainer+s.config.SekaiDebFileName)
	errors.HandleErr("Installing dep package in container", err)

	err = s.RunSekaidContainer(ctx)
	errors.HandleErr("Setup container", err)
}

func (i *InterxManager) InitAndRun(ctx context.Context) {
	check, err := i.dockerClient.CheckForContainersName(ctx, i.config.InterxContainerName)
	errors.HandleErr("Checking container names", err)
	if check {
		i.dockerClient.StopAndDeleteContainer(ctx, i.config.InterxContainerName)
		errors.HandleErr("Deleting container", err)
	}

	errors.HandleErr("Interx managing", err)
	err = i.dockerClient.InitAndCreateContainer(ctx, i.ContainerConfig, i.SekaidNetworkingConfig, i.SekaiHostConfig, i.config.InterxContainerName)
	errors.HandleErr("Interx initialization", err)

	err = i.dockerClient.SendFileToContainer(ctx, i.config.InterxDebFileName, debFileDestInContainer, i.config.InterxContainerName)
	errors.HandleErr("Sending file to container", err)

	err = i.dockerClient.InstallDebPackage(ctx, i.config.InterxContainerName, debFileDestInContainer+i.config.InterxDebFileName)
	errors.HandleErr("Installing dep package in container", err)

	err = i.RunInterxContainer(ctx)
	errors.HandleErr("Setup container", err)
}
