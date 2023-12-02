package tabs

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeWelcomeScreen(_ fyne.Window) fyne.CanvasObject {
	fmt.Println("makeWalcomeScreen")
	return container.NewStack(widget.NewLabel("Welcome fren"))
}
