package manager

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/mrlutik/kira2.0/internal/errors"
	"github.com/mrlutik/kira2.0/internal/logging"
)

var log = logging.Log

func (s *SekaidManager) MustStopSekaid(ctx context.Context) {
	p := "sekaid"
	code := "15"
	log.Printf("Stoping <%s> proccess\n", p)
	out, err := s.containerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{"pkill", "-" + code, p})
	if err != nil {
		errors.HandleFatalErr(fmt.Sprintf("cannot kill <%s> process inside  <%s> container\nout: %s", p, s.config.SekaidContainerName, string(out)), err)
	}
	log.Printf("<%s> proccess was successfully stoped\n", p)
	log.Printf("Stoping <%s> container\n", s.config.SekaidContainerName)
	err = s.dockerManager.Cli.ContainerStop(ctx, s.config.SekaidContainerName, container.StopOptions{})
	errors.HandleFatalErr("cannot stop container", err)
}

func (i *InterxManager) MustStopInterx(ctx context.Context) {
	p := "interx"
	code := "9"
	log.Printf("Stoping <%s> proccess\n", p)
	out, err := i.containerManager.ExecCommandInContainer(ctx, i.config.InterxContainerName, []string{"pkill", "-" + code, p})
	if err != nil {
		errors.HandleFatalErr(fmt.Sprintf("cannot kill <%s> process inside <%s> container\nout: %s", p, i.config.InterxContainerName, string(out)), err)
	}
	log.Printf("<%s> proccess was successfully stoped\n", p)
	log.Printf("Stoping <%s> container\n", i.config.InterxContainerName)
	err = i.containerManager.Cli.ContainerStop(ctx, i.config.InterxContainerName, container.StopOptions{})
	errors.HandleFatalErr("cannot stop container", err)
}
