package gui

import (
	"bytes"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/mrlutik/kira2.0/internal/gui/dialogs"
	"github.com/mrlutik/kira2.0/internal/gui/tabs"
	"golang.org/x/crypto/ssh"
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

	// a := fyne.CurrentApp()
	// a.Lifecycle().SetOnStarted(func() {
	g.showConnect()
	// })

	tab := container.NewBorder(container.NewVBox(title, info), nil, nil, nil, mainWindow)

	setTab := func(t tabs.Tab) {
		title.SetText(t.Title)
		info.SetText(t.Info)
		mainWindow.Objects = []fyne.CanvasObject{t.View(g.Window)}
	}
	menuAndTab := container.NewHSplit(g.makeNav(setTab), tab)
	menuAndTab.Offset = 0.2
	return menuAndTab

}

func (g *Gui) makeNav(setTab func(t tabs.Tab)) fyne.CanvasObject {
	a := fyne.CurrentApp()
	const preferenceCurrentTutorial = "currentTutorial"

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return tabs.TabsIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := tabs.TabsIndex[uid]

			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := tabs.Tabs[uid]
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
			if t, ok := tabs.Tabs[uid]; ok {
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

func (g *Gui) showConnect() {
	var wizard *dialogs.Wizard
	userEntry := widget.NewEntry()
	ipEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()
	errorLabel := widget.NewLabel("")
	errorLabel.Wrapping = 2
	submitFunc := func() {
		err := tabs.TryToRunSSHSessionForTerminal(ipEntry.Text, userEntry.Text, passwordEntry.Text)
		if err != nil {
			errorLabel.SetText(err.Error())
		} else {
			wizard.Hide()
		}
	}
	passwordEntry.OnSubmitted = func(s string) { submitFunc() }
	connectButton := widget.NewButton("connect to remote host", func() { submitFunc() })

	loging := container.NewVBox(
		widget.NewLabel("ip and port"),
		ipEntry,
		widget.NewLabel("user"),
		userEntry,
		widget.NewLabel("password"),
		passwordEntry,
		connectButton,
		errorLabel,
	)

	wizard = dialogs.NewWizard("Create ssh connection", loging)
	wizard.Show(g.Window)
}

func runCommand(session *ssh.Session, command string) (string, error) {
	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf

	if err := session.Run(command); err != nil {
		return "", err
	}

	return stdoutBuf.String(), nil
}
