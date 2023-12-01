package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/mrlutik/kira2.0/internal/gui"
	// "github.com/mrlutik/kira2.0/internal/gui"
)

func main() {
	a := app.NewWithID("kira manager 2.0")
	w := a.NewWindow("Title")
	w.SetMaster()
	w.Resize(fyne.NewSize(1024, 768))
	g := gui.Gui{
		Window: w,
	}

	content := g.MakeGui()

	w.SetContent(content)
	w.ShowAndRun()
}
