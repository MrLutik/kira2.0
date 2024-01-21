package deploy

import (
	"fmt"

	"github.com/mrlutik/kira2.0/internal/logging"
	"golang.org/x/crypto/ssh"
)

func checkCosignInstalled(session *ssh.Session, log *logging.Logger) error {
	log.Info("Checking if Cosign is installed on the remote machine...")
	if err := session.Run("which cosign"); err != nil {
		log.Warn("Cosign not found on the remote machine.")
		return err
	}
	return nil
}

func installCosign(session *ssh.Session, log *logging.Logger) error {
	log.Info("Attempting to install Cosign...")
	if err := session.Run("wget -q -O- https://github.com/sigstore/cosign/releases/download/v1.2.1/cosign-linux-amd64 | sudo tee /usr/local/bin/cosign > /dev/null && sudo chmod +x /usr/local/bin/cosign"); err != nil {
		log.Error("Failed to install Cosign.")
		return err
	}
	log.Info("Cosign installed successfully.")
	return nil
}

func verifyCosignSignature(client *ssh.Client, pkgPath, sigPath string, log *logging.Logger) error {
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create SSH session: %w", err)
	}
	defer session.Close()

	// Check if cosign is installed
	if err := checkCosignInstalled(session, log); err != nil {
		if installErr := installCosign(session, log); installErr != nil {
			return fmt.Errorf("failed to install Cosign: %w", installErr)
		}
	}

	// Define your command string
	cmdString := fmt.Sprintf("cosign verify -key /path/to/cosign/key -signature %s %s", sigPath, pkgPath)

	// Run the command on the remote host
	log.Info("Verifying Cosign signature...")
	if err := session.Run(cmdString); err != nil {
		return fmt.Errorf("failed to verify Cosign signature: %w", err)
	}

	return nil
}
