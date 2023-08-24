package osutils

import (
	"os/exec"
	"strconv"

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
