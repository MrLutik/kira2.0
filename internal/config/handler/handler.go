package handler

import (
	"errors"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/osutils"
)

type Handler struct {
	utils *osutils.OSUtils

	log *logging.Logger
}

var ErrFileDoesNotExist = errors.New("file does NOT exist")

func NewHandler(utils *osutils.OSUtils, logger *logging.Logger) *Handler {
	return &Handler{
		utils: utils,
		log:   logger,
	}
}

func (h *Handler) GetConfigFilePath() (string, error) {
	usr, err := h.utils.GetCurrentOSUser()
	if err != nil {
		return "", err
	}

	const pathPattern = "%s/.config/km2/config.toml"
	return fmt.Sprintf(pathPattern, usr.HomeDir), nil
}

func (h *Handler) ReadConfigFile(filePath string) (*config.KiraConfig, error) {
	isFileExist, err := h.utils.CheckIfFileExist(filePath)
	if err != nil {
		return nil, err
	}

	if !isFileExist {
		return nil, fmt.Errorf("'%s' %w", filePath, ErrFileDoesNotExist)
	}

	h.log.Debugf("Reading config from '%s' file", filePath)
	var cfg *config.KiraConfig
	if _, err = toml.DecodeFile(filePath, &cfg); err != nil {
		return nil, fmt.Errorf("error reading config from file '%s': %w", filePath, err)
	}

	h.log.Debugf("Config: %+v", cfg)
	return cfg, nil
}

func (h *Handler) WriteConfigFile(filePath string, cfg *config.KiraConfig) error {
	h.log.Debugf("Writing %+v to '%s'", cfg, filePath)
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
