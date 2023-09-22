package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/mrlutik/kira2.0/internal/logging"
)

var log = logging.Log

func ReadConfigFile(filePath string) (cfg *KiraConfig, err error) {
	//todo
	//check if path exist
	log.Debugf("Reding cofig from %s file", filePath)
	if _, err := toml.DecodeFile(filePath, cfg); err != nil {
		return nil, err
	}
	log.Debugf("Config:\n %+v", cfg)

	return cfg, nil
}

func WriteConfigFile(filePath string, cfg *KiraConfig) error {
	//todo
	//check if path exist from os utils (firewall__controller branch)
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

func ReadOrCreateConfig(filePath string) (cfg *KiraConfig, err error) {
	log.Infof("Reading config from <%s>\n", filePath)
	cfg, err = ReadConfigFile(filePath)
	if err != nil {
		return cfg, fmt.Errorf("cannot read KiraConfig from file: <%s>", filePath)
	}

	if cfg == nil {
		log.Infof("Cannot read config from  <%s> file, creating new config\n", filePath)
		defaultCfg := NewKiraConfig()
		err = WriteConfigFile(filePath, defaultCfg)
		if err != nil {
			return cfg, fmt.Errorf("cannot create new KiraConfig in: <%s>", filePath)
		}
		log.Infof("Reading %s\n second time", filePath)
		cfg, err = ReadConfigFile(filePath)
		if err != nil {
			return cfg, fmt.Errorf("cannot read KiraConfig from file: <%s>", filePath)
		}
		if cfg == nil {
			return cfg, fmt.Errorf("error, cannot read config from: <%s>", filePath)

		}
	}

	return cfg, nil
}
