package deploy

import (
	"fmt"
	"os/exec"
	"strings"

	"golang.org/x/crypto/ssh"
)

func checkPkg(client *ssh.Client, packageName string) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("Failed to create SSH session: %v", err)
	}
	defer session.Close()

	// Command to get the package version
	cmdString := fmt.Sprintf("dpkg -s %s | awk '/^Version/{print $2}'", packageName)

	log.Infof("Checking if package %s is installed on the remote machine...", packageName)
	pkgVersionBytes, err := session.Output(cmdString)
	if err != nil {
		return "", fmt.Errorf("Package %s not found on the remote machine: %v", packageName, err)
	}

	pkgVersion := strings.Trim(string(pkgVersionBytes), "\n")

	return pkgVersion, nil
}

func installPkg(path string) error {
	installCmd := exec.Command("dpkg", "-i", path)
	output, err := installCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("Failed to install package: %v", err)
	}

	log.Println(string(output))
	return nil
}

func downloadPkg(client *ssh.Client, url, path string) error {
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("Failed to create SSH session: %v", err)
	}
	defer session.Close()

	// Define the command string to download the file with wget
	cmdString := fmt.Sprintf("wget -O %s %s", path, url)

	log.Infof("Downloading package from URL: %s", url)
	if err := session.Run(cmdString); err != nil {
		return fmt.Errorf("Failed to download package: %v", err)
	}

	return nil
}
