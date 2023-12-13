package config

import (
	"time"

	vlg "github.com/PeepoFrog/validator-key-gen/MnemonicsGenerator"
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
	NetworkName         string                 `json:"NetworkName"`         // name of a blockchain name (chain-ID)
	SekaidHome          string                 `json:"SekaidHome"`          // home folder for sekai bin
	InterxHome          string                 `json:"InterxHome"`          // home folder for interx bin
	KeyringBackend      string                 `json:"KeyringBackend"`      // name of keyring backend
	DockerImageName     string                 `json:"DockerImageName"`     // name of a docker image that will be used to create containers for sekai and interx
	DockerImageVersion  string                 `json:"DockerImageVersion"`  // version of a docker image that will be used to create containers for sekai and interx
	DockerNetworkName   string                 `json:"DockerNetworkName"`   // the name of docker network that will be create and used for sekaid and interx containers
	SekaiVersion        string                 `json:"SekaiVersion"`        // version of sekai binary
	InterxVersion       string                 `json:"InterxVersion"`       // version of interx binary
	SekaidContainerName string                 `json:"SekaidContainerName"` // name for sekai container
	InterxContainerName string                 `json:"InterxContainerName"` // name for interx container
	VolumeName          string                 `json:"VolumeName"`          // the name of a docker's volume that will be SekaidContainerName and InterxContainerName will be using
	MnemonicDir         string                 `json:"MnemonicDir"`         // destination where mnemonics file will be saved
	RpcPort             string                 `json:"RpcPort"`             // sekaid's rpc port
	GrpcPort            string                 `json:"GrpcPort"`            // sekaid's grpc port
	P2PPort             string                 `json:"P2PPort"`             // sekaid's p2p port
	PrometheusPort      string                 `json:"PrometheusPort"`      // prometheus port
	InterxPort          string                 `json:"InterxPort"`          // interx endpoint port
	Moniker             string                 `json:"Moniker"`             // Moniker
	SekaiDebFileName    string                 `json:"SekaiDebFileName"`    // fileName of sekai deb file
	InterxDebFileName   string                 `json:"InterxDebFileName"`   // fileName of interx deb file
	SecretsFolder       string                 `json:"SecretsFolder"`       // path where mnemonics.env and node keys located
	TimeBetweenBlocks   time.Duration          `json:"TimeBetweenBlocks"`   // Awaiting time between blocks
	KiraConfigFilePath  string                 `json:"KiraConfigFilePath"`  //string to toml km2 config file //default /home/$USER/.config/kira2/kiraConfig.toml
	ConfigTomlValues    []TomlValue            `json:"ConfigTomlValues"`    //`toml:"-"` // List of configs for update
	Recover             bool                   `json:"-" toml:"-"`          // switch for recover mode
	MasterMnamonicSet   *vlg.MasterMnemonicSet `json:"-" toml:"-"`
	// NOTE Default time of block is ~5 seconds!
	// Check (m *MonitoringService) GetConsensusInfo method
	// from cmd/monitoring/main.go
}
