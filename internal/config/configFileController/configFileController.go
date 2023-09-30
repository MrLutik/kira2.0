package configFileController

import (
	"fmt"

	"github.com/mrlutik/kira2.0/internal/config"
	configHandler "github.com/mrlutik/kira2.0/internal/config/configFileController/configFileHandler"

	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/osutils"
)

var log = logging.Log

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
		defaultCfg := config.NewDefaultKiraConfig()
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
