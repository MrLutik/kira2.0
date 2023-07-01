package manager

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/logging"
)

// InterxManager represents a manager for Interx container and its associated configurations.
type InterxManager struct {
	ContainerConfig        *container.Config
	SekaiHostConfig        *container.HostConfig
	SekaidNetworkingConfig *network.NetworkingConfig
	DockerClient           *docker.DockerManager
}

// Returns configured InterxManager.
// *docker.DockerManager: The poiner for docker.DockerManager instance.
// interxPort: endpoing of interx.
// imageName: The name of a image that interx container will be using.
// volumeName: The name of a volume that interx container will be using.
// dockerNetworkName: The name of a docker network that interx container will be using.
// containerName: The name of a container that interx will have.
func NewInterxManager(dockerClient *docker.DockerManager, interxPort, imageName, volumeName, dockerNetworkName, containerName string) (*InterxManager, error) {
	log := logging.Log
	log.Infof("Creating interx manager with port: %s, image: '%s', volume: '%s' in '%s' network\n", interxPort, imageName, volumeName, dockerNetworkName)

	natInterxPort, err := nat.NewPort("tcp", interxPort)
	if err != nil {
		log.Errorf("Creating NAT interx port error: %s\n", err)
		return nil, err
	}

	interxContainerConfig := &container.Config{
		Image:        imageName,
		Cmd:          []string{"/bin/bash"},
		Tty:          true,
		AttachStdin:  true,
		OpenStdin:    true,
		StdinOnce:    true,
		Hostname:     fmt.Sprintf("%s.local", containerName),
		ExposedPorts: nat.PortSet{natInterxPort: struct{}{}},
	}
	interxNetworkingConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			dockerNetworkName: {},
		},
	}
	interxHostConfig := &container.HostConfig{
		Binds: []string{
			volumeName,
		},
		PortBindings: nat.PortMap{
			natInterxPort: []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: interxPort}},
		},
		Privileged: true,
	}

	return &InterxManager{interxContainerConfig, interxHostConfig, interxNetworkingConfig, dockerClient}, err
}

// InitInterxBinInContainer sets up the 'interx' container with the specified configurations.
// ctx: The context for the operation.
// sekaidContainerName: The name of the 'sekaid' container.
// interxContainerName: The name of the 'interx' container.
// rpc_port: The RPC port for 'interx' to connect to 'sekaid'.
// grpc_port: The gRPC port for 'interx' to connect to 'sekaid'.
// Returns an error if any issue occurs during the init process.
func (i *InterxManager) InitInterxBinInContainer(ctx context.Context, sekaidContainerName, interxContainerName, rpc_port, grpc_port, interxHome string) error {
	log := logging.Log
	log.Infoln("Setting up 'interx' container")

	command := fmt.Sprintf(`interx init --rpc="http://%s:%s" --grpc="dns:///%s:%s" -home=%s`, sekaidContainerName, rpc_port, sekaidContainerName, grpc_port, interxHome)
	_, err := i.DockerClient.ExecCommandInContainer(ctx, interxContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}

	log.Infoln("interx inited")
	return err
}

// StartInterxBinInContainer starts interx binary inside interxContainerName
// ctx: The context for the operation.
// interxContainerName: The name of the 'interx' container.
// rpc_port: The RPC port for 'interx' to connect to 'sekaid'.
// grpc_port: The gRPC port for 'interx' to connect to 'sekaid'.
// interxHome: The home directory for 'interx'.
// Returns an error if any issue occurs during the start process.
func (i *InterxManager) StartInterxBinInContainer(ctx context.Context, interxContainerName, rpc_port, grpc_port, interxHome string) error {
	log := logging.Log
	command := fmt.Sprintf(`interx start -home=%s`, interxHome)
	_, err := i.DockerClient.ExecCommandInContainerInDetachMode(ctx, interxContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}
	log.Infoln("interx started")
	return nil
}

// Combine SetupInterxContainer and StartInterxBinInContainer together.
// First trying to run interx bin from previus state if exist.
// Then checking if interx bin running inside container.
// If no, initing new one.
// Then starting again.
// If no interx bin running inside container second time - return error.
// ctx: The context for the operation.
// sekaidContainerName: The name of the 'sekaid" container.
// interxContainerName: The name of the 'interx' container.
// rpc_port: The RPC port for 'interx' to connect to 'sekaid'.
// grpc_port: The gRPC port for 'interx' to connect to 'sekaid'.
// interxHome: The home directory for 'interx'.
// Returns an error if any issue occurs during the run process.
func (i *InterxManager) RunInterxContainer(ctx context.Context, sekaidContainerName, interxContainerName, rpc_port, grpc_port, interxHome string) error {
	log := logging.Log
	err := i.StartInterxBinInContainer(ctx, interxContainerName, rpc_port, grpc_port, interxHome)
	if err != nil {
		log.Errorf("Error while starting interex bin in '%s' container: %s\n", interxContainerName, err)
	}
	time.Sleep(time.Second * 1)
	check, _, err := i.DockerClient.CheckIfProcessIsRunningInContainer(ctx, "interx", interxContainerName)
	if err != nil {
		log.Errorf("Error while starting '%s' container: %s\n", interxContainerName, err)
		return err
	}
	if !check {
		log.Infof("Error starting interx binary first time in '%s' container, initing new instance\n", interxContainerName)
		err = i.InitInterxBinInContainer(ctx, sekaidContainerName, interxContainerName, rpc_port, grpc_port, interxHome)
		if err != nil {
			log.Errorf("Error while initing '%s' container: %s\n", interxContainerName, err)
			return err
		}
		err := i.StartInterxBinInContainer(ctx, interxContainerName, rpc_port, grpc_port, interxHome)
		if err != nil {
			log.Errorf("Error while running interx bin in '%s' container: %s\n", interxContainerName, err)
		}
		time.Sleep(time.Second * 1)
		check, _, err := i.DockerClient.CheckIfProcessIsRunningInContainer(ctx, "interx", interxContainerName)
		if err != nil {
			log.Errorf("Error while checking if procces is running in '%s' container: %s\n", interxContainerName, err)
			return err
		}
		if !check {
			log.Errorf("Error starting interex binary second time in '%s' container\n", interxContainerName)
			return fmt.Errorf("cannot start interex bin in %s container", interxContainerName)
		}
	}
	return nil
}
