package tabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/mrlutik/kira2.0/internal/gui/guiHelper"
)

func makeNodeControllScreen(_ fyne.Window, g *Gui) fyne.CanvasObject {
	return container.NewVScroll(
		container.NewVBox(
			widget.NewButton("start node", func() {
				startKM2(g)
			}),
			widget.NewButton("test", func() {
				runLScmd(g)
			}),
		),
	)
}

func runLScmd(g *Gui) {
	cmd := "sleep 1  && lss"
	go showCmdExecDialogAndRunCmdV3(g, "executing ls", cmd)
}

func startKM2(g *Gui) {
	// cmd := "export GITHUB_TOKEN=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 && sudo -E ~/main start   --log-level debug"

	// need to make km2 non dependet from sudo user
	cmd := `echo 'd' | sudo -S -E sh -c 'export GITHUB_TOKEN=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 && ./main start --log-level debug'`
	resultChan := make(chan guiHelper.Result)
	go func() {
		output, err := guiHelper.ExecuteSSHCommand(g.sshClient, cmd)
		resultChan <- guiHelper.Result{Output: output, Err: err}
		close(resultChan)
	}()
	go showCmdExecDialogV2(g, "starting node", resultChan)
}
