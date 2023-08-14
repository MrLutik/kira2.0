package utils

import (
	"fmt"

	"github.com/mrlutik/kira2.0/internal/types"
)

func (h *HelperManager) GenerateMnemonics(masterMnemonic string) *types.MasterMnemonicSet {

	defaultprefix := "kira"
	defaultPath := "44'/118'/0'/0/0"

	// mnemonic := "want vanish frown filter resemble purchase trial baby equal never cinnamon claim wrap cash snake cable head tray few daring shine clip loyal series"

	mnemonicSet := MasterKeysGen([]byte(masterMnemonic), defaultprefix, defaultPath, h.config.SecretsFolder)
	str := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", mnemonicSet.SignerAddrMnemonic, mnemonicSet.ValidatorNodeMnemonic, mnemonicSet.ValidatorNodeId, mnemonicSet.ValidatorAddrMnemonic, mnemonicSet.ValidatorValMnemonic)
	fmt.Println((str))
	return mnemonicSet
}

func (h *HelperManager) SetSekaidKeys() {

}

func (h *HelperManager) SetInterxKeys() {

}

func (h *HelperManager) ReadMnemonicFromEnvFile() {

}
