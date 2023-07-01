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

// SekaidManager represents a manager for Sekaid container and its associated configurations.
type SekaidManager struct {
	ContainerConfig        *container.Config
	SekaiHostConfig        *container.HostConfig
	SekaidNetworkingConfig *network.NetworkingConfig
	DockerManager          *docker.DockerManager
}

// Returns configured SekaidManager.
// *docker.DockerManager: The poiner for docker.DockerManager instance.
// grpcPort: The grpc port for sekaid.
// rpcPort: The rpc port for sekaid.
// imageName: The name of a image that sekaid container will be using.
// volumeName: The name of a volume that sekaid container will be using.
// dockerNetworkName: The name of a docker network that sekaid container will be using.
// containerName: The name of a container that sekaid will have.
func NewSekaidManager(dockerManager *docker.DockerManager, grpcPort, rpcPort, imageName, volumeName, dockerNetworkName, containerName string) (*SekaidManager, error) {
	log := logging.Log
	log.Infof("Creating sekaid manager with ports: %s, %s, image: '%s', volume: '%s' in '%s' network\n", grpcPort, rpcPort, imageName, volumeName, dockerNetworkName)

	natGrcpPort, err := nat.NewPort("tcp", grpcPort)
	if err != nil {
		log.Errorf("Creating NAT GRPC port error: %s\n", err)
		return nil, err
	}

	natRpcPort, err := nat.NewPort("tcp", rpcPort)
	if err != nil {
		log.Errorf("Creating NAT RPC port error: %s\n", err)
		return nil, err
	}

	sekaiContainerConfig := &container.Config{
		Image:       imageName,
		Cmd:         []string{"/bin/bash"},
		Tty:         true,
		AttachStdin: true,
		OpenStdin:   true,
		StdinOnce:   true,
		Hostname:    fmt.Sprintf("%s.local", containerName),
		ExposedPorts: nat.PortSet{
			natGrcpPort: struct{}{},
			natRpcPort:  struct{}{},
		},
	}

	sekaiHostConfig := &container.HostConfig{
		Binds: []string{
			volumeName,
		},
		PortBindings: nat.PortMap{
			natGrcpPort: []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: grpcPort}},
			natRpcPort:  []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: rpcPort}},
		},
		Privileged: true,
	}

	sekaidNetworkingConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			dockerNetworkName: {},
		},
	}

	return &SekaidManager{sekaiContainerConfig, sekaiHostConfig, sekaidNetworkingConfig, dockerManager}, err
}

// InitSekaidBinInContainer sets up the 'sekaid' container with the specified configurations.
// ctx: The context for the operation.
// moniker: The moniker for the 'sekaid' container.
// sekaidContainerName: The name of the 'sekaid' container.
// sekaidNetworkName: The name of the network associated with the 'sekaid' container.
// sekaidHome: The home directory for 'sekaid'.
// keyringBackend: The keyring backend to use.
// rcpPort: The RPC port for 'sekaid'.
// mnemonicDir: The directory to store the generated mnemonics.
// Returns an error if any issue occurs during the init process.
func (s *SekaidManager) InitSekaidBinInContainer(ctx context.Context, moniker, sekaidContainerName, sekaidNetworkName, sekaidHome, keyringBackend, rcpPort, mnemonicDir string) error {
	log := logging.Log
	log.Infoln("Setting up 'sekaid' container")

	command := fmt.Sprintf(`sekaid init  --overwrite --chain-id=%s --home=%s "%s"`, sekaidNetworkName, sekaidHome, moniker)
	_, err := s.DockerManager.ExecCommandInContainer(ctx, sekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}

	command = fmt.Sprintf(`mkdir %s`, mnemonicDir)
	_, err = s.DockerManager.ExecCommandInContainer(ctx, sekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}

	command = fmt.Sprintf(`sekaid keys add "validator" --keyring-backend=%s --home=%s --output=json | jq .mnemonic > %s/sekai.mnemonic`, keyringBackend, sekaidHome, mnemonicDir)
	_, err = s.DockerManager.ExecCommandInContainer(ctx, sekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}

	command = fmt.Sprintf(`sekaid keys add "faucet" --keyring-backend=%s --home=%s --output=json | jq .mnemonic > %s/faucet.mnemonic`, keyringBackend, sekaidHome, mnemonicDir)
	_, err = s.DockerManager.ExecCommandInContainer(ctx, sekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}

	command = fmt.Sprintf(`sekaid add-genesis-account validator 150000000000000ukex,300000000000000test,2000000000000000000000000000samolean,1000000lol --keyring-backend=%v --home=%v`, keyringBackend, sekaidHome)
	_, err = s.DockerManager.ExecCommandInContainer(ctx, sekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}

	command = fmt.Sprintf(`sekaid gentx-claim validator --keyring-backend=%s --moniker="%s" --home=%s`, keyringBackend, moniker, sekaidHome)
	_, err = s.DockerManager.ExecCommandInContainer(ctx, sekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
		return err
	}

	log.Infoln("'sekaid' container started")
	return nil
}

// StartSekaidBinInContainer starts sekaid binary inside sekaidContainerName var
// ctx: The context for the operation.
// moniker: The moniker for the 'sekaid' container.
// sekaidContainerName: The name of the 'sekaid' container.
// sekaidNetworkName: The name of the network associated with the 'sekaid' container.
// sekaidHome: The home directory for 'sekaid'.
// keyringBackend: The keyring backend to use.
// rcpPort: The RPC port for 'sekaid'.
// mnemonicDir: The directory to store the generated mnemonics.
// Returns an error if any issue occurs during the start process.
func (s *SekaidManager) StartSekaidBinInContainer(ctx context.Context, moniker, sekaidContainerName, sekaidNetworkName, sekaidHome, keyringBackend, rcpPort, mnemonicDir string) error {
	log := logging.Log
	log.Infoln("Starting 'sekaid' container")
	command := fmt.Sprintf(`sekaid start --rpc.laddr "tcp://0.0.0.0:%s" --home=%s`, rcpPort, sekaidHome)
	_, err := s.DockerManager.ExecCommandInContainerInDetachMode(ctx, sekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s\n", command, err)
	}

	return nil
}

// Combine SetupSekaidBinInContainer and StartSekaidBinInContainer together.
// First trying to run sekaid bin from previus state if exist.
// Then checking if sekaid bin running inside container.
// If no initing new one.
// Then starting again.
// If no sekaid bin running inside container second time - return error.
// ctx: The context for the operation.
// moniker: The moniker for the 'sekaid' container.
// sekaidContainerName: The name of the 'sekaid' container.
// sekaidNetworkName: The name of the network associated with the 'sekaid' container.
// sekaidHome: The home directory for 'sekaid'.
// keyringBackend: The keyring backend to use.
// rcpPort: The RPC port for 'sekaid'.
// mnemonicDir: The directory to store the generated mnemonics.
// Returns an error if any issue occurs during the run process.
func (s *SekaidManager) RunSekaidContainer(ctx context.Context, moniker, sekaidContainerName, sekaidNetworkName, sekaidHome, keyringBackend, rcpPort, mnemonicDir string) error {
	log := logging.Log
	err := s.StartSekaidBinInContainer(ctx, moniker, sekaidContainerName, sekaidNetworkName, sekaidHome, keyringBackend, rcpPort, mnemonicDir)
	if err != nil {
		log.Errorf("Cannot start sekaid bin in %s container", sekaidContainerName)
	}
	time.Sleep(time.Second * 1)
	check, _, err := s.DockerManager.CheckIfProcessIsRunningInContainer(ctx, "sekaid", sekaidContainerName)
	if err != nil {
		log.Infof("Error while setup '%s' container: %s\n", sekaidContainerName, err)
		return err
	}
	if !check {
		log.Infof("Error starting sekaid binary first time in '%s' container, initing new instance\n", sekaidContainerName)
		err = s.InitSekaidBinInContainer(ctx, moniker, sekaidContainerName, sekaidNetworkName, sekaidHome, keyringBackend, rcpPort, mnemonicDir)
		if err != nil {
			log.Errorf("Error while setup '%s' container: %s\n", sekaidContainerName, err)
			return err
		}
		err := s.StartSekaidBinInContainer(ctx, moniker, sekaidContainerName, sekaidNetworkName, sekaidHome, keyringBackend, rcpPort, mnemonicDir)
		if err != nil {
			log.Errorf("Cannot start sekaid bin in %s container", sekaidContainerName)
		}

		time.Sleep(time.Second * 1)

		check, _, err = s.DockerManager.CheckIfProcessIsRunningInContainer(ctx, "sekaid", sekaidContainerName)
		if err != nil {
			log.Errorf("Error while setup '%s' container: %s\n", sekaidContainerName, err)
			return err
		}
		if !check {
			log.Errorf("Error starting sekaid bin second time in '%s' container\n", sekaidContainerName)
			return fmt.Errorf("coudnt start sekaid bin second time")
		}
	}
	return nil
}
