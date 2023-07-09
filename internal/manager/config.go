package manager

// # Create new config file for sekaidManager
//
//	NetworkName // name of a blockchain name (chandID)
//	SekaidHome // home folder for sekai bin
//	InterxHome // home folder for interx bin
//	KeyringBackend // name of keyring
//	DockerImageName // name of a docker image that will be used to create containers for sekai and interx
//	DockerImageVersion // version of a docker image that will be used to create containers for sekai and interx
//	DockerNetworkName // the name of docker network that will be create and used for sekaid and interx containers
//	SekaiVersion // version of sekai binary
//	InterxVersion // version of interx binary
//	SekaidContainerName // name for sekai container
//	InterxContainerName // name for interx container
//	VolumeName // the name of a docker's volume that will be SekaidContainerName and InterxContainerName will be using
//	MnemonicDir // destination where mnemonics file will be saved
//	RpcPort // sekaid's rpc port
//	GrpcPort // sekaid's grpc port
//	InterxPort // interx endpoint port
//	Moniker // Moniker
func NewConfig(
	NetworkName,
	SekaidHome,
	InterxHome,
	KeyringBackend,
	DockerImageName,
	DockerImageVersion,
	DockerNetworkName,
	SekaiVersion,
	InterxVersion,
	SekaidContainerName,
	InterxContainerName,
	VolumeName,
	MnemonicDir,
	RpcPort,
	GrpcPort,
	InterxPort,
	Moniker string,
) *Config {
	return &Config{
		NetworkName:         NetworkName,
		SekaidHome:          SekaidHome,
		InterxHome:          InterxHome,
		KeyringBackend:      KeyringBackend,
		DockerImageName:     DockerImageName,
		DockerImageVersion:  DockerImageVersion,
		DockerNetworkName:   DockerNetworkName,
		SekaiVersion:        SekaiVersion,
		InterxVersion:       InterxVersion,
		SekaidContainerName: SekaidContainerName,
		InterxContainerName: InterxContainerName,
		VolumeName:          VolumeName,
		MnemonicDir:         MnemonicDir,
		RpcPort:             RpcPort,
		GrpcPort:            GrpcPort,
		InterxPort:          InterxPort,
		Moniker:             Moniker,
	}
}
