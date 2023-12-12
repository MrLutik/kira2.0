package tabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
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
	cmd := "ls && sleep 1  && ls && sleep 1 && ls && ls && ls && ls && ls "
	showCmdExecDialogAndRunCmdV4(g, "executing ls", cmd)
}

func startKM2(g *Gui) {
	// cmd := "export GITHUB_TOKEN=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 && sudo -E ~/main start   --log-level debug"

	// need to make km2 non dependet from sudo user
	cmd := `echo 'd' | sudo -S -E sh -c 'export GITHUB_TOKEN=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 && ./main start --log-level debug'`
	// resultChan := make(chan guiHelper.Result)
	// go func() {
	// 	output, err := guiHelper.ExecuteSSHCommand(g.sshClient, cmd)
	// 	resultChan <- guiHelper.Result{Output: output, Err: err}
	// 	close(resultChan)
	// }()
	go showCmdExecDialogAndRunCmdV4(g, "starting node", cmd)
}
