package tabs

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/mrlutik/kira2.0/internal/gui/dialogs"
	"github.com/mrlutik/kira2.0/internal/gui/guiHelper"
	"github.com/mrlutik/kira2.0/internal/gui/sshC"
	"github.com/sirupsen/logrus"
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
	log.SetLevel(logrus.DebugLevel)
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
		// g.sshClient, err = sshC.MakeSHH_Client(ipEntry.Text, userEntry.Text, passwordEntry.Text)
		g.sshClient, err = sshC.MakeSHH_Client("192.168.1.101:22", "d", "d")
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

func showCmdExecDialogV2(g *Gui, infoMSG string, outputChan chan guiHelper.Result) {
	var wizard *dialogs.Wizard
	outputMsg := binding.NewString()
	statusMsg := binding.NewString()
	statusMsg.Set("loading...")
	loadiningWidget := widget.NewProgressBarInfinite()

	label := widget.NewLabelWithData(outputMsg)
	closeButton := widget.NewButton("CLOSE", func() { wizard.Hide() })

	loadingDialog := container.NewBorder(
		widget.NewLabelWithData(statusMsg),
		container.NewVBox(loadiningWidget, closeButton),
		nil,
		nil,
		container.NewHScroll(container.NewVScroll(label)),
	)
	closeButton.Hide()
	wizard = dialogs.NewWizard(infoMSG, loadingDialog)
	wizard.Show(g.Window)
	wizard.Resize(fyne.NewSize(300, 400))
	wizard.ChangeTitle(infoMSG)

	output := <-outputChan
	log.Printf("Command Output: %s", output)
	outputMsg.Set(output.Output)
	// loadiningWidget.Stop()
	loadiningWidget.Hide()
	closeButton.Show()
	if output.Err != nil {
		statusMsg.Set(fmt.Sprintf("Error:\n%s", output.Err))
	} else {
		statusMsg.Set("Seccusess")
	}
}

func showCmdExecDialogAndRunCmdV3(g *Gui, infoMSG string, cmd string) {
	resultChan := make(chan guiHelper.Result)
	go func() {
		output, err := guiHelper.ExecuteSSHCommand(g.sshClient, cmd)
		resultChan <- guiHelper.Result{Output: output, Err: err}
		close(resultChan)
	}()

	var wizard *dialogs.Wizard
	outputMsg := binding.NewString()
	statusMsg := binding.NewString()
	statusMsg.Set("loading...")
	loadiningWidget := widget.NewProgressBarInfinite()

	label := widget.NewLabelWithData(outputMsg)
	closeButton := widget.NewButton("CLOSE", func() { wizard.Hide() })

	loadingDialog := container.NewBorder(
		widget.NewLabelWithData(statusMsg),
		container.NewVBox(loadiningWidget, closeButton),
		nil,
		nil,
		container.NewHScroll(container.NewVScroll(label)),
	)
	closeButton.Hide()
	wizard = dialogs.NewWizard(infoMSG, loadingDialog)
	wizard.Show(g.Window)
	wizard.Resize(fyne.NewSize(300, 400))
	wizard.ChangeTitle(infoMSG)

	output := <-resultChan
	log.Printf("Command Output: %s", output)
	outputMsg.Set(output.Output)
	// loadiningWidget.Stop()
	loadiningWidget.Hide()
	closeButton.Show()
	if output.Err != nil {
		statusMsg.Set(fmt.Sprintf("Error:\n%s", output.Err))
	} else {
		statusMsg.Set("Seccusess")
	}
}
