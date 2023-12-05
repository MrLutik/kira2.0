package guiHelper

import (
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
)

func GetIPFromSshClient(sshClient *ssh.Client) (net.IP, error) {
	if sshClient == nil {
		return nil, fmt.Errorf("sshClient is nil")
	}

	remoteAddr := sshClient.RemoteAddr()
	fmt.Println("remote addr:", remoteAddr)
	tcpAddr, ok := remoteAddr.(*net.TCPAddr)
	fmt.Println("tcpAddr", tcpAddr)
	if !ok {
		return nil, fmt.Errorf("could not type assert to *net.TCPAddr")
	}
	return tcpAddr.IP, nil
}
