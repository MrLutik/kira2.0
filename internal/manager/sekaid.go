package manager

import (
	"context"
	"fmt"
	"time"

	mnemonicsgenerator "github.com/PeepoFrog/validator-key-gen/MnemonicsGenerator"
	dockerTypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	"github.com/kiracore/tools/bip39gen/pkg/bip39"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/monitoring"
	"github.com/mrlutik/kira2.0/internal/types"
	"github.com/mrlutik/kira2.0/internal/utils"
)

// SekaidManager represents a manager for Sekaid container and its associated configurations.
type (
	SekaidManager struct {
		ContainerConfig        *container.Config
		SekaiHostConfig        *container.HostConfig
		SekaidNetworkingConfig *network.NetworkingConfig
		config                 *config.KiraConfig
		dockerManager          DockerManager

		containerManager ContainerManager
		helper           HelperManager
		tomlEditor       TOMLEditor

		log *logging.Logger
	}
	ContainerManager interface {
		InitAndCreateContainer(ctx context.Context, containerConfig *container.Config, networkConfig *network.NetworkingConfig, hostConfig *container.HostConfig, containerName string) error
		SendFileToContainer(ctx context.Context, filePathOnHostMachine, directoryPathOnContainer, containerID string) error
		InstallDebPackage(ctx context.Context, containerID, debDestPath string) error
		ExecCommandInContainer(ctx context.Context, containerID string, command []string) ([]byte, error)
		GetFileFromContainer(ctx context.Context, folderPathOnContainer, fileName, containerID string) ([]byte, error)
		WriteFileDataToContainer(ctx context.Context, fileData []byte, fileName, destPath, containerID string) error
		ExecCommandInContainerInDetachMode(ctx context.Context, containerID string, command []string) ([]byte, error)
		GetInspectOfContainer(ctx context.Context, containerIdentification string) (*dockerTypes.ContainerJSON, error)
		CheckIfProcessIsRunningInContainer(ctx context.Context, processName, containerName string) (bool, string, error)
		StopProcessInsideContainer(ctx context.Context, processName string, codeToStopWith int, containerName string) error
		StartContainer(ctx context.Context, containerName string) error
		StopContainer(ctx context.Context, containerName string) error
	}

	DockerManager interface {
		GetNetworksInfo(ctx context.Context) ([]dockerTypes.NetworkResource, error)
	}

	HelperManager interface {
		MnemonicReader() (masterMnemonic string)
		ReadMnemonicsFromFile(pathToFile string) (masterMnemonic string, err error)
		GenerateMnemonic() (masterMnemonic bip39.Mnemonic, err error)
		GenerateMnemonicsFromMaster(masterMnemonic string) (*mnemonicsgenerator.MasterMnemonicSet, error)
		SetSekaidKeys(ctx context.Context) error
		SetEmptyValidatorState(ctx context.Context) error
		GetAddressByName(ctx context.Context, addressName string) (string, error)
		GivePermissionToAddress(ctx context.Context, permissionToAdd int, address string) error
		CheckAccountPermission(ctx context.Context, permissionToCheck int, address string) (bool, error)
	}

	TOMLEditor interface {
		SetTomlVar(config *config.TomlValue, configStr string) (string, error)
	}
)

// NewSekaidManager initializes and returns a new instance of SekaidManager.
// This function is responsible for setting up the SekaidManager with the necessary
// Docker container configuration, host configuration, and networking settings
// for running a Sekai node in a Docker container. It logs the creation process and
// handles the NAT port mappings for the Sekai node's RPC, P2P, and Prometheus ports.
func NewSekaidManager(containerManager ContainerManager, helper HelperManager, dockerManager DockerManager, config *config.KiraConfig, logger *logging.Logger) (*SekaidManager, error) {
	logger.Infof("Creating sekaid manager with ports: %s, %s, image: '%s', volume: '%s' in '%s' network\n",
		config.P2PPort, config.RpcPort, config.DockerImageName, config.GetVolumeMountPoint(), config.DockerNetworkName)

	natRpcPort, err := nat.NewPort("tcp", config.RpcPort)
	if err != nil {
		logger.Errorf("Creating NAT RPC port error: %s", err)
		return nil, err
	}

	natP2PPort, err := nat.NewPort("tcp", config.P2PPort)
	if err != nil {
		logger.Errorf("Creating NAT P2P port error: %s", err)
		return nil, err
	}

	natPrometheusPort, err := nat.NewPort("tcp", config.PrometheusPort)
	if err != nil {
		logger.Errorf("Creating NAT Prometheus port error: %s", err)
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
			natRpcPort:        struct{}{},
			natP2PPort:        struct{}{},
			natPrometheusPort: struct{}{},
		},
	}

	sekaiHostConfig := &container.HostConfig{
		Binds: []string{
			config.GetVolumeMountPoint(),
		},
		PortBindings: nat.PortMap{
			natRpcPort:        []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: config.RpcPort}},
			natP2PPort:        []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: config.P2PPort}},
			natPrometheusPort: []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: config.PrometheusPort}},
		},
		Privileged: true,
	}

	sekaidNetworkingConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			config.DockerNetworkName: {},
		},
	}

	tomlEditor := utils.NewTOMLEditor(logger)

	return &SekaidManager{
		ContainerConfig:        sekaiContainerConfig,
		SekaiHostConfig:        sekaiHostConfig,
		SekaidNetworkingConfig: sekaidNetworkingConfig,
		config:                 config,
		dockerManager:          dockerManager,
		containerManager:       containerManager,
		helper:                 helper,
		tomlEditor:             tomlEditor,
		log:                    logger,
	}, err
}

// runCommands executes a list of shell commands inside the Sekaid container
func (s *SekaidManager) runCommands(ctx context.Context, commands []string) error {
	for _, command := range commands {
		_, err := s.containerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{"bash", "-c", command})
		if err != nil {
			s.log.Errorf("Command '%s' execution error: %s", command, err)
			return err
		}
	}

	return nil
}

// getStandardConfigPack returns a slice of toml value representing the standard configurations to apply to the 'sekaid' application.
func (s *SekaidManager) getStandardConfigPack() []config.TomlValue {
	configs := []config.TomlValue{
		// # CFG [base]
		{Tag: "", Name: "moniker", Value: s.config.Moniker},
		{Tag: "", Name: "fast_sync", Value: "true"},
		// # CFG [FASTSYNC]
		{Tag: "fastsync", Name: "version", Value: "v1"},
		// # CFG [MEMPOOL]
		{Tag: "mempool", Name: "max_txs_bytes", Value: "131072000"},
		{Tag: "mempool", Name: "max_tx_bytes", Value: "131072"},
		// # CFG [CONSENSUS]
		{Tag: "consensus", Name: "timeout_commit", Value: "10000ms"},
		{Tag: "consensus", Name: "create_empty_blocks_interval", Value: "20s"},
		{Tag: "consensus", Name: "skip_timeout_commit", Value: "false"},
		// # CFG [INSTRUMENTATION]
		{Tag: "instrumentation", Name: "prometheus", Value: "true"},
		// # CFG [P2P]
		{Tag: "p2p", Name: "pex", Value: "true"},
		{Tag: "p2p", Name: "private_peer_ids", Value: ""},
		{Tag: "p2p", Name: "unconditional_peer_ids", Value: ""},
		{Tag: "p2p", Name: "persistent_peers", Value: ""},
		{Tag: "p2p", Name: "seeds", Value: ""},
		{Tag: "p2p", Name: "laddr", Value: fmt.Sprintf("tcp://0.0.0.0:%s", s.config.P2PPort)},
		{Tag: "p2p", Name: "seed_mode", Value: "false"},
		{Tag: "p2p", Name: "max_num_outbound_peers", Value: "32"},
		{Tag: "p2p", Name: "max_num_inbound_peers", Value: "128"},
		{Tag: "p2p", Name: "send_rate", Value: "65536000"},
		{Tag: "p2p", Name: "recv_rate", Value: "65536000"},
		{Tag: "p2p", Name: "max_packet_msg_payload_size", Value: "131072"},
		{Tag: "p2p", Name: "handshake_timeout", Value: "60s"},
		{Tag: "p2p", Name: "dial_timeout", Value: "30s"},
		{Tag: "p2p", Name: "allow_duplicate_ip", Value: "true"},
		{Tag: "p2p", Name: "addr_book_strict", Value: "true"},
		// # CFG [RPC]
		{Tag: "rpc", Name: "laddr", Value: fmt.Sprintf("tcp://0.0.0.0:%s", s.config.RpcPort)},
		{Tag: "rpc", Name: "cors_allowed_origins", Value: "[ \"*\" ]"},
	}

	return configs
}

// getGenesisAppConfig returns a slice of toml value representing the genesis app configurations to apply to the 'sekaid' app.toml
func (s *SekaidManager) getGenesisAppConfig() []config.TomlValue {
	return []config.TomlValue{
		{Tag: "state-sync", Name: "snapshot-interval", Value: "1000"},
		{Tag: "state-sync", Name: "snapshot-keep-recent", Value: "2"},
		{Tag: "", Name: "pruning", Value: "nothing"},
		{Tag: "", Name: "pruning-keep-recent", Value: "2"},
		{Tag: "", Name: "pruning-keep-every", Value: "100"},
	}
}

// getJoinerAppConfig returns a slice of toml value representing the joiner app configurations to apply to the 'sekaid' app.toml
func (s *SekaidManager) getJoinerAppConfig() []config.TomlValue {
	return []config.TomlValue{
		{Tag: "state-sync", Name: "snapshot-interval", Value: "200"},
		{Tag: "state-sync", Name: "snapshot-keep-recent", Value: "2"},
		{Tag: "", Name: "pruning", Value: "custom"},
		{Tag: "", Name: "pruning-keep-recent", Value: "2"},
		{Tag: "", Name: "pruning-keep-every", Value: "100"},
		{Tag: "", Name: "pruning-interval", Value: "10"},
	}
}

// applyNewConfig applies a set of configurations to the 'sekaid' application running in the SekaidManager's container.
func (s *SekaidManager) applyNewConfig(ctx context.Context, configsToml []config.TomlValue, filename string) error {
	configDir := fmt.Sprintf("%s/config", s.config.SekaidHome)

	s.log.Infof("Applying new configs to '%s/%s'", configDir, filename)

	configFileContent, err := s.containerManager.GetFileFromContainer(ctx, configDir, filename, s.config.SekaidContainerName)
	if err != nil {
		s.log.Errorf("Can't get '%s' file of sekaid application. Error: %s", filename, err)
		return fmt.Errorf("getting '%s' file from sekaid container error: %w", filename, err)
	}

	config := string(configFileContent)
	var newConfig string
	for _, update := range configsToml {
		newConfig, err = s.tomlEditor.SetTomlVar(&update, config)
		if err != nil {
			s.log.Errorf("Updating ([%s] %s = %s) error: %s", update.Tag, update.Name, update.Value, err)

			// TODO What can we do if updating value is not successful?

			continue
		}

		s.log.Infof("Value ([%s] %s = %s) updated successfully", update.Tag, update.Name, update.Value)

		config = newConfig
	}

	err = s.containerManager.WriteFileDataToContainer(ctx, []byte(config), filename, configDir, s.config.SekaidContainerName)
	if err != nil {
		return fmt.Errorf("writing file '%s' into container error: %w", filename, err)
	}

	return nil
}

func (s *SekaidManager) getExternalP2PAddress() (config.TomlValue, error) {
	monitoringService := monitoring.NewMonitoringService(s.dockerManager, s.containerManager, s.log)
	publicIp, err := monitoringService.GetPublicIP() // TODO move method to other package?
	if err != nil {
		s.log.Errorf("Getting public IP address error: %s", err)
		return config.TomlValue{}, err
	}

	return config.TomlValue{
		Tag:   "p2p",
		Name:  "external_address",
		Value: fmt.Sprintf("tcp://%s:%s", publicIp, s.config.P2PPort),
	}, nil
}

// This function allows modifying specific values in the 'config.toml' file of the 'sekaid' application by updating its content.
func (s *SekaidManager) applyNewConfigToml(ctx context.Context, configsToml []config.TomlValue) error {
	// Adding external p2p address to config
	// This action performed here due to avoiding duplication
	// Genesis and Joiner should both have this configuration
	externalP2PConfig, err := s.getExternalP2PAddress()
	if err != nil {
		s.log.Errorf("Getting external P2P address error: %s", err)
		return err
	}
	configsToml = append(configsToml, externalP2PConfig)

	return s.applyNewConfig(ctx, configsToml, "config.toml")
}

// This function allows modifying specific values in the 'app.toml' file of the 'sekaid' application by updating its content.
func (s *SekaidManager) applyNewAppToml(ctx context.Context, configsToml []config.TomlValue) error {
	return s.applyNewConfig(ctx, configsToml, "app.toml")
}

func (s *SekaidManager) ReadOrGenerateMasterMnemonic() error {
	var (
		masterMnemonic string
		err            error
	)
	if s.config.Recover {
		masterMnemonic = s.helper.MnemonicReader()
	} else {
		masterMnemonic, err = s.helper.ReadMnemonicsFromFile(s.config.SecretsFolder + "/mnemonics.env")

		if masterMnemonic == "" || err != nil {
			s.log.Warningf("Could not read master mnemonic from file, trying to generate new one \nError: %s\n", err)
			bip39mn, err := s.helper.GenerateMnemonic()
			if err != nil {
				return err
			}
			masterMnemonic = bip39mn.String()
		} else {
			s.log.Info("Master mnemonic was found and restored")
		}
	}

	s.log.Debugf("Master mnemonic is:\n%s\n", masterMnemonic)
	s.config.MasterMnamonicSet, err = s.helper.GenerateMnemonicsFromMaster(string(masterMnemonic))
	if err != nil {
		return err
	}
	return nil
}

// initGenesisSekaidBinInContainer sets up the 'sekaid' Genesis container and initializes it with necessary configurations.
func (s *SekaidManager) initGenesisSekaidBinInContainer(ctx context.Context) error {
	s.log.Infof("Setting up '%s' (sekaid) genesis container", s.config.SekaidContainerName)

	// Have to do this because need to initialize sekaid folder
	initcmd := fmt.Sprintf(`sekaid init  --overwrite --chain-id=%s --home=%s "%s"`, s.config.NetworkName, s.config.SekaidHome, s.config.Moniker)
	out, err := s.containerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{"bash", "-c", initcmd})
	s.log.Tracef("out: %s, err:%v\n", string(out), err)

	err = s.helper.SetSekaidKeys(ctx)
	if err != nil {
		s.log.Errorf("Can't set sekaid keys: %s", err)
		return fmt.Errorf("can't set sekaid keys %w", err)
	}

	err = s.helper.SetEmptyValidatorState(ctx)
	if err != nil {
		s.log.Errorf("Setting empty validator state error: %s", err)
		return err
	}

	commands := []string{
		fmt.Sprintf(`sekaid init  --overwrite --chain-id=%s --home=%s "%s"`,
			s.config.NetworkName, s.config.SekaidHome, s.config.Moniker),
		fmt.Sprintf("mkdir %s", s.config.MnemonicDir),
		fmt.Sprintf(`yes %s | sekaid keys add "%s" --keyring-backend=%s --home=%s --output=json --recover | jq .mnemonic > %s/sekai.mnemonic`,
			s.config.MasterMnamonicSet.ValidatorAddrMnemonic, types.ValidatorAccountName, s.config.KeyringBackend, s.config.SekaidHome, s.config.MnemonicDir),
		fmt.Sprintf(`yes %s | sekaid keys add "signer" --keyring-backend=%s --home=%s --output=json --recover | jq .mnemonic > %s/sekai.mnemonic`,
			s.config.MasterMnamonicSet.SignerAddrMnemonic, s.config.KeyringBackend, s.config.SekaidHome, s.config.MnemonicDir),
		fmt.Sprintf(`sekaid keys add "faucet" --keyring-backend=%s --home=%s --output=json | jq .mnemonic > %s/faucet.mnemonic`,
			s.config.KeyringBackend, s.config.SekaidHome, s.config.MnemonicDir),
		fmt.Sprintf("sekaid add-genesis-account %s 150000000000000ukex,300000000000000test,2000000000000000000000000000samolean,1000000lol --keyring-backend=%s --home=%s",
			types.ValidatorAccountName, s.config.KeyringBackend, s.config.SekaidHome),
		fmt.Sprintf(`sekaid gentx-claim %s --keyring-backend=%s --moniker="%s" --home=%s`,
			types.ValidatorAccountName, s.config.KeyringBackend, s.config.Moniker, s.config.SekaidHome),
	}

	err = s.runCommands(ctx, commands)
	if err != nil {
		s.log.Errorf("Initialized container error: %s", err)
		return err
	}
	err = s.applyNewConfigToml(ctx, s.getStandardConfigPack())
	if err != nil {
		s.log.Errorf("Can't apply new config, error: %s", err)
		return fmt.Errorf("applying new config error: %w", err)
	}

	err = s.applyNewAppToml(ctx, s.getGenesisAppConfig())
	if err != nil {
		s.log.Errorf("Can't apply new app config, error: %s", err)
		return fmt.Errorf("applying new app config error: %w", err)
	}

	s.log.Infof("'sekaid' genesis container '%s' initialized", s.config.SekaidContainerName)
	return nil
}

// initJoinerSekaidBinInContainer sets up the 'sekaid' joiner container and initializes it with necessary configurations.
func (s *SekaidManager) initJoinerSekaidBinInContainer(ctx context.Context, genesisData []byte) error {
	s.log.Infof("Setting up '%s' joiner container", s.config.SekaidContainerName)
	err := s.helper.SetSekaidKeys(ctx)
	if err != nil {
		s.log.Errorf("Unable to set sekaid keys: %s", err)
		return err
	}
	err = s.helper.SetEmptyValidatorState(ctx)
	if err != nil {
		s.log.Errorf("Unable to set empty validator state: %s", err)
		return err
	}
	commands := []string{
		fmt.Sprintf("mkdir %s", s.config.MnemonicDir),
		fmt.Sprintf(`yes %s | sekaid keys add "%s" --keyring-backend=%s --home=%s --output=json --recover | jq .mnemonic > %s/sekai.mnemonic`,
			s.config.MasterMnamonicSet.ValidatorAddrMnemonic, types.ValidatorAccountName, s.config.KeyringBackend, s.config.SekaidHome, s.config.MnemonicDir),
	}

	err = s.runCommands(ctx, commands)
	if err != nil {
		s.log.Errorf("Initialized container error: %s", err)
		return err
	}

	// TODO Do we need to validate genesisData here?!

	err = s.containerManager.WriteFileDataToContainer(ctx, genesisData, types.GenesisFileName,
		fmt.Sprintf("%s/config", s.config.SekaidHome), s.config.SekaidContainerName)
	if err != nil {
		s.log.Errorf("Write genesis file error: %s", err)
		return err
	}

	updates := s.getStandardConfigPack()
	if len(s.config.ConfigTomlValues) == 0 {
		s.log.Errorf("There is no provided configs for joiner")
		return ErrEmptyNecessaryConfigs
	}
	updates = append(updates, s.config.ConfigTomlValues...)

	err = s.applyNewConfigToml(ctx, updates)
	if err != nil {
		s.log.Errorf("Can't apply new config, error: %s", err)
		return fmt.Errorf("applying new config error: %w", err)
	}

	err = s.applyNewAppToml(ctx, s.getJoinerAppConfig())
	if err != nil {
		s.log.Errorf("Can't apply new app config, error: %s", err)
		return fmt.Errorf("applying new app config error: %w", err)
	}

	s.log.Infof("'sekaid' joiner container '%s' initialized", s.config.SekaidContainerName)
	return nil
}

// startSekaidBinInContainer starts the 'sekaid' binary inside the Sekaid container.
func (s *SekaidManager) startSekaidBinInContainer(ctx context.Context) error {
	s.log.Infof("Setting up '%s' genesis container", s.config.SekaidContainerName)
	const processName = "sekaid"
	// TODO move all args to config.toml
	command := fmt.Sprintf(`%s start --home=%s --grpc.address "0.0.0.0:%s" --trace`, processName, s.config.SekaidHome, s.config.GrpcPort)
	_, err := s.containerManager.ExecCommandInContainerInDetachMode(ctx, s.config.SekaidContainerName, []string{"bash", "-c", command})
	if err != nil {
		s.log.Errorf("Command '%s' execution error: %s", command, err)
	}
	const delay = time.Second * 3
	s.log.Warningf("Waiting to start '%s' for %0.0f seconds", processName, delay.Seconds())
	time.Sleep(delay)

	check, _, err := s.containerManager.CheckIfProcessIsRunningInContainer(ctx, processName, s.config.SekaidContainerName)
	if err != nil {
		s.log.Errorf("Starting '%s' bin second time in '%s' container error: %s", processName, s.config.SekaidContainerName, err)
		return fmt.Errorf("starting '%s' bin second time in '%s' container error: %w", processName, s.config.SekaidContainerName, err)
	}
	if !check {
		s.log.Errorf("Process '%s' is not running in '%s' container", processName, s.config.SekaidContainerName)
		return &ProcessNotRunningError{
			ProcessName:   processName,
			ContainerName: s.config.SekaidContainerName,
		}

	}
	return nil
}

// postGenesisProposals posts genesis proposals by giving permissions to the validator account.
// It retrieves the address of the validator account and adds a set of predefined permissions to it.
// The method waits for a specified duration before the first block is propagated.
// For each permission, it gives the permission to the address and checks if the permission is assigned successfully.
// If any errors occur during the process, an error is returned.
func (s *SekaidManager) postGenesisProposals(ctx context.Context) error {
	address, err := s.helper.GetAddressByName(ctx, types.ValidatorAccountName)
	if err != nil {
		s.log.Errorf("Getting address in '%s' container error: %s", s.config.SekaidContainerName, err)
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
	s.log.Infof("Permissions to add: '%d' for: '%s'", permissions, address)

	// Waiting for first block when it's going to be propagated
	s.log.Infof("Waiting for %0.0f seconds before first block be propagated", time.Duration.Seconds(s.config.TimeBetweenBlocks))
	time.Sleep(s.config.TimeBetweenBlocks)

	for _, perm := range permissions {
		s.log.Infof("Adding permission: '%d'", perm)

		err = s.helper.GivePermissionToAddress(ctx, perm, address)
		if err != nil {
			s.log.Errorf("Giving permission '%d' error: %s", perm, err)
			return fmt.Errorf("giving permission '%d' error: %w", perm, err)
		}

		s.log.Infof("Checking if '%s' address has '%d' permission", address, perm)
		check, err := s.helper.CheckAccountPermission(ctx, perm, address)
		if err != nil {
			s.log.Errorf("Checking account permission error: %s", err)

			// TODO skip error?
		}
		if !check {
			s.log.Errorf("Could not find '%d' permission for '%s'", perm, address)

			// TODO skip error?
		}

	}
	return nil
}
