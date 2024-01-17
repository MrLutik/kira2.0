package osutils

import (
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

var log = logging.Log

func CopyFile(src, dst string) error {
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

func CreateDirPath(dirPath string) error {
	log.Debugf("Creating dir path: %s", dirPath)
	err := os.MkdirAll(dirPath, 0o755) // 0755 are the standard permissions for directories.
	if err != nil {
		return err
	}
	return nil
}

func CheckIfFileExist(filePath string) (bool, error) {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false, nil
	}
	log.Debugf("Checking if '%s' exist: %t", filePath, !info.IsDir())
	return !info.IsDir(), nil
}

func GetCurrentOSUser() *user.User {
	// Getting current user home folder even if it run by sudo
	sudoUser := os.Getenv("SUDO_USER")

	if sudoUser != "" {
		usr, err := user.Lookup(sudoUser)
		if err != nil {
			// TODO leave panic?
			panic(err)
		}
		log.Debugf("Getting current user: %+v\n", usr)
		return usr
	} else {
		// Fallback to the current user if not running via sudo
		usr, err := user.Current()
		if err != nil {
			// TODO leave panic?
			panic(err)
		}
		log.Debugf("Getting current user: %+v\n", usr)
		return usr
	}
}

func CheckItPathExist(path string) (bool, error) {
	log.Debugf("Checking if path exist: %s\n", path)

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CheckIfIPIsValid(input string) (bool, error) {
	ipCheck := net.ParseIP(input)
	if ipCheck == nil {
		return false, &InvalidIPError{Input: input}
	}
	return true, nil
}

// CheckIfPortIsValid checks if the input string is a valid port (0-65535) and returns true if valid, false otherwise.
func CheckIfPortIsValid(input string) bool {
	port, err := strconv.Atoi(input)
	if err != nil {
		return false
	}

	if port < 0 || port > 65535 {
		return false
	}

	return true
}

// Run command
func RunCommand(command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	log.Debugf("Running: %s ", cmd.String())
	output, err := cmd.CombinedOutput()
	if err != nil {
		return output, err
	}
	return output, err
}

func RunCommandV2(commandStr string) ([]byte, error) {
	args, err := shlex.Split(commandStr)
	if err != nil {
		return []byte{}, err
	}
	log.Debugf("Running: %s ", commandStr)
	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return out, err
	}
	return (out), nil
}

func GetInternetInterface() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp != 0 && iface.Name != "lo" && hasInternetAccess(iface.Name) {
			return iface.Name
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
	return err == nil && !strings.Contains(string(out), "100% packet loss")
}

// accepts file path with file name for example "/tmp/hello.txt" and data to write in []byte
func CreateFileWithData(filePath string, data []byte) error {
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
