package manager

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/mrlutik/kira2.0/internal/errors"
	"github.com/mrlutik/kira2.0/internal/logging"
)

var log = logging.Log

// MustStopSekaid is stoping sekaid process with StopProcessInsideContainer func (signal code - 15) and stoping sekaid container
func (s *SekaidManager) MustStopSekaid(ctx context.Context) {
	err := s.containerManager.StopProcessInsideContainer(ctx, "sekaid", 15, s.config.SekaidContainerName)
	errors.HandleFatalErr("Stoping sekaid bin in container", err)
	log.Printf("Stoping <%s> container\n", s.config.SekaidContainerName)
	err = s.dockerManager.Cli.ContainerStop(ctx, s.config.SekaidContainerName, container.StopOptions{})
	errors.HandleFatalErr(fmt.Sprintf("cannot stop %s container", s.config.SekaidContainerName), err)
}

// MustStopInterx is stoping interx process with StopProcessInsideContainer func (signal code - 9) and stoping interx container
func (i *InterxManager) MustStopInterx(ctx context.Context) {
	err := i.containerManager.StopProcessInsideContainer(ctx, interxProcessName, 9, i.config.InterxContainerName)
	errors.HandleFatalErr("Stoping interx bin in container", err)
	log.Printf("Stoping <%s> container\n", i.config.InterxContainerName)
	err = i.containerManager.Cli.ContainerStop(ctx, i.config.InterxContainerName, container.StopOptions{})
	errors.HandleFatalErr(fmt.Sprintf("cannot stop %s container", i.config.InterxContainerName), err)
}
