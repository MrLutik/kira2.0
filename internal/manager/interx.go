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

// NewInterxManager returns configured InterxManager.
func NewInterxManager(containerManager *docker.ContainerManager, config *config.KiraConfig) (*InterxManager, error) {
	log := logging.Log
	log.Infof("Creating interx manager with port: %s, image: '%s', volume: '%s' in '%s' network",
		config.InterxPort, config.DockerImageName, config.VolumeName, config.DockerNetworkName)

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
			config.VolumeName,
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

	command := fmt.Sprintf(`interx init --home=%s`, i.config.InterxHome)
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
	log := logging.Log
	command := fmt.Sprintf("interx start -home=%s", i.config.InterxHome)
	_, err := i.containerManager.ExecCommandInContainerInDetachMode(ctx, i.config.InterxContainerName, []string{"bash", "-c", command})
	if err != nil {
		log.Errorf("Command '%s' execution error: %s", command, err)
		return err
	}
	log.Infoln("'interx' started")
	return nil
}

// runInterxContainer starts the 'interx' container and checks if the process is running.
// If the 'interx' process is not running, it initializes the 'interx' binary in the container
// and starts it again. It checks if the process is running after the initialization.
// The method waits for a specified duration before checking if the process is running.
// If any errors occur during the process, an error is returned.
func (i *InterxManager) runInterxContainer(ctx context.Context) error {
	log := logging.Log
	const delay = time.Second

	err := i.startInterxBinInContainer(ctx)
	if err != nil {
		log.Errorf("Starting 'interx' bin in '%s' container error: %s", i.config.InterxContainerName, err)
		return err
	}

	log.Warningf("Waiting for %0.0f seconds for process", delay.Seconds())
	time.Sleep(time.Second * 1)

	check, _, err := i.containerManager.CheckIfProcessIsRunningInContainer(ctx, "interx", i.config.InterxContainerName)
	if err != nil {
		log.Errorf("Starting '%s' container error: %s", i.config.InterxContainerName, err)
		return err
	}

	if !check {
		log.Warningf("Error starting 'interx' binary first time in '%s' container, initialization new instance", i.config.InterxContainerName)
		err = i.initInterxBinInContainer(ctx)
		if err != nil {
			log.Errorf("Initialization '%s' in container error: %s", i.config.InterxContainerName, err)
			return err
		}

		err := i.startInterxBinInContainer(ctx)
		if err != nil {
			log.Errorf("Running 'interx' bin in '%s' container error: %s", i.config.InterxContainerName, err)
			return fmt.Errorf("running 'interx' bin in '%s' container error: %w", i.config.InterxContainerName, err)
		}

		log.Warningf("Waiting for %0.0f seconds for process", delay.Seconds())
		time.Sleep(delay)

		check, _, err := i.containerManager.CheckIfProcessIsRunningInContainer(ctx, "interx", i.config.InterxContainerName)
		if err != nil {
			log.Errorf("Checking 'interx' process in '%s' container error: %s", i.config.InterxContainerName, err)
			return fmt.Errorf("checking 'interx' process in '%s' container error: %w", i.config.InterxContainerName, err)
		}
		if !check {
			log.Errorf("Error starting 'interx' binary second time in '%s' container", i.config.InterxContainerName)
			return fmt.Errorf("cannot start 'interx' bin 'in' %s container", i.config.InterxContainerName)
		}
	}

	return nil
}
