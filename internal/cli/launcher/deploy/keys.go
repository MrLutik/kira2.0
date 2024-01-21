package deploy

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func installKeys(client *ssh.Client, _, pubKey string) error {
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Close()

	// Create .ssh directory if it doesn't exist
	_, err = session.Output("mkdir -p ~/.ssh")
	if err != nil {
		return fmt.Errorf("failed to create .ssh directory: %w", err)
	}

	// Write the public key to authorized_keys
	_, err = session.Output(fmt.Sprintf("echo '%s' >> ~/.ssh/authorized_keys", pubKey))
	if err != nil {
		return fmt.Errorf("failed to add public key to authorized_keys: %w", err)
	}

	// Set permissions for the .ssh directory and authorized_keys file
	_, err = session.Output("chmod 700 ~/.ssh; chmod 600 ~/.ssh/authorized_keys")
	if err != nil {
		return fmt.Errorf("failed to set permissions: %w", err)
	}

	return nil
}
