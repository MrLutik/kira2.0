package monitoring

import (
	"context"
	"time"

	"github.com/docker/docker/client"
	"github.com/mrlutik/kira2.0/internal/config/controller"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/monitoring"
	"github.com/mrlutik/kira2.0/internal/osutils"
	"github.com/spf13/cobra"
)

const (
	use   = "monitoring"
	short = "Monitoring sekaid network"
	long  = "Monitoring sekaid network"
)

func Monitoring(log *logging.Logger) *cobra.Command {
	log.Info("Adding `monitoring` command...")
	monitoringCmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Run: func(_ *cobra.Command, _ []string) {
			mainMonitoring(log)
		},
	}

	return monitoringCmd
}

func mainMonitoring(log *logging.Logger) {
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Can't initialize the Docker client: %s", err)
	}

	utilsOS := osutils.NewOSUtils(log)

	dockerManager := docker.NewTestDockerManager(client, utilsOS, log)
	if err != nil {
		log.Fatalf("Can't create instance of docker manager: %s", err)
	}
	defer dockerManager.CloseClient()

	containerManager := docker.NewTestContainerManager(client, log)
	if err != nil {
		log.Fatalf("Can't create instance of container manager: %s", err)
	}
	defer containerManager.CloseClient()

	configController := controller.NewConfigController(utilsOS, log)
	kiraCfg, _ := configController.ReadOrCreateConfig()

	// TODO make flexible setting timeout
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancelFunc()

	_ = dockerManager.VerifyDockerInstallation(ctx)

	monitoring := monitoring.NewMonitoringService(dockerManager, containerManager, log)

	networkResource, _ := monitoring.GetDockerNetwork(ctx, kiraCfg.DockerNetworkName)
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

	validatorAddress, _ := monitoring.GetValidatorAddress(ctx, kiraCfg.SekaidContainerName, kiraCfg.KeyringBackend, kiraCfg.SekaidHome)
	log.Infof("Validator address: %s", validatorAddress)

	topOfValidator, _ := monitoring.GetTopForValidator(ctx, kiraCfg.InterxPort, validatorAddress)
	log.Infof("Validator top: %s", topOfValidator)

	valopersInfo, _ := monitoring.GetValopersInfo(ctx, kiraCfg.InterxPort)
	log.Infof("Valopers info: %+v", valopersInfo)

	consensusInfo, _ := monitoring.GetConsensusInfo(ctx, kiraCfg.InterxPort)
	log.Infof("Consensus info: %+v", consensusInfo)

	sekaidContainerInfo, _ := monitoring.GetContainerInfo(ctx, kiraCfg.SekaidContainerName, kiraCfg.DockerNetworkName)
	log.Infof("%+v", sekaidContainerInfo)

	interxContainerInfo, _ := monitoring.GetContainerInfo(ctx, kiraCfg.InterxContainerName, kiraCfg.DockerNetworkName)
	log.Infof("%+v", interxContainerInfo)

	sekaidNetworkInfo, _ := monitoring.GetSekaidInfo(ctx, kiraCfg.RpcPort)
	log.Infof("Sekaid network info: %+v", sekaidNetworkInfo)

	interxNetworkInfo, _ := monitoring.GetInterxInfo(ctx, kiraCfg.InterxPort)
	log.Infof("Interx network info: %+v", interxNetworkInfo)
}
