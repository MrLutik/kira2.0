package manager

import (
	"context"
	"fmt"
	"os"

	"github.com/mrlutik/kira2.0/internal/osutils"
)

const debFileDestInContainer = "/tmp/"

// InitJoiner inits new node which connects to existing network
// It delegates the setup process to initializeContainer method with 'isGenesisValidator' - false and with provided genesis file.
func (s *SekaidManager) InitJoiner(ctx context.Context, genesis []byte) error {
	err := s.initializeContainer(ctx, genesis, false)
	if err != nil {
		return fmt.Errorf("initialization joiner node failed, error: %w", err)
	}

	return nil
}

// InitNew inits new network with genesis validator
// It delegates the setup process to mustInitializeAndRunContainer method with 'isGenesisValidator' - true
func (s *SekaidManager) InitNew(ctx context.Context) error {
	err := s.initializeContainer(ctx, nil, true)
	if err != nil {
		return fmt.Errorf("initialization new node failed, error: %w", err)
	}

	return nil
}

// mustInitializeJoiner sets up containers based on the provided configuration and runs the Sekaid container.
// If the 'isGenesisValidator' flag is set to true, it sets up the container for the genesis validator, otherwise for the joiner.
// The method will terminate the program with a fatal error if any step encounters an error.
func (s *SekaidManager) initializeContainer(ctx context.Context, genesis []byte, isGenesisValidator bool) error {
	utils := osutils.NewOSUtils(s.log)

	usr, err := utils.GetCurrentOSUser()
	if err != nil {
		return fmt.Errorf("getting current user, error: %w", err)
	}

	// TODO this field is filled dynamically, is it GOOD?
	s.config.SecretsFolder = usr.HomeDir + "/.secrets"
	check, err := utils.CheckIfPathExists(s.config.SecretsFolder)
	if err != nil {
		return fmt.Errorf("checking secrets folder path error: %w", err)
	}
	if !check {
		err = os.Mkdir(s.config.SecretsFolder, os.ModePerm)
		if err != nil {
			return fmt.Errorf("creating folder error: %w", err)
		}
	}

	err = s.ReadOrGenerateMasterMnemonic()
	if err != nil {
		return fmt.Errorf("reading or generating master mnemonic, error: %w", err)
	}

	err = s.containerManager.InitAndCreateContainer(ctx, s.ContainerConfig, s.SekaidNetworkingConfig, s.SekaiHostConfig, s.config.SekaidContainerName)
	if err != nil {
		return fmt.Errorf("sekaid initialization, error: %w", err)
	}

	err = s.containerManager.SendFileToContainer(ctx, s.config.SekaiDebFileName, debFileDestInContainer, s.config.SekaidContainerName)
	if err != nil {
		return fmt.Errorf("sending file to container '%s', error: %w", s.config.SekaidContainerName, err)
	}

	// TODO Do we need to delete file after sending?
	err = s.containerManager.InstallDebPackage(ctx, s.config.SekaidContainerName, debFileDestInContainer+s.config.SekaiDebFileName)
	if err != nil {
		return fmt.Errorf("installing dep package '%s' in container '%s', error: %w", s.config.SekaidContainerName, s.config.SekaiDebFileName, err)
	}

	if isGenesisValidator {
		err = s.initGenesisSekaidBinInContainer(ctx)
	} else {
		err = s.initJoinerSekaidBinInContainer(ctx, genesis)
	}

	if err != nil {
		return fmt.Errorf("setup container, error: %w", err)
	}

	return nil
}

// InitInterx initializes the Interx environment within a container. It performs several steps
// including creating and initializing the container, sending necessary files to the container,
// and installing a Debian package. It concludes with setting up the Interx binary in the container.
// Returns an error if any of the initialization steps fail. The error includes specific details
// about the step that failed and the underlying cause.
func (i *InterxManager) InitInterx(ctx context.Context) error {
	err := i.containerManager.InitAndCreateContainer(ctx, i.ContainerConfig, i.InterxNetworkConfig, i.InterxHostConfig, i.config.InterxContainerName)
	if err != nil {
		return fmt.Errorf("initialization Interx, error: %w", err)
	}

	err = i.containerManager.SendFileToContainer(ctx, i.config.InterxDebFileName, debFileDestInContainer, i.config.InterxContainerName)
	if err != nil {
		return fmt.Errorf("sending file to container '%s', error: %w", i.config.InterxContainerName, err)
	}

	err = i.containerManager.InstallDebPackage(ctx, i.config.InterxContainerName, debFileDestInContainer+i.config.InterxDebFileName)
	if err != nil {
		return fmt.Errorf("installing dep package '%s' in container '%s', error: %w", i.config.InterxContainerName, i.config.InterxDebFileName, err)
	}

	err = i.initInterxBinInContainer(ctx)
	if err != nil {
		return fmt.Errorf("setup container, error: %w", err)
	}
	return nil
}
