package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/crypto/ssh"
	// "github.com/mrlutik/kira2.0/internal/gui"
)

type Gui struct {
	// term *terminal.Terminal
	// sshConnection
	Window fyne.Window
}

func (g *Gui) MakeGui() fyne.CanvasObject {

	title := widget.NewLabel("Component name")
	info := widget.NewLabel("An introduction would probably go\nhere, as well as a")
	// g.content = container.NewStack()
	mainWindow := container.NewStack()

	sshS, err := makeSSHsession()
	if err != nil {
		panic(err)
	}
	sshSession = sshS
	go sshSession.Shell()
	sshIn, _ = sshSession.StdinPipe()
	sshOut, _ = sshSession.StdoutPipe()

	tab := container.NewBorder(container.NewVBox(title, info), nil, nil, nil, mainWindow)

	setTab := func(t Tab) {
		title.SetText(t.Title)
		info.SetText(t.Info)
		mainWindow.Objects = []fyne.CanvasObject{t.View(g.Window)}
	}
	menuAndTab := container.NewHSplit(g.makeNav(setTab), tab)
	menuAndTab.Offset = 0.2
	return menuAndTab

}

func (g *Gui) makeNav(setTab func(t Tab)) fyne.CanvasObject {
	a := fyne.CurrentApp()
	const preferenceCurrentTutorial = "currentTutorial"

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return TabsIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := TabsIndex[uid]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := Tabs[uid]
			if !ok {
				fyne.LogError("Missing tutorial panel: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.Title)
			// if unsupportedTutorial(t) {
			// 	obj.(*widget.Label).TextStyle = fyne.TextStyle{Italic: true}
			// } else {
			// 	obj.(*widget.Label).TextStyle = fyne.TextStyle{}
			// }
			obj.(*widget.Label).TextStyle = fyne.TextStyle{}
		},
		OnSelected: func(uid string) {
			if t, ok := Tabs[uid]; ok {
				// if unsupportedTutorial(t) {
				// 	return
				// }
				fmt.Println(uid)
				a.Preferences().SetString(preferenceCurrentTutorial, uid)
				setTab(t)
			}
		},
	}

	return tree
}

func makeSSHsession() (*ssh.Session, error) {
	config := &ssh.ClientConfig{
		User: "d",
		Auth: []ssh.AuthMethod{
			ssh.Password("d"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Connect to the SSH server
	client, err := ssh.Dial("tcp", "192.168.1.104:22", config)
	if err != nil {
		return nil, err
	}

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

	// // Get stdin and stdout for the session
	// stdin, err := session.StdinPipe()
	// if err != nil {
	// 	session.Close()
	// 	client.Close()
	// 	return nil, nil, nil, err
	// }

	// stdout, err := session.StdoutPipe()
	// if err != nil {
	// 	stdin.Close()
	// 	session.Close()
	// 	client.Close()
	// 	return nil, nil, nil, err
	// }

	// // Start a shell
	// if err := session.Shell(); err != nil {

	// 	stdin.Close()
	// 	session.Close()
	// 	client.Close()
	// 	return nil, nil, nil, err
	// }

	// Return the stdin and stdout
	// return stdin, stdout, session, nil
	return session, nil
}
