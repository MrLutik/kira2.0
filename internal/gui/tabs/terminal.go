package tabs

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/fyne-io/terminal"
	"golang.org/x/crypto/ssh"
)

var term = terminal.New()

// var term *terminal.Terminal
var SshSession *ssh.Session
var SshIn io.WriteCloser
var SshOut io.Reader

func makeTerminalScreen(_ fyne.Window) fyne.CanvasObject {
	go term.RunWithConnection(SshIn, SshOut)
	return container.NewVScroll(term)
	// return fyne.NewContainer()
}

func MakeSSH() {}
