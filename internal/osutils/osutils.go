package osutils

import (
	"net"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/mrlutik/kira2.0/internal/logging"
)

var log = logging.Log

// Checks if input string is a valid port  0-65535
func CheckIfPortIsValid(input string) bool {
	// Convert string to integer
	port, err := strconv.Atoi(input)
	if err != nil {
		return false
	}

	// Check if the port is in the valid range
	return port >= 0 && port <= 65535
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
