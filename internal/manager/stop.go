package manager

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/mrlutik/kira2.0/internal/errors"
	"github.com/mrlutik/kira2.0/internal/logging"
)

var log = logging.Log

// MustStopSekaid is stopping sekaid process with StopProcessInsideContainer func (signal code - 15) and stopping sekaid container
func (s *SekaidManager) MustStopSekaid(ctx context.Context) {
	err := s.containerManager.StopProcessInsideContainer(ctx, "sekaid", 15, s.config.SekaidContainerName)
	errors.HandleFatalErr("Stopping sekaid bin in container", err)

	// TODO change method for dockerManager method instead of Cli
	log.Infof("Stopping <%s> container\n", s.config.SekaidContainerName)
	err = s.dockerManager.Cli.ContainerStop(ctx, s.config.SekaidContainerName, container.StopOptions{})
	errors.HandleFatalErr(fmt.Sprintf("cannot stop %s container", s.config.SekaidContainerName), err)
}

// MustStopInterx is stopping interx process with StopProcessInsideContainer func (signal code - 9) and stopping interx container
func (i *InterxManager) MustStopInterx(ctx context.Context) {
	err := i.containerManager.StopProcessInsideContainer(ctx, interxProcessName, 9, i.config.InterxContainerName)
	errors.HandleFatalErr("Stopping interx bin in container", err)

	// TODO change method for dockerManager method instead of Cli
	log.Infof("Stopping <%s> container\n", i.config.InterxContainerName)
	err = i.containerManager.Cli.ContainerStop(ctx, i.config.InterxContainerName, container.StopOptions{})
	errors.HandleFatalErr(fmt.Sprintf("cannot stop %s container", i.config.InterxContainerName), err)
}
