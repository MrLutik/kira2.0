package configFileController

import (
	"fmt"
	"time"

	"github.com/mrlutik/kira2.0/internal/config"
	configHandler "github.com/mrlutik/kira2.0/internal/config/configFileController/configFileHandler"

	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/osutils"
)

var log = logging.Log

func ChangeConfigFile(cfg *config.KiraConfig) error {
	log.Infof("Changing config file\n")
	filePath, configPath := configHandler.GetConfigFilePath()
	okPath, err := osutils.CheckItPathExist(configPath)
	if err != nil {
		return fmt.Errorf("error while checking if %s exist, error:%s", configPath, err)
	}
	if !okPath {
		return fmt.Errorf("config path <%s> does not exist", configPath)
	}
	err = configHandler.WriteConfigFile(filePath, cfg)
	if err != nil {
		return fmt.Errorf("error while writing cfg file: %s", err)
	}
	return nil
}

func ReadOrCreateConfig() (cfg *config.KiraConfig, err error) {
	filePath, configPath := configHandler.GetConfigFilePath()
	// filePath := configPath + "/" + fileName
	log.Infof("Reading config from <%s>\n", filePath)
	okPath, err := osutils.CheckItPathExist(configPath)
	if err != nil {
		return cfg, err
	}

	if !okPath {
		err = osutils.CreateDirPath(configPath)
		if err != nil {
			return cfg, err
		}
	}

	okFile, err := osutils.CheckIfFileExist(filePath)
	if err != nil {
		return cfg, err
	}
	log.Debugf("%s exist?:%v\n", filePath, okFile)
	if !okFile {
		log.Infof("cannot find file <%s> file, creating new config file with default values\n", filePath)
		defaultCfg := newDefaultKiraConfig()
		defaultCfg.KiraConfigFilePath = filePath
		err = configHandler.WriteConfigFile(filePath, defaultCfg)
		if err != nil {
			return cfg, fmt.Errorf("cannot create new KiraConfig in: <%s>", filePath)
		}
	} else {
		log.Infof("file <%s> exist, trying to read values\n", filePath)
	}
	cfg, err = configHandler.ReadConfigFile(filePath)
	if err != nil {
		return cfg, fmt.Errorf("cannot read KiraConfig from file: <%s>, %s", filePath, err)
	}
	if cfg == nil {
		return cfg, fmt.Errorf("cannot read config from: <%s>, kiraconfig is NIL", filePath)
	}
	log.Debugf("RETURTING %+v \n", cfg)
	return cfg, nil
}

func newDefaultKiraConfig() *config.KiraConfig {
	return &config.KiraConfig{
		NetworkName:    "testnet-1",
		SekaidHome:     "/data/.sekai",
		InterxHome:     "/data/.interx",
		KeyringBackend: "test",
		// DockerImageName:     "ghcr.io/kiracore/docker/kira-base",
		// DockerImageVersion:  "v0.13.11",
		DockerImageName:     "ubuntu",
		DockerImageVersion:  "latest",
		DockerNetworkName:   "kira_network",
		SekaiVersion:        "latest", // or v0.3.32 or latest
		InterxVersion:       "latest", // or v0.4.33 or latest
		SekaidContainerName: "sekaid",
		InterxContainerName: "interx",
		VolumeName:          "kira_volume",
		VolumeMoutPath:      "/data",
		MnemonicDir:         "~/mnemonics",
		RpcPort:             "26657",
		P2PPort:             "26656",
		GrpcPort:            "9090",
		PrometheusPort:      "26660",
		InterxPort:          "11000",
		Moniker:             "VALIDATOR",
		SekaiDebFileName:    "sekai-linux-amd64.deb",
		InterxDebFileName:   "interx-linux-amd64.deb",
		TimeBetweenBlocks:   time.Second * 10,
		KiraConfigFilePath:  "/home/$USER/.config/km2",
	}
}
