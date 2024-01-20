package docker

import (
	"context"
	"fmt"

	"github.com/mrlutik/kira2.0/internal/config"
)

func VerifyingDockerEnvironment(ctx context.Context, dockerManager *DockerManager, cfg *config.KiraConfig) error {
	err := dockerManager.VerifyDockerInstallation(ctx)
	if err != nil {
		return fmt.Errorf("docker is not available, error: %w", err)
	}

	dockerImage := fmt.Sprintf("%s:%s", cfg.DockerImageName, cfg.DockerImageVersion)
	err = dockerManager.PullImage(ctx, dockerImage)
	if err != nil {
		return fmt.Errorf("pulling image error: %w", err)
	}

	err = dockerManager.CheckOrCreateNetwork(ctx, cfg.DockerNetworkName)
	if err != nil {
		return fmt.Errorf("docker networking setup error: %w", err)
	}

	return nil
}
