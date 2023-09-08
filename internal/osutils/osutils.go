package osutils

import (
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/google/shlex"
	"github.com/mrlutik/kira2.0/internal/logging"
)

var log = logging.Log

func CheckIfIPIsValid(input string) (bool, error) {
	ipCheck := net.ParseIP(input)
	if ipCheck == nil {
		return false, fmt.Errorf("<%s> is not a valid ip", input)
	}
	return true, nil
}

// Checks if input string is a valid port  0-65535
func CheckIfPortIsValid(input string) (bool, error) {
	// Convert string to integer
	port, err := strconv.Atoi(input)
	if err != nil {
		return false, err
	}
	// Check if the port is in the valid range
	if port < 0 || port > 65535 {
		return false, fmt.Errorf("%v port in not in valid range", port)
	}

	return true, nil
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

func RunCommandV2(commandStr string) (string, error) {
	log.Printf("RUNNING V2 COMMAND RUNNER\n")
	args, err := shlex.Split(commandStr)
	if err != nil {
		return "", err
	}
	log.Debugf("Running: %s ", commandStr)
	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(out), nil
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
