package manager

import (
	"context"
	"os"

	"github.com/mrlutik/kira2.0/internal/errors"
	"github.com/mrlutik/kira2.0/internal/osutils"
)

const debFileDestInContainer = "/tmp/"

// mustInitializeAndRunContainer sets up containers based on the provided configuration and runs the Sekaid container.
// If the 'isGenesisValidator' flag is set to true, it sets up the container for the genesis validator, otherwise for the joiner.
// The method will terminate the program with a fatal error if any step encounters an error.
func (s *SekaidManager) mustInitializeAndRunContainer(ctx context.Context, genesis []byte, isGenesisValidator bool) {
	//removing and creating again .secrets folder
	usr := osutils.GetCurrentOSUser()
	s.config.SecretsFolder = usr.HomeDir + "/.secrets"

	check, err := osutils.CheckItPathExist(s.config.SecretsFolder)
	errors.HandleFatalErr("Error checking secrets folder path", err)
	if !check {
		// os.RemoveAll(s.config.SecretsFolder)
		os.Mkdir(s.config.SecretsFolder, os.ModePerm)
	}

	err = s.ReadOrGenerateMasterMnemonic()
	errors.HandleFatalErr("Reading or generating master mnemonic", err)
	check, err = s.containerManager.CheckForContainersName(ctx, s.config.SekaidContainerName)
	errors.HandleFatalErr("Checking container names", err)
	if check {
		err = s.containerManager.StopAndDeleteContainer(ctx, s.config.SekaidContainerName)
		errors.HandleFatalErr("Deleting container", err)
	}

	err = s.containerManager.InitAndCreateContainer(ctx, s.ContainerConfig, s.SekaidNetworkingConfig, s.SekaiHostConfig, s.config.SekaidContainerName)
	errors.HandleFatalErr("Sekaid initialization", err)

	err = s.containerManager.SendFileToContainer(ctx, s.config.SekaiDebFileName, debFileDestInContainer, s.config.SekaidContainerName)
	errors.HandleFatalErr("Sending file to container", err)

	// TODO Do we need to delete file after sending?

	err = s.containerManager.InstallDebPackage(ctx, s.config.SekaidContainerName, debFileDestInContainer+s.config.SekaiDebFileName)
	errors.HandleFatalErr("Installing dep package in container", err)

	if isGenesisValidator {
		err = s.runGenesisSekaidContainer(ctx)
	} else {
		err = s.runJoinerSekaidContainer(ctx, genesis)
	}
	errors.HandleFatalErr("Setup container", err)
}

// MustInitAndRunGenesisValidator initializes and runs the Sekaid container for the genesis validator.
// It delegates the setup process to mustInitializeAndRunContainer method with 'genesis' - nil and 'isGenesisValidator' - true.
func (s *SekaidManager) MustInitAndRunGenesisValidator(ctx context.Context) {
	s.mustInitializeAndRunContainer(ctx, nil, true)
}

// MustInitAndRunJoiner initializes and runs the Sekaid container for the joiner.
// It delegates the setup process to mustInitializeAndRunContainer method with 'isGenesisValidator' - false and with provided genesis file.
func (s *SekaidManager) MustInitAndRunJoiner(ctx context.Context, genesis []byte) {
	s.mustInitializeAndRunContainer(ctx, genesis, false)
}

// MustInitAndRunInterx initializes and runs the Interx container.
// The method will terminate the program with a fatal error if any step encounters an error.
func (i *InterxManager) MustInitAndRunInterx(ctx context.Context) {
	check, err := i.containerManager.CheckForContainersName(ctx, i.config.InterxContainerName)
	errors.HandleFatalErr("Checking container names", err)
	if check {
		i.containerManager.StopAndDeleteContainer(ctx, i.config.InterxContainerName)
		errors.HandleFatalErr("Deleting container", err)
	}

	errors.HandleFatalErr("Interx managing", err)
	err = i.containerManager.InitAndCreateContainer(ctx, i.ContainerConfig, i.InterxNetworkConfig, i.InterxHostConfig, i.config.InterxContainerName)
	errors.HandleFatalErr("Interx initialization", err)

	err = i.containerManager.SendFileToContainer(ctx, i.config.InterxDebFileName, debFileDestInContainer, i.config.InterxContainerName)
	errors.HandleFatalErr("Sending file to container", err)

	err = i.containerManager.InstallDebPackage(ctx, i.config.InterxContainerName, debFileDestInContainer+i.config.InterxDebFileName)
	errors.HandleFatalErr("Installing dep package in container", err)

	err = i.runInterxContainer(ctx)
	errors.HandleFatalErr("Setup container", err)
}

//run_stop_init branch

// initilazing new node which connects to existing network
// It delegates the setup process to mustInitializeAndRunContainer method with 'isGenesisValidator' - false and with provided genesis file.
func (s *SekaidManager) MustInitJoiner(ctx context.Context, genesis []byte) {
	s.mustInitializeContainer(ctx, genesis, false)
}

func (s *SekaidManager) MustInitNew(ctx context.Context) {
	s.mustInitializeContainer(ctx, nil, true)
}

// mustInitializeJoiner sets up containers based on the provided configuration and runs the Sekaid container.
// If the 'isGenesisValidator' flag is set to true, it sets up the container for the genesis validator, otherwise for the joiner.
// The method will terminate the program with a fatal error if any step encounters an error.
func (s *SekaidManager) mustInitializeContainer(ctx context.Context, genesis []byte, isGenesisValidator bool) {
	usr := osutils.GetCurrentOSUser()
	s.config.SecretsFolder = usr.HomeDir + "/.secrets"

	check, err := osutils.CheckItPathExist(s.config.SecretsFolder)
	errors.HandleFatalErr("Error checking secrets folder path", err)
	if !check {
		os.Mkdir(s.config.SecretsFolder, os.ModePerm)
	}

	err = s.ReadOrGenerateMasterMnemonic()
	errors.HandleFatalErr("Reading or generating master mnemonic", err)
	check, err = s.containerManager.CheckForContainersName(ctx, s.config.SekaidContainerName)
	errors.HandleFatalErr("Checking container names", err)
	if check {
		err = s.containerManager.StopAndDeleteContainer(ctx, s.config.SekaidContainerName)
		errors.HandleFatalErr("Deleting container", err)
	}

	err = s.containerManager.InitAndCreateContainer(ctx, s.ContainerConfig, s.SekaidNetworkingConfig, s.SekaiHostConfig, s.config.SekaidContainerName)
	errors.HandleFatalErr("Sekaid initialization", err)

	err = s.containerManager.SendFileToContainer(ctx, s.config.SekaiDebFileName, debFileDestInContainer, s.config.SekaidContainerName)
	errors.HandleFatalErr("Sending file to container", err)

	// TODO Do we need to delete file after sending?

	err = s.containerManager.InstallDebPackage(ctx, s.config.SekaidContainerName, debFileDestInContainer+s.config.SekaiDebFileName)
	errors.HandleFatalErr("Installing dep package in container", err)

	if isGenesisValidator {
		err = s.initGenesisSekaidBinInContainer(ctx)
	} else {
		// err = s.runJoinerSekaidContainer(ctx, genesis)
		err = s.initJoinerSekaidBinInContainer(ctx, genesis)
	}
	errors.HandleFatalErr("Setup container", err)
}

// MustInitAndRunInterx initializes and runs the Interx container.
// The method will terminate the program with a fatal error if any step encounters an error.
func (i *InterxManager) MustInitInterx(ctx context.Context) {
	check, err := i.containerManager.CheckForContainersName(ctx, i.config.InterxContainerName)
	errors.HandleFatalErr("Checking container names", err)
	if check {
		i.containerManager.StopAndDeleteContainer(ctx, i.config.InterxContainerName)
		errors.HandleFatalErr("Deleting container", err)
	}

	errors.HandleFatalErr("Interx managing", err)
	err = i.containerManager.InitAndCreateContainer(ctx, i.ContainerConfig, i.InterxNetworkConfig, i.InterxHostConfig, i.config.InterxContainerName)
	errors.HandleFatalErr("Interx initialization", err)

	err = i.containerManager.SendFileToContainer(ctx, i.config.InterxDebFileName, debFileDestInContainer, i.config.InterxContainerName)
	errors.HandleFatalErr("Sending file to container", err)

	err = i.containerManager.InstallDebPackage(ctx, i.config.InterxContainerName, debFileDestInContainer+i.config.InterxDebFileName)
	errors.HandleFatalErr("Installing dep package in container", err)

	err = i.initInterxBinInContainer(ctx)
	errors.HandleFatalErr("Setup container", err)
}
