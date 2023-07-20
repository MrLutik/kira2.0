package manager

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/logging"
)

// InterxManager represents a manager for Interx container and its associated configurations.
type InterxManager struct {
	ContainerConfig     *container.Config
	InterxHostConfig    *container.HostConfig
	InterxNetworkConfig *network.NetworkingConfig
	containerManager    *docker.ContainerManager
	config              *config.KiraConfig
}

// NewInterxManager returns configured InterxManager.
func NewInterxManager(containerManager *docker.ContainerManager, config *config.KiraConfig) (*InterxManager, error) {
	log := logging.Log
	log.Infof("Creating interx manager with port: %s, image: '%s', volume: '%s' in '%s' network",
		config.InterxPort, config.DockerImageName, config.VolumeName, config.DockerNetworkName)

	natInterxPort, err := nat.NewPort("tcp", config.InterxPort)
	if err != nil {
		log.Errorf("Creating NAT interx port error: %s", err)
		return nil, err
	}

	interxContainerConfig := &container.Config{
		Image:        fmt.Sprintf("%s:%s", config.DockerImageName, config.DockerImageVersion),
		Cmd:          []string{"/bin/bash"},
		Tty:          true,
		AttachStdin:  true,
		OpenStdin:    true,
		StdinOnce:    true,
		Hostname:     fmt.Sprintf("%s.local", config.InterxContainerName),
		ExposedPorts: nat.PortSet{natInterxPort: struct{}{}},
	}
	interxNetworkingConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			config.DockerNetworkName: {},
		},
	}
	interxHostConfig := &container.HostConfig{
		Binds: []string{
			config.VolumeName,
		},
		PortBindings: nat.PortMap{
			natInterxPort: []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: config.InterxPort}},
		},
		Privileged: true,
	}

	return &InterxManager{
		ContainerConfig:     interxContainerConfig,
		InterxHostConfig:    interxHostConfig,
		InterxNetworkConfig: interxNetworkingConfig,
		containerManager:    containerManager,
		config:              config,
	}, err
}

// initInterxBinInContainer sets up the 'interx' container with the specified configurations.
// Returns an error if any issue occurs during the init process.
func (i *InterxManager) initInterxBinInContainer(ctx context.Context) error {
	log := logging.Log
	log.Infof("Setting up '%s' (interx) container", i.config.InterxContainerName)

	command := fmt.Sprintf(`interx init --rpc="http://%s:%s" --grpc="dns:///%s:%s" -home=%s`,
		i.config.SekaidContainerName, i.config.RpcPort, i.config.SekaidContainerName, i.config.GrpcPort, i.config.InterxHome)
	_, err := i.containerManager.ExecCommandInContainer(ctx, i.config.InterxContainerName, []string{"bash", "-c", command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s", command, err)
		return err
	}

	log.Infoln("'interx' is initialized")
	return err
}

// startInterxBinInContainer starts interx binary inside InterxContainerName
// Returns an error if any issue occurs during the start process.
func (i *InterxManager) startInterxBinInContainer(ctx context.Context) error {
	log := logging.Log
	command := fmt.Sprintf("interx start -home=%s", i.config.InterxHome)
	_, err := i.containerManager.ExecCommandInContainerInDetachMode(ctx, i.config.InterxContainerName, []string{"bash", "-c", command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s", command, err)
		return err
	}
	log.Infoln("'interx' started")
	return nil
}

// runInterxContainer starts the 'interx' container and checks if the process is running.
// If the 'interx' process is not running, it initializes the 'interx' binary in the container
// and starts it again. It checks if the process is running after the initialization.
// The method waits for a specified duration before checking if the process is running.
// If any errors occur during the process, an error is returned.
func (i *InterxManager) runInterxContainer(ctx context.Context) error {
	log := logging.Log
	const delay = time.Second

	err := i.startInterxBinInContainer(ctx)
	if err != nil {
		log.Errorf("Starting 'interx' bin in '%s' container error: %s", i.config.InterxContainerName, err)
		return err
	}

	log.Warningf("Waiting for %0.0f seconds for process", delay.Seconds())
	time.Sleep(time.Second * 1)

	check, _, err := i.containerManager.CheckIfProcessIsRunningInContainer(ctx, "interx", i.config.InterxContainerName)
	if err != nil {
		log.Errorf("Starting '%s' container error: %s", i.config.InterxContainerName, err)
		return err
	}

	if !check {
		log.Warningf("Error starting 'interx' binary first time in '%s' container, initialization new instance", i.config.InterxContainerName)
		err = i.initInterxBinInContainer(ctx)
		if err != nil {
			log.Errorf("Initialization '%s' in container error: %s", i.config.InterxContainerName, err)
			return err
		}

		err := i.startInterxBinInContainer(ctx)
		if err != nil {
			log.Errorf("Running 'interx' bin in '%s' container error: %s", i.config.InterxContainerName, err)
			return fmt.Errorf("running 'interx' bin in '%s' container error: %w", i.config.InterxContainerName, err)
		}

		log.Warningf("Waiting for %0.0f seconds for process", delay.Seconds())
		time.Sleep(delay)

		check, _, err := i.containerManager.CheckIfProcessIsRunningInContainer(ctx, "interx", i.config.InterxContainerName)
		if err != nil {
			log.Errorf("Checking 'interx' process in '%s' container error: %s", i.config.InterxContainerName, err)
			return fmt.Errorf("checking 'interx' process in '%s' container error: %w", i.config.InterxContainerName, err)
		}
		if !check {
			log.Errorf("Error starting 'interx' binary second time in '%s' container", i.config.InterxContainerName)
			return fmt.Errorf("cannot start 'interx' bin 'in' %s container", i.config.InterxContainerName)
		}
	}

	return nil
}
