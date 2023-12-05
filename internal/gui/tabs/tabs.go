package tabs

import "fyne.io/fyne/v2"

type Tab struct {
	Title, Info string
	View        func(w fyne.Window, g *Gui) fyne.CanvasObject
}

var (
	Tabs = map[string]Tab{
		"welcome": {
			Title: "Welcome",
			Info:  "SomeInfo",
			View:  makeWelcomeScreen,
		},
		"terminal": {
			Title: "Host Terminal",
			View:  makeTerminalScreen,
		},
		"status": {
			Title: "Node Status",
			View:  makeStatusScreen,
		},
	}

	TabsIndex = map[string][]string{
		"": {"welcome", "terminal", "status"},
	}
)
