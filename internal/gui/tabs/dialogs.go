package tabs

import (
	"fmt"
	"regexp"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/mrlutik/kira2.0/internal/gui/dialogs"
	"github.com/mrlutik/kira2.0/internal/gui/guiHelper"
	"github.com/mrlutik/kira2.0/internal/gui/sshC"
	"github.com/mrlutik/kira2.0/internal/osutils"
)

func (g *Gui) showConnect() {
	var wizard *dialogs.Wizard
	userEntry := widget.NewEntry()
	ipEntry := widget.NewEntry()
	passwordEntry := widget.NewPasswordEntry()
	errorLabel := widget.NewLabel("")

	errorLabel.Wrapping = 2
	submitFunc := func() {
		var err error
		g.sshClient, err = sshC.MakeSHH_Client(ipEntry.Text, userEntry.Text, passwordEntry.Text)
		// g.sshClient, err = sshC.MakeSHH_Client("192.168.1.103:22", "d", "d")
		if err != nil {

			errorLabel.SetText(fmt.Sprintf("ERROR: %s", err.Error()))
		} else {
			err = TryToRunSSHSessionForTerminal(g.sshClient)
			if err != nil {
			} else {
				wizard.Hide()

			}
		}

	}
	ipEntry.OnSubmitted = func(s string) { submitFunc() }
	userEntry.OnSubmitted = func(s string) { submitFunc() }
	passwordEntry.OnSubmitted = func(s string) { submitFunc() }
	connectButton := widget.NewButton("connect to remote host", func() { submitFunc() })

	logging := container.NewVBox(
		widget.NewLabel("ip and port"),
		ipEntry,
		widget.NewLabel("user"),
		userEntry,
		widget.NewLabel("password"),
		passwordEntry,
		connectButton,
		errorLabel,

		widget.NewButton("test", func() {
			wizard.Push("set up your node", widget.NewLabel("Create node or smth"))

		}),
	)

	wizard = dialogs.NewWizard("Create ssh connection", logging)
	wizard.Show(g.Window)
	wizard.Resize(fyne.NewSize(300, 200))
}

func showCmdExecDialogAndRunCmdV4(g *Gui, infoMSG string, cmd string) {
	outputChannel := make(chan string)
	errorChannel := make(chan guiHelper.ResultV2)
	go guiHelper.ExecuteSSHCommandV2(g.sshClient, cmd, outputChannel, errorChannel)

	var wizard *dialogs.Wizard
	outputMsg := binding.NewString()
	statusMsg := binding.NewString()
	statusMsg.Set("loading...")
	loadingWidget := widget.NewProgressBarInfinite()

	label := widget.NewLabelWithData(outputMsg)
	closeButton := widget.NewButton("CLOSE", func() { wizard.Hide() })
	outputScroll := container.NewVScroll(label)
	loadingDialog := container.NewBorder(
		widget.NewLabelWithData(statusMsg),
		container.NewVBox(loadingWidget, closeButton),
		nil,
		nil,
		container.NewHScroll(outputScroll),
	)
	closeButton.Hide()
	wizard = dialogs.NewWizard(infoMSG, loadingDialog)
	wizard.Show(g.Window)
	wizard.Resize(fyne.NewSize(300, 400))
	wizard.ChangeTitle(infoMSG)
	var out string
	for line := range outputChannel {
		cleanLine := cleanString(line)
		out = fmt.Sprintf("%s\n%s", out, cleanLine)
		outputMsg.Set(out)
		outputScroll.ScrollToBottom()
	}
	outputScroll.ScrollToBottom()
	loadingWidget.Hide()
	closeButton.Show()
	errcheck := <-errorChannel
	if errcheck.Err != nil {
		statusMsg.Set(fmt.Sprintf("Error:\n%s", errcheck.Err))
	} else {
		statusMsg.Set("Successes")
	}
}

func cleanString(s string) string {
	re := regexp.MustCompile("[^\x20-\x7E\n]+")
	return re.ReplaceAllString(s, "")
}

func showInfoDialog(g *Gui, infoTitle, infoString string) {
	var wizard *dialogs.Wizard
	closeButton := widget.NewButton("Close", func() { wizard.Hide() })
	infoLabel := widget.NewLabel(infoString)
	infoLabel.Wrapping = 2
	content := container.NewBorder(nil, closeButton, nil, nil,
		container.NewHScroll(
			container.NewVScroll(
				infoLabel,
			),
		),
	)

	wizard = dialogs.NewWizard(infoTitle, content)
	wizard.Show(g.Window)
	wizard.Resize(fyne.NewSize(400, 400))
}

func showInitDialog(g *Gui) {
	// var wizard *dialogs.Wizard
	// open new dialog with choices to init new or join to existing (check point)
	// if checkpoint is true, opens new text entries for ip, ports, etc... and switch main button to start\\join
	// after main button pressed exec showCmdExecDialogAndRunCmdV4 with constructed cmd from previous step
	joinExistingNetworkCheck := binding.NewBool()
	newOrJoinCHeckButton := widget.NewCheckWithData("Join to existing network", joinExistingNetworkCheck)

	var wizard *dialogs.Wizard
	mainScreen := container.NewStack()
	const defaultInterxPort, defaultSekaidRpcPort, defaultSekaiP2PPort int = 11000, 26657, 26656

	IPBinding := binding.NewString()
	// interxPortBinding := binding.NewString()
	// sekaidRPCPortBinding := binding.NewString()
	// sekaidP2PPortBinding := binding.NewString()
	ipEntry := widget.NewEntryWithData(IPBinding)
	ipEntry.SetPlaceHolder("ip of the node to connect to")
	interxPortEntry := widget.NewEntry()
	interxPortEntry.SetPlaceHolder(fmt.Sprintf("interx port (default %v)", defaultInterxPort))
	sekaidRPCPortEntry := widget.NewEntry()
	sekaidRPCPortEntry.SetPlaceHolder(fmt.Sprintf("sekaid rpc port (default %v)", defaultSekaidRpcPort))
	sekaidP2PPortEntry := widget.NewEntry()
	sekaidP2PPortEntry.SetPlaceHolder(fmt.Sprintf("sekaid p2p port (default %v)", defaultSekaiP2PPort))

	validatePort := func(port *int, defaultValue int, portEntryText string) error {
		if portEntryText == "" {
			log.Debugln("switching to default value", port, defaultValue)
			*port = defaultValue
			return nil
		}
		if ok, err := osutils.CheckIfPortIsValid(portEntryText); ok {
			*port, _ = strconv.Atoi(portEntryText)
			return nil
		} else {
			return err
		}
	}

	var interxPort int
	var RPCport int
	var P2PPort int
	//check flags
	checkEntries := func() error {
		ip, err := IPBinding.Get()
		if err != nil {
			return err
		}
		ok, err := osutils.CheckIfIPIsValid(ip)
		if err != nil {
			return err
		}
		if !ok {
			return fmt.Errorf("ip is not valid")
		}

		if err := validatePort(&interxPort, defaultInterxPort, interxPortEntry.Text); err != nil {
			return fmt.Errorf("interx port <%s> is invalid: %w", interxPortEntry.Text, err)
		}
		if err := validatePort(&RPCport, defaultSekaidRpcPort, sekaidRPCPortEntry.Text); err != nil {
			return fmt.Errorf("rpc port <%s> is invalid: %w", sekaidRPCPortEntry.Text, err)
		}
		if err := validatePort(&P2PPort, defaultSekaiP2PPort, sekaidP2PPortEntry.Text); err != nil {
			return fmt.Errorf("p2p port <%s> is invalid: %w", sekaidP2PPortEntry.Text, err)
		}
		return nil
	}
	errorMsgBinding := binding.NewString()
	errorMsgLabel := widget.NewLabelWithData(errorMsgBinding)
	joinScreen := container.NewVScroll(container.NewVBox(
		ipEntry,
		interxPortEntry,
		sekaidP2PPortEntry,
		sekaidRPCPortEntry,
		errorMsgLabel,
	))
	// mainScreen := joinScreen
	closeButton := widget.NewButton("Close", func() { wizard.Hide() })
	joinOrCreateButton := widget.NewButton("Create",
		func() {
			b, err := joinExistingNetworkCheck.Get()
			if err != nil {
				log.Fatalln(err)
			}
			err = checkEntries()
			if err != nil {
				err = errorMsgBinding.Set(err.Error())
				if err != nil {
					log.Fatalln(err)
				}
			} else {
				switch b {
				case true:
					fmt.Println("joining")
					log.Debugf("%v %v %v", interxPort, P2PPort, RPCport)
				default:
					fmt.Println("creating new")
				}
			}

		},
	)
	switchFunc := func() {
		b, err := joinExistingNetworkCheck.Get()
		if err != nil {
			log.Fatalln(err)
		}
		if b {
			mainScreen.Objects = []fyne.CanvasObject{joinScreen}
			mainScreen.Refresh()
			joinOrCreateButton.SetText("Join to existing network")

		} else {
			mainScreen.Objects = []fyne.CanvasObject{}
			mainScreen.Refresh()
			joinOrCreateButton.SetText("Initialize new network")
		}

	}

	switchFunc()
	joinExistingNetworkCheck.AddListener(binding.NewDataListener(switchFunc))
	content := container.NewBorder(newOrJoinCHeckButton, container.NewVBox(joinOrCreateButton, closeButton), nil, nil, mainScreen)
	wizard = dialogs.NewWizard("Node initializing", content)
	wizard.Show(g.Window)
	wizard.Resize(fyne.NewSize(400, 400))
}
