package deploy

import (
	"fmt"
	"strings"

	"github.com/mrlutik/kira2.0/internal/logging"
	"golang.org/x/crypto/ssh"
)

func installDocker(client *ssh.Client, log *logging.Logger) error {
	session, err := client.NewSession()
	if err != nil {
		log.Errorln("Failed to create session: ", err)
		return fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Close()

	// Check if Docker is already installed
	_, err = session.Output("docker -v")
	if err == nil {
		log.Info("Docker is already installed")
		return nil
	}

	// Install Docker
	installCmd := "curl -fsSL https://get.docker.com -o get-docker.sh && sh get-docker.sh"
	_, err = session.Output(installCmd)
	if err != nil {
		log.Errorln("Failed to install Docker: ", err)
		return fmt.Errorf("failed to install Docker: %w", err)
	}

	return nil
}

// checkDocker checks if Docker is installed on the remote machine and returns its version or an error.
func checkDocker(client *ssh.Client, log *logging.Logger) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create SSH session: %w", err)
	}
	defer session.Close()

	// Execute 'docker --version' command
	log.Info("Checking if Docker is installed on the remote machine...")
	dockerVersionBytes, err := session.Output("docker --version")
	if err != nil {
		return "", fmt.Errorf("docker not found on the remote machine: %w", err)
	}

	dockerVersion := strings.Trim(string(dockerVersionBytes), "\n")

	return dockerVersion, nil
}
