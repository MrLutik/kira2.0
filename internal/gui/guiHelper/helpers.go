package guiHelper

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"sync"

	"github.com/mrlutik/kira2.0/internal/config"
	"github.com/mrlutik/kira2.0/internal/logging"
	cmdConfig "github.com/mrlutik/kira2.0/internal/manager/cli/commands/config"
	"golang.org/x/crypto/ssh"
)

var log = logging.Log

const KM2BinaryPath string = "~/main"
const SudoPassword string = "d"

type Result struct {
	Output string
	Err    error
}

type ResultV2 struct {
	Err error
}

func GetIPFromSshClient(sshClient *ssh.Client) (net.IP, error) {
	if sshClient == nil {
		return nil, fmt.Errorf("sshClient is nil")
	}

	remoteAddr := sshClient.RemoteAddr()
	tcpAddr, ok := remoteAddr.(*net.TCPAddr)
	if !ok {
		return nil, fmt.Errorf("could not type assert to *net.TCPAddr")
	}
	return tcpAddr.IP, nil
}

// exec cmd on remote host, returns the outpot of execution or error
func ExecuteSSHCommand(client *ssh.Client, command string) (string, error) {
	// Create a session. It is important to defer closing the session.
	log.Printf("RUNNING CMD:\n%s", command)
	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()

	// Run the command and capture the output.
	output, err := session.CombinedOutput(command)
	if err != nil {
		log.Printf("OUT OF CMD: %s\n ERROR OUT: %s", string(output), err)
		return string(output), fmt.Errorf("failed to run command: %v", err)
	}
	log.Printf("OUT OF CMD: %s\n ERROR OUT: %s", string(output), err)

	return string(output), nil
}
func ExecuteSSHCommandV2(client *ssh.Client, command string, outputChan chan<- string, resultChan chan<- ResultV2) {
	log.Printf("RUNNING CMD:\n%s", command)
	session, err := client.NewSession()
	if err != nil {
		resultChan <- ResultV2{Err: err}
		return
	}
	defer session.Close()

	// Setting up stdout and stderr pipes
	stdoutPipe, err := session.StdoutPipe()
	if err != nil {
		resultChan <- ResultV2{Err: err}
		return
	}
	stderrPipe, err := session.StderrPipe()
	if err != nil {
		resultChan <- ResultV2{Err: err}
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // This will be called if an early return occurs
	var wg sync.WaitGroup
	wg.Add(2)
	// Start the command
	err = session.Start(command)
	if err != nil {
		cancel()
		close(outputChan) // Close the channel on error
		resultChan <- ResultV2{Err: err}
		return
	}

	// Read from stdout and stderr concurrently
	go streamOutput(ctx, stdoutPipe, outputChan, &wg)
	go streamOutput(ctx, stderrPipe, outputChan, &wg)

	err = session.Wait()
	cancel()
	wg.Wait()
	close(outputChan) // Close the channel when done
	if err != nil {
		resultChan <- ResultV2{Err: err}
		return
	}

	resultChan <- ResultV2{Err: err}
}

func streamOutput(ctx context.Context, reader io.Reader, outputChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	scanner := bufio.NewScanner(reader)
	for {
		select {
		case <-ctx.Done():
			return // Exit if context is cancelled
		default:
			if scanner.Scan() {
				outputChan <- scanner.Text()
			} else {
				return // Exit if there's nothing more to read
			}
		}
	}
}

func MakeHttpRequest(url string) ([]byte, error) {
	// Make a GET request
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

// reads cfg file from kira config on remote host
func ReadKiraConfigFromKM2cfgFile(client *ssh.Client) (*config.KiraConfig, error) {
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %s", err)
		return nil, err
	}
	defer session.Close()

	// var b bytes.Buffer
	// session.Stdout = &b
	// if err := session.Run(fmt.Sprintf("cat %s", configFileController.KiraCfgFilePath)); err != nil { // replace with the path to your config file
	// 	log.Fatalf("Failed to run command: %s", err)
	// 	return nil, err
	// }
	out, err := ExecuteSSHCommand(client, fmt.Sprintf("%s %s --%s", KM2BinaryPath, cmdConfig.Use, cmdConfig.PrintFlag))
	if err != nil {
		return nil, fmt.Errorf(out, err)
	}

	// Parse the TOML content
	var config config.KiraConfig
	err = json.Unmarshal([]byte(out), &config)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %s", err)
	}

	log.Println(config)
	return &config, nil
}
