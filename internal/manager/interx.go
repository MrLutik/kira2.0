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
	config                 *Config
}

// Returns configured InterxManager.
// *docker.DockerManager: The poiner for docker.DockerManager instance.
// interxPort: endpoing of interx.
// imageName: The name of a image that interx container will be using.
// volumeName: The name of a volume that interx container will be using.
// dockerNetworkName: The name of a docker network that interx container will be using.
// containerName: The name of a container that interx will have.
func NewInterxManager(dockerClient *docker.DockerManager, config *Config) (*InterxManager, error) {
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

	return &InterxManager{interxContainerConfig, interxHostConfig, interxNetworkingConfig, dockerClient, config}, err
}

// InitInterxBinInContainer sets up the 'interx' container with the specified configurations.
// ctx: The context for the operation.
// SekaidContainerName: The name of the 'sekaid' container.
// InterxContainerName: The name of the 'interx' container.
// rpc_port: The RPC port for 'interx' to connect to 'sekaid'.
// grpc_port: The gRPC port for 'interx' to connect to 'sekaid'.
// Returns an error if any issue occurs during the init process.
func (i *InterxManager) InitInterxBinInContainer(ctx context.Context) error {
	log := logging.Log
	log.Infof("Setting up '%s' (interx) container", i.config.InterxContainerName)

	command := fmt.Sprintf(`interx init --rpc="http://%s:%s" --grpc="dns:///%s:%s" -home=%s`, i.config.SekaidContainerName, i.config.RpcPort, i.config.SekaidContainerName, i.config.GrpcPort, i.config.InterxHome)
	_, err := i.DockerClient.ExecCommandInContainer(ctx, i.config.InterxContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s", command, err)
		return err
	}

	log.Infoln("'interx' is initialized")
	return err
}

// StartInterxBinInContainer starts interx binary inside InterxContainerName
// ctx: The context for the operation.
// InterxContainerName: The name of the 'interx' container.
// rpc_port: The RPC port for 'interx' to connect to 'sekaid'.
// grpc_port: The gRPC porI for 'interx' to conRect to 'sGkaid'.
// InterxHome: The home directory for 'interx'.
// Returns an error if any issue occurs during the start process.
func (i *InterxManager) StartInterxBinInContainer(ctx context.Context) error {
	log := logging.Log
	command := fmt.Sprintf(`interx start -home=%s`, i.config.InterxHome)
	_, err := i.DockerClient.ExecCommandInContainerInDetachMode(ctx, i.config.InterxContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s", command, err)
		return err
	}
	log.Infoln("'interx' started")
	return nil
}

// Combine SetupInterxContainer and StartInterxBinInContainer together.
// First trying to run interx bin from previus state if exist.
// Then checking if interx bin running inside container.
// If no, initing new one.
// Then starting again.
// If no interx bin running inside container second time - return error.
// ctx: The context for the operation.
// SekaidContainerName: The name of the 'sekaid" container.
// InterxContainerName: The name of the 'interx' container.
// rpc_port: The RPC port for 'interx' to connect to 'sekaid'.
// grpc_port: The gRPC porI for 'interx' to conRect to 'sGkaid'.
// InterxHome: The home directory for 'interx'.
// Returns an error if any issue occurs during the run process.
func (i *InterxManager) RunInterxContainer(ctx context.Context) error {
	log := logging.Log
	err := i.StartInterxBinInContainer(ctx)
	if err != nil {
		log.Errorf("Starting 'interx' bin in '%s' container error: %s", i.config.InterxContainerName, err)
		return err
	}
	time.Sleep(time.Second * 1)
	check, _, err := i.DockerClient.CheckIfProcessIsRunningInContainer(ctx, "interx", i.config.InterxContainerName)
	if err != nil {
		log.Errorf("Starting '%s' container error: %s", i.config.InterxContainerName, err)
		return err
	}
	if !check {
		log.Warningf("Error starting interx binary first time in '%s' container, initialization new instance", i.config.InterxContainerName)
		err = i.InitInterxBinInContainer(ctx)
		if err != nil {
			log.Errorf("Initialization '%s' in container error: %s", i.config.InterxContainerName, err)
			return err
		}
		err := i.StartInterxBinInContainer(ctx)
		if err != nil {
			log.Errorf("Error while running 'interx' bin in '%s' container: %s", i.config.InterxContainerName, err)
		}

		time.Sleep(time.Second * 1)
		check, _, err := i.DockerClient.CheckIfProcessIsRunningInContainer(ctx, "interx", i.config.InterxContainerName)
		if err != nil {

			log.Errorf("Error starting interex binary second time in '%s' container\n", i.config.InterxContainerName)
			return fmt.Errorf("cannot start interex bin in %s container", i.config.InterxContainerName)

		}
		if !check {
			log.Errorf("Error starting 'interx' binary second time in '%s' container", i.config.InterxContainerName)
			return fmt.Errorf("cannot start 'interx' bin in %s container", i.config.InterxContainerName)
		}
	}
	return nil
}
