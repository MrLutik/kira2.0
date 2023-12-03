package tabs

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/fyne-io/terminal"
	"github.com/mrlutik/kira2.0/internal/gui/sshC"
	"golang.org/x/crypto/ssh"
)

// var term *terminal.Terminal
var term = terminal.New()
var sshSessionForTerminal *ssh.Session

var sshIn io.WriteCloser
var sshOut io.Reader

func TryToRunSSHSessionForTerminal(ipAndPort string, user string, psswrd string) (err error) {
	c, err := sshC.MakeSHH_Client(ipAndPort, user, psswrd)
	if err != nil {
		return err
	}
	s, err := sshC.MakeSSHsessionForTerminal(c)
	if err != nil {
		return err
	}
	sshSessionForTerminal = s
	go sshSessionForTerminal.Shell()
	sshIn, err = sshSessionForTerminal.StdinPipe()
	if err != nil {
		return err
	}
	sshOut, err = sshSessionForTerminal.StdoutPipe()
	if err != nil {
		return err
	}
	return nil
}

func makeTerminalScreen(_ fyne.Window) fyne.CanvasObject {
	go term.RunWithConnection(sshIn, sshOut)
	return container.NewVScroll(term)
}

func MakeSSH() {}
