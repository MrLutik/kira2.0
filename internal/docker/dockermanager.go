package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/mrlutik/kira2.0/internal/logging"
)

var log = logging.Log

type DockerConfig struct {
	Host       string `json:"Host"`
	APIVersion string `json:"APIVersion,omitempty"`
	CertPath   string `json:"CertPath"`
	CacertPath string `json:"CacertPath"`
	KeyPath    string `json:"KeyPath"`
}

func (dm *DockerConfig) SetVersion(version string) {
	dm.APIVersion = version
}

type DockerManager struct {
	Cli *client.Client
}

// NewDockerManager creates a new DockerManager instance based on the provided configuration.
// r: An io.Reader containing the configuration data in JSON format.
// Returns a new DockerManager instance and an error if there is an issue with decoding the configuration or creating the Docker client.
func NewDockerManager(r io.Reader) (*DockerManager, error) {
	log.Infoln("Creating new DockerManager from config")

	decoder := json.NewDecoder(r)
	config := DockerConfig{}
	err := decoder.Decode(&config)
	if err != nil {
		log.Errorf("Failed to decode config: %s\n", err)
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}

	cli, err := client.NewClientWithOpts(client.WithHost(config.Host), client.WithAPIVersionNegotiation(), client.WithTLSClientConfig(config.CacertPath, config.CertPath, config.KeyPath))

	// Add API version to config. For future use in debug log etc.
	config.SetVersion(cli.ClientVersion())
	log.Infof("Docker API version set to: %v\n", config.APIVersion)

	if err != nil {
		log.Errorf("Failed to create Docker client: %s", err)
		return nil, fmt.Errorf("failed to create Docker client: %w", err)
	}

	log.Infoln("Successfully created DockerManager")
	return &DockerManager{Cli: cli}, nil
}

func NewTestDockerManager() (*DockerManager, error) {
	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}
	return &DockerManager{Cli: client}, err
}

func NewTestDockerManagerWithVersion(version string) (*DockerManager, error) {
	client, err := client.NewClientWithOpts(client.WithVersion(version))
	if err != nil {
		return nil, err
	}
	return &DockerManager{Cli: client}, err
}

// VerifyDockerInstallation verifies if Docker is installed and running by pinging the Docker daemon.
// ctx: The context.Context to use for the ping operation.
// Returns an error if the Docker daemon is not reachable or if there is an error in the ping operation.
func (dm *DockerManager) VerifyDockerInstallation(ctx context.Context) error {
	// Try to ping the Docker daemon to check if it's running
	_, err := dm.Cli.Ping(ctx)
	if err != nil {
		log.Errorf("Error pinging Docker daemon: %s\n", err)
		return fmt.Errorf("failed to ping Docker daemon: %w", err)
	}

	// If we got here, Docker is installed and running
	log.Infoln("Docker is installed and running!")
	return nil
}

// PullImage pulls the specified Docker image using the Docker client associated with the DockerManager.
// It streams the image pull output to a buffer and logs the prettified output.
// ctx: The context.Context to use for the image pull operation.
// imageName: The name of the Docker image to pull.
// Returns an error if the image pull fails or if there is an error in copying the image pull output.
func (dm *DockerManager) PullImage(ctx context.Context, imageName string) error {
	options := types.ImagePullOptions{}
	reader, err := dm.Cli.ImagePull(ctx, imageName, options)
	if err != nil {
		log.Errorf("failed to pull image: %s\n", err)
		return fmt.Errorf("failed to pull image: %w", err)
	}
	defer reader.Close()

	// Create a buffer for the reader
	buf := new(bytes.Buffer)

	// Copy the image pull output to the buffer
	_, err = io.Copy(buf, reader)
	if err != nil {
		log.Errorf("failed to copy image pull output: %s", err)
		return fmt.Errorf("failed to copy image pull output: %w", err)
	}

	// Print the prettified output from the buffer
	log.Infof("Image pull output: %s\n", buf.String())

	return nil
}

// GetFileFromContainer allows you to retrieve a file from a Docker container and save it to the host machine.
// ctx (context.Context): The context for the operation.
// filePathOnHostMachine (string): The file path on the host machine where the file will be saved.
// filePathOnContainer (string): The file path inside the Docker container from which the file will be retrieved.
// containerID (string): The ID or name of the Docker container from which the file will be retrieved.
func (dm *DockerManager) GetFileFromContainer(ctx context.Context, filePathOnHostMachine, filePathOnContainer, containerID string) error {
	log.Infof("Getting file from container '%s' to '%s'", filePathOnContainer, filePathOnHostMachine)
	rc, _, err := dm.Cli.CopyFromContainer(ctx, containerID, filePathOnContainer)
	if err != nil {
		log.Errorf("Error during copying file from container: %s\n", err)
		return err
	}
	defer rc.Close()

	contents, err := io.ReadAll(rc)
	if err != nil {
		log.Errorf("Reading error: %s\n", err)
		return err
	}

	err = os.WriteFile(filePathOnHostMachine, contents, 0o644)
	if err != nil {
		log.Errorf("Writing file error: %s\n", err)
		return err
	}

	log.Infof("Successfully got file '%s' to the host!\n", filePathOnHostMachine)

	return nil
}

// SendFileToContainer sends a file from the host machine to a specified directory inside a Docker container.
// ctx: The context for the operation.
// filePathOnHostMachine: The path of the file on the host machine.
// directoryPathOnContainer: The path of the directory inside the container where the file will be copied.
// containerID: The ID or name of the Docker container.
// Returns an error if any issue occurs during the file sending process.
func (dm *DockerManager) SendFileToContainer(ctx context.Context, filePathOnHostMachine, directoryPathOnContainer, containerID string) error {
	log.Infof("Sending file '%s' to container '%s' to '%s'", filePathOnHostMachine, containerID, directoryPathOnContainer)
	file, err := os.Open(filePathOnHostMachine)
	if err != nil {
		log.Errorf("Opening file error: %s\n", err)
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Errorf("Can't open file stat: %s\n", err)
		return err
	}

	var buf bytes.Buffer
	tarWriter := tar.NewWriter(&buf)

	err = addFileToTar(fileInfo, file, tarWriter)
	if err != nil {
		log.Errorf("Adding file to tar error: %s\n", err)
		return err
	}

	err = tarWriter.Close()
	if err != nil {
		log.Errorf("Closing tar error: %s\n", err)
		return err
	}

	tarContent := buf.Bytes()
	tarReader := bytes.NewReader(tarContent)
	copyOptions := types.CopyToContainerOptions{
		AllowOverwriteDirWithFile: false,
	}

	err = dm.Cli.CopyToContainer(ctx, containerID, directoryPathOnContainer, tarReader, copyOptions)
	if err != nil {
		log.Errorf("Copying tar to container error: %s\n", err)
		return err
	}

	log.Infof("Successfully copied '%s' to '%s' in '%s' container\n", filePathOnHostMachine, directoryPathOnContainer, containerID)
	return nil
}

// addFileToTar adds a file to a tar archive.
// fileInfo: The file information.
// file: The reader for the file data.
// tarWriter: The tar writer.
// Returns an error if any issue occurs during the file writing process.
func addFileToTar(fileInfo os.FileInfo, file io.Reader, tarWriter *tar.Writer) error {
	log.Infof("Writing file '%s' to tar archive\n", fileInfo.Name())

	header := &tar.Header{
		Name: fileInfo.Name(),
		Mode: int64(fileInfo.Mode()),
		Size: fileInfo.Size(),
	}

	if err := tarWriter.WriteHeader(header); err != nil {
		log.Errorf("Writing tar header error: %s\n", err)
		return err
	}

	if _, err := io.Copy(tarWriter, file); err != nil {
		log.Errorf("Copying error: %s\n", err)
		return err
	}

	return nil
}

// InstallDebPackage installs a Debian package (.deb) inside a specified container.
// ctx: The context for the operation.
// containerID: The ID or name of the container where the package will be installed.
// debDestPath: The destination path of the .deb package inside the container.
// Returns an error if any issue occurs during the package installation process.
func (dm *DockerManager) InstallDebPackage(ctx context.Context, containerID, debDestPath string) error {
	log.Infof("Installing '%s'\n", debDestPath)

	installCmd := []string{"dpkg", "-i", debDestPath}
	execOptions := types.ExecConfig{
		Cmd:          installCmd,
		AttachStdout: true,
		AttachStderr: true,
	}
	resp, err := dm.Cli.ContainerExecCreate(ctx, containerID, execOptions)
	if err != nil {
		log.Errorf("Creating exec configuration error: %s\n", err)
		return err
	}

	attachOptions := types.ExecStartCheck{
		Detach: false,
		Tty:    false,
	}
	respConn, err := dm.Cli.ContainerExecAttach(ctx, resp.ID, attachOptions)
	if err != nil {
		log.Errorf("Attaching error: %s\n", err)
		return err
	}
	defer respConn.Close()

	// Capture the output from the container
	output, err := io.ReadAll(respConn.Reader)
	if err != nil {
		log.Errorf("Reading output error: %s\n", err)
		return err
	}

	// Wait for the execution to complete
	waitResponse, err := dm.Cli.ContainerExecInspect(ctx, resp.ID)
	if err != nil {
		log.Errorf("Inspecting process '%s' error: %s\n", resp.ID, err)
		return err
	}

	if waitResponse.ExitCode != 0 {
		err = fmt.Errorf("package installation failed: %s", string(output))
		log.Errorf("Installation error: %s\n", err)
		return err
	}

	log.Infof("Package '%s' installed successfully\n", debDestPath)

	return nil
}

// ExecCommandInContainerInDetachMode runs a command inside a specified container in detach mode.
// ctx: The context for the operation.
// containerID: The ID or name of the container.
// command: The command to execute inside the container.
// Returns the output of the command as a byte slice and an error if any issue occurs during the command execution.
func (dm *DockerManager) ExecCommandInContainerInDetachMode(ctx context.Context, containerID string, command []string) ([]byte, error) {
	log.Infof("Running command '%s' in detach mode in '%s'\n", strings.Join(command, " "), containerID)

	execCreateResponse, err := dm.Cli.ContainerExecCreate(ctx, containerID, types.ExecConfig{
		Cmd:          command,
		AttachStdout: false,
		AttachStderr: false,
		Detach:       true,
	})
	if err != nil {
		log.Errorf("Exec configuration error: %s\n", err)
		return nil, err
	}

	execAttachConfig := types.ExecStartCheck{}
	resp, err := dm.Cli.ContainerExecAttach(ctx, execCreateResponse.ID, execAttachConfig)
	if err != nil {
		log.Errorf("Attaching to container '%s' error: %s\n", containerID, err)
		return nil, err
	}
	defer resp.Close()

	var outBuf, errBuf bytes.Buffer
	_, err = stdcopy.StdCopy(&outBuf, &errBuf, resp.Reader)
	if err != nil {
		log.Printf("Reading response error: %s\n", err)
		return nil, err
	}

	output := outBuf.Bytes()

	log.Infoln("Reading successfully")
	return output, err
}

// ExecCommandInContainer executes a command inside a specified container.
// ctx: The context for the operation.
// containerID: The ID or name of the container.
// command: The command to execute inside the container.
// Returns the output of the command as a byte slice and an error if any issue occurs during the command execution.
func (dm *DockerManager) ExecCommandInContainer(ctx context.Context, containerID string, command []string) ([]byte, error) {
	log.Infof("Running command '%s' in '%s'\n", strings.Join(command, " "), containerID)

	execCreateResponse, err := dm.Cli.ContainerExecCreate(ctx, containerID, types.ExecConfig{
		Cmd:          command,
		AttachStdout: true,
		AttachStderr: true,
	})
	if err != nil {
		log.Errorf("Exec configuration error: %s\n", err)
		return nil, err
	}

	execAttachConfig := types.ExecStartCheck{}
	resp, err := dm.Cli.ContainerExecAttach(ctx, execCreateResponse.ID, execAttachConfig)
	if err != nil {
		log.Errorf("Attaching to container '%s' error: %s\n", containerID, err)
		return nil, err
	}
	defer resp.Close()

	var outBuf, errBuf bytes.Buffer
	_, err = stdcopy.StdCopy(&outBuf, &errBuf, resp.Reader)
	if err != nil {
		log.Errorf("Reading response error: %s\n", err)
		return nil, err
	}

	output := outBuf.Bytes()
	log.Infof("Running '%s' successfully\n", strings.Join(command, " "))
	return output, err
}

// CheckForContainersName checks if a container with the specified name exists.
// ctx: The context for the operation.
// containerNameToCheck: The name of the container to check.
// Returns true if a container with the specified name is found, false otherwise, and an error if any issue occurs during the process.
func (dm *DockerManager) CheckForContainersName(ctx context.Context, containerNameToCheck string) (bool, error) {
	log.Infof("Checking container with name: %s\n", containerNameToCheck)

	containers, err := dm.Cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		log.Errorf("Cannot get the list of containers: %s\n", err)
		return false, err
	}

	for _, c := range containers {
		for _, name := range c.Names {
			if name == `/`+containerNameToCheck {
				log.Infof("Container '%s' detected\n", name)
				return true, nil
			}
		}
	}

	log.Infof("Container '%s' is not detected\n", containerNameToCheck)
	return false, nil
}

// StopAndDeleteContainer stops and deletes a container with the specified name.
// ctx: The context for the operation.
// containerNameToStop: The name of the container to stop and delete.
// Returns an error if any issue occurs during the process.
func (dm *DockerManager) StopAndDeleteContainer(ctx context.Context, containerNameToStop string) error {
	log.Infof("Stopping '%s' container...\n", containerNameToStop)

	err := dm.Cli.ContainerStop(ctx, containerNameToStop, container.StopOptions{})
	if err != nil {
		log.Errorf("Stopping container error: %s\n", err)
		return err
	}

	log.Infof("Deleting %s container...\n", containerNameToStop)
	err = dm.Cli.ContainerRemove(ctx, containerNameToStop, types.ContainerRemoveOptions{})
	if err != nil {
		log.Println(err)
		return err
	}

	log.Infof("Container %s is deleted\n", containerNameToStop)
	return nil
}

// InitAndCreateContainer initializes and creates a new container.
// ctx: The context for the operation.
// containerConfig: The container configuration.
// networkConfig: The network configuration.
// hostConfig: The host configuration.
// containerName: The name of the container.
// Returns an error if any issue occurs during the container initialization and creation process.
func (dm *DockerManager) InitAndCreateContainer(
	ctx context.Context,
	containerConfig *container.Config,
	networkConfig *network.NetworkingConfig,
	hostConfig *container.HostConfig,
	containerName string,
) error {
	log.Infof("Starting new container '%s'\n", containerName)

	resp, err := dm.Cli.ContainerCreate(ctx, containerConfig, hostConfig, networkConfig, nil, containerName)
	if err != nil {
		log.Errorf("Creating container error: %s\n", err)
		return err
	}

	err = dm.Cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		log.Errorf("Starting container error: %s\n", err)
		return err
	}

	log.Infof("'%s' container started successfully! ID: %s\n", containerName, resp.ID)
	return err
}

// CheckAndCreateNetwork checks if a network with the specified name exists, and creates it if it doesn't.
// ctx: The context for the operation.
// networkName: The name of the network to check and create.
// Returns an error if any issue occurs during the network checking and creation process.
func (dm *DockerManager) CheckOrCreateNetwork(ctx context.Context, networkName string) error {
	log.Infof("Checking network '%s'", networkName)

	networkList, err := dm.Cli.NetworkList(ctx, types.NetworkListOptions{})
	if err != nil {
		log.Errorf("Getting list of networks error: %s\n", err)
		return err
	}

	for _, network := range networkList {
		if network.Name == networkName {
			log.Printf("Network '%s' already exists\n", networkName)
			return nil
		}
	}

	log.Infof("Creating network '%s'\n", networkName)

	_, err = dm.Cli.NetworkCreate(ctx, networkName, types.NetworkCreate{})
	if err != nil {
		log.Errorf("Creating Docker network error: %s", err)
		return err
	}

	return nil
}

// CheckIfProcessIsRunningInContainer checks if a process with the specified name is running inside a container.
// ctx: The context for the operation.
// processName: The name of the process to check.
// containerName: The name of the container.
// Returns a boolean indicating if the process is running, the output of the process, and an error if any issue occurs.
func (dm *DockerManager) CheckIfProcessIsRunningInContainer(ctx context.Context, processName, containerName string) (bool, string, error) {
	log.Infof("Checking if sekaid is running inside a '%s' container", containerName)
	// Create exec configuration
	execConfig := types.ExecConfig{
		Cmd:          []string{"sh", "-c", fmt.Sprintf("pgrep %s", processName)},
		AttachStdout: true,
		AttachStderr: true,
		Detach:       false,
		Tty:          false,
	}

	// Create exec
	resp, err := dm.Cli.ContainerExecCreate(ctx, containerName, execConfig)
	if err != nil {
		return false, "", err
	}

	// Attach to exec
	attach, err := dm.Cli.ContainerExecAttach(ctx, resp.ID, types.ExecStartCheck{})
	if err != nil {
		return false, "", err
	}
	defer attach.Close()

	// Create buffers to save stdout and stderr
	var stdout, stderr bytes.Buffer

	// Use stdcopy to demultiplex attach.Reader into stdout and stderr
	if _, err = stdcopy.StdCopy(&stdout, &stderr, attach.Reader); err != nil {
		return false, "", err
	}

	output := stdout.String()
	if errOutput := stderr.String(); errOutput != "" {
		fmt.Println("Stderr:", errOutput)
	}

	if strings.TrimSpace(output) != "" {
		log.Infof("Process with name '%s' running inside '%s' container with id: %s ", processName, containerName, string(output))
	} else {
		log.Infof("Process with name '%s' is not running inside '%s' container ", processName, containerName)
	}
	// If the output is not empty, the process is running
	return strings.TrimSpace(output) != "", string(output), nil
}

// GetInspectOfContainer inspects the Docker container with the given containerIdentification and returns
// the detailed information in the form of types.ContainerJSON struct.
// The containerIdentification parameter is the identifier of the container to inspect, such as the container ID or name.
// The function returns the docker package types.ContainerJSON struct containing the detailed information about the container,
// or an error if the inspection fails.
func (dm *DockerManager) GetInspectOfContainer(ctx context.Context, containerIdentification string) (types.ContainerJSON, error) {
	log.Infof("Inspecting container '%s'\n", containerIdentification)

	containerInfo, err := dm.Cli.ContainerInspect(ctx, containerIdentification)
	if err != nil {
		log.Errorf("Inspection container error: %s\n", err)
		return types.ContainerJSON{}, err
	}

	return containerInfo, nil
}

// GetNetworksInfo retrieves the list of Docker networks and returns the information
// about each network as a slice of types.NetworkResource containing the information about each network,
// or an error if there was a problem retrieving the networks.
func (dm *DockerManager) GetNetworksInfo(ctx context.Context) ([]types.NetworkResource, error) {
	resources, err := dm.Cli.NetworkList(ctx, types.NetworkListOptions{})
	if err != nil {
		log.Errorf("Getting networks info error: %s\n", err)
		return nil, err
	}

	return resources, nil
}
