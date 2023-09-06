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
	"github.com/mrlutik/kira2.0/internal/firewall/firewallManager"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/manager/utils"
	"github.com/mrlutik/kira2.0/internal/monitoring"
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
	dockerManager          *docker.DockerManager
}

const (
	validatorAccountName = "validator"
	genesisFileName      = "genesis.json"
)

// Returns configured SekaidManager.
//
//	*docker.DockerManager // The pointer for docker.DockerManager instance.
//	*config	// Config of Kira application struct
func NewSekaidManager(containerManager *docker.ContainerManager, dockerManager *docker.DockerManager, config *config.KiraConfig) (*SekaidManager, error) {
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

	natPrometheusPort, err := nat.NewPort("tcp", config.PrometheusPort)
	if err != nil {
		log.Errorf("Creating NAT Prometheus port error: %s", err)
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
			config.VolumeName,
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
	helper := utils.NewHelperManager(containerManager, config)
	return &SekaidManager{
		ContainerConfig:        sekaiContainerConfig,
		SekaiHostConfig:        sekaiHostConfig,
		SekaidNetworkingConfig: sekaidNetworkingConfig,
		containerManager:       containerManager,
		dockerManager:          dockerManager,
		config:                 config,
		helper:                 helper,
	}, err
}

// runCommands executes a list of shell commands inside the Sekaid container
func (s *SekaidManager) runCommands(ctx context.Context, commands []string) error {
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
	log := logging.Log

	configDir := fmt.Sprintf("%s/config", s.config.SekaidHome)

	log.Infof("Applying new configs to '%s/%s'", configDir, filename)

	configFileContent, err := s.containerManager.GetFileFromContainer(ctx, configDir, filename, s.config.SekaidContainerName)
	if err != nil {
		log.Errorf("Can't get '%s' file of sekaid application. Error: %s", filename, err)
		return fmt.Errorf("getting '%s' file from sekaid container error: %w", filename, err)
	}

	config := string(configFileContent)
	var newConfig string
	for _, update := range configsToml {
		newConfig, err = utils.SetTomlVar(&update, config)
		if err != nil {
			log.Errorf("Updating ([%s] %s = %s) error: %s\n", update.Tag, update.Name, update.Value, err)

			// TODO What can we do if updating value is not successful?

			continue
		}

		log.Printf("Value ([%s] %s = %s) updated successfully\n", update.Tag, update.Name, update.Value)

		config = newConfig
	}

	err = s.containerManager.WriteFileDataToContainer(ctx, []byte(config), filename, configDir, s.config.SekaidContainerName)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

func (s *SekaidManager) getExternalP2PAddress() (config.TomlValue, error) {
	log := logging.Log

	publicIp, err := monitoring.GetPublicIP() // TODO move method to other package?
	if err != nil {
		log.Errorf("Getting public IP address error: %s", err)
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
	log := logging.Log

	// Adding external p2p address to config
	// This action performed here due to avoiding duplication
	// Genesis and Joiner should both have this configuration
	externalP2PConfig, err := s.getExternalP2PAddress()
	if err != nil {
		log.Errorf("Getting external P2P address error: %s", err)
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
	var masterMnemonic string
	log := logging.Log
	var err error
	if s.config.Recover {
		masterMnemonic = s.helper.MnemonicReader()
	} else {
		bip39mn, err := s.helper.GenerateMnemonic()
		if err != nil {
			return err
		}
		masterMnemonic = bip39mn.String()
	}
	log.Printf("MASTER MNEMONIC IS:\n%s\n", masterMnemonic)
	s.config.MasterMnamonicSet, err = s.helper.GenerateMnemonicsFromMaster(string(masterMnemonic))
	if err != nil {
		return err
	}
	return nil
}

// initGenesisSekaidBinInContainer sets up the 'sekaid' Genesis container and initializes it with necessary configurations.
func (s *SekaidManager) initGenesisSekaidBinInContainer(ctx context.Context) error {
	log := logging.Log
	log.Infof("Setting up '%s' (sekaid) genesis container", s.config.SekaidContainerName)

	// initcmd := fmt.Sprintf(`sekaid init  --overwrite --chain-id=%s --home=%s "%s"`, s.config.NetworkName, s.config.SekaidHome, s.config.Moniker)
	err := s.helper.SetSekaidKeys(ctx)
	if err != nil {
		log.Errorf("Can't set sekaid keys: %s", err)
		return fmt.Errorf("Can't set sekaid keys %w", err)
	}
	// s.containerManager.ExecCommandInContainer(ctx, s.config.SekaidContainerName, []string{initcmd})

	commands := []string{
		fmt.Sprintf(`sekaid init  --overwrite --chain-id=%s --home=%s "%s"`,
			s.config.NetworkName, s.config.SekaidHome, s.config.Moniker),
		fmt.Sprintf("mkdir %s", s.config.MnemonicDir),
		fmt.Sprintf(`yes %s | sekaid keys add "%s" --keyring-backend=%s --home=%s --output=json --recover | jq .mnemonic > %s/sekai.mnemonic`,
			s.config.MasterMnamonicSet.ValidatorAddrMnemonic, validatorAccountName, s.config.KeyringBackend, s.config.SekaidHome, s.config.MnemonicDir),
		fmt.Sprintf(`yes %s | sekaid keys add "signer" --keyring-backend=%s --home=%s --output=json --recover | jq .mnemonic > %s/sekai.mnemonic`,
			s.config.MasterMnamonicSet.SignerAddrMnemonic, s.config.KeyringBackend, s.config.SekaidHome, s.config.MnemonicDir),
		fmt.Sprintf(`sekaid keys add "faucet" --keyring-backend=%s --home=%s --output=json | jq .mnemonic > %s/faucet.mnemonic`,
			s.config.KeyringBackend, s.config.SekaidHome, s.config.MnemonicDir),
		fmt.Sprintf("sekaid add-genesis-account %s 150000000000000ukex,300000000000000test,2000000000000000000000000000samolean,1000000lol --keyring-backend=%s --home=%s",
			validatorAccountName, s.config.KeyringBackend, s.config.SekaidHome),
		fmt.Sprintf(`sekaid gentx-claim %s --keyring-backend=%s --moniker="%s" --home=%s`,
			validatorAccountName, s.config.KeyringBackend, s.config.Moniker, s.config.SekaidHome),
	}

	err = s.runCommands(ctx, commands)
	if err != nil {
		log.Errorf("Initialized container error: %s", err)
		return err
	}
	err = s.applyNewConfigToml(ctx, s.getStandardConfigPack())
	if err != nil {
		log.Errorf("Can't apply new config, error: %s", err)
		return fmt.Errorf("applying new config error: %w", err)
	}

	err = s.applyNewAppToml(ctx, s.getGenesisAppConfig())
	if err != nil {
		log.Errorf("Can't apply new app config, error: %s", err)
		return fmt.Errorf("applying new app config error: %w", err)
	}

	log.Infof("'sekaid' genesis container '%s' initialized", s.config.SekaidContainerName)
	return nil
}

// initJoinerSekaidBinInContainer sets up the 'sekaid' joiner container and initializes it with necessary configurations.
func (s *SekaidManager) initJoinerSekaidBinInContainer(ctx context.Context, genesisData []byte) error {
	log := logging.Log
	log.Infof("Setting up '%s' joiner container", s.config.SekaidContainerName)
	s.helper.SetSekaidKeys(ctx)
	commands := []string{
		fmt.Sprintf("mkdir %s", s.config.MnemonicDir),
		fmt.Sprintf(`yes %s | sekaid keys add "%s" --keyring-backend=%s --home=%s --output=json --recover | jq .mnemonic > %s/sekai.mnemonic`,
			s.config.MasterMnamonicSet.ValidatorAddrMnemonic, validatorAccountName, s.config.KeyringBackend, s.config.SekaidHome, s.config.MnemonicDir),
	}

	err := s.runCommands(ctx, commands)
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

	updates := s.getStandardConfigPack()
	if len(s.config.ConfigTomlValues) == 0 {
		log.Errorf("There is no provided configs for joiner")
		return fmt.Errorf("cannot apply empty necessary configs for joiner")
	}
	updates = append(updates, s.config.ConfigTomlValues...)

	err = s.applyNewConfigToml(ctx, updates)
	if err != nil {
		log.Errorf("Can't apply new config, error: %s", err)
		return fmt.Errorf("applying new config error: %w", err)
	}

	err = s.applyNewAppToml(ctx, s.getJoinerAppConfig())
	if err != nil {
		log.Errorf("Can't apply new app config, error: %s", err)
		return fmt.Errorf("applying new app config error: %w", err)
	}

	log.Infof("'sekaid' joiner container '%s' initialized", s.config.SekaidContainerName)
	return nil
}

// startSekaidBinInContainer starts the 'sekaid' binary inside the Sekaid container.
func (s *SekaidManager) startSekaidBinInContainer(ctx context.Context) error {
	log := logging.Log
	log.Infof("Setting up '%s' genesis container", s.config.SekaidContainerName)

	// TODO move all args to config.toml
	command := fmt.Sprintf("sekaid start --home=%s --trace", s.config.SekaidHome)
	_, err := s.containerManager.ExecCommandInContainerInDetachMode(ctx, s.config.SekaidContainerName, []string{"bash", "-c", command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s", command, err)
	}

	return nil
}

// runGenesisSekaidContainer starts the 'sekaid' container and checks if the process is running.
// If the 'sekaid' process is not running, it initializes the Sekaid node using the `initializeGenesisSekaid` method.
// The method waits for a specified duration before checking if the 'sekaid' process is running.
func (s *SekaidManager) runGenesisSekaidContainer(ctx context.Context) error {
	log := logging.Log

	if err := s.startSekaidBinInContainer(ctx); err != nil {
		log.Errorf("Cannot start 'sekaid' bin in '%s' container, error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("cannot start 'sekaid' bin in '%s' container, error: %w", s.config.SekaidContainerName, err)
	}

	const delay = time.Second * 3
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

// runJoinerSekaidContainer starts the 'sekaid' container and checks if the process is running.
// If the 'sekaid' process is not running, it initializes the Sekaid joiner node using the `initializeJoinerSekaid` method.
// The method waits for a specified duration before checking if the 'sekaid' process is running.
func (s *SekaidManager) runJoinerSekaidContainer(ctx context.Context, genesis []byte) error {
	log := logging.Log

	if err := s.startSekaidBinInContainer(ctx); err != nil {
		log.Errorf("Cannot start 'sekaid' bin in '%s' container, error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("cannot start 'sekaid' bin in '%s' container, error: %w", s.config.SekaidContainerName, err)
	}

	const delay = time.Second * 3
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

// initializeGenesisSekaid initializes the Sekaid node by performing several setup steps.
// It starts the 'sekaid' binary for the first time in the specified container,
// initializes a new instance, and waits for it to start.
// If any errors occur during the initialization process, an error is returned.
// The method also propagates genesis proposals and updates the identity registrar from the validator account.
func (s *SekaidManager) initializeGenesisSekaid(ctx context.Context) error {
	log := logging.Log

	ports := []types.Port{
		{Port: s.config.RpcPort, Type: "tcp"},
		{Port: s.config.P2PPort, Type: "tcp"},
		{Port: s.config.GrpcPort, Type: "tcp"},
		{Port: s.config.InterxPort, Type: "tcp"},
		{Port: "26660", Type: "tcp"},
		{Port: "22", Type: "tcp"},
		{Port: "53", Type: "udp"},
		{Port: "4789", Type: "udp"},
		{Port: "7946", Type: "udp"},
		{Port: "7946", Type: "tcp"},
	}

	firewallManager := firewallManager.NewFirewallManager(s.dockerManager, "validator", ports)
	err := firewallManager.SetUpFirewall(ctx, s.config)

	log.Warningf("Starting sekaid binary first time in '%s' container, initializing new instance of genesis validator", s.config.SekaidContainerName)

	if err := s.initGenesisSekaidBinInContainer(ctx); err != nil {
		log.Errorf("Setup '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("setup '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	if err := s.startSekaidBinInContainer(ctx); err != nil {
		log.Errorf("Starting 'sekaid' bin in '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("starting 'sekaid' bin in '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	const delay = time.Second * 3
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

// initializeJoinerSekaid initializes the Sekaid joiner node by performing several setup steps.
// It starts the 'sekaid' binary for the first time in the specified container,
// initializes a new instance using the provided Genesis data, and waits for it to start.
// If any errors occur during the initialization process, an error is returned.
func (s *SekaidManager) initializeJoinerSekaid(ctx context.Context, genesis []byte) error {
	log := logging.Log

	ports := []types.Port{
		{Port: s.config.RpcPort, Type: "tcp"},
		{Port: s.config.P2PPort, Type: "tcp"},
		{Port: s.config.GrpcPort, Type: "tcp"},
		{Port: s.config.InterxPort, Type: "tcp"},
		{Port: "26660", Type: "tcp"},
		{Port: "22", Type: "tcp"},
		{Port: "53", Type: "udp"},
		{Port: "4789", Type: "udp"},
		{Port: "7946", Type: "udp"},
		{Port: "7946", Type: "tcp"},
	}

	firewallManager := firewallManager.NewFirewallManager(s.dockerManager, "validator", ports)
	err := firewallManager.SetUpFirewall(ctx, s.config)

	log.Warningf("Starting sekaid binary first time in '%s' container, initializing new instance of joiner", s.config.SekaidContainerName)

	if err := s.initJoinerSekaidBinInContainer(ctx, genesis); err != nil {
		log.Errorf("Setup '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("setup '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	if err := s.startSekaidBinInContainer(ctx); err != nil {
		log.Errorf("Starting 'sekaid' bin in '%s' container error: %s", s.config.SekaidContainerName, err)
		return fmt.Errorf("starting 'sekaid' bin in '%s' container error: %w", s.config.SekaidContainerName, err)
	}

	const delay = time.Second * 3
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
