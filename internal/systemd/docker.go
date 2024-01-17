package systemd

import (
	"context"
	"fmt"
	"time"

	"github.com/mrlutik/kira2.0/internal/logging"
)

const (
	allTimeForDockerServiceOperation = 10 * time.Second
	waitingForServiceStatusTime      = 3 * time.Second
	dockerServiceName                = "docker.service"
)

func DockerServiceManagement(logger *logging.Logger) error {
	dockerServiceContext, cancel := context.WithTimeout(context.Background(), allTimeForDockerServiceOperation)
	defer cancel()

	dockerServiceManager, err := NewServiceManager(dockerServiceContext, logger, dockerServiceName, "replace")
	if err != nil {
		return fmt.Errorf("can't create instance of service manager, error: %w", err)
	}
	defer dockerServiceManager.Close()

	exists, err := dockerServiceManager.CheckServiceExists(dockerServiceContext)
	if err != nil {
		return fmt.Errorf("can't reach the service, error: %w", err)
	}
	if !exists {
		logger.Fatalf("'%s' is not available", dockerServiceName)
	}

	status, err := dockerServiceManager.GetServiceStatus(dockerServiceContext)
	if err != nil {
		return fmt.Errorf("can't get the '%s' status, error: %w", dockerServiceName, err)
	}
	if status != "active" {
		logger.Errorf("'%s' is not active", dockerServiceName)
		logger.Infof("Trying to restart '%s'", dockerServiceName)
		err = dockerServiceManager.RestartService(dockerServiceContext)
		if err != nil {
			return fmt.Errorf("can't restart '%s', error: %w", dockerServiceName, err)
		}
	}

	err = dockerServiceManager.WaitForServiceStatus(dockerServiceContext, "active", waitingForServiceStatusTime)
	if err != nil {
		return fmt.Errorf("waiting for status, error: %w", err)
	}

	return nil
}
