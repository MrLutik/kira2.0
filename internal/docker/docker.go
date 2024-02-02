package docker

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/osutils"
)

type (
	DockerConfig struct {
		Host       string `json:"Host"`
		APIVersion string `json:"APIVersion,omitempty"`
		CertPath   string `json:"CertPath"`
		CacertPath string `json:"CacertPath"`
		KeyPath    string `json:"KeyPath"`
	}
	DockerManager struct {
		cli           *client.Client
		commandRunner CommandRunner

		log *logging.Logger
	}

	CommandRunner interface {
		RunCommandV2(commandStr string) ([]byte, error)
	}
)

func (dm *DockerConfig) SetVersion(version string) {
	dm.APIVersion = version
}

func NewTestDockerManager(client *client.Client, utils *osutils.OSUtils, logger *logging.Logger) *DockerManager {
	return &DockerManager{
		cli:           client,
		commandRunner: utils,
		log:           logger,
	}
}

// VerifyDockerInstallation verifies if Docker is installed and running by pinging the Docker daemon.
// ctx: The context.Context to use for the ping operation.
// Returns an error if the Docker daemon is not reachable or if there is an error in the ping operation.
func (d *DockerManager) VerifyDockerInstallation(ctx context.Context) error {
	// Try to ping the Docker daemon to check if it's running
	_, err := d.cli.Ping(ctx)
	if err != nil {
		d.log.Errorf("Pinging Docker daemon error: %s", err)
		return fmt.Errorf("failed to ping Docker daemon: %w", err)
	}

	// If we got here, Docker is installed and running
	d.log.Infoln("Docker is installed and running!")
	return nil
}

// PullImage pulls the specified Docker image using the Docker client associated with the DockerManager.
// It streams the image pull output to a buffer and logs the prettified output.
// ctx: The context.Context to use for the image pull operation.
// imageName: The name of the Docker image to pull.
// Returns an error if the image pull fails or if there is an error in copying the image pull output.
func (d *DockerManager) PullImage(ctx context.Context, imageName string) error {
	options := types.ImagePullOptions{}
	reader, err := d.cli.ImagePull(ctx, imageName, options)
	if err != nil {
		d.log.Errorf("Failed to pull image: %s", err)
		return fmt.Errorf("failed to pull image: %w", err)
	}
	defer reader.Close()

	// Create a buffer for the reader
	buf := new(bytes.Buffer)

	// Copy the image pull output to the buffer
	_, err = io.Copy(buf, reader)
	if err != nil {
		d.log.Errorf("Failed to copy image pull output: %s", err)
		return fmt.Errorf("failed to copy image pull output: %w", err)
	}

	// Print the prettified output from the buffer
	d.log.Infof("Image pull output: %s", buf.String())

	return nil
}

// CheckAndCreateNetwork checks if a network with the specified name exists, and creates it if it doesn't.
// ctx: The context for the operation.
// networkName: The name of the network to check and create.
// Returns an error if any issue occurs during the network checking and creation process.
func (d *DockerManager) CheckOrCreateNetwork(ctx context.Context, networkName string) error {
	d.log.Infof("Checking network '%s'", networkName)

	networkList, err := d.cli.NetworkList(ctx, types.NetworkListOptions{})
	if err != nil {
		d.log.Errorf("Getting list of networks error: %s", err)
		return err
	}

	for _, network := range networkList {
		if network.Name == networkName {
			d.log.Infof("Network '%s' already exists", networkName)
			return nil
		}
	}

	d.log.Infof("Creating network '%s'", networkName)

	_, err = d.cli.NetworkCreate(ctx, networkName, types.NetworkCreate{})
	if err != nil {
		d.log.Errorf("Creating Docker network error: %s", err)
		return err
	}

	return nil
}

// GetNetworksInfo retrieves the list of Docker networks and returns the information
// about each network as a slice of types.NetworkResource containing the information about each network,
// or an error if there was a problem retrieving the networks.
func (d *DockerManager) GetNetworksInfo(ctx context.Context) ([]types.NetworkResource, error) {
	resources, err := d.cli.NetworkList(ctx, types.NetworkListOptions{})
	if err != nil {
		d.log.Errorf("Getting networks info error: %s", err)
		return nil, err
	}

	return resources, nil
}

func (d *DockerManager) RestartDockerService() error {
	d.log.Info("Restarting docker service")
	out, err := d.commandRunner.RunCommandV2("sudo systemctl restart docker")
	if err != nil {
		return fmt.Errorf("failed to restart:\n %s\n%w", string(out), err)
	}
	return nil
}

func (d *DockerManager) DisableIpTablesForDocker() error {
	d.log.Info("Disabling iptables for docker")
	filepath := "/etc/docker/daemon.json"

	type dockerServiceConfig struct {
		Iptables bool `json:"iptables"`
	}

	var config dockerServiceConfig
	file, err := os.Open(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			config.Iptables = false
		} else {
			return err
		}
	} else {
		defer file.Close()
		if err = json.NewDecoder(file).Decode(&config); err != nil {
			return err
		}
		config.Iptables = false
	}

	outFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	encoder := json.NewEncoder(outFile)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(config); err != nil {
		return err
	}

	return nil
}

func (d *DockerManager) NetworkInspect(ctx context.Context, dockerNetwork string) (*types.NetworkResource, error) {
	network, err := d.cli.NetworkInspect(ctx, dockerNetwork, types.NetworkInspectOptions{})
	if err != nil {
		return nil, fmt.Errorf("cannot get docker network info: %w", err)
	}
	return &network, nil
}

func (d *DockerManager) CloseClient() {
	d.cli.Close()
}
