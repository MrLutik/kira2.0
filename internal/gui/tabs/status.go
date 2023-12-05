package tabs

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/mrlutik/kira2.0/internal/gui/guiHelper"
)

func makeStatusScreen(_ fyne.Window, g *Gui) fyne.CanvasObject {
	// var h helper
	// h.GetIP(g.sshClient)
	ip, err := guiHelper.GetIPFromSshClient(g.sshClient)
	if err != nil {
		return container.NewStack(
			widget.NewLabel(fmt.Sprintf("%s", err.Error())),
			// widget.NewLabel(fmt.Sprintf("%s,", ip)),
		)
	}
	fmt.Println(g.sshClient.SessionID(), "iP::::", ip)

	// }
	// ip, _ := "pepeg", "pepeg"
	return container.NewStack(
		widget.NewLabel(fmt.Sprintf("%s", ip.String())),
		// widget.NewLabel(fmt.Sprintf("%s,", ip)),
	)
}
