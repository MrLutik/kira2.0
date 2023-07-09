package main

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/mrlutik/kira2.0/internal/adapters"
	"github.com/mrlutik/kira2.0/internal/cosign"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/manager"
	"github.com/mrlutik/kira2.0/internal/systemd"
	"github.com/sirupsen/logrus"
)

const (
	NETWORK_NAME          = "testnet-1"
	SEKAID_HOME           = `/data/.sekai`
	INTERXD_HOME          = `/data/.interx`
	KEYRING_BACKEND       = "test"
	DOCKER_IMAGE_NAME     = "ghcr.io/kiracore/docker/kira-base"
	DOCKER_IMAGE_VERSION  = "v0.13.11"
	DOCKER_NETWORK_NAME   = "kira_network"
	SEKAI_VERSION         = "latest" // or v0.3.16
	INTERX_VERSION        = "latest" // or v0.4.33
	SEKAID_CONTAINER_NAME = "sekaid"
	INTERX_CONTAINER_NAME = "interx"
	VOLUME_NAME           = "kira_volume:/data"
	MNEMONIC_FOLDER       = "~/mnemonics"
	RPC_PORT              = 26657
	GRPC_PORT             = 9090
	INTERX_PORT           = 11000
	MONIKER               = "VALIDATOR"
)

const DockerImagePubKey = `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE/IrzBQYeMwvKa44/DF/HB7XDpnE+
f+mU9F/Qbfq25bBWV2+NlYMJv3KvKHNtu3Jknt6yizZjUV4b8WGfKBzFYw==
-----END PUBLIC KEY-----`

const waitingForServiceStatusTime = 3 * time.Second

var log = logging.Log

func main() {
	// TODO: Instead of consts - using config file

	// TODO: change level by flag
	log.SetLevel(logrus.DebugLevel)

	// 'docker.service' management
	dockerServiceManager, err := systemd.NewServiceManager(context.Background(), "docker.service", "replace")
	if err != nil {
		log.Fatalf("Can't create instance of service manager: %s", err)
	}

	dockerServiceContext, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	exists, err := dockerServiceManager.CheckServiceExists(dockerServiceContext)
	if err != nil {
		log.Fatalf("Can't reach the service, error: %s", err)
	}
	if !exists {
		log.Fatalln("'docker.service' is not available")
	}

	status, err := dockerServiceManager.GetServiceStatus(dockerServiceContext)
	if err != nil {
		log.Fatalf("Can't get the 'docker.service' status, error: %s", err)
	}
	if status != "active" {
		log.Errorln("'docker.service' is not active")
		log.Infof("Trying to restart it")
		err = dockerServiceManager.RestartService(dockerServiceContext)
		if err != nil {
			log.Fatalf("Can't restart 'docker.service', error: %s", err)
		}
	}

	err = dockerServiceManager.WaitForServiceStatus(dockerServiceContext, "active", waitingForServiceStatusTime)
	if err != nil {
		log.Fatalf("Waiting for status error: %s", err)
	}

	// End of 'docker.service' management

	dockerManager, err := docker.NewTestDockerManager()
	if err != nil {
		log.Fatalln("Can't create instance of docker manager", err)
	}
	defer dockerManager.Cli.Close()

	ctx := context.Background()

	err = dockerManager.VerifyDockerInstallation(ctx)
	if err != nil {
		log.Fatalf("Docker is not available: %s\n", err)
	}

	dockerBaseImageName := DOCKER_IMAGE_NAME + ":" + DOCKER_IMAGE_VERSION
	err = dockerManager.PullImage(ctx, dockerBaseImageName)
	if err != nil {
		log.Fatalf("Pulling image error: %s\n", err)
	}

	checkBool, err := cosign.VerifyImageSignature(ctx, dockerBaseImageName, DockerImagePubKey)
	if err != nil {
		log.Fatalln(err)
	}

	log.Infoln("Verified:", checkBool)

	repositories := adapters.Repositories{}
	kiraGit := "KiraCore"
	sekaiRepo := "sekai"
	interxRepo := "interx"
	repositories.Set(kiraGit, sekaiRepo, SEKAI_VERSION)
	repositories.Set(kiraGit, interxRepo, INTERX_VERSION)
	log.Infof("Getting repositories: %v", repositories.Get())

	token, exists := os.LookupEnv("GITHUB_TOKEN")
	if !exists {
		log.Fatalln("'GITHUB_TOKEN' variable is not set")
	}

	repositories = adapters.Fetch(repositories, token)

	gitHubAdapter := adapters.NewGitHubAdapter(token)
	sekaiDebFileName := "sekai-linux-amd64.deb"
	interxDebFileName := "interx-linux-amd64.deb"
	// goto F
	gitHubAdapter.DownloadBinaryFromRepo(ctx, kiraGit, sekaiRepo, sekaiDebFileName, SEKAI_VERSION)
	gitHubAdapter.DownloadBinaryFromRepo(ctx, kiraGit, interxRepo, interxDebFileName, INTERX_VERSION)
	// F:
	check, err := dockerManager.CheckForContainersName(ctx, SEKAID_CONTAINER_NAME)
	if err != nil {
		log.Fatalln(err)
	}
	if check {
		err = dockerManager.StopAndDeleteContainer(ctx, SEKAID_CONTAINER_NAME)
		if err != nil {
			log.Fatalln(err)
		}
	}

	check, err = dockerManager.CheckForContainersName(ctx, INTERX_CONTAINER_NAME)
	if err != nil {
		log.Fatalln(err)
	}
	if check {
		dockerManager.StopAndDeleteContainer(ctx, INTERX_CONTAINER_NAME)
		if err != nil {
			log.Fatalln(err)
		}
	}

	err = dockerManager.CheckOrCreateNetwork(ctx, DOCKER_NETWORK_NAME)
	if err != nil {
		log.Fatalln(err)
	}
	config := manager.NewConfig(
		NETWORK_NAME,
		SEKAID_HOME,
		INTERXD_HOME,
		KEYRING_BACKEND,
		DOCKER_IMAGE_NAME,
		DOCKER_IMAGE_VERSION,
		DOCKER_NETWORK_NAME,
		SEKAI_VERSION,
		INTERX_VERSION,
		SEKAID_CONTAINER_NAME,
		INTERX_CONTAINER_NAME,
		VOLUME_NAME,
		MNEMONIC_FOLDER,
		strconv.Itoa(RPC_PORT),
		strconv.Itoa(GRPC_PORT),
		strconv.Itoa(INTERX_PORT),
		MONIKER)

	sekaidManager, err := manager.NewSekaidManager(dockerManager, config)
	if err != nil {
		log.Fatalln(err)
	}

	err = dockerManager.InitAndCreateContainer(ctx, sekaidManager.ContainerConfig, sekaidManager.SekaidNetworkingConfig, sekaidManager.SekaiHostConfig, SEKAID_CONTAINER_NAME)
	if err != nil {
		log.Fatalln(err)
	}

	interxManager, err := manager.NewInterxManager(dockerManager, config)
	if err != nil {
		log.Fatalln(err)
	}
	err = dockerManager.InitAndCreateContainer(ctx, interxManager.ContainerConfig, interxManager.SekaidNetworkingConfig, interxManager.SekaiHostConfig, INTERX_CONTAINER_NAME)
	if err != nil {
		log.Fatalln(err)
	}

	debFileDestInContainer := "/tmp/"
	err = dockerManager.SendFileToContainer(ctx, sekaiDebFileName, debFileDestInContainer, SEKAID_CONTAINER_NAME)
	if err != nil {
		log.Fatalln("Error while sending file to container", err)
	}
	err = dockerManager.SendFileToContainer(ctx, interxDebFileName, debFileDestInContainer, INTERX_CONTAINER_NAME)
	if err != nil {
		log.Fatalln("Error while sending file to container:", err)
	}
	err = dockerManager.InstallDebPackage(ctx, SEKAID_CONTAINER_NAME, debFileDestInContainer+sekaiDebFileName)
	if err != nil {
		log.Fatalln("Error while installing dep package:", err)
	}
	err = dockerManager.InstallDebPackage(ctx, INTERX_CONTAINER_NAME, debFileDestInContainer+interxDebFileName)
	if err != nil {
		log.Fatalln("Error while installing dep package:", err)
	}

	err = sekaidManager.RunSekaidContainer(ctx)
	if err != nil {
		log.Fatalf("Error while setup '%s' container: %s\n", SEKAID_CONTAINER_NAME, err)
	}
	err = interxManager.RunInterxContainer(ctx)
	if err != nil {
		log.Fatalf("Error while setup '%s' container: %s\n", INTERX_CONTAINER_NAME, err)
	}
}
