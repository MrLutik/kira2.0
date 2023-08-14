package types

type MasterMnemonicSet struct {
	ValidatorAddrMnemonic []byte
	ValidatorNodeMnemonic []byte
	ValidatorNodeId       []byte
	ValidatorValMnemonic  []byte
	SignerAddrMnemonic    []byte
}
