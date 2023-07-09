package manager

// Config is a configuration for sekaid or interx manager.
type Config struct {
	NetworkName         string // name of a blockchain name (chain-ID)
	SekaidHome          string // home folder for sekai bin
	InterxHome          string // home folder for interx bin
	KeyringBackend      string // name of keyring backend
	DockerImageName     string // name of a docker image that will be used to create containers for sekai and interx
	DockerImageVersion  string // version of a docker image that will be used to create containers for sekai and interx
	DockerNetworkName   string // the name of docker network that will be create and used for sekaid and interx containers
	SekaiVersion        string // version of sekai binary
	InterxVersion       string // version of interx binary
	SekaidContainerName string // name for sekai container
	InterxContainerName string // name for interx container
	VolumeName          string // the name of a docker's volume that will be SekaidContainerName and InterxContainerName will be using
	MnemonicDir         string // destination where mnemonics file will be saved
	RpcPort             string // sekaid's rpc port
	GrpcPort            string // sekaid's grpc port
	InterxPort          string // interx endpoint port
	Moniker             string // Moniker
}

// NewConfig creates a new Config with provided values.
func NewConfig(
	networkName,
	sekaidHome,
	interxHome,
	keyringBackend,
	dockerImageName,
	dockerImageVersion,
	dockerNetworkName,
	sekaiVersion,
	interxVersion,
	sekaidContainerName,
	interxContainerName,
	volumeName,
	mnemonicDir,
	rpcPort,
	grpcPort,
	interxPort,
	moniker string,
) *Config {
	return &Config{
		NetworkName:         networkName,
		SekaidHome:          sekaidHome,
		InterxHome:          interxHome,
		KeyringBackend:      keyringBackend,
		DockerImageName:     dockerImageName,
		DockerImageVersion:  dockerImageVersion,
		DockerNetworkName:   dockerNetworkName,
		SekaiVersion:        sekaiVersion,
		InterxVersion:       interxVersion,
		SekaidContainerName: sekaidContainerName,
		InterxContainerName: interxContainerName,
		VolumeName:          volumeName,
		MnemonicDir:         mnemonicDir,
		RpcPort:             rpcPort,
		GrpcPort:            grpcPort,
		InterxPort:          interxPort,
		Moniker:             moniker,
	}
}
