package tabs

import (
	"fmt"
	"reflect"

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
			widget.NewButton("stop node", func() {
				stopKM2(g)
			}),
			widget.NewButton("init node", func() {
				initKM2(g)
			}),
			widget.NewButton("show curent config", func() {
				showKiraConfig(g)
			}),
			widget.NewButton("With no error", func() {
				runLScmd(g, false)
			}),
			widget.NewButton("With  error", func() {
				runLScmd(g, true)
			}),
		),
	)
}

func runLScmd(g *Gui, e bool) {
	var cmd string
	if e {
		cmd = "lss && sleep 1  && ls && sleep 1 && ls && ls && ls && ls && ls "
	} else {
		cmd = "ls && sleep 1  && ls && sleep 1 && ls && ls && ls && ls && ls "
	}
	showCmdExecDialogAndRunCmdV4(g, "executing ls", cmd)
}

func showKiraConfig(g *Gui) {
	cfg, err := guiHelper.ReadKiraConfigFromKM2cfgFile(g.sshClient)
	val := reflect.ValueOf(cfg)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	var outputString string
	t := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		// fmt.Printf("%s: %v\n", t.Field(i).Name, field.Interface())
		outputString += fmt.Sprintf("%s: %v\n", t.Field(i).Name, field.Interface())
	}
	showInfoDialog(g, "Kira Config", fmt.Sprintf("cfg:\n%s\nerr:\n%s", outputString, err))
	fmt.Println(cfg, err)
}

func stopKM2(g *Gui) {
	cmd := fmt.Sprintf(`echo '%s'  | sudo -S -E sh -c 'export GITHUB_TOKEN=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 && %s stop --log-level debug'`, guiHelper.SudoPassword, guiHelper.KM2BinaryPath)
	go showCmdExecDialogAndRunCmdV4(g, "stoping node", cmd)
}

func initKM2(g *Gui) {
	go showInitDialog(g)
}

func startKM2(g *Gui) {
	// need to make km2 non dependet from sudo user
	cmd := fmt.Sprintf(`echo '%s' | sudo -S -E sh -c 'export GITHUB_TOKEN=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 && %s start --log-level debug'`, guiHelper.SudoPassword, guiHelper.KM2BinaryPath)
	// resultChan := make(chan guiHelper.Result)
	// go func() {
	// 	output, err := guiHelper.ExecuteSSHCommand(g.sshClient, cmd)
	// 	resultChan <- guiHelper.Result{Output: output, Err: err}
	// 	close(resultChan)
	// }()
	go showCmdExecDialogAndRunCmdV4(g, "starting node", cmd)
}
