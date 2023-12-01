package gui

import (
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/fyne-io/terminal"
	"golang.org/x/crypto/ssh"
)

var term = terminal.New()

// var term *terminal.Terminal
var sshSession *ssh.Session
var sshIn io.WriteCloser
var sshOut io.Reader

func makeTerminalScreen(_ fyne.Window) fyne.CanvasObject {
	go term.RunWithConnection(sshIn, sshOut)
	return container.NewVScroll(term)
	// return fyne.NewContainer()
}

func MakeSSH() {}
