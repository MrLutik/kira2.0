package manager

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/errors"
	"github.com/mrlutik/kira2.0/internal/logging"
)

var log = logging.Log

func (s *SekaidManager) MustStopSekaid(ctx context.Context) {
	err := stopProcessInsideContainer(ctx, "sekaid", 15, s.config.SekaidContainerName, s.containerManager)
	errors.HandleFatalErr("Stoping sekaid bin in container", err)
	log.Printf("Stoping <%s> container\n", s.config.SekaidContainerName)
	err = s.dockerManager.Cli.ContainerStop(ctx, s.config.SekaidContainerName, container.StopOptions{})
	errors.HandleFatalErr(fmt.Sprintf("cannot stop %s container", s.config.SekaidContainerName), err)
}

func (i *InterxManager) MustStopInterx(ctx context.Context) {
	err := stopProcessInsideContainer(ctx, interxProcessName, 9, i.config.InterxContainerName, i.containerManager)
	errors.HandleFatalErr("Stoping interx bin in container", err)
	log.Printf("Stoping <%s> container\n", i.config.InterxContainerName)
	err = i.containerManager.Cli.ContainerStop(ctx, i.config.InterxContainerName, container.StopOptions{})
	errors.HandleFatalErr(fmt.Sprintf("cannot stop %s container", i.config.InterxContainerName), err)
}

func stopProcessInsideContainer(ctx context.Context, processName string, codeTopStopWIth int, containerName string, containerManager *docker.ContainerManager) error {
	log.Printf("Checking if %s is running inside container", processName)
	check, _, err := containerManager.CheckIfProcessIsRunningInContainer(ctx, processName, containerName)
	if err != nil {
		return fmt.Errorf("cant check if procces is running inside container, %s", err)
	}
	if !check {
		log.Warnf("process <%s> is not running inside <%s> container\n", processName, containerName)
		return nil
	}
	log.Printf("Stoping <%s> proccess\n", processName)
	out, err := containerManager.ExecCommandInContainer(ctx, containerName, []string{"pkill", fmt.Sprintf("-%v", codeTopStopWIth), processName})
	if err != nil {
		log.Errorf("cannot kill <%s> process inside <%s> container\nout: %s\nerr: %v\n", processName, containerName, string(out), err)
		return fmt.Errorf("cannot kill <%s> process inside <%s> container\nout: %s\nerr: %s", processName, containerName, string(out), err)
	}

	check, _, err = containerManager.CheckIfProcessIsRunningInContainer(ctx, processName, containerName)
	if err != nil {
		return fmt.Errorf("cant check if procces is running inside container, %s", err)
	}
	if check {
		log.Errorf("Process <%s> is still running inside <%s> container\n", processName, containerName)
		return err
	}
	log.Printf("<%s> proccess was successfully stoped\n", processName)
	return nil
}
