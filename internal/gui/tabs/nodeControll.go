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
			widget.NewButton("show curent config", func() {
				showKiraConfig(g)
			}),
			widget.NewButton(",", func() {
				runLScmd(g)
			}),
		),
	)
}

func runLScmd(g *Gui) {
	cmd := "ls && sleep 1  && ls && sleep 1 && ls && ls && ls && ls && ls "
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

func startKM2(g *Gui) {
	// cmd := "export GITHUB_TOKEN=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 && sudo -E ~/main start   --log-level debug"

	// need to make km2 non dependet from sudo user
	cmd := fmt.Sprintf(`echo 'd' | sudo -S -E sh -c 'export GITHUB_TOKEN=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 && %s start --log-level debug'`, guiHelper.KM2BinaryPath)
	// resultChan := make(chan guiHelper.Result)
	// go func() {
	// 	output, err := guiHelper.ExecuteSSHCommand(g.sshClient, cmd)
	// 	resultChan <- guiHelper.Result{Output: output, Err: err}
	// 	close(resultChan)
	// }()
	go showCmdExecDialogAndRunCmdV4(g, "starting node", cmd)
}
