package tabs

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/mrlutik/kira2.0/internal/gui/dialogs"
	"github.com/mrlutik/kira2.0/internal/gui/guiHelper"
)

func makeNodeControllScreen(_ fyne.Window, g *Gui) fyne.CanvasObject {
	return container.NewVScroll(
		container.NewVBox(
			widget.NewButton("start node", func() {}),
			widget.NewButton("test", func() {
				startKM2(g)
			}),
		),
	)
}

func startKM2(g *Gui) {
	// cmd := "export GITHUB_TOKEN=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 && sudo -E ~/main start   --log-level debug"
	cmd := "/home/d/main start   --log-level debug"
	// cmd = "ls"
	out, err := guiHelper.ExecuteSSHCommand(g.sshClient, cmd)

	showLoadingDialog(g, string(out), err)
	// showLoadingDialog(g, "string(out)", errors.New("eror"))

}

func showLoadingDialog(g *Gui, out string, outErr error) {
	var wizard *dialogs.Wizard

	dialogMSG := binding.NewString()

	if outErr != nil {
		dialogMSG.Set(fmt.Sprintf("%s", outErr))
	} else {
		dialogMSG.Set(out)
	}
	loadingDialog := container.NewVBox(
		widget.NewLabelWithData(dialogMSG),
		widget.NewButton("CLOSE", func() { wizard.Hide() }),
	)
	wizard = dialogs.NewWizard("Create ssh connection", loadingDialog)
	wizard.Show(g.Window)
	wizard.Resize(fyne.NewSize(300, 200))
}
