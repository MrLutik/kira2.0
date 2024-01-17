package configFileController

import (
	"errors"
	"fmt"
	"time"

	"github.com/mrlutik/kira2.0/internal/config"
	configHandler "github.com/mrlutik/kira2.0/internal/config/configFileController/configFileHandler"

	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/osutils"
)

var (
	log = logging.Log

	ErrConfigPathNotExist = errors.New("config path does not exist")
)

func ChangeConfigFile(cfg *config.KiraConfig) error {
	log.Infof("Changing config file\n")
	filePath, configPath := configHandler.GetConfigFilePath()
	isPathExist, err := osutils.CheckItPathExist(configPath)
	if err != nil {
		return fmt.Errorf("error checking if '%s' exist: %w", configPath, err)
	}
	if !isPathExist {
		return fmt.Errorf("%w: '%s'", ErrConfigPathNotExist, configPath)
	}
	err = configHandler.WriteConfigFile(filePath, cfg)
	if err != nil {
		return fmt.Errorf("error while writing cfg file: %w", err)
	}
	return nil
}

func ReadOrCreateConfig() (cfg *config.KiraConfig, err error) {
	filePath, configPath := configHandler.GetConfigFilePath()
	log.Infof("Reading config from '%s'", filePath)
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
	log.Debugf("'%s' exist: %t", filePath, okFile)
	if !okFile {
		log.Infof("Cannot find file '%s' file\nCreating new config file with default values", filePath)
		defaultCfg := newDefaultKiraConfig()
		defaultCfg.KiraConfigFilePath = filePath
		err = configHandler.WriteConfigFile(filePath, defaultCfg)
		if err != nil {
			return cfg, fmt.Errorf("cannot create new KiraConfig in: '%s': %w", filePath, err)
		}
	} else {
		log.Infof("file <%s> exist, trying to read values\n", filePath)
	}
	cfg, err = configHandler.ReadConfigFile(filePath)
	if err != nil {
		return cfg, fmt.Errorf("cannot read Kira Config from file: '%s': %w", filePath, err)
	}
	log.Debugf("Returning %+v \n", cfg)
	return cfg, nil
}

func newDefaultKiraConfig() *config.KiraConfig {
	return &config.KiraConfig{
		NetworkName:         "testnet-1",
		SekaidHome:          "/data/.sekai",
		InterxHome:          "/data/.interx",
		KeyringBackend:      "test",
		DockerImageName:     "ubuntu",
		DockerImageVersion:  "latest",
		DockerNetworkName:   "kira_network",
		SekaiVersion:        "latest",
		InterxVersion:       "latest",
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
