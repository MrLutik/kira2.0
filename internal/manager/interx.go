package manager

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/docker"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/manager/utils"
	"github.com/mrlutik/kira2.0/internal/types"
)

// InterxManager represents a manager for Interx container and its associated configurations.
type InterxManager struct {
	ContainerConfig     *container.Config
	InterxHostConfig    *container.HostConfig
	InterxNetworkConfig *network.NetworkingConfig
	containerManager    *docker.ContainerManager
	config              *config.KiraConfig
}

const interxProcessName = "interx"

// NewInterxManager returns configured InterxManager.
func NewInterxManager(containerManager *docker.ContainerManager, config *config.KiraConfig) (*InterxManager, error) {
	log := logging.Log
	log.Infof("Creating interx manager with port: %s, image: '%s', volume: '%s' in '%s' network",
		config.InterxPort, config.DockerImageName, config.GetVolumeMountPoint(), config.DockerNetworkName)

	natInterxPort, err := nat.NewPort("tcp", config.InterxPort)
	if err != nil {
		log.Errorf("Creating NAT interx port error: %s", err)
		return nil, err
	}

	interxContainerConfig := &container.Config{
		Image:        fmt.Sprintf("%s:%s", config.DockerImageName, config.DockerImageVersion),
		Cmd:          []string{"/bin/bash"},
		Tty:          true,
		AttachStdin:  true,
		OpenStdin:    true,
		StdinOnce:    true,
		Hostname:     fmt.Sprintf("%s.local", config.InterxContainerName),
		ExposedPorts: nat.PortSet{natInterxPort: struct{}{}},
	}
	interxNetworkingConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			config.DockerNetworkName: {},
		},
	}
	interxHostConfig := &container.HostConfig{
		Binds: []string{
			config.GetVolumeMountPoint(),
		},
		PortBindings: nat.PortMap{
			natInterxPort: []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: config.InterxPort}},
		},
		Privileged: true,
	}

	return &InterxManager{
		ContainerConfig:     interxContainerConfig,
		InterxHostConfig:    interxHostConfig,
		InterxNetworkConfig: interxNetworkingConfig,
		containerManager:    containerManager,
		config:              config,
	}, err
}

// initInterxBinInContainer sets up the 'interx' container with the specified configurations.
// Returns an error if any issue occurs during the init process.
func (i *InterxManager) initInterxBinInContainer(ctx context.Context) error {
	log := logging.Log
	log.Infof("Setting up '%s' (interx) container", i.config.InterxContainerName)

	command := fmt.Sprintf(`%s init --home=%s`, interxProcessName, i.config.InterxHome)
	_, err := i.containerManager.ExecCommandInContainer(ctx, i.config.InterxContainerName, []string{"bash", "-c", command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s", command, err)
		return err
	}

	updates, err := i.getConfigPacks(ctx)
	if err != nil {
		log.Errorf("Can't get config pack based on sekaid application, error: %s", err)
		return fmt.Errorf("config pack sekaid initialization error: %w", err)
	}

	err = i.applyNewConfigs(ctx, updates)
	if err != nil {
		log.Errorf("Can't apply new config, error: %s", err)
		return fmt.Errorf("applying new config error: %w", err)
	}

	log.Infoln("'interx' is initialized")
	return err
}

func (i *InterxManager) applyNewConfigs(ctx context.Context, updates []config.JsonValue) error {
	log := logging.Log
	filename := "config.json"

	log.Infof("Applying new configs to '%s/%s'", i.config.InterxHome, filename)

	configFileContent, err := i.containerManager.GetFileFromContainer(ctx, i.config.InterxHome, filename, i.config.InterxContainerName)
	if err != nil {
		log.Errorf("Can't get '%s' file of interx application. Error: %s", filename, err)
		return fmt.Errorf("getting '%s' file from interx container error: %w", filename, err)
	}

	var newFileContent []byte
	for _, update := range updates {
		newFileContent, err = utils.UpdateJsonValue(configFileContent, &update)
		if err != nil {
			log.Errorf("Updating: (%s = %v) error: %s\n", update.Key, update.Value, err)

			// TODO What can we do if updating value is not successful?

			continue
		}

		log.Printf("(%s = %v) updated successfully\n", update.Key, update.Value)

		configFileContent = newFileContent
	}

	err = i.containerManager.WriteFileDataToContainer(ctx, configFileContent, filename, i.config.InterxHome, i.config.InterxContainerName)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

func (i *InterxManager) getConfigPacks(ctx context.Context) ([]config.JsonValue, error) {
	log := logging.Log

	configs := make([]config.JsonValue, 0)

	node_id, err := getLocalSekaidNodeID(i.config.RpcPort)
	if err != nil {
		log.Errorf("Getting sekaid node status error: %s", err)
		return nil, err
	}

	configs = append(configs,
		// node type: validator
		config.JsonValue{Key: "node.validator_node_id", Value: node_id},
		config.JsonValue{Key: "node.node_type", Value: "validator"},

		// TODO seed: node.seed_node_id & node.node_type seed
		// TODO sentry: node.sentry_node_id & node.node_type sentry

		config.JsonValue{Key: "grpc", Value: fmt.Sprintf("dns:///%s:%s", i.config.SekaidContainerName, i.config.GrpcPort)},
		config.JsonValue{Key: "rpc", Value: fmt.Sprintf("http://%s:%s", i.config.SekaidContainerName, i.config.RpcPort)},
		config.JsonValue{Key: "port", Value: i.config.InterxPort},

		config.JsonValue{Key: "mnemonic", Value: string(i.config.MasterMnamonicSet.SignerAddrMnemonic)},
		config.JsonValue{Key: "faucet.mnemonic", Value: string(i.config.MasterMnamonicSet.SignerAddrMnemonic)},
	)

	// TODO other needed configurations

	return configs, nil
}

func getLocalSekaidNodeID(port string) (string, error) {
	log := logging.Log
	var responseStatus types.ResponseSekaidStatus

	url := fmt.Sprintf("http://localhost:%s/status", port)
	response, err := http.Get(url)
	if err != nil {
		log.Errorf("Can't reach sekaid RPC status, error: %s", err)
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Errorf("Can't read the response body")
		return "", err
	}

	err = json.Unmarshal(body, &responseStatus)
	if err != nil {
		log.Errorf("Can't parse JSON response: %s", err)
		return "", err
	}

	return responseStatus.Result.NodeInfo.ID, nil
}

// startInterxBinInContainer starts interx binary inside InterxContainerName
// Returns an error if any issue occurs during the start process.
func (i *InterxManager) startInterxBinInContainer(ctx context.Context) error {
	command := fmt.Sprintf("%s start -home=%s", interxProcessName, i.config.InterxHome)
	_, err := i.containerManager.ExecCommandInContainerInDetachMode(ctx, i.config.InterxContainerName, []string{"bash", "-c", command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s", command, err)
		return err
	}
	const delay = time.Second * 3
	log.Warningf("Waiting to start '%s' for %0.0f seconds", interxProcessName, delay.Seconds())
	time.Sleep(delay)

	check, _, err := i.containerManager.CheckIfProcessIsRunningInContainer(ctx, interxProcessName, i.config.InterxContainerName)
	if err != nil {
		log.Errorf("Starting '%s' bin second time in '%s' container error: %s", interxProcessName, i.config.InterxContainerName, err)
		return fmt.Errorf("starting '%s' bin second time in '%s' container error: %w", interxProcessName, i.config.InterxContainerName, err)
	}
	if !check {
		log.Errorf("Process '%s' is not running in '%s' container", interxProcessName, i.config.InterxContainerName)
		return fmt.Errorf("process '%s' is not running in '%s' container", interxProcessName, i.config.InterxContainerName)
	}
	log.Infof("'%s' started\n", interxProcessName)
	return nil
}
