package tabs

import (
	"encoding/hex"
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
	"github.com/tyler-smith/go-bip39"
)

var sshMnemonicsMap = make(map[string]string)

func (g *Gui) showConnect() {
	var wizard *dialogs.Wizard

	joinToInitializedNode := func() *fyne.Container {

		updateMnemonicKeysList := func() []string {
			var list []string
			encryptedMnemonics, _ := guiHelper.GetKeys()
			for _, m := range encryptedMnemonics {
				list = append(list, m.Name)
			}
			return list
		}
		mnemonics := updateMnemonicKeysList()

		var selectedKeyToJoin guiHelper.EncryptedMnemonic
		userEntry := widget.NewEntry()
		ipEntry := widget.NewEntry()
		errorLabel := widget.NewLabel("")
		decryptionPasswordEntry := widget.NewPasswordEntry()
		encryptedMnemonicsSelect := widget.NewSelect(mnemonics, func(s string) {
			k, err := guiHelper.GetKey(s)
			if err != nil {
				errorLabel.Text = err.Error()
				return
			}
			log.Println("selected to join: ", s, k)
			selectedKeyToJoin = k
		})
		restoreButton := widget.NewButton("restore SSH key", func() {

			g.showAddKeyDialog(encryptedMnemonicsSelect)

			selectList := updateMnemonicKeysList()
			log.Printf("updating list\n")
			encryptedMnemonicsSelect.SetOptions(selectList)
		})
		bNonce := []byte(guiHelper.Nonce)
		log.Printf("nonce: %s hexNonce %s", guiHelper.Nonce, hex.EncodeToString([]byte(guiHelper.Nonce)))
		log.Printf("bnonce: %s hexbnonce %s", bNonce, hex.EncodeToString(bNonce))
		log.Printf("bnonce: %v", bNonce)

		submitFunc := func() {
			var err error
			convertedPassword, err := guiHelper.Set32BytePassword(decryptionPasswordEntry.Text)
			if err != nil {
				errorLabel.Text = err.Error()
				return
			}
			sshMnemonic := guiHelper.DecryptMnemonic(selectedKeyToJoin.EncoderMnemonicHex, hex.EncodeToString(convertedPassword), hex.EncodeToString([]byte(guiHelper.Nonce)))
			log.Printf("HexPassword: %s, hexNonce%s", hex.EncodeToString(convertedPassword), hex.EncodeToString([]byte(guiHelper.Nonce)))
			log.Println("decrypted mnemonic", string(sshMnemonic))
			privateKey, err := guiHelper.GeneratePrivateP256KeyFromMnemonic(string(sshMnemonic))

			if err != nil {
				errorLabel.Text = err.Error()
				return
			}
			g.sshClient, err = sshC.MakeSHH_ClientWithKey(ipEntry.Text, userEntry.Text, privateKey)
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
		connectButton := widget.NewButton("connect", func() { submitFunc() })

		authScreen := container.NewVBox(
			widget.NewLabel("ip and port"),
			ipEntry,
			widget.NewLabel("user"),
			userEntry,
			widget.NewLabel("select SSH key"),
			encryptedMnemonicsSelect,
			widget.NewLabel("Password to unlock SSH key"),
			decryptionPasswordEntry,
			connectButton,
			errorLabel,
			restoreButton,
		)
		return authScreen
	}

	//join to new host tab
	joinToNewHost := func() *fyne.Container {
		userEntry := widget.NewEntry()
		ipEntry := widget.NewEntry()
		passwordEntry := widget.NewPasswordEntry()
		errorLabel := widget.NewLabel("")

		errorLabel.Wrapping = 2
		submitFunc := func() {
			var err error
			g.sshClient, err = sshC.MakeSHH_ClientWithPassword(ipEntry.Text, userEntry.Text, passwordEntry.Text)
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
		return logging
	}
	mainDialogScreen := container.NewAppTabs(
		container.NewTabItem("Existing Node", joinToInitializedNode()),
		container.NewTabItem("New Host", joinToNewHost()),
	)
	wizard = dialogs.NewWizard("Create ssh connection", mainDialogScreen)
	wizard.Show(g.Window)
	wizard.Resize(fyne.NewSize(300, 200))
}

func (g *Gui) showErrorDialog(err error) {
	var wizard *dialogs.Wizard
	mainDialogScreen := container.NewVBox(
		widget.NewLabel(err.Error()),
		widget.NewButton("Close", func() { wizard.Hide() }),
	)
	wizard = dialogs.NewWizard("Create ssh connection", mainDialogScreen)
	wizard.Show(g.Window)
	wizard.Resize(fyne.NewSize(300, 200))

}

func (g *Gui) showAddKeyDialog(sW *widget.Select) {
	var wizard *dialogs.Wizard
	updateMnemonicKeysList := func() []string {
		var list []string
		encryptedMnemonics, _ := guiHelper.GetKeys()
		for _, m := range encryptedMnemonics {
			list = append(list, m.Name)
		}
		return list
	}
	sshMnemonicChoice := "sshMnemonic"
	masterMnemonicChoice := "masterMnemonic"
	var masterMnemonicCheck bool
	choices := []string{sshMnemonicChoice, masterMnemonicChoice}
	mnemonicNameEntering := widget.NewEntry()
	mnemonicEntering := widget.NewEntry()
	encryptionPassword := widget.NewPasswordEntry()
	radioCheck := widget.NewRadioGroup(choices, func(s string) {
		if s == sshMnemonicChoice {
			masterMnemonicCheck = false
		} else if s == masterMnemonicChoice {
			masterMnemonicCheck = true
		}
	})
	radioCheck.SetSelected(sshMnemonicChoice)
	addButton := widget.NewButton("Add", func() {
		var gErr error
		if masterMnemonicCheck {
		} else {

			mnemonicCheck := bip39.IsMnemonicValid(mnemonicEntering.Text)
			if !mnemonicCheck {
				err := fmt.Errorf("Mnemonic is not valid")
				g.showErrorDialog(err)
				gErr = err
				return
			}
			formattedPassword, err := guiHelper.Set32BytePassword(encryptionPassword.Text)
			if err != nil {
				gErr = err
				g.showErrorDialog(err)
				return
			}
			eMnemonic, err := guiHelper.EncryptMnemonic(mnemonicEntering.Text, formattedPassword, []byte(guiHelper.Nonce))
			if err != nil {
				gErr = err
				g.showErrorDialog(err)
				return
			}
			err = guiHelper.AddKey(guiHelper.EncryptedMnemonic{Name: mnemonicNameEntering.Text, EncoderMnemonicHex: hex.EncodeToString(eMnemonic)})
			if err != nil {
				gErr = err
				g.showErrorDialog(err)
				return
			}
		}
		if gErr == nil {
			sW.SetOptions(updateMnemonicKeysList())
			wizard.Hide()

		}
	})
	closeButton := widget.NewButton("Close", func() {
		sW.SetOptions(updateMnemonicKeysList())
		wizard.Hide()

	})
	mainDialogScreen := container.NewVBox(
		radioCheck,
		widget.NewLabel("name of mnemonic key"),
		mnemonicNameEntering,
		widget.NewLabel("bip39 mnemonic"),
		mnemonicEntering,
		widget.NewLabel("encryption password"),
		encryptionPassword,
		container.NewVBox(addButton, closeButton),
	)
	wizard = dialogs.NewWizard("Restore key with mnemonic", mainDialogScreen)
	wizard.Show(g.Window)
	wizard.Resize(fyne.NewSize(500, 200))

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
