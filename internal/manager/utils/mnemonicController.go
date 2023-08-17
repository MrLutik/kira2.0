package utils

import (
	"bufio"
	"context"
	"fmt"
	"os"

	kiraMnemonicGen "github.com/kiracore/tools/bip39gen/cmd"
	"github.com/kiracore/tools/bip39gen/pkg/bip39"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/mrlutik/kira2.0/internal/types"
)

func (h *HelperManager) GenerateMnemonicsFromMaster(masterMnemonic string) *types.MasterMnemonicSet {
	log := logging.Log
	log.Debugf("GenerateMnemonicFromMaster: masterMnemonic:\n%s", masterMnemonic)
	defaultprefix := "kira"
	defaultPath := "44'/118'/0'/0/0"

	// masterMnemonic = "want vanish frown filter resemble purchase trial baby equal never cinnamon claim wrap cash snake cable head tray few daring shine clip loyal series"

	mnemonicSet := MasterKeysGen([]byte(masterMnemonic), defaultprefix, defaultPath, h.config.SecretsFolder)
	str := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", mnemonicSet.SignerAddrMnemonic, mnemonicSet.ValidatorNodeMnemonic, mnemonicSet.ValidatorNodeId, mnemonicSet.ValidatorAddrMnemonic, mnemonicSet.ValidatorValMnemonic)
	fmt.Println((str))
	return mnemonicSet
}

func (h *HelperManager) MnemonicReader() (masterMnemonic string) {
	log := logging.Log
	log.Printf("\nENTER YOUR MASTER MNEMONIC:\n")
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
	err := h.containerManager.SendFileToContainer(ctx, h.config.SecretsFolder+"/priv_validator_key.json", sekaidConfigFolder, h.config.SekaidContainerName)
	if err != nil {
		log.Errorf("cannot send priv_validator_key.json to container\n")
		return err
	}

	h.copyFile(h.config.SecretsFolder+"/validator_node_key.json", h.config.SecretsFolder+"/node_key.json")
	err = h.containerManager.SendFileToContainer(ctx, h.config.SecretsFolder+"/node_key.json", sekaidConfigFolder, h.config.SekaidContainerName)
	if err != nil {
		log.Errorf("cannot send validator_node_key.json to container\n")
		return err
	}
	return nil
}
