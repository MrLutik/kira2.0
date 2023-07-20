package monitoring

import (
	"context"
	"strconv"

	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/errors"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/monitoring"
	"github.com/spf13/cobra"
)

const (
	use   = "monitoring"
	short = "Monitoring sekaid network"
	long  = "Monitoring sekaid network"
)

const (
	// SEKAID_HOME           = `/root/.sekai`
	// DOCKER_NETWORK_NAME   = "kiranet"
	// SEKAID_CONTAINER_NAME = "validator"
	// RPC_PORT              = 36657

	SEKAID_HOME           = `/data/.sekai`
	DOCKER_NETWORK_NAME   = "kira_network"
	SEKAID_CONTAINER_NAME = "sekaid"
	RPC_PORT              = 26657
	KEYRING_BACKEND       = "test"
	INTERX_CONTAINER_NAME = "interx"
	INTERX_PORT           = 11000
)

var log = logging.Log

func Monitoring() *cobra.Command {
	log.Debugln("Adding `monitoring` command...")
	monitoringCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(_ *cobra.Command, _ []string) {
			mainMonitoring()
		},
	}

	return monitoringCmd
}

func mainMonitoring() {
	dockerManager, err := docker.NewTestDockerManager()
	errors.HandleFatalErr("Can't create instance of docker manager", err)
	defer dockerManager.Cli.Close()

	containerManager, err := docker.NewTestContainerManager()
	errors.HandleFatalErr("Can't create instance of container docker manager", err)
	defer dockerManager.Cli.Close()

	ctx := context.Background()

	err = dockerManager.VerifyDockerInstallation(ctx)
	errors.HandleFatalErr("Docker is not available", err)

	monitoring := monitoring.NewMonitoringService(dockerManager, containerManager)

	networkResource, _ := monitoring.GetDockerNetwork(ctx, DOCKER_NETWORK_NAME)
	log.Infof("%+v", networkResource)

	cpuLoadPercentage, _ := monitoring.GetCPULoadPercentage()
	log.Infof("CPU Load: %.2f%%", cpuLoadPercentage)

	ramUsageInfo, _ := monitoring.GetRAMUsage()
	log.Infof("Ram usage: %+v", ramUsageInfo)

	diskUsageInfo, _ := monitoring.GetDiskUsage()
	log.Infof("Disk usage: %+v", diskUsageInfo)

	publicIpAddress, _ := monitoring.GetPublicIP()
	log.Infof("Public IP: %s", publicIpAddress)

	interfacesIPaddresses, _ := monitoring.GetInterfacesIP()
	log.Infof("Interfaces IP: %+v", interfacesIPaddresses)

	validatorAddress, _ := monitoring.GetValidatorAddress(ctx, SEKAID_CONTAINER_NAME, KEYRING_BACKEND, SEKAID_HOME)
	log.Infof("Validator address: %s", validatorAddress)

	topOfValidator, _ := monitoring.GetTopForValidator(ctx, strconv.Itoa(INTERX_PORT), validatorAddress)
	log.Infof("Validator top: %s", topOfValidator)

	valopersInfo, _ := monitoring.GetValopersInfo(ctx, strconv.Itoa(INTERX_PORT))
	log.Infof("Valopers info: %+v", valopersInfo)

	consensusInfo, _ := monitoring.GetConsensusInfo(ctx, strconv.Itoa(INTERX_PORT))
	log.Infof("Consensus info: %+v", consensusInfo)

	sekaidContainerInfo, _ := monitoring.GetContainerInfo(ctx, SEKAID_CONTAINER_NAME, DOCKER_NETWORK_NAME)
	log.Infof("%+v", sekaidContainerInfo)

	interxContainerInfo, _ := monitoring.GetContainerInfo(ctx, INTERX_CONTAINER_NAME, DOCKER_NETWORK_NAME)
	log.Infof("%+v", interxContainerInfo)

	sekaidNetworkInfo, _ := monitoring.GetSekaidInfo(ctx, strconv.Itoa(RPC_PORT))
	log.Infof("Sekaid network info: %+v", sekaidNetworkInfo)

	interxNetworkInfo, _ := monitoring.GetInterxInfo(ctx, strconv.Itoa(INTERX_PORT))
	log.Infof("Interx network info: %+v", interxNetworkInfo)
}
