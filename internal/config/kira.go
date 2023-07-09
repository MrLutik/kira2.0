package config

// TODO future config file
type KiraConfig struct {
	NetworkName         string
	SekaidHome          string
	InterxdHome         string
	KeyringBackend      string
	DockerImageName     string
	DockerImageVersion  string
	DockerNetworkName   string
	SekaiVersion        string
	InterxVersion       string
	SekaidContainerName string
	InterxContainerName string
	VolumeName          string
	MnemonicFolder      string
	RpcPort             string
	GrpcPort            string
	InterxPort          string
	Moniker             string
}

func NewKiraConfig() *KiraConfig {
	return &KiraConfig{
		NetworkName:         "testnet-1",
		SekaidHome:          `/data/.sekai`,
		InterxdHome:         `/data/.interx`,
		KeyringBackend:      "test",
		DockerImageName:     "ghcr.io/kiracore/docker/kira-base",
		DockerImageVersion:  "v0.13.11",
		DockerNetworkName:   "kira_network",
		SekaiVersion:        "latest", // or v0.3.16
		InterxVersion:       "latest", // or v0.4.33
		SekaidContainerName: "sekaid",
		InterxContainerName: "interx",
		VolumeName:          "kira_volume:/data",
		MnemonicFolder:      "~/mnemonics",
		RpcPort:             "26657",
		GrpcPort:            "9090",
		InterxPort:          "11000",
		Moniker:             "VALIDATOR",
	}
}
