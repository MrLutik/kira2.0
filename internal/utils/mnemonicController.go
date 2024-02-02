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
)

func (h *HelperManager) ReadMnemonicsFromFile(pathToFile string) (masterMnemonic string, err error) {
	h.log.Infof("Checking if path exist: %s", pathToFile)
	check, err := h.utils.CheckIfPathExists(pathToFile)
	if err != nil {
		h.log.Errorf("Checking path to '%s', error: %s", pathToFile, err)
	}
	if check {
		h.log.Infof("Path exist, trying to read mnemonic from mnemonics.env file")
		if err = godotenv.Load(pathToFile); err != nil {
			err = fmt.Errorf("error loading .env file: %w", err)
			return "", err
		}
		// Retrieve the MASTER_MNEMONIC value
		const masterMnemonicEnv = "MASTER_MNEMONIC"
		masterMnemonic = os.Getenv(masterMnemonicEnv)
		if masterMnemonic == "" {
			err = &EnvVariableNotFoundError{VariableName: masterMnemonicEnv}
			return masterMnemonic, err
		} else {
			h.log.Debugf("MASTER_MNEMONIC: %s", masterMnemonic)
		}
	}

	return masterMnemonic, nil
}

func (h *HelperManager) GenerateMnemonicsFromMaster(masterMnemonic string) (*vlg.MasterMnemonicSet, error) {
	h.log.Debugf("GenerateMnemonicFromMaster: masterMnemonic:\n%s", masterMnemonic)
	defaultPrefix := "kira"
	defaultPath := "44'/118'/0'/0/0"

	mnemonicSet, err := vlg.MasterKeysGen([]byte(masterMnemonic), defaultPrefix, defaultPath, h.config.SecretsFolder)
	if err != nil {
		return &vlg.MasterMnemonicSet{}, err
	}
	str := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", mnemonicSet.SignerAddrMnemonic, mnemonicSet.ValidatorNodeMnemonic, mnemonicSet.ValidatorNodeId, mnemonicSet.ValidatorAddrMnemonic, mnemonicSet.ValidatorValMnemonic)
	h.log.Infof("Master mnemonic:\n%s", str)
	return &mnemonicSet, nil
}

func (h *HelperManager) MnemonicReader() (masterMnemonic string) {
	h.log.Infoln("ENTER YOUR MASTER MNEMONIC:")

	reader := bufio.NewReader(os.Stdin)
	//nolint:forbidigo // reading user input
	fmt.Println("Enter mnemonic: ")

	text, err := reader.ReadString('\n')
	if err != nil {
		h.log.Errorf("An error occurred: %s", err)
		return
	}
	mnemonicBytes := []byte(text)
	mnemonicBytes = mnemonicBytes[0 : len(mnemonicBytes)-1]
	masterMnemonic = string(mnemonicBytes)
	return masterMnemonic
}

// GenerateMnemonic generates random bip 24 word mnemonic
func (h *HelperManager) GenerateMnemonic() (masterMnemonic bip39.Mnemonic, err error) {
	masterMnemonic = kiraMnemonicGen.NewMnemonic()
	masterMnemonic.SetRandomEntropy(24)
	masterMnemonic.Generate()

	return masterMnemonic, nil
}

func (h *HelperManager) SetSekaidKeys(ctx context.Context) error {
	// TODO path set as variables or constants
	sekaidConfigFolder := h.config.SekaidHome + "/config"
	_, err := h.containerExecutor.ExecCommandInContainer(ctx, h.config.SekaidContainerName, []string{"bash", "-c", fmt.Sprintf(`mkdir %s`, h.config.SekaidHome)})
	if err != nil {
		return fmt.Errorf("unable to create <%s> folder, err: %w", h.config.SekaidHome, err)
	}
	_, err = h.containerExecutor.ExecCommandInContainer(ctx, h.config.SekaidContainerName, []string{"bash", "-c", fmt.Sprintf(`mkdir %s`, sekaidConfigFolder)})
	if err != nil {
		return fmt.Errorf("unable to create <%s> folder, err: %w", sekaidConfigFolder, err)
	}
	err = h.containerFileManager.SendFileToContainer(ctx, h.config.SecretsFolder+"/priv_validator_key.json", sekaidConfigFolder, h.config.SekaidContainerName)
	if err != nil {
		h.log.Errorf("cannot send priv_validator_key.json to container")
		return err
	}

	err = h.utils.CopyFile(h.config.SecretsFolder+"/validator_node_key.json", h.config.SecretsFolder+"/node_key.json")
	if err != nil {
		h.log.Errorf("copying file error: %s", err)
		return err
	}

	err = h.containerFileManager.SendFileToContainer(ctx, h.config.SecretsFolder+"/node_key.json", sekaidConfigFolder, h.config.SekaidContainerName)
	if err != nil {
		h.log.Errorf("cannot send node_key.json to container")
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
	err := h.utils.CreateFileWithData(tmpFilePath, []byte(emptyState))
	if err != nil {
		return fmt.Errorf("unable to create file <%s>, error: %w", tmpFilePath, err)
	}
	sekaidDataFolder := h.config.SekaidHome + "/data"
	_, err = h.containerExecutor.ExecCommandInContainer(ctx, h.config.SekaidContainerName, []string{"bash", "-c", fmt.Sprintf(`mkdir %s`, sekaidDataFolder)})
	if err != nil {
		return fmt.Errorf("unable to create folder <%s>, error: %w", sekaidDataFolder, err)
	}
	err = h.containerFileManager.SendFileToContainer(ctx, tmpFilePath, sekaidDataFolder, h.config.SekaidContainerName)
	if err != nil {
		return fmt.Errorf("cannot send %s to container, err: %w", tmpFilePath, err)
	}
	return nil
}
