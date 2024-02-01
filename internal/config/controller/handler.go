package controller

import (
	"errors"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/mrlutik/kira2.0/internal/config"
)

var ErrFileDoesNotExist = errors.New("file does NOT exist")

func (c *Controller) getConfigFolderPath() (string, error) {
	usr, err := c.utils.GetCurrentOSUser()
	if err != nil {
		return "", err
	}

	const pathPattern = "%s/.config/km2"
	return fmt.Sprintf(pathPattern, usr.HomeDir), nil
}

func (c *Controller) ReadConfigFile(filePath string) (*config.KiraConfig, error) {
	isFileExist, err := c.utils.CheckIfFileExist(filePath)
	if err != nil {
		return nil, err
	}

	if !isFileExist {
		return nil, fmt.Errorf("'%s' %w", filePath, ErrFileDoesNotExist)
	}

	c.log.Debugf("Reading config from '%s' file", filePath)
	var cfg *config.KiraConfig
	if _, err = toml.DecodeFile(filePath, &cfg); err != nil {
		return nil, fmt.Errorf("error reading config from file '%s': %w", filePath, err)
	}

	c.log.Debugf("Config: %+v", cfg)
	return cfg, nil
}

func (c *Controller) WriteConfigFile(filePath string, cfg *config.KiraConfig) error {
	c.log.Debugf("Writing %+v to '%s'", cfg, filePath)
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
