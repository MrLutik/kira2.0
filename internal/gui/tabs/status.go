package tabs

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/mrlutik/kira2.0/internal/gui/guiHelper"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/sirupsen/logrus"
)

var log = logging.Log

func makeStatusScreen(_ fyne.Window, g *Gui) fyne.CanvasObject {
	log.SetLevel(logrus.DebugLevel)
	ip, err := guiHelper.GetIPFromSshClient(g.sshClient)

	fmt.Println(g.sshClient.SessionID(), "iP::::", ip, err)

	tabs := container.NewAppTabs(
		container.NewTabItem("api/status", makeTab1Status(ip.String())),
		container.NewTabItem("api/dashboard", makeTab2Dashboard(ip.String())),
	)

	return tabs

}

func makeTab1Status(ip string) fyne.CanvasObject {
	log.Debugln("maketab1")
	status := binding.NewString()
	errStatus := binding.NewString()

	s, e := guiHelper.MakeHttpRequest(fmt.Sprintf("http://%s:11000/api/status", ip))
	status.Set(string(s))
	errStatus.Set(fmt.Sprintf("%s", e))

	return container.NewBorder(
		widget.NewButton("REFRESH", func() {
			s, e = guiHelper.MakeHttpRequest(fmt.Sprintf("http://%s:11000/api/status", ip))
			status.Set("status:" + string(s))
			errStatus.Set("error:" + fmt.Sprintf("%s", e))
		}),
		nil,
		nil,
		nil,
		container.NewVScroll(
			container.NewBorder(
				nil,
				widget.NewLabelWithData(errStatus),
				nil,
				nil,
				widget.NewLabelWithData(status),
			),
		),
	)

}

func makeTab2Dashboard(ip string) fyne.CanvasObject {
	log.Debugln("maketab2")
	out, err := guiHelper.MakeHttpRequest(fmt.Sprintf("http://%s:11000/api/dashboard", ip))
	data := binding.NewString()
	data.Set(string(out) + fmt.Sprintf("%s", err))

	return container.NewBorder(
		widget.NewButton("REFRESH", func() {
			out, err = guiHelper.MakeHttpRequest(fmt.Sprintf("http://%s:11000/api/dashboard", ip))
			data.Set(string(out) + fmt.Sprintf("%s", err))
		}),
		nil,
		nil,
		nil,
		// widget.NewLabelWithData(fmt.Sprintf("status: %s, errs: %s", status, errStatus)),
		container.NewVScroll(
			container.NewBorder(
				nil,
				nil,
				nil,
				nil,
				widget.NewLabelWithData(data),
			),
		),
	)
}
