package deploy

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

func installDocker(client *ssh.Client) error {
	session, err := client.NewSession()
	if err != nil {
		log.Errorln("Failed to create session: ", err)
		return fmt.Errorf("Failed to create session: %v", err)
	}
	defer session.Close()

	// Check if Docker is already installed
	_, err = session.Output("docker -v")
	if err == nil {
		log.Println("Docker is already installed")
		return nil
	}

	// Install Docker
	installCmd := "curl -fsSL https://get.docker.com -o get-docker.sh && sh get-docker.sh"
	_, err = session.Output(installCmd)
	if err != nil {
		log.Errorln("Failed to install Docker: ", err)
		return fmt.Errorf("Failed to install Docker: %v", err)
	}

	return nil
}

// checkDocker checks if Docker is installed on the remote machine and returns its version or an error.
func checkDocker(client *ssh.Client) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("Failed to create SSH session: %v", err)
	}
	defer session.Close()

	// Execute 'docker --version' command
	log.Info("Checking if Docker is installed on the remote machine...")
	dockerVersionBytes, err := session.Output("docker --version")
	if err != nil {
		return "", fmt.Errorf("Docker not found on the remote machine: %v", err)
	}

	dockerVersion := strings.Trim(string(dockerVersionBytes), "\n")

	return dockerVersion, nil
}
