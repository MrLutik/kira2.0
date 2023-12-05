package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/mrlutik/kira2.0/internal/gui/tabs"
)

func RunGui() {
	a := app.NewWithID("kira manager 2.0")
	w := a.NewWindow("Title")
	w.SetMaster()
	w.Resize(fyne.NewSize(1024, 768))
	g := tabs.Gui{
		Window: w,
	}

	content := g.MakeGui()

	g.Window.SetContent(content)
	g.Window.ShowAndRun()
}
