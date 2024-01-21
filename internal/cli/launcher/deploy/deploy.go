package deploy

import (
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

const (
	// Command information
	use   = "deploy [ip address]"
	short = "Short description of deploy command"
	long  = "Long description of deploy command"

	// Flags
	privateKeyFlag = "priv-key"
	publicKeyFlag  = "pub-key"
)

var (
	log   = logging.Log
	nodes = []string{"interx", "sekai"}
)

func Node() *cobra.Command {
	log.Debugln("Adding `deploy` command...")
	nodeCmd := &cobra.Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Args:    cobra.ExactArgs(1),
		Example: "deploy 127.0.0.1 --priv-key=path/to/priv-key --pub-key=path/to/pub-key --interx=v0.3.16 --sekai=v0.3.46",
		Run: func(cmd *cobra.Command, args []string) {
			privKey, err := cmd.Flags().GetString(privateKeyFlag)
			if err != nil {
				log.Fatalf("Failed to get private key: %s", err)
			}
			pubKey, err := cmd.Flags().GetString(publicKeyFlag)
			if err != nil {
				log.Fatalf("Failed to get public key: %s", err)
			}

			client, err := createSSHClient(args[0], privKey)
			if err != nil {
				log.Fatalf("Failed to create SSH client: %v", err)
			}
			defer client.Close()

			if err = installKeys(client, privKey, pubKey); err != nil {
				log.Fatalf("Failed to install keys: %v", err)
			}

			if err = forbidRootLogin(client); err != nil {
				log.Fatalf("Failed to forbid root login: %v", err)
			}

			osAndHardwareInfo, err := checkOSAndHardware(client)
			if err != nil {
				log.Fatalf("Failed to check OS and hardware: %v", err)
			}
			log.Infof("OS and Hardware info: %s", osAndHardwareInfo)
		},
	}
	for _, node := range nodes {
		nodeCmd.PersistentFlags().String(node, "", "Provide version to deploy")
	}
	nodeCmd.PersistentFlags().String(privateKeyFlag, "", "Path to private key")
	nodeCmd.PersistentFlags().String(publicKeyFlag, "", "Path to pub key") // !Can be generated from private

	return nodeCmd
}

func createSSHClient(host, privKeyPath string) (*ssh.Client, error) {
	key, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read private key: %w", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("unable to parse private key: %w", err)
	}

	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // Be cautious with this in a production environment
	}

	client, err := ssh.Dial("tcp", host+":22", config)
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %w", err)
	}

	return client, nil
}

func forbidRootLogin(client *ssh.Client) error {
	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Close()

	// Disable root login in sshd_config
	_, err = session.Output("sed -i 's/^PermitRootLogin yes/PermitRootLogin no/g' /etc/ssh/sshd_config")
	if err != nil {
		return fmt.Errorf("failed to disable root login: %w", err)
	}

	// Restart SSH service to apply changes
	_, err = session.Output("service ssh restart")
	if err != nil {
		return fmt.Errorf("failed to restart SSH service: %w", err)
	}

	return nil
}

func checkOSAndHardware(client *ssh.Client) (string, error) {
	// Check the operating system
	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Close()

	var os string
	if output, err := session.Output("uname"); err == nil {
		os = string(output)
	} else if output, err := session.Output("cmd /c ver"); err == nil {
		os = string(output)
	} else {
		return "", fmt.Errorf("failed to determine operating system: %w", err)
	}

	// Check hardware resources
	hardwareCommands := []string{
		"lscpu",
		"lspci",
		"lshw",
		"lsscsi",
		"lsusb",
		"df",
		"free",
		"dmidecode",
		"hdparm",
	}

	var (
		hardwareInfo string
		mutex        sync.Mutex
		wg           sync.WaitGroup
	)

	for _, cmd := range hardwareCommands {
		wg.Add(1)
		go func(command string) {
			defer wg.Done()
			session, err := client.NewSession()
			if err != nil {
				log.Errorf("Failed to create session: %s", err)
				return
			}
			defer session.Close()
			if output, err := session.Output(command); err == nil {
				mutex.Lock()
				hardwareInfo += string(output) + "\n"
				mutex.Unlock()
			} else {
				log.Errorf("Failed to execute command %s: %s", command, err)
			}
		}(cmd)
	}

	wg.Wait()

	return fmt.Sprintf("OS: %s\nHardware Info:\n%s", os, hardwareInfo), nil
}
