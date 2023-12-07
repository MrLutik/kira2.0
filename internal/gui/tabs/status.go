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
	// var h helper
	// h.GetIP(g.sshClient)
	ip, err := guiHelper.GetIPFromSshClient(g.sshClient)

	fmt.Println(g.sshClient.SessionID(), "iP::::", ip)

	// }
	// ip, _ := "pepeg", "pepeg"
	// content := container.NewStack()
	tabs := container.NewAppTabs(
		container.NewTabItem("api/status", makeTab1Status(ip.String())),
		container.NewTabItem("api/dashboard", makeTab2Dashboard(ip.String())),
	)
	// normalLs, err1 := guiHelper.ExecuteSSHCommand(g.sshClient, "ls")
	// errorLS, err2 := guiHelper.ExecuteSSHCommand(g.sshClient, "lsssssss")
	return container.NewVScroll(container.NewVBox(
		// widget.NewLabel(fmt.Sprintf("Host ip: %s, err: %s", ip.String(), err)),
		// widget.NewLabel(fmt.Sprintf("normalLS: %s, err1: %s", normalLs, err1)),
		// widget.NewLabel(fmt.Sprintf("errorLS: %s, err1: %s", errorLS, err2)),
		// widget.NewLabel(fmt.Sprintf("status: %s, errs: %s", status2, errStatus2)),
		tabs,
		widget.NewLabel(fmt.Sprintf("%s", err)),
	),
	)

	// widget.NewLabel(fmt.Sprintf("%s,", ip)),

}

func makeTab1Status(ip string) fyne.CanvasObject {
	log.Debugln("maketab1")
	status := binding.NewString()
	errStatus := binding.NewString()

	s, e := guiHelper.MakeHttpRequest(fmt.Sprintf("http://%s:11000/api/status", ip))
	status.Set(string(s))
	errStatus.Set(fmt.Sprintf("%s", e))

	return container.NewVBox(
		widget.NewButton("REFRESH", func() {
			s, e = guiHelper.MakeHttpRequest(fmt.Sprintf("http://%s:11000/api/status", ip))
			status.Set("status:" + string(s))
			errStatus.Set("error:" + fmt.Sprintf("%s", e))
		}),
		// widget.NewLabelWithData(fmt.Sprintf("status: %s, errs: %s", status, errStatus)),
		widget.NewLabelWithData(status),
		widget.NewLabelWithData(errStatus),
	)

}

func makeTab2Dashboard(ip string) fyne.CanvasObject {
	log.Debugln("maketab2")
	out, err := guiHelper.MakeHttpRequest(fmt.Sprintf("http://%s:11000/api/dashboard", ip))
	data := binding.NewString()
	data.Set(string(out) + fmt.Sprintf("%s", err))

	// return widget.NewLabel(fmt.Sprintf("status: %s, errs: %s", out, err))
	return container.NewVBox(
		widget.NewButton("REFRESH", func() {
			out, err = guiHelper.MakeHttpRequest(fmt.Sprintf("http://%s:11000/api/dashboard", ip))
			data.Set(string(out) + fmt.Sprintf("%s", err))
		}),
		// widget.NewLabelWithData(fmt.Sprintf("status: %s, errs: %s", status, errStatus)),
		widget.NewLabelWithData(data),
	)
}
