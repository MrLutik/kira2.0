package manager

import (
	"context"

	"github.com/mrlutik/kira2.0/internal/errors"
)

const debFileDestInContainer = "/tmp/"

// ContainerRunner describe public methods of container managers
type ContainerRunner interface {
	InitAndRun(context.Context)
}

func (s *SekaidManager) InitAndRun(ctx context.Context) {
	check, err := s.containerManager.CheckForContainersName(ctx, s.config.SekaidContainerName)
	errors.HandleFatalErr("Checking container names", err)
	if check {
		err = s.containerManager.StopAndDeleteContainer(ctx, s.config.SekaidContainerName)
		errors.HandleFatalErr("Deleting container", err)
	}

	errors.HandleFatalErr("Sekaid managing", err)
	err = s.containerManager.InitAndCreateContainer(ctx, s.ContainerConfig, s.SekaidNetworkingConfig, s.SekaiHostConfig, s.config.SekaidContainerName)
	errors.HandleFatalErr("Sekaid initialization", err)

	err = s.containerManager.SendFileToContainer(ctx, s.config.SekaiDebFileName, debFileDestInContainer, s.config.SekaidContainerName)
	errors.HandleFatalErr("Sending file to container", err)

	// TODO Do we need to delete file after sending?

	err = s.containerManager.InstallDebPackage(ctx, s.config.SekaidContainerName, debFileDestInContainer+s.config.SekaiDebFileName)
	errors.HandleFatalErr("Installing dep package in container", err)

	err = s.runSekaidContainer(ctx)
	errors.HandleFatalErr("Setup container", err)
}

func (i *InterxManager) InitAndRun(ctx context.Context) {
	check, err := i.containerManager.CheckForContainersName(ctx, i.config.InterxContainerName)
	errors.HandleFatalErr("Checking container names", err)
	if check {
		i.containerManager.StopAndDeleteContainer(ctx, i.config.InterxContainerName)
		errors.HandleFatalErr("Deleting container", err)
	}

	errors.HandleFatalErr("Interx managing", err)
	err = i.containerManager.InitAndCreateContainer(ctx, i.ContainerConfig, i.InterxNetworkConfig, i.InterxHostConfig, i.config.InterxContainerName)
	errors.HandleFatalErr("Interx initialization", err)

	err = i.containerManager.SendFileToContainer(ctx, i.config.InterxDebFileName, debFileDestInContainer, i.config.InterxContainerName)
	errors.HandleFatalErr("Sending file to container", err)

	err = i.containerManager.InstallDebPackage(ctx, i.config.InterxContainerName, debFileDestInContainer+i.config.InterxDebFileName)
	errors.HandleFatalErr("Installing dep package in container", err)

	err = i.runInterxContainer(ctx)
	errors.HandleFatalErr("Setup container", err)
}
