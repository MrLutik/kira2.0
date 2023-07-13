package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

type ContainerManager struct {
	Cli *client.Client
}

func NewTestContainerManager() (*ContainerManager, error) {
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return &ContainerManager{Cli: client}, err
}

// CheckForContainersName checks if a container with the specified name exists.
// ctx: The context for the operation.
// containerNameToCheck: The name of the container to check.
// Returns true if a container with the specified name is found, false otherwise, and an error if any issue occurs during the process.
func (dm *ContainerManager) CheckForContainersName(ctx context.Context, containerNameToCheck string) (bool, error) {
	log.Infof("Checking container with name: %s", containerNameToCheck)

	containers, err := dm.Cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		log.Errorf("Cannot get the list of containers: %s", err)
		return false, err
	}

	for _, c := range containers {
		for _, name := range c.Names {
			if name == `/`+containerNameToCheck {
				log.Infof("Container '%s' detected", name)
				return true, nil
			}
		}
	}

	log.Infof("Container '%s' is not detected", containerNameToCheck)
	return false, nil
}

// CheckIfProcessIsRunningInContainer checks if a process with the specified name is running inside a container.
// ctx: The context for the operation.
// processName: The name of the process to check.
// containerName: The name of the container.
// Returns a boolean indicating if the process is running, the output of the process, and an error if any issue occurs.
func (dm *ContainerManager) CheckIfProcessIsRunningInContainer(ctx context.Context, processName, containerName string) (bool, string, error) {
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
		log.Infof("Process with name '%s' running inside '%s' container with id: %s", processName, containerName, string(output))
	} else {
		log.Infof("Process with name '%s' is not running inside '%s' container ", processName, containerName)
	}
	// If the output is not empty, the process is running
	return strings.TrimSpace(output) != "", string(output), nil
}

// ExecCommandInContainerInDetachMode runs a command inside a specified container in detach mode.
// ctx: The context for the operation.
// containerID: The ID or name of the container.
// command: The command to execute inside the container.
// Returns the output of the command as a byte slice and an error if any issue occurs during the command execution.
func (dm *ContainerManager) ExecCommandInContainerInDetachMode(ctx context.Context, containerID string, command []string) ([]byte, error) {
	log.Infof("Running command '%s' in detach mode in '%s'", strings.Join(command, " "), containerID)

	execCreateResponse, err := dm.Cli.ContainerExecCreate(ctx, containerID, types.ExecConfig{
		Cmd:          command,
		AttachStdout: false,
		AttachStderr: false,
		Detach:       true,
	})
	if err != nil {
		log.Errorf("Exec configuration error: %s", err)
		return nil, err
	}

	execAttachConfig := types.ExecStartCheck{}
	resp, err := dm.Cli.ContainerExecAttach(ctx, execCreateResponse.ID, execAttachConfig)
	if err != nil {
		log.Errorf("Attaching to container '%s' error: %s", containerID, err)
		return nil, err
	}
	defer resp.Close()

	var outBuf, errBuf bytes.Buffer
	_, err = stdcopy.StdCopy(&outBuf, &errBuf, resp.Reader)
	if err != nil {
		log.Printf("Reading response error: %s", err)
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
func (dm *ContainerManager) ExecCommandInContainer(ctx context.Context, containerID string, command []string) ([]byte, error) {
	log.Infof("Running command '%s' in '%s'", strings.Join(command, " "), containerID)

	execCreateResponse, err := dm.Cli.ContainerExecCreate(ctx, containerID, types.ExecConfig{
		Cmd:          command,
		AttachStdout: true,
		AttachStderr: true,
	})
	if err != nil {
		log.Errorf("Exec configuration error: %s", err)
		return nil, err
	}

	execAttachConfig := types.ExecStartCheck{}
	resp, err := dm.Cli.ContainerExecAttach(ctx, execCreateResponse.ID, execAttachConfig)
	if err != nil {
		log.Errorf("Attaching to container '%s' error: %s", containerID, err)
		return nil, err
	}
	defer resp.Close()

	var outBuf, errBuf bytes.Buffer
	_, err = stdcopy.StdCopy(&outBuf, &errBuf, resp.Reader)
	if err != nil {
		log.Errorf("Reading response error: %s", err)
		return nil, err
	}

	output := outBuf.Bytes()
	log.Infof("Running '%s' successfully", strings.Join(command, " "))
	return output, err
}

// GetFileFromContainer allows you to retrieve a file from a Docker container and save it to the host machine.
// ctx (context.Context): The context for the operation.
// filePathOnHostMachine (string): The file path on the host machine where the file will be saved.
// filePathOnContainer (string): The file path inside the Docker container from which the file will be retrieved.
// containerID (string): The ID or name of the Docker container from which the file will be retrieved.
func (dm *ContainerManager) GetFileFromContainer(ctx context.Context, filePathOnHostMachine, filePathOnContainer, containerID string) error {
	log.Infof("Getting file from container '%s' to '%s'", filePathOnContainer, filePathOnHostMachine)
	rc, _, err := dm.Cli.CopyFromContainer(ctx, containerID, filePathOnContainer)
	if err != nil {
		log.Errorf("Error during copying file from container: %s", err)
		return err
	}
	defer rc.Close()

	contents, err := io.ReadAll(rc)
	if err != nil {
		log.Errorf("Reading error: %s", err)
		return err
	}

	err = os.WriteFile(filePathOnHostMachine, contents, 0o644)
	if err != nil {
		log.Errorf("Writing file error: %s", err)
		return err
	}

	log.Infof("Successfully got file '%s' to the host!", filePathOnHostMachine)

	return nil
}

// GetInspectOfContainer inspects the Docker container with the given containerIdentification and returns
// the detailed information in the form of types.ContainerJSON struct.
// The containerIdentification parameter is the identifier of the container to inspect, such as the container ID or name.
// The function returns the docker package types.ContainerJSON struct containing the detailed information about the container,
// or an error if the inspection fails.
func (dm *ContainerManager) GetInspectOfContainer(ctx context.Context, containerIdentification string) (types.ContainerJSON, error) {
	log.Infof("Inspecting container '%s'", containerIdentification)

	containerInfo, err := dm.Cli.ContainerInspect(ctx, containerIdentification)
	if err != nil {
		log.Errorf("Inspection container error: %s", err)
		return types.ContainerJSON{}, err
	}

	return containerInfo, nil
}

// InitAndCreateContainer initializes and creates a new container.
// ctx: The context for the operation.
// containerConfig: The container configuration.
// networkConfig: The network configuration.
// hostConfig: The host configuration.
// containerName: The name of the container.
// Returns an error if any issue occurs during the container initialization and creation process.
func (dm *ContainerManager) InitAndCreateContainer(
	ctx context.Context,
	containerConfig *container.Config,
	networkConfig *network.NetworkingConfig,
	hostConfig *container.HostConfig,
	containerName string,
) error {
	log.Infof("Starting new container '%s'", containerName)

	resp, err := dm.Cli.ContainerCreate(ctx, containerConfig, hostConfig, networkConfig, nil, containerName)
	if err != nil {
		log.Errorf("Creating container error: %s", err)
		return err
	}

	err = dm.Cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		log.Errorf("Starting container error: %s", err)
		return err
	}

	log.Infof("'%s' container started successfully! ID: %s", containerName, resp.ID)
	return err
}

// InstallDebPackage installs a Debian package (.deb) inside a specified container.
// ctx: The context for the operation.
// containerID: The ID or name of the container where the package will be installed.
// debDestPath: The destination path of the .deb package inside the container.
// Returns an error if any issue occurs during the package installation process.
func (dm *ContainerManager) InstallDebPackage(ctx context.Context, containerID, debDestPath string) error {
	log.Infof("Installing '%s'", debDestPath)

	installCmd := []string{"dpkg", "-i", debDestPath}
	execOptions := types.ExecConfig{
		Cmd:          installCmd,
		AttachStdout: true,
		AttachStderr: true,
	}
	resp, err := dm.Cli.ContainerExecCreate(ctx, containerID, execOptions)
	if err != nil {
		log.Errorf("Creating exec configuration error: %s", err)
		return err
	}

	attachOptions := types.ExecStartCheck{
		Detach: false,
		Tty:    false,
	}
	respConn, err := dm.Cli.ContainerExecAttach(ctx, resp.ID, attachOptions)
	if err != nil {
		log.Errorf("Attaching error: %s", err)
		return err
	}
	defer respConn.Close()

	// Capture the output from the container
	output, err := io.ReadAll(respConn.Reader)
	if err != nil {
		log.Errorf("Reading output error: %s", err)
		return err
	}

	// Wait for the execution to complete
	waitResponse, err := dm.Cli.ContainerExecInspect(ctx, resp.ID)
	if err != nil {
		log.Errorf("Inspecting process '%s' error: %s", resp.ID, err)
		return err
	}

	if waitResponse.ExitCode != 0 {
		err = fmt.Errorf("package installation failed: %s", string(output))
		log.Errorf("Installation error: %s", err)
		return err
	}

	log.Infof("Package '%s' installed successfully", debDestPath)

	return nil
}

// SendFileToContainer sends a file from the host machine to a specified directory inside a Docker container.
// ctx: The context for the operation.
// filePathOnHostMachine: The path of the file on the host machine.
// directoryPathOnContainer: The path of the directory inside the container where the file will be copied.
// containerID: The ID or name of the Docker container.
// Returns an error if any issue occurs during the file sending process.
func (dm *ContainerManager) SendFileToContainer(ctx context.Context, filePathOnHostMachine, directoryPathOnContainer, containerID string) error {
	log.Infof("Sending file '%s' to container '%s' to '%s'", filePathOnHostMachine, containerID, directoryPathOnContainer)
	file, err := os.Open(filePathOnHostMachine)
	if err != nil {
		log.Errorf("Opening file error: %s", err)
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Errorf("Can't open file stat: %s", err)
		return err
	}

	var buf bytes.Buffer
	tarWriter := tar.NewWriter(&buf)

	err = addFileToTar(fileInfo, file, tarWriter)
	if err != nil {
		log.Errorf("Adding file to tar error: %s", err)
		return err
	}

	err = tarWriter.Close()
	if err != nil {
		log.Errorf("Closing tar error: %s", err)
		return err
	}

	tarContent := buf.Bytes()
	tarReader := bytes.NewReader(tarContent)
	copyOptions := types.CopyToContainerOptions{
		AllowOverwriteDirWithFile: false,
	}

	err = dm.Cli.CopyToContainer(ctx, containerID, directoryPathOnContainer, tarReader, copyOptions)
	if err != nil {
		log.Errorf("Copying tar to container error: %s", err)
		return err
	}

	log.Infof("Successfully copied '%s' to '%s' in '%s' container", filePathOnHostMachine, directoryPathOnContainer, containerID)
	return nil
}

// addFileToTar adds a file to a tar archive.
// fileInfo: The file information.
// file: The reader for the file data.
// tarWriter: The tar writer.
// Returns an error if any issue occurs during the file writing process.
func addFileToTar(fileInfo os.FileInfo, file io.Reader, tarWriter *tar.Writer) error {
	log.Infof("Writing file '%s' to tar archive", fileInfo.Name())

	header := &tar.Header{
		Name: fileInfo.Name(),
		Mode: int64(fileInfo.Mode()),
		Size: fileInfo.Size(),
	}

	if err := tarWriter.WriteHeader(header); err != nil {
		log.Errorf("Writing tar header error: %s", err)
		return err
	}

	if _, err := io.Copy(tarWriter, file); err != nil {
		log.Errorf("Copying error: %s", err)
		return err
	}

	return nil
}

// StopAndDeleteContainer stops and deletes a container with the specified name.
// ctx: The context for the operation.
// containerNameToStop: The name of the container to stop and delete.
// Returns an error if any issue occurs during the process.
func (dm *ContainerManager) StopAndDeleteContainer(ctx context.Context, containerNameToStop string) error {
	log.Infof("Stopping '%s' container...", containerNameToStop)

	err := dm.Cli.ContainerStop(ctx, containerNameToStop, container.StopOptions{})
	if err != nil {
		log.Errorf("Stopping container error: %s", err)
		return err
	}

	log.Infof("Deleting %s container...", containerNameToStop)
	err = dm.Cli.ContainerRemove(ctx, containerNameToStop, types.ContainerRemoveOptions{})
	if err != nil {
		log.Println(err)
		return err
	}

	log.Infof("Container %s is deleted", containerNameToStop)
	return nil
}
