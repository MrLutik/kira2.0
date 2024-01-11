package sshC

import "golang.org/x/crypto/ssh"

func MakeSHH_Client(ipAndPort, user, psswrd string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(psswrd),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the SSH server
	client, err := ssh.Dial("tcp", ipAndPort, config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// create session with pty request for fyne term
func MakeSSHsessionForTerminal(client *ssh.Client) (*ssh.Session, error) {
	// Create a session
	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, err
	}

	// Request a pty (pseudo-terminal)
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // Enable echoing
		ssh.TTY_OP_ISPEED: 14400, // Input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // Output speed = 14.4kbaud
	}

	if err := session.RequestPty("ansi", 80, 40, modes); err != nil {
		session.Close()
		client.Close()
		return nil, err
	}

	return session, nil
}

func MakeSSHsessionForCommands(client *ssh.Client) (*ssh.Session, error) {
	// Create a session
	session, err := client.NewSession()
	if err != nil {
		client.Close()
		session.Close()
		return nil, err
	}

	return session, nil
}
