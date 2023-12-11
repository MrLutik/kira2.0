package guiHelper

import (
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/mrlutik/kira2.0/internal/logging"
	"golang.org/x/crypto/ssh"
)

var log = logging.Log

type Result struct {
	Output string
	Err    error
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
