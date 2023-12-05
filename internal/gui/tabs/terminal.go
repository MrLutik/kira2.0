package tabs

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/fyne-io/terminal"
	"github.com/mrlutik/kira2.0/internal/gui/sshC"
	"golang.org/x/crypto/ssh"
)

var term *terminal.Terminal

// term = terminal.New()
var sshSessionForTerminal *ssh.Session

var sshIn io.WriteCloser
var sshOut io.Reader

func TryToRunSSHSessionForTerminal(c *ssh.Client) (err error) {
	sshSessionForTerminal, err = sshC.MakeSSHsessionForTerminal(c)
	if err != nil {
		return err
	}
	go sshSessionForTerminal.Shell()
	sshIn, err = sshSessionForTerminal.StdinPipe()
	if err != nil {
		return err
	}
	sshOut, err = sshSessionForTerminal.StdoutPipe()
	if err != nil {
		return err
	}
	term = terminal.New()
	go term.RunWithConnection(sshIn, sshOut)

	return nil
}

func makeTerminalScreen(_ fyne.Window) fyne.CanvasObject {

	return container.NewVScroll(term)
}

func MakeSSH() {}
