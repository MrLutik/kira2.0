package config

import (
	"time"

	"github.com/mrlutik/kira2.0/internal/types"
)

// TomlValue represents a configuration value to be updated in the '*.toml' file of the 'sekaid' application.
type TomlValue struct {
	Tag   string
	Name  string
	Value string
}

// JsonValue represents a configuration value to be updated in the '*.json' file of the 'interx' application
type JsonValue struct {
	Key   string // dot-separated keys by nesting
	Value any
}

// KiraConfig is a configuration for sekaid or interx manager.
type KiraConfig struct {
	NetworkName         string        // name of a blockchain name (chain-ID)
	SekaidHome          string        // home folder for sekai bin
	InterxHome          string        // home folder for interx bin
	KeyringBackend      string        // name of keyring backend
	DockerImageName     string        // name of a docker image that will be used to create containers for sekai and interx
	DockerImageVersion  string        // version of a docker image that will be used to create containers for sekai and interx
	DockerNetworkName   string        // the name of docker network that will be create and used for sekaid and interx containers
	SekaiVersion        string        // version of sekai binary
	InterxVersion       string        // version of interx binary
	SekaidContainerName string        // name for sekai container
	InterxContainerName string        // name for interx container
	VolumeName          string        // the name of a docker's volume that will be SekaidContainerName and InterxContainerName will be using
	MnemonicDir         string        // destination where mnemonics file will be saved
	RpcPort             string        // sekaid's rpc port
	GrpcPort            string        // sekaid's grpc port
	P2PPort             string        // sekaid's p2p port
	InterxPort          string        // interx endpoint port
	Moniker             string        // Moniker
	SekaiDebFileName    string        // fileName of sekai deb file
	InterxDebFileName   string        // fileName of interx deb file
	SecretsFolder       string        // path where mnemonics.env and node keys located
	TimeBetweenBlocks   time.Duration // Awaiting time between blocks
	ConfigTomlValues    []TomlValue   // List of configs for update
	Recover             bool          // switch for recover mode
	MasterMnamonicSet   *types.MasterMnemonicSet
	// NOTE Default time of block is ~5 seconds!
	// Check (m *MonitoringService) GetConsensusInfo method
	// from cmd/monitoring/main.go
}
