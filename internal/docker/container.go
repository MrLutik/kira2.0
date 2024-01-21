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
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/logging"
)

type (
	ContainerManager struct {
		cli *client.Client
		log *logging.Logger
	}
)

func NewTestContainerManager(dockerClient *client.Client, logger *logging.Logger) *ContainerManager {
	return &ContainerManager{
		cli: dockerClient,
		log: logger,
	}
}

// CheckForContainersName checks if a container with the specified name exists.
// ctx: The context for the operation.
// containerNameToCheck: The name of the container to check.
// Returns true if a container with the specified name is found, false otherwise, and an error if any issue occurs during the process.
func (c *ContainerManager) CheckForContainersName(ctx context.Context, containerNameToCheck string) (bool, error) {
	c.log.Infof("Checking container with name: %s", containerNameToCheck)

	containers, err := c.cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		c.log.Errorf("Cannot get the list of containers: %s", err)
		return false, err
	}

	for _, container := range containers {
		for _, name := range container.Names {
			if name == `/`+containerNameToCheck {
				c.log.Infof("Container '%s' detected", name)
				return true, nil
			}
		}
	}

	c.log.Infof("Container '%s' is not detected", containerNameToCheck)
	return false, nil
}

// CheckIfProcessIsRunningInContainer checks if a process with the specified name is running inside a container.
// - ctx: The context for the operation.
// - processName: The name of the process to check.
// - containerName: The name of the container.
// Returns a boolean indicating if the process is running, the output of the process, and an error if any issue occurs.
func (c *ContainerManager) CheckIfProcessIsRunningInContainer(ctx context.Context, processName, containerName string) (bool, string, error) {
	c.log.Infof("Checking if sekaid is running inside a '%s' container", containerName)
	// Create exec configuration
	execConfig := types.ExecConfig{
		Cmd:          []string{"sh", "-c", fmt.Sprintf("pgrep %s", processName)},
		AttachStdout: true,
		AttachStderr: true,
		Detach:       false,
		Tty:          false,
	}

	// Create exec
	resp, err := c.cli.ContainerExecCreate(ctx, containerName, execConfig)
	if err != nil {
		return false, "", err
	}

	// Attach to exec
	attach, err := c.cli.ContainerExecAttach(ctx, resp.ID, types.ExecStartCheck{})
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

	if errOutput := stderr.String(); errOutput != "" {
		c.log.Infof("Stderr: %s", errOutput)
		return false, "", fmt.Errorf("%w:\nOutput: %s", ErrStderrNotEmpty, errOutput)
	}

	output := stdout.String()
	if strings.TrimSpace(output) != "" {
		c.log.Infof("Process with name '%s' running inside '%s' container with id: %s", processName, containerName, string(output))
	} else {
		c.log.Infof("Process with name '%s' is not running inside '%s' container ", processName, containerName)
	}
	// If the output is not empty, the process is running
	return strings.TrimSpace(output) != "", string(output), nil
}

// ExecCommandInContainerInDetachMode runs a command inside a specified container in detach mode.
// ctx: The context for the operation.
// containerID: The ID or name of the container.
// command: The command to execute inside the container.
// Returns the output of the command as a byte slice and an error if any issue occurs during the command execution.
func (c *ContainerManager) ExecCommandInContainerInDetachMode(ctx context.Context, containerID string, command []string) ([]byte, error) {
	c.log.Infof("Running command '%s' in detach mode in '%s'", strings.Join(command, " "), containerID)

	execCreateResponse, err := c.cli.ContainerExecCreate(ctx, containerID, types.ExecConfig{
		Cmd:          command,
		AttachStdout: false,
		AttachStderr: false,
		Detach:       true,
	})
	if err != nil {
		c.log.Errorf("Exec configuration error: %s", err)
		return nil, err
	}

	execAttachConfig := types.ExecStartCheck{}
	resp, err := c.cli.ContainerExecAttach(ctx, execCreateResponse.ID, execAttachConfig)
	if err != nil {
		c.log.Errorf("Attaching to container '%s' error: %s", containerID, err)
		return nil, err
	}
	defer resp.Close()

	var outBuf, errBuf bytes.Buffer
	_, err = stdcopy.StdCopy(&outBuf, &errBuf, resp.Reader)
	if err != nil {
		c.log.Errorf("Reading response error: %s", err)
		return nil, err
	}

	output := outBuf.Bytes()

	c.log.Infoln("Reading successfully")
	return output, err
}

// ExecCommandInContainer executes a command inside a specified container.
// ctx: The context for the operation.
// containerID: The ID or name of the container.
// command: The command to execute inside the container.
// Returns the output of the command as a byte slice and an error if any issue occurs during the command execution.
func (c *ContainerManager) ExecCommandInContainer(ctx context.Context, containerID string, command []string) ([]byte, error) {
	c.log.Infof("Running command '%s' in '%s'", strings.Join(command, " "), containerID)

	execCreateResponse, err := c.cli.ContainerExecCreate(ctx, containerID, types.ExecConfig{
		Cmd:          command,
		AttachStdout: true,
		AttachStderr: true,
	})
	if err != nil {
		c.log.Errorf("Exec configuration error: %s", err)
		return nil, err
	}

	execAttachConfig := types.ExecStartCheck{}
	resp, err := c.cli.ContainerExecAttach(ctx, execCreateResponse.ID, execAttachConfig)
	if err != nil {
		c.log.Errorf("Attaching to container '%s' error: %s", containerID, err)
		return nil, err
	}
	defer resp.Close()

	var outBuf, errBuf bytes.Buffer
	_, err = stdcopy.StdCopy(&outBuf, &errBuf, resp.Reader)
	if err != nil {
		c.log.Errorf("Reading response error: %s", err)
		return nil, err
	}

	output := outBuf.Bytes()
	c.log.Infof("Running '%s' successfully", strings.Join(command, " "))
	return output, err
}

// readTarArchive reads a file from the TAR archive stream
// and returns the file content as a byte slice.
func readTarArchive(tr *tar.Reader, fileName string) ([]byte, error) {
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if hdr.Name == fileName {
			b, err := io.ReadAll(tr)
			if err != nil {
				return nil, err
			}
			return b, nil
		}
	}
	return nil, fmt.Errorf("%w: %s", ErrFileNotFoundInTarBase, fileName)
}

// GetFileFromContainer retrieves a file from a specified container using the Docker API.
// It copies the TAR archive with file from the specified folder path in the container,
// read file from TAR archive and returns the file content as a byte slice.
func (c *ContainerManager) GetFileFromContainer(ctx context.Context, folderPathOnContainer, fileName, containerID string) ([]byte, error) {
	c.log.Infof("Getting file '%s' from container '%s'", fileName, folderPathOnContainer)

	rc, _, err := c.cli.CopyFromContainer(ctx, containerID, folderPathOnContainer+"/"+fileName)
	if err != nil {
		c.log.Errorf("Copying from container error: %s", err)
		return nil, err
	}
	defer rc.Close()

	tr := tar.NewReader(rc)
	b, err := readTarArchive(tr, fileName)
	if err != nil {
		c.log.Errorf("Reading Tar archive error: %s", err)
		return nil, err
	}

	return b, nil
}

// GetInspectOfContainer inspects the Docker container with the given containerIdentification and returns
// the detailed information in the form of types.ContainerJSON struct.
// The containerIdentification parameter is the identifier of the container to inspect, such as the container ID or name.
// The function returns the docker package types.ContainerJSON struct containing the detailed information about the container,
// or an error if the inspection fails.
func (c *ContainerManager) GetInspectOfContainer(ctx context.Context, containerIdentification string) (*types.ContainerJSON, error) {
	c.log.Infof("Inspecting container '%s'", containerIdentification)

	containerInfo, err := c.cli.ContainerInspect(ctx, containerIdentification)
	if err != nil {
		c.log.Errorf("Inspection container error: %s", err)
		return &types.ContainerJSON{}, err
	}

	return &containerInfo, nil
}

// InitAndCreateContainer initializes and creates a new container.
// - ctx: The context for the operation.
// - containerConfig: The container configuration.
// - networkConfig: The network configuration.
// - hostConfig: The host configuration.
// - containerName: The name of the container.
// Returns an error if any issue occurs during the container initialization and creation process.
func (c *ContainerManager) InitAndCreateContainer(
	ctx context.Context,
	containerConfig *container.Config,
	networkConfig *network.NetworkingConfig,
	hostConfig *container.HostConfig,
	containerName string,
) error {
	c.log.Infof("Starting new container '%s'", containerName)

	resp, err := c.cli.ContainerCreate(ctx, containerConfig, hostConfig, networkConfig, nil, containerName)
	if err != nil {
		c.log.Errorf("Creating container error: %s", err)
		return err
	}

	err = c.cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		c.log.Errorf("Starting container error: %s", err)
		return err
	}

	c.log.Infof("'%s' container started successfully! ID: %s", containerName, resp.ID)
	return err
}

func (c *ContainerManager) StartContainer(ctx context.Context, containerName string) error {
	err := c.cli.ContainerStart(ctx, containerName, types.ContainerStartOptions{})
	if err != nil {
		c.log.Errorf("Starting container error: %s", err)
		return err
	}

	return nil
}

func (c *ContainerManager) StopContainer(ctx context.Context, containerName string) error {
	err := c.cli.ContainerStop(ctx, containerName, container.StopOptions{})
	if err != nil {
		c.log.Errorf("Stopping container error: %s", err)
		return err
	}

	return nil
}

// InstallDebPackage installs a Debian package (.deb) inside a specified container.
// - ctx: The context for the operation.
// - containerID: The ID or name of the container where the package will be installed.
// - debDestPath: The destination path of the .deb package inside the container.
// Returns an error if any issue occurs during the package installation process.
func (c *ContainerManager) InstallDebPackage(ctx context.Context, containerID, debDestPath string) error {
	c.log.Infof("Installing '%s'", debDestPath)

	installCmd := []string{"dpkg", "-i", debDestPath}
	execOptions := types.ExecConfig{
		Cmd:          installCmd,
		AttachStdout: true,
		AttachStderr: true,
	}
	resp, err := c.cli.ContainerExecCreate(ctx, containerID, execOptions)
	if err != nil {
		c.log.Errorf("Creating exec configuration error: %s", err)
		return err
	}

	attachOptions := types.ExecStartCheck{
		Detach: false,
		Tty:    false,
	}
	respConn, err := c.cli.ContainerExecAttach(ctx, resp.ID, attachOptions)
	if err != nil {
		c.log.Errorf("Attaching error: %s", err)
		return err
	}
	defer respConn.Close()

	// Capture the output from the container
	output, err := io.ReadAll(respConn.Reader)
	if err != nil {
		c.log.Errorf("Reading output error: %s", err)
		return err
	}

	// Wait for the execution to complete
	waitResponse, err := c.cli.ContainerExecInspect(ctx, resp.ID)
	if err != nil {
		c.log.Errorf("Inspecting process '%s' error: %s", resp.ID, err)
		return err
	}

	if waitResponse.ExitCode != 0 {
		c.log.Errorf("Installation error: %s", string(output))
		return fmt.Errorf("%w:\nOutput: %s", ErrPackageInstallationFailed, string(output))
	}

	c.log.Infof("Package '%s' installed successfully", debDestPath)

	return nil
}

// WriteFileDataToContainer writes the provided fileData as a file with the given fileName into the specified container.
// It creates a tar archive containing the file data and sends it to the container using the Docker client's CopyToContainer method.
// The destination path in the container is determined by the destPath parameter.
func (c *ContainerManager) WriteFileDataToContainer(ctx context.Context, fileData []byte, fileName, destPath, containerID string) error {
	c.log.Infof("Writing file to container '%s'", containerID)

	tarBuffer := new(bytes.Buffer)
	tw := tar.NewWriter(tarBuffer)

	header := &tar.Header{
		Name: fileName,
		Mode: 0o644,
		Size: int64(len(fileData)),
	}
	if err := tw.WriteHeader(header); err != nil {
		c.log.Errorf("Writing tar header error: %s", err)
		return err
	}

	if _, err := tw.Write(fileData); err != nil {
		c.log.Errorf("Writing file data to tar error: %s", err)
		return err
	}

	if err := tw.Close(); err != nil {
		c.log.Errorf("Closing tar writer error: %s", err)
		return err
	}

	err := c.cli.CopyToContainer(ctx, containerID, destPath, tarBuffer, types.CopyToContainerOptions{
		AllowOverwriteDirWithFile: true,
	})
	if err != nil {
		c.log.Errorf("Failed to copy file to container '%s': %s", containerID, err)
		return err
	}

	c.log.Infof("File '%s' is successfully written on '%s' in container '%s'", fileName, destPath, containerID)

	return nil
}

// SendFileToContainer sends a file from the host machine to a specified directory inside a Docker container.
// - ctx: The context for the operation.
// - filePathOnHostMachine: The path of the file on the host machine.
// - directoryPathOnContainer: The path of the directory inside the container where the file will be copied.
// - containerID: The ID or name of the Docker container.
// Returns an error if any issue occurs during the file sending process.
func (c *ContainerManager) SendFileToContainer(ctx context.Context, filePathOnHostMachine, directoryPathOnContainer, containerID string) error {
	c.log.Infof("Sending file '%s' to container '%s' to '%s'", filePathOnHostMachine, containerID, directoryPathOnContainer)
	file, err := os.Open(filePathOnHostMachine)
	if err != nil {
		c.log.Errorf("Opening file error: %s", err)
		return err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		c.log.Errorf("Can't open file stat: %s", err)
		return err
	}

	var buf bytes.Buffer
	tarWriter := tar.NewWriter(&buf)

	c.log.Infof("Writing file '%s' to tar archive", fileInfo.Name())
	err = addFileToTar(fileInfo, file, tarWriter)
	if err != nil {
		c.log.Errorf("Adding file to tar error: %s", err)
		return err
	}

	err = tarWriter.Close()
	if err != nil {
		c.log.Errorf("Closing tar error: %s", err)
		return err
	}

	tarContent := buf.Bytes()
	tarReader := bytes.NewReader(tarContent)
	copyOptions := types.CopyToContainerOptions{
		AllowOverwriteDirWithFile: false,
	}

	err = c.cli.CopyToContainer(ctx, containerID, directoryPathOnContainer, tarReader, copyOptions)
	if err != nil {
		c.log.Errorf("Copying tar to container error: %s", err)
		return err
	}

	c.log.Infof("Successfully copied '%s' to '%s' in '%s' container", filePathOnHostMachine, directoryPathOnContainer, containerID)
	return nil
}

// addFileToTar adds a file to a tar archive.
// - fileInfo: The file information.
// - file: The reader for the file data.
// - tarWriter: The tar writer.
// Returns an error if any issue occurs during the file writing process.
func addFileToTar(fileInfo os.FileInfo, file io.Reader, tarWriter *tar.Writer) error {
	header := &tar.Header{
		Name: fileInfo.Name(),
		Mode: int64(fileInfo.Mode()),
		Size: fileInfo.Size(),
	}

	if err := tarWriter.WriteHeader(header); err != nil {
		return fmt.Errorf("writing tar header error: %w", err)
	}

	if _, err := io.Copy(tarWriter, file); err != nil {
		return fmt.Errorf("copying error: %w", err)
	}

	return nil
}

// StopAndDeleteContainer stops and deletes a container with the specified name.
// ctx: The context for the operation.
// containerNameToStop: The name of the container to stop and delete.
// Returns an error if any issue occurs during the process.
func (c *ContainerManager) StopAndDeleteContainer(ctx context.Context, containerNameToStop string) error {
	c.log.Infof("Stopping '%s' container...", containerNameToStop)

	err := c.cli.ContainerStop(ctx, containerNameToStop, container.StopOptions{})
	if err != nil {
		c.log.Errorf("Stopping container error: %s", err)
		return err
	}

	c.log.Infof("Deleting %s container...", containerNameToStop)
	err = c.cli.ContainerRemove(ctx, containerNameToStop, types.ContainerRemoveOptions{})
	if err != nil {
		c.log.Errorf("Removing container '%s' error: %s", containerNameToStop, err)
		return err
	}

	c.log.Infof("Container %s is deleted", containerNameToStop)
	return nil
}

// CheckForVolumeName is checking if docker volume with volumeName exist, if do - returns true
func (c *ContainerManager) CheckForVolumeName(ctx context.Context, volumeName string) (bool, error) {
	c.log.Info("Getting volumes list")
	volumes, err := c.cli.VolumeList(ctx, volume.ListOptions{})
	if err != nil {
		c.log.Errorf("cannot get list of volumes: %s", err)
		return false, err
	}
	c.log.Debugf("Volumes list %v\n", volumes.Volumes)

	for _, volume := range volumes.Volumes {
		c.log.Tracef("searching for %s, curent: %s\n", volumeName, volume.Name)
		if volume.Name == volumeName {
			c.log.Debugf("Volume with <%s> name was found\n", volumeName)
			return true, nil
		}
	}
	c.log.Debugf("Volume with <%s> name was not found\n", volumeName)
	return false, nil
}

// CleanupContainersAndVolumes is cleaning up container and volumes (needed for new node initializing),
// accepts only *KiraConfig and takes all values from it
func (c *ContainerManager) CleanupContainersAndVolumes(ctx context.Context, kiraCfg *config.KiraConfig) error {
	check, err := c.CheckForContainersName(ctx, kiraCfg.SekaidContainerName)
	if err != nil {
		return err
	}
	if check {
		err = c.StopAndDeleteContainer(ctx, kiraCfg.SekaidContainerName)
		if err != nil {
			return err
		}
	}
	check, err = c.CheckForContainersName(ctx, kiraCfg.InterxContainerName)
	if err != nil {
		return err
	}
	if check {
		err = c.StopAndDeleteContainer(ctx, kiraCfg.InterxContainerName)
		if err != nil {
			return err
		}
	}
	check, err = c.CheckForVolumeName(ctx, kiraCfg.VolumeName)
	if err != nil {
		return err
	}
	if check {
		c.log.Infof("Removing '%s' volume\n", kiraCfg.VolumeName)
		err = c.cli.VolumeRemove(ctx, kiraCfg.VolumeName, true)
		if err != nil {
			return err
		}
	}
	return nil
}

// StopProcessInsideContainer is checking if process is running inside container, then executing p-kill, then checking again if process exist
//
// processName - process to kill,
// codeToStopWith - signal code to stop with,
// containerName - container name in which pkill will be executed
func (c *ContainerManager) StopProcessInsideContainer(ctx context.Context, processName string, codeToStopWith int, containerName string) error {
	c.log.Infof("Checking if '%s' is running inside container", processName)
	check, _, err := c.CheckIfProcessIsRunningInContainer(ctx, processName, containerName)
	if err != nil {
		return fmt.Errorf("cant check if procces is running inside container, %w", err)
	}
	if !check {
		c.log.Warnf("process <%s> is not running inside <%s> container\n", processName, containerName)
		return nil
	}
	c.log.Infof("Stopping <%s> process\n", processName)
	out, err := c.ExecCommandInContainer(ctx, containerName, []string{"pkill", fmt.Sprintf("-%v", codeToStopWith), processName})
	if err != nil {
		c.log.Errorf("cannot kill <%s> process inside <%s> container\nout: %s\nerr: %v\n", processName, containerName, string(out), err)
		return fmt.Errorf("cannot kill <%s> process inside <%s> container\nout: %s\nerr: %w", processName, containerName, string(out), err)
	}

	check, _, err = c.CheckIfProcessIsRunningInContainer(ctx, processName, containerName)
	if err != nil {
		return fmt.Errorf("cant check if procces is running inside container, %w", err)
	}
	if check {
		c.log.Errorf("Process <%s> is still running inside <%s> container\n", processName, containerName)
		return err
	}
	c.log.Infof("<%s> process was successfully stopped\n", processName)
	return nil
}

func (c *ContainerManager) CloseClient() {
	c.cli.Close()
}
