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
	"github.com/mrlutik/kira2.0/internal/manager/utils"
	"github.com/mrlutik/kira2.0/internal/types"
)

// SekaidManager represents a manager for Sekaid container and its associated configurations.
type SekaidManager struct {
	ContainerConfig        *container.Config
	SekaiHostConfig        *container.HostConfig
	SekaidNetworkingConfig *network.NetworkingConfig
	dockerManager          *docker.DockerManager
	config                 *config.KiraConfig
	helper                 *utils.HelperManager
}

const (
	validatorAccountName = "validator"
)

func NewSekaiInterface(sekaidManager *SekaidManager) Repository {
	return sekaidManager
}

// Returns configured SekaidManager.
//
//	*docker.DockerManager // The pointer for docker.DockerManager instance.
//	*config	// Pointer to config struct, can create new instance by calling NewConfig() function
func NewSekaidManager(dockerManager *docker.DockerManager, config *config.KiraConfig) (*SekaidManager, error) {
	log := logging.Log
	log.Infof("Creating sekaid manager with ports: %s, %s, image: '%s', volume: '%s' in '%s' network\n",
		config.GrpcPort, config.RpcPort, config.DockerImageName, config.VolumeName, config.DockerNetworkName)

	natGrpcPort, err := nat.NewPort("tcp", config.GrpcPort)
	if err != nil {
		log.Errorf("Creating NAT GRPC port error: %s", err)
		return nil, err
	}

	natRpcPort, err := nat.NewPort("tcp", config.RpcPort)
	if err != nil {
		log.Errorf("Creating NAT RPC port error: %s", err)
		return nil, err
	}

	sekaiContainerConfig := &container.Config{
		Image:       fmt.Sprintf("%s:%s", config.DockerImageName, config.DockerImageVersion),
		Cmd:         []string{"/bin/bash"},
		Tty:         true,
		AttachStdin: true,
		OpenStdin:   true,
		StdinOnce:   true,
		Hostname:    fmt.Sprintf("%s.local", config.SekaidContainerName),
		ExposedPorts: nat.PortSet{
			natGrpcPort: struct{}{},
			natRpcPort:  struct{}{},
		},
	}

	sekaiHostConfig := &container.HostConfig{
		Binds: []string{
			config.VolumeName,
		},
		PortBindings: nat.PortMap{
			natGrpcPort: []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: config.GrpcPort}},
			natRpcPort:  []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: config.RpcPort}},
		},
		Privileged: true,
	}

	sekaidNetworkingConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			config.DockerNetworkName: {},
		},
	}
	helper := utils.NewHelperManager(dockerManager, config)
	return &SekaidManager{
		sekaiContainerConfig,
		sekaiHostConfig,
		sekaidNetworkingConfig,
		dockerManager,
		config,
		helper,
	}, err
}

// InitSekaidBinInContainer sets up the 'sekaid' container with the specified configurations.
// ctx: The context for the operation.
// Moniker: The Moniker for the 'sekaid' container.
// SekaidContainerName: The name of the 'sekaid' container.
// sekaidNetworkName: The name of the network associated with the 'sekaid' container.
// SekaidHome: The home directory for 'sekaid'.
// KeyringBackend: The keyring backend to use.
// RpcPort: The RPC port for 'sekaid'.
// MnemonicDir: The directory to store the generated mnemonics.
// Returns an error if any issue occurs during the init process.
func (s *SekaidManager) InitSekaidBinInContainer(ctx context.Context) error {
	log := logging.Log
	log.Infof("Setting up '%s' (sekaid) container", s.config.SekaidContainerName)

	commands := []string{
		fmt.Sprintf(`sekaid init  --overwrite --chain-id=%s --home=%s "%s"`,
			s.config.NetworkName, s.config.SekaidHome, s.config.Moniker),
		fmt.Sprintf(`mkdir %s`, s.config.MnemonicDir),
		fmt.Sprintf(`sekaid keys add "%s" --keyring-backend=%s --home=%s --output=json | jq .mnemonic > %s/sekai.mnemonic`,
			validatorAccountName, s.config.KeyringBackend, s.config.SekaidHome, s.config.MnemonicDir),
		fmt.Sprintf(`sekaid keys add "faucet" --keyring-backend=%s --home=%s --output=json | jq .mnemonic > %s/faucet.mnemonic`,
			s.config.KeyringBackend, s.config.SekaidHome, s.config.MnemonicDir),
		fmt.Sprintf(`sekaid add-genesis-account %s 150000000000000ukex,300000000000000test,2000000000000000000000000000samolean,1000000lol --keyring-backend=%v --home=%v`,
			validatorAccountName, s.config.KeyringBackend, s.config.SekaidHome),
		fmt.Sprintf(`sekaid gentx-claim %s --keyring-backend=%s --moniker="%s" --home=%s`,
			validatorAccountName, s.config.KeyringBackend, s.config.Moniker, s.config.SekaidHome),
	}

	for _, command := range commands {
		_, err := s.dockerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{`bash`, `-c`, command})
		if err != nil {
			log.Errorf("Command '%s' execution error: %s", command, err)
			return err
		}
	}

	log.Infoln("'sekaid' container started")
	return nil
}

// StartSekaidBinInContainer starts sekaid binary inside sekaid container name
// Returns an error if any issue occurs during the start process.
func (s *SekaidManager) StartSekaidBinInContainer(ctx context.Context) error {
	log := logging.Log
	log.Infoln("Starting 'sekaid' container")
	command := fmt.Sprintf(`sekaid start --rpc.laddr "tcp://0.0.0.0:%s" --home=%s`, s.config.RpcPort, s.config.SekaidHome)
	_, err := s.dockerManager.ExecCommandInContainerInDetachMode(ctx, s.config.SekaidContainerName, []string{`bash`, `-c`, command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s", command, err)
	}

	return nil
}

// Combine SetupSekaidBinInContainer and StartSekaidBinInContainer together.
// First trying to run sekaid bin from previous state if exist.
// Then checking if sekaid bin running inside container.
// If not initialized new one, then starting again.
// If no sekaid bin running inside container second time - return error.
// Then starting propagating transactions for permissions as in sekai-env.sh
// Returns an error if any issue occurs during the run process.
func (s *SekaidManager) RunSekaidContainer(ctx context.Context) error {
	log := logging.Log

	if err := s.StartSekaidBinInContainer(ctx); err != nil {
		log.Errorf("Cannot start 'sekaid' bin in '%s' container, error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("cannot start 'sekaid' bin in '%s' container, error: %w", s.config.SekaidContainerName, err)
	}

	const delay = time.Second * 1
	log.Warningf("Waiting to start 'sekaid' for %0.0f seconds", delay.Seconds())
	time.Sleep(delay)

	check, _, err := s.dockerManager.CheckIfProcessIsRunningInContainer(ctx, "sekaid", s.config.SekaidContainerName)
	if err != nil {
		log.Errorf("Setup '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("setup '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	if !check {
		if err := s.initializeSekaid(ctx); err != nil {
			return err
		}
	}

	log.Printf("SEKAID CONTAINER STARTED")
	return nil
}

// initializeSekaid initializes the sekaid node by performing several setup steps.
// It starts the sekaid binary for the first time in the specified container,
// initializes a new instance, and waits for it to start.
// If any errors occur during the initialization process, an error is returned.
// The method also propagates genesis proposals and updates the identity registrar from the validator account.
func (s *SekaidManager) initializeSekaid(ctx context.Context) error {
	log := logging.Log

	log.Warningf("Starting sekaid binary first time in '%s' container, initializing new instance", s.config.SekaidContainerName)

	if err := s.InitSekaidBinInContainer(ctx); err != nil {
		log.Errorf("Setup '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("setup '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	if err := s.StartSekaidBinInContainer(ctx); err != nil {
		log.Errorf("Starting 'sekaid' bin in '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("starting 'sekaid' bin in '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	const delay = time.Second * 1
	log.Warningf("Waiting to start 'sekaid' for %0.0f seconds", delay.Seconds())
	time.Sleep(delay)

	check, _, err := s.dockerManager.CheckIfProcessIsRunningInContainer(ctx, "sekaid", s.config.SekaidContainerName)
	if err != nil || !check {
		log.Errorf("Starting 'sekaid' bin second time in '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("starting 'sekaid' bin second time in '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	err = s.PostGenesisProposals(ctx)
	if err != nil {
		log.Errorf("propagating transaction error: %s", err)
		return fmt.Errorf("propagating transaction error: %w", err)
	}

	err = s.helper.UpdateIdentityRegistrarFromValidator(ctx, validatorAccountName)
	if err != nil {
		log.Errorf("updating identity registrar error: %s", err)
		return fmt.Errorf("updating identity registrar error: %w", err)
	}

	return nil
}

// Post genesis proposals after launching new network from KM1 await-validator-init.sh file.
// Adding required permissions for validator.
// First getting validator address with GetAddressByName.
// Then in loop calling GivePermissionsToAddress func with delay between calls 10 sec because tx can be propagated once per 10 sec
func (s *SekaidManager) PostGenesisProposals(ctx context.Context) error {
	log := logging.Log

	address, err := s.helper.GetAddressByName(ctx, validatorAccountName)
	if err != nil {
		log.Errorf("Getting address in '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("getting address in '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	permissions := []int{
		types.PermWhitelistAccountPermissionProposal,
		types.PermRemoveWhitelistedAccountPermissionProposal,
		types.PermCreateUpsertTokenAliasProposal,
		types.PermCreateSoftwareUpgradeProposal,
		types.PermVoteWhitelistAccountPermissionProposal,
		types.PermVoteRemoveWhitelistedAccountPermissionProposal,
		types.PermVoteUpsertTokenAliasProposal,
		types.PermVoteSoftwareUpgradeProposal,
	}
	log.Printf("Permissions to add: '%d' for: '%s'", permissions, address)

	// waiting 10 sec to first block to be propagated
	log.Infof("Waiting for %0.0f seconds before first block be propagated", time.Duration.Seconds(s.config.TimeBetweenBlocks))
	time.Sleep(s.config.TimeBetweenBlocks)

	for _, perm := range permissions {
		log.Printf("Adding permission: '%d'", perm)

		err = s.helper.GivePermissionToAddress(ctx, perm, address)
		if err != nil {
			log.Errorf("Giving permission '%d' error: %s", perm, err)
			return fmt.Errorf("giving permission '%d' error: %w", perm, err)
		}

		log.Printf("Checking if '%s' address has '%d' permission", address, perm)
		check, err := s.helper.CheckAccountPermission(ctx, perm, address)
		if err != nil {
			log.Errorf("Checking account permission error: %s", err)

			// TODO skip error?
		}
		if !check {
			log.Errorf("Could not find '%d' permission for '%s'", perm, address)

			// TODO skip error?
		}

	}
	return nil
}
