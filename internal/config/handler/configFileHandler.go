package configFileHandler

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/mrlutik/kira2.0/internal/config"
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

func ReadConfigFile(filePath string) (*config.KiraConfig, error) {
	isFileExist, err := osutils.CheckIfFileExist(filePath)
	if err != nil {
		return nil, err
	}

	var cfg *config.KiraConfig
	log.Debugf("File '%s' exist: %t\n", filePath, isFileExist)
	if isFileExist {
		log.Debugf("Reading config from '%s' file", filePath)
		if _, err = toml.DecodeFile(filePath, &cfg); err != nil {
			return nil, err
		}
		log.Debugf("Config:\n %+v", cfg)
	} else {
		return nil, fmt.Errorf("file '%s' does not exist: %w", filePath, err)
	}
	return cfg, nil
}

func WriteConfigFile(filePath string, cfg *config.KiraConfig) error {
	log.Debugf("Writing %+v to <%s>\n", cfg, filePath)
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
