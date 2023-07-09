package systemd

import (
	"context"
	"fmt"
	"time"

	"github.com/mrlutik/kira2.0/internal/errors"
)

const (
	allTimeForDockerServiceOperation = 10 * time.Second
	waitingForServiceStatusTime      = 3 * time.Second
	dockerServiceName                = "docker.service"
)

func DockerServiceManagement() {
	dockerServiceManager, err := NewServiceManager(context.Background(), dockerServiceName, "replace")
	errors.HandleErr("Can't create instance of service manager", err)

	dockerServiceContext, cancel := context.WithTimeout(context.Background(), allTimeForDockerServiceOperation)
	defer cancel()

	exists, err := dockerServiceManager.CheckServiceExists(dockerServiceContext)
	errors.HandleErr("Can't reach the service", err)
	if !exists {
		log.Fatalf("'%s' is not available", dockerServiceName)
	}

	status, err := dockerServiceManager.GetServiceStatus(dockerServiceContext)
	errors.HandleErr(fmt.Sprintf("Can't get the '%s' status", dockerServiceName), err)
	if status != "active" {
		log.Errorf("'%s' is not active", dockerServiceName)
		log.Infof("Trying to restart it")
		err = dockerServiceManager.RestartService(dockerServiceContext)
		errors.HandleErr(fmt.Sprintf("Can't restart '%s'", dockerServiceName), err)
	}

	err = dockerServiceManager.WaitForServiceStatus(dockerServiceContext, "active", waitingForServiceStatusTime)
	errors.HandleErr("Waiting for status", err)
}
