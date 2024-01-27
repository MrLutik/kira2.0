package controller

import (
	"errors"
	"fmt"
	"time"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/config/handler"

	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/osutils"
)

type Controller struct {
	utils   *osutils.OSUtils
	handler *handler.Handler
	log     *logging.Logger
}

var ErrConfigPathNotExist = errors.New("config path does not exist")

func NewConfigController(handler *handler.Handler, utils *osutils.OSUtils, logger *logging.Logger) *Controller {
	return &Controller{
		utils:   utils,
		handler: handler,
		log:     logger,
	}
}

// TODO reorganize config package
// Ask Dmytro why we have handler + controller logic separation
func (c *Controller) GetConfigFilePath() (string, error) {
	return c.handler.GetConfigFilePath()
}

func (c *Controller) WriteConfigFile(filePath string, cfg *config.KiraConfig) error {
	return c.handler.WriteConfigFile(filePath, cfg)
}

func (c *Controller) ChangeConfigFile(cfg *config.KiraConfig) error {
	c.log.Infof("Changing config file")

	configPath, err := c.handler.GetConfigFilePath()
	if err != nil {
		return fmt.Errorf("getting config file error: %w", err)
	}

	isPathExist, err := c.utils.CheckIfPathExists(configPath)
	if err != nil {
		return fmt.Errorf("error checking if '%s' exist: %w", configPath, err)
	}
	if !isPathExist {
		return fmt.Errorf("%w: '%s'", ErrConfigPathNotExist, configPath)
	}

	err = c.handler.WriteConfigFile(configPath, cfg)
	if err != nil {
		return fmt.Errorf("writing cfg file error: %w", err)
	}

	return nil
}

func (c *Controller) ReadOrCreateConfig() (cfg *config.KiraConfig, err error) {
	c.log.Info("Reading config file")

	configPath, err := c.handler.GetConfigFilePath()
	if err != nil {
		return nil, fmt.Errorf("getting config path error: %w", err)
	}

	c.log.Infof("Config path is '%s', checking if exists", configPath)
	okPath, err := c.utils.CheckIfPathExists(configPath)
	if err != nil {
		return cfg, err
	}

	if !okPath {
		err = c.utils.CreateDirPath(configPath)
		if err != nil {
			return cfg, err
		}
	}

	okFile, err := c.utils.CheckIfFileExist(configPath)
	if err != nil {
		return cfg, err
	}

	c.log.Debugf("Config file '%s' exist: %t", configPath, okFile)
	if !okFile {
		c.log.Infof("Cannot find file '%s' file", configPath)
		c.log.Info("Creating new config file with default values")

		defaultCfg := newDefaultKiraConfig()
		defaultCfg.KiraConfigFilePath = configPath

		err = c.handler.WriteConfigFile(configPath, defaultCfg)
		if err != nil {
			return cfg, fmt.Errorf("cannot create new KiraConfig in: '%s': %w", configPath, err)
		}
	} else {
		c.log.Infof("File '%s' exist, trying to read values", configPath)
	}

	cfg, err = c.handler.ReadConfigFile(configPath)
	if err != nil {
		return cfg, fmt.Errorf("cannot read Kira Config from file: '%s': %w", configPath, err)
	}

	c.log.Debugf("Returning %+v", cfg)
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
