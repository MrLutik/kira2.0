package tabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeWelcomeScreen(_ fyne.Window, _ *Gui) fyne.CanvasObject {
	// fmt.Println("makeWalcomeScreen")
	return container.NewStack(widget.NewLabel("Welcome fren"))
}
