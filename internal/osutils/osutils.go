package osutils

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"strconv"
	"strings"

	"github.com/google/shlex"
	"github.com/mrlutik/kira2.0/internal/logging"
)

type OSUtils struct {
	log *logging.Logger
}

var ErrNotAFile = errors.New("path exists but is not a file")

// CopyFile copies a file from a source path to a destination path.
// - src: Source file path.
// - dst: Destination file path.
// Returns an error if the copying process fails.
func (o *OSUtils) CopyFile(src, dst string) error {
	o.log.Debugf("Copying '%s' to '%s", src, dst)

	// Open source file for reading
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create destination file for writing
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	// Copy the contents from srcFile to dstFile
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

// CreateDirPath creates a directory at the specified path including any necessary parent directories.
// - dirPath: The path of the directory to create.
// Returns an error if the directory creation fails.
func (o *OSUtils) CreateDirPath(dirPath string) error {
	o.log.Debugf("Creating dir path: %s", dirPath)
	err := os.MkdirAll(dirPath, 0o755) // 0755 are the standard permissions for directories.
	if err != nil {
		return err
	}
	return nil
}

// CheckIfFileExist checks whether a file exists at the given file path.
// - filePath: Path of the file to check.
// Returns true if the file exists, false otherwise, along with an error if the check fails.
func (o *OSUtils) CheckIfFileExist(filePath string) (bool, error) {
	o.log.Debugf("Checking if '%s' exist", filePath)

	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	if info.IsDir() {
		return false, ErrNotAFile
	}

	return true, nil
}

// GetCurrentOSUser retrieves the current OS user. If the program is run with sudo, it gets the user who invoked sudo.
// Returns the user and an error if the retrieval fails.
func (o *OSUtils) GetCurrentOSUser() (*user.User, error) {
	// Getting current user home folder even if it run by sudo
	sudoUser := os.Getenv("SUDO_USER")

	if sudoUser != "" {
		usr, err := user.Lookup(sudoUser)
		if err != nil {
			return nil, err
		}
		o.log.Debugf("Getting current user: %+v", usr)
		return usr, nil
	}

	// Fallback to the current user if not running via sudo
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}
	o.log.Debugf("Getting current user: %+v", usr)
	return usr, nil
}

// CheckIfPathExists checks whether a path exists in the file system.
// - path: The path to check.
// Returns true if the path exists, false otherwise, along with an error if the check fails.
func (o *OSUtils) CheckIfPathExists(path string) (bool, error) {
	o.log.Debugf("Checking if path '%s' exists", path)

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CheckIfIPIsValid checks whether the given string is a valid IP address.
// - input: The IP address string to validate.
// Returns true if the IP address is valid, false otherwise, along with an error if the validation fails.
func (o *OSUtils) CheckIfIPIsValid(input string) (bool, error) {
	o.log.Debugf("Checking if '%s' is valid IP address", input)

	ipCheck := net.ParseIP(input)
	if ipCheck == nil {
		return false, &InvalidIPError{Input: input}
	}

	return true, nil
}

// CheckIfPortIsValid checks whether the given string is a valid port number (0-65535).
// - input: The port number string to validate.
// Returns true if the port number is valid, false otherwise.
func (o *OSUtils) CheckIfPortIsValid(input string) bool {
	o.log.Debugf("Checking if '%s' is valid port", input)

	port, err := strconv.Atoi(input)
	if err != nil {
		return false
	}

	if port < 0 || port > 65535 {
		return false
	}

	return true
}

// RunCommand executes a command with the provided arguments and returns the combined standard output and standard error.
// - command: The command to execute.
// - args: Arguments for the command.
// Returns the output of the command and an error if the execution fails.
func (o *OSUtils) RunCommand(command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	o.log.Debugf("Running: %s", cmd.String())

	output, err := cmd.CombinedOutput()
	if err != nil {
		return output, err
	}
	return output, err
}

// RunCommandV2 executes a command given as a single string and returns the combined standard output and standard error.
// - commandStr: The entire command with arguments as a single string.
// Returns the output of the command and an error if the execution fails.
func (o *OSUtils) RunCommandV2(commandStr string) ([]byte, error) {
	o.log.Debugf("Running: %s", commandStr)

	args, err := shlex.Split(commandStr)
	if err != nil {
		return []byte{}, err
	}

	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return out, err
	}
	return (out), nil
}

// GetInternetInterface attempts to find an active internet interface on the system.
// Returns the name of the internet interface if found, an empty string otherwise.
func (o *OSUtils) GetInternetInterface() string {
	o.log.Debug("Getting internet interfaces")

	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, internetInterface := range interfaces {
		if internetInterface.Flags&net.FlagUp != 0 && internetInterface.Name != "lo" && hasInternetAccess(internetInterface.Name) {
			return internetInterface.Name
		}
	}

	return ""
}

func hasInternetAccess(ifaceName string) bool {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("ping", "-n", "1", "-w", "2000", "8.8.8.8")
	case "darwin":
		cmd = exec.Command("ping", "-c", "1", "-W", "2000", "8.8.8.8")
	default:
		cmd = exec.Command("ping", "-c", "1", "-W", "2", "-I", ifaceName, "8.8.8.8")
	}

	out, err := cmd.CombinedOutput()

	// TODO "100% package loss" case does not cover bad connection issues
	return err == nil && !strings.Contains(string(out), "100%% packet loss")
}

// CreateFileWithData creates a new file at the specified path and writes the provided data into it.
// - filePath: The path where the file will be created.
// - data: The byte array of data to write into the file.
// Returns an error if the file creation or data writing fails.
func (o *OSUtils) CreateFileWithData(filePath string, data []byte) error {
	o.log.Debugf("Creating file '%s' with provided data", filePath)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write data to file: %w", err)
	}

	return nil
}
