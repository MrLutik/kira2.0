package mnemonicsgenerator_test

import (
	"reflect"
	"testing"

	mnemonicsgenerator "github.com/mrlutik/kira2.0/pkg/validatorKeyGen/MnemonicsGenerator"
)

// Test mnemonic:
//
// MASTER_MNEMONIC=bargain erosion electric skill extend aunt unfold cricket spice sudden insane shock purpose trumpet holiday tornado fiction check pony acoustic strike side gold resemble
// VALIDATOR_ADDR_MNEMONIC=result tank riot circle cost hundred exotic soft angle bulb sunset margin virus simple bean topic next initial embody sample ordinary what pulp engage
// VALIDATOR_NODE_MNEMONIC=shed history misery describe sail sight know snake route humor soda gossip lonely torch state drama salmon jungle possible lock runway wild cross tank
// VALIDATOR_NODE_ID=935ea41280fa8754a35bd2916d935f222b559488
// VALIDATOR_VAL_MNEMONIC=stick about junk liberty same envelope boy machine zoo wide shrimp clutch oval mango diary strike round divorce toilet cross guard appear govern chief
// SIGNER_ADDR_MNEMONIC=near spirit dial february access song panda clean diesel legend clock remind name pupil drum general trap afford tuition side dune address alpha stool

func TestMasterKeysGen(t *testing.T) {
	tests := []struct {
		name           string
		masterMnemonic []byte
		want           mnemonicsgenerator.MasterMnemonicSet
		wantErr        bool
	}{
		{
			name:           "working mnemonic",
			masterMnemonic: []byte("bargain erosion electric skill extend aunt unfold cricket spice sudden insane shock purpose trumpet holiday tornado fiction check pony acoustic strike side gold resemble"),
			want: mnemonicsgenerator.MasterMnemonicSet{
				ValidatorAddrMnemonic: []byte("result tank riot circle cost hundred exotic soft angle bulb sunset margin virus simple bean topic next initial embody sample ordinary what pulp engage"),
				ValidatorNodeMnemonic: []byte("shed history misery describe sail sight know snake route humor soda gossip lonely torch state drama salmon jungle possible lock runway wild cross tank"),
				ValidatorValMnemonic:  []byte("stick about junk liberty same envelope boy machine zoo wide shrimp clutch oval mango diary strike round divorce toilet cross guard appear govern chief"),
				SignerAddrMnemonic:    []byte("near spirit dial february access song panda clean diesel legend clock remind name pupil drum general trap afford tuition side dune address alpha stool"),
				ValidatorNodeId:       []byte("935ea41280fa8754a35bd2916d935f222b559488"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mnemonicsgenerator.MasterKeysGen(tt.masterMnemonic, mnemonicsgenerator.DefaultPrefix, mnemonicsgenerator.DefaultPath, "")
			if (err != nil) != tt.wantErr {
				t.Errorf("MasterKeysGen(%s) error = %v, wantErr %v", tt.masterMnemonic, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessData(%d) = %+v, want %+v", tt.masterMnemonic, got, tt.want)
			}
		})
	}
}
