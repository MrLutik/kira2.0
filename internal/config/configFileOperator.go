package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/osutils"
)

var log = logging.Log

func GetConfigFilePath() (filePath, folderPath string) {
	fileName := "config.toml"
	usr := osutils.GetCurrentOSUser()
	folderPath = fmt.Sprintf("%s/.config/km2", usr.HomeDir)
	filePath = folderPath + "/" + fileName
	return filePath, folderPath
}

func ReadConfigFile(filePath string) (cfg *KiraConfig, err error) {

	ok, err := osutils.CheckIfFileExist(filePath)
	if err != nil {
		return cfg, err
	}
	log.Debugf("File <%s> exist:%v\n", filePath, ok)
	if ok {
		log.Debugf("Reding config from %s file", filePath)
		if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
			return nil, err
		}
		log.Debugf("Config:\n %+v", cfg)
	} else {
		return cfg, fmt.Errorf("file <%s> does not exist, %w", filePath, err)
	}
	return cfg, nil
}

func WriteConfigFile(filePath string, cfg *KiraConfig) error {

	log.Infof("creating <%s>\n", filePath)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}

func ReadOrCreateConfig() (cfg *KiraConfig, err error) {
	filePath, configPath := GetConfigFilePath()
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
		defaultCfg := NewDefaultKiraConfig()
		defaultCfg.KiraConfigFilePath = filePath
		err = WriteConfigFile(filePath, defaultCfg)
		if err != nil {
			return cfg, fmt.Errorf("cannot create new KiraConfig in: <%s>", filePath)
		}
	} else {
		log.Infof("file <%s> exist, trying to read values\n", filePath)
	}
	cfg, err = ReadConfigFile(filePath)
	if err != nil {
		return cfg, fmt.Errorf("cannot read KiraConfig from file: <%s>, %s", filePath, err)
	}
	if cfg == nil {
		return cfg, fmt.Errorf("cannot read config from: <%s>, kiraconfig is NIL", filePath)
	}
	log.Debugf("RETURTING %+v \n", cfg)
	return cfg, nil
}
