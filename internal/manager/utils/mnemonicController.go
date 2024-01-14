package utils

import (
	"bufio"
	"context"
	"fmt"
	"os"

	vlg "github.com/PeepoFrog/validator-key-gen/MnemonicsGenerator"
	"github.com/joho/godotenv"
	kiraMnemonicGen "github.com/kiracore/tools/bip39gen/cmd"
	"github.com/kiracore/tools/bip39gen/pkg/bip39"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/osutils"
)

func (h *HelperManager) ReadMnemonicsFromFile(pathToFile string) (mastermnemonic string, err error) {
	log := logging.Log
	log.Println("checking if path exist: ", pathToFile)
	check, err := osutils.CheckItPathExist(pathToFile)
	if err != nil {
		log.Printf("error while checkin path to %s, error: %s", pathToFile, err)
	}
	if check {
		log.Println("path exist, trying to read mnemonic from mnemonics.env file")
		if err := godotenv.Load(pathToFile); err != nil {
			err = fmt.Errorf("error loading .env file: %v", err)
			return mastermnemonic, err
		}
		// Retrieve the MASTER_MNEMONIC value
		mastermnemonic = os.Getenv("MASTER_MNEMONIC")
		if mastermnemonic == "" {
			err = fmt.Errorf("MASTER_MNEMONIC not found")
			return mastermnemonic, err
		} else {
			log.Debugln("MASTER_MNEMONIC:", mastermnemonic)
		}
	}

	return mastermnemonic, nil
}

func (h *HelperManager) GenerateMnemonicsFromMaster(masterMnemonic string) (*vlg.MasterMnemonicSet, error) {
	log := logging.Log
	log.Debugf("GenerateMnemonicFromMaster: masterMnemonic:\n%s", masterMnemonic)
	defaultprefix := "kira"
	defaultPath := "44'/118'/0'/0/0"

	// masterMnemonic = "want vanish frown filter resemble purchase trial baby equal never cinnamon claim wrap cash snake cable head tray few daring shine clip loyal series"

	mnemonicSet, err := vlg.MasterKeysGen([]byte(masterMnemonic), defaultprefix, defaultPath, h.config.SecretsFolder)
	if err != nil {
		return &vlg.MasterMnemonicSet{}, err
	}
	str := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", mnemonicSet.SignerAddrMnemonic, mnemonicSet.ValidatorNodeMnemonic, mnemonicSet.ValidatorNodeId, mnemonicSet.ValidatorAddrMnemonic, mnemonicSet.ValidatorValMnemonic)
	fmt.Println((str))
	return &mnemonicSet, nil
}

func (h *HelperManager) MnemonicReader() (masterMnemonic string) {
	log := logging.Log
	fmt.Printf("\nENTER YOUR MASTER MNEMONIC:\n")
	// var input string
	reader := bufio.NewReader(os.Stdin)
	log.Println("Enter mnemonic: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred:", err)
		return
	}
	mnemonicBytes := []byte(text)
	mnemonicBytes = mnemonicBytes[0 : len(mnemonicBytes)-1]
	masterMnemonic = string(mnemonicBytes)
	return masterMnemonic
}

// generates random bip 24 word mnemonic
func (h *HelperManager) GenerateMnemonic() (masterMnemonic bip39.Mnemonic, err error) {
	// bip39.Mnemonic
	masterMnemonic = kiraMnemonicGen.NewMnemonic()
	masterMnemonic.SetRandomEntropy(24)
	masterMnemonic.Generate()

	return masterMnemonic, nil
}

func (h *HelperManager) SetSekaidKeys(ctx context.Context) error {
	log := logging.Log
	sekaidConfigFolder := h.config.SekaidHome + "/config"
	// err := h.containerManager.SendFileToContainer(ctx, h.config.SecretsFolder+"/priv_validator_key.json", sekaidConfigFolder+"/priv_validator_key.json", h.config.SekaidContainerName)
	_, err := h.containerManager.ExecCommandInContainer(ctx, h.config.SekaidContainerName, []string{"bash", "-c", fmt.Sprintf(`mkdir %s`, h.config.SekaidHome)})
	if err != nil {
		return fmt.Errorf("unable to create <%s> folder, err: %s", h.config.SekaidHome, err)
	}
	_, err = h.containerManager.ExecCommandInContainer(ctx, h.config.SekaidContainerName, []string{"bash", "-c", fmt.Sprintf(`mkdir %s`, sekaidConfigFolder)})
	if err != nil {
		return fmt.Errorf("unable to create <%s> folder, err: %s", sekaidConfigFolder, err)
	}
	err = h.containerManager.SendFileToContainer(ctx, h.config.SecretsFolder+"/priv_validator_key.json", sekaidConfigFolder, h.config.SekaidContainerName)
	if err != nil {
		log.Errorf("cannot send priv_validator_key.json to container\n")
		return err
	}

	osutils.CopyFile(h.config.SecretsFolder+"/validator_node_key.json", h.config.SecretsFolder+"/node_key.json")
	err = h.containerManager.SendFileToContainer(ctx, h.config.SecretsFolder+"/node_key.json", sekaidConfigFolder, h.config.SekaidContainerName)
	if err != nil {
		log.Errorf("cannot send validator_node_key.json to container\n")
		return err
	}
	return nil
}

// sets empty state of validator into $sekaidHome/data/priv_validator_state.json
func (h *HelperManager) SetEmptyValidatorState(ctx context.Context) error {
	emptyState := `
	{
		"height": "0",
		"round": 0,
		"step": 0
	}`
	// TODO
	// mount docker volume to the folder on host machine and do file manipulations inside this folder
	tmpFilePath := "/tmp/priv_validator_state.json"
	err := osutils.CreateFileWithData(tmpFilePath, []byte(emptyState))
	if err != nil {
		return fmt.Errorf("unable to create file <%s>, error: %s", tmpFilePath, err)
	}
	sekaidDataFoder := h.config.SekaidHome + "/data"
	_, err = h.containerManager.ExecCommandInContainer(ctx, h.config.SekaidContainerName, []string{"bash", "-c", fmt.Sprintf(`mkdir %s`, sekaidDataFoder)})
	if err != nil {
		return fmt.Errorf("unable to create folder <%s>, error: %s", sekaidDataFoder, err)
	}
	err = h.containerManager.SendFileToContainer(ctx, tmpFilePath, sekaidDataFoder, h.config.SekaidContainerName)
	if err != nil {
		return fmt.Errorf("cannot send %s to container, err: %s", tmpFilePath, err)
	}
	return nil
}
