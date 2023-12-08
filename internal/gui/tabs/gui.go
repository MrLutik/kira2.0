package tabs

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/mrlutik/kira2.0/internal/gui/dialogs"
	"github.com/mrlutik/kira2.0/internal/gui/sshC"
	"golang.org/x/crypto/ssh"
)

type Gui struct {
	sshClient *ssh.Client
	Window    fyne.Window
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

	setTab := func(t Tab) {
		title.SetText(t.Title)
		info.SetText(t.Info)
		mainWindow.Objects = []fyne.CanvasObject{t.View(g.Window, g)}
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
				// fmt.Println(uid)
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
		var err error
		g.sshClient, err = sshC.MakeSHH_Client(ipEntry.Text, userEntry.Text, passwordEntry.Text)
		// g.sshClient, err = sshC.MakeSHH_Client("192.168.0.44:22", "d", "d")
		if err != nil {

			errorLabel.SetText(fmt.Sprintf("ERROR: %s", err.Error()))
		} else {
			err = TryToRunSSHSessionForTerminal(g.sshClient)
			if err != nil {
			} else {
				wizard.Hide()

			}
		}

	}
	ipEntry.OnSubmitted = func(s string) { submitFunc() }
	userEntry.OnSubmitted = func(s string) { submitFunc() }
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

		widget.NewButton("test", func() {
			wizard.Push("set up your node", widget.NewLabel("Create node or smth"))

		}),
	)

	wizard = dialogs.NewWizard("Create ssh connection", loging)
	wizard.Show(g.Window)
	wizard.Resize(fyne.NewSize(300, 200))
}
