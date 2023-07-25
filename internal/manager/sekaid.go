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
	containerManager       *docker.ContainerManager
	config                 *config.KiraConfig
	helper                 *utils.HelperManager
}

const (
	validatorAccountName = "validator"
	genesisFileName      = "genesis.json"
)

// Returns configured SekaidManager.
//
//	*docker.DockerManager // The pointer for docker.DockerManager instance.
//	*config	// Config of Kira application struct
func NewSekaidManager(containerManager *docker.ContainerManager, config *config.KiraConfig) (*SekaidManager, error) {
	log := logging.Log
	log.Infof("Creating sekaid manager with ports: %s, %s, image: '%s', volume: '%s' in '%s' network\n",
		config.P2PPort, config.RpcPort, config.DockerImageName, config.VolumeName, config.DockerNetworkName)

	natRpcPort, err := nat.NewPort("tcp", config.RpcPort)
	if err != nil {
		log.Errorf("Creating NAT RPC port error: %s", err)
		return nil, err
	}

	natP2PPort, err := nat.NewPort("tcp", config.P2PPort)
	if err != nil {
		log.Errorf("Creating NAT P2P port error: %s", err)
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
			natRpcPort: struct{}{},
			natP2PPort: struct{}{},
		},
	}

	sekaiHostConfig := &container.HostConfig{
		Binds: []string{
			config.VolumeName,
		},
		PortBindings: nat.PortMap{
			natRpcPort: []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: config.RpcPort}},
			natP2PPort: []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: config.P2PPort}},
		},
		Privileged: true,
	}

	sekaidNetworkingConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			config.DockerNetworkName: {},
		},
	}
	helper := utils.NewHelperManager(containerManager, config)
	return &SekaidManager{
		ContainerConfig:        sekaiContainerConfig,
		SekaiHostConfig:        sekaiHostConfig,
		SekaidNetworkingConfig: sekaidNetworkingConfig,
		containerManager:       containerManager,
		config:                 config,
		helper:                 helper,
	}, err
}

func (s *SekaidManager) initializeSekaid(ctx context.Context, commands []string) error {
	log := logging.Log
	for _, command := range commands {
		_, err := s.containerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{"bash", "-c", command})
		if err != nil {
			log.Errorf("Command '%s' execution error: %s", command, err)
			return err
		}
	}

	return nil
}

func (s *SekaidManager) initGenesisSekaidBinInContainer(ctx context.Context) error {
	log := logging.Log
	log.Infof("Setting up '%s' (sekaid) genesis container", s.config.SekaidContainerName)

	commands := []string{
		fmt.Sprintf(`sekaid init  --overwrite --chain-id=%s --home=%s "%s"`,
			s.config.NetworkName, s.config.SekaidHome, s.config.Moniker),
		fmt.Sprintf("mkdir %s", s.config.MnemonicDir),
		fmt.Sprintf(`sekaid keys add "%s" --keyring-backend=%s --home=%s --output=json | jq .mnemonic > %s/sekai.mnemonic`,
			validatorAccountName, s.config.KeyringBackend, s.config.SekaidHome, s.config.MnemonicDir),
		fmt.Sprintf(`sekaid keys add "faucet" --keyring-backend=%s --home=%s --output=json | jq .mnemonic > %s/faucet.mnemonic`,
			s.config.KeyringBackend, s.config.SekaidHome, s.config.MnemonicDir),
		fmt.Sprintf("sekaid add-genesis-account %s 150000000000000ukex,300000000000000test,2000000000000000000000000000samolean,1000000lol --keyring-backend=%s --home=%s",
			validatorAccountName, s.config.KeyringBackend, s.config.SekaidHome),
		fmt.Sprintf(`sekaid gentx-claim %s --keyring-backend=%s --moniker="%s" --home=%s`,
			validatorAccountName, s.config.KeyringBackend, s.config.Moniker, s.config.SekaidHome),
	}

	err := s.initializeSekaid(ctx, commands)
	if err != nil {
		log.Errorf("Initialized container error: %s", err)
		return err
	}

	log.Infof("'sekaid' genesis container '%s' initialized", s.config.SekaidContainerName)
	return nil
}

func (s *SekaidManager) initJoinerSekaidBinInContainer(ctx context.Context, genesisData []byte) error {
	log := logging.Log
	log.Infof("Setting up '%s' joiner container", s.config.SekaidContainerName)

	commands := []string{
		fmt.Sprintf(`sekaid init --overwrite --chain-id=%s --home=%s "%s"`,
			s.config.NetworkName, s.config.SekaidHome, s.config.Moniker),
		fmt.Sprintf("mkdir %s", s.config.MnemonicDir),
		fmt.Sprintf(`sekaid keys add "%s" --keyring-backend=%s --home=%s --output=json | jq .mnemonic > %s/sekai.mnemonic`,
			validatorAccountName, s.config.KeyringBackend, s.config.SekaidHome, s.config.MnemonicDir),
	}

	err := s.initializeSekaid(ctx, commands)
	if err != nil {
		log.Errorf("Initialized container error: %s", err)
		return err
	}

	// TODO Do we need to validate genesisData here?!

	err = s.containerManager.WriteFileDataToContainer(ctx, genesisData, genesisFileName,
		fmt.Sprintf("%s/config", s.config.SekaidHome), s.config.SekaidContainerName)
	if err != nil {
		log.Errorf("Write genesis file error: %s", err)
		return err
	}

	log.Infof("'sekaid' joiner container '%s' initialized", s.config.SekaidContainerName)
	return nil
}

// startGenesisSekaidBinInContainer starts sekaid binary inside sekaid container name
// Returns an error if any issue occurs during the start process.
func (s *SekaidManager) startGenesisSekaidBinInContainer(ctx context.Context) error {
	log := logging.Log
	log.Infof("Setting up '%s' genesis container", s.config.SekaidContainerName)

	command := fmt.Sprintf(`sekaid start --rpc.laddr="tcp://0.0.0.0:%s" --p2p.laddr="tcp://0.0.0.0:%s" --home=%s`,
		s.config.RpcPort, s.config.P2PPort, s.config.SekaidHome)
	_, err := s.containerManager.ExecCommandInContainerInDetachMode(ctx, s.config.SekaidContainerName, []string{"bash", "-c", command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s", command, err)
	}

	return nil
}

func (s *SekaidManager) startJoinerSekaidBinInContainer(ctx context.Context) error {
	log := logging.Log
	log.Infof("Setting up '%s' joiner container", s.config.SekaidContainerName)

	command := fmt.Sprintf(`sekaid start --home=%s --fast_sync=true --p2p.seeds=%s --rpc.laddr="tcp://0.0.0.0:%s" --p2p.laddr="tcp://0.0.0.0:%s"`,
		s.config.SekaidHome, s.config.Seed, s.config.RpcPort, s.config.P2PPort)
	_, err := s.containerManager.ExecCommandInContainerInDetachMode(ctx, s.config.SekaidContainerName, []string{"bash", "-c", command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s", command, err)
	}

	return nil
}

// runGenesisSekaidContainer starts the 'sekaid' container and checks if the process is running.
// If the 'sekaid' process is not running, it initializes the sekaid node using the `initializeSekaid` method.
// The method waits for a specified duration before checking if the 'sekaid' process is running.
// If any errors occur during the process, an error is returned.
// Upon successful start of the 'sekaid' container, the method indicates that the container has started.
func (s *SekaidManager) runGenesisSekaidContainer(ctx context.Context) error {
	log := logging.Log

	if err := s.startGenesisSekaidBinInContainer(ctx); err != nil {
		log.Errorf("Cannot start 'sekaid' bin in '%s' container, error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("cannot start 'sekaid' bin in '%s' container, error: %w", s.config.SekaidContainerName, err)
	}

	const delay = time.Second * 1
	log.Warningf("Waiting to start 'sekaid' for %0.0f seconds", delay.Seconds())
	time.Sleep(delay)

	check, _, err := s.containerManager.CheckIfProcessIsRunningInContainer(ctx, "sekaid", s.config.SekaidContainerName)
	if err != nil {
		log.Errorf("Setup '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("setup '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	if !check {
		if err := s.initializeGenesisSekaid(ctx); err != nil {
			return err
		}
	}

	log.Printf("SEKAID GENESIS CONTAINER '%s' STARTED", s.config.SekaidContainerName)
	return nil
}

func (s *SekaidManager) runJoinerSekaidContainer(ctx context.Context, genesis []byte) error {
	log := logging.Log

	if err := s.startJoinerSekaidBinInContainer(ctx); err != nil {
		log.Errorf("Cannot start 'sekaid' bin in '%s' container, error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("cannot start 'sekaid' bin in '%s' container, error: %w", s.config.SekaidContainerName, err)
	}

	const delay = time.Second * 1
	log.Warningf("Waiting to start 'sekaid' for %0.0f seconds", delay.Seconds())
	time.Sleep(delay)

	check, _, err := s.containerManager.CheckIfProcessIsRunningInContainer(ctx, "sekaid", s.config.SekaidContainerName)
	if err != nil {
		log.Errorf("Setup '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("setup '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	if !check {
		if err := s.initializeJoinerSekaid(ctx, genesis); err != nil {
			return err
		}
	}

	log.Printf("SEKAID JOINER CONTAINER '%s' STARTED", s.config.SekaidContainerName)
	return nil
}

// initializeGenesisSekaid initializes the sekaid node by performing several setup steps.
// It starts the sekaid binary for the first time in the specified container,
// initializes a new instance, and waits for it to start.
// If any errors occur during the initialization process, an error is returned.
// The method also propagates genesis proposals and updates the identity registrar from the validator account.
func (s *SekaidManager) initializeGenesisSekaid(ctx context.Context) error {
	log := logging.Log

	log.Warningf("Starting sekaid binary first time in '%s' container, initializing new instance of genesis validator", s.config.SekaidContainerName)

	if err := s.initGenesisSekaidBinInContainer(ctx); err != nil {
		log.Errorf("Setup '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("setup '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	if err := s.startGenesisSekaidBinInContainer(ctx); err != nil {
		log.Errorf("Starting 'sekaid' bin in '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("starting 'sekaid' bin in '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	const delay = time.Second * 1
	log.Warningf("Waiting to start 'sekaid' for %0.0f seconds", delay.Seconds())
	time.Sleep(delay)

	check, _, err := s.containerManager.CheckIfProcessIsRunningInContainer(ctx, "sekaid", s.config.SekaidContainerName)
	if err != nil {
		log.Errorf("Starting 'sekaid' bin second time in '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("starting 'sekaid' bin second time in '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	if !check {
		log.Errorf("Process 'sekaid' is not running in '%s' container", s.config.SekaidContainerName)
		return fmt.Errorf("process 'sekaid' is not running in '%s' container", s.config.SekaidContainerName)
	}

	// TODO temporary skip this part

	// err = s.postGenesisProposals(ctx)
	// if err != nil {
	// 	log.Errorf("propagating transaction error: %s", err)
	// 	return fmt.Errorf("propagating transaction error: %w", err)
	// }

	// err = s.helper.UpdateIdentityRegistrarFromValidator(ctx, validatorAccountName)
	// if err != nil {
	// 	log.Errorf("updating identity registrar error: %s", err)
	// 	return fmt.Errorf("updating identity registrar error: %w", err)
	// }

	return nil
}

func (s *SekaidManager) initializeJoinerSekaid(ctx context.Context, genesis []byte) error {
	log := logging.Log

	log.Warningf("Starting sekaid binary first time in '%s' container, initializing new instance of joiner", s.config.SekaidContainerName)

	if err := s.initJoinerSekaidBinInContainer(ctx, genesis); err != nil {
		log.Errorf("Setup '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("setup '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	if err := s.startJoinerSekaidBinInContainer(ctx); err != nil {
		log.Errorf("Starting 'sekaid' bin in '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("starting 'sekaid' bin in '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	const delay = time.Second * 1
	log.Warningf("Waiting to start 'sekaid' for %0.0f seconds", delay.Seconds())
	time.Sleep(delay)

	check, _, err := s.containerManager.CheckIfProcessIsRunningInContainer(ctx, "sekaid", s.config.SekaidContainerName)
	if err != nil {
		log.Errorf("Starting 'sekaid' bin second time in '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("starting 'sekaid' bin second time in '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	if !check {
		log.Errorf("Process 'sekaid' is not running in '%s' container", s.config.SekaidContainerName)
		return fmt.Errorf("process 'sekaid' is not running in '%s' container", s.config.SekaidContainerName)
	}

	return nil
}

// postGenesisProposals posts genesis proposals by giving permissions to the validator account.
// It retrieves the address of the validator account and adds a set of predefined permissions to it.
// The method waits for a specified duration before the first block is propagated.
// For each permission, it gives the permission to the address and checks if the permission is assigned successfully.
// If any errors occur during the process, an error is returned.
func (s *SekaidManager) postGenesisProposals(ctx context.Context) error {
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
