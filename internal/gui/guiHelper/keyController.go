package guiHelper

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"os"

	"github.com/cosmos/go-bip39"
	"github.com/mrlutik/kira2.0/internal/osutils"
	"golang.org/x/crypto/chacha20"
)

// todo Create more fancy nonce and place it in types
const Nonce string = "24CharacterNonce!!!!!!!!"

type EncryptedMnemonic struct {
	Name               string `json:"name"`
	EncoderMnemonicHex string `json:"encoderMnemonic"`
}

func getEncryptedMnemonicsFilePath() (string, error) {
	homeFolder, err := GetAndSetHomeFolderForKMUI()
	if err != nil {
		return "", err
	}
	fileToWrite := homeFolder + "/encryptedMnemonics.json"
	check, err := osutils.CheckIfFileExist(fileToWrite)
	if err != nil {
		log.Errorf("error when checking if %s file exist: %s", fileToWrite, err)
		return "", err
	}
	if !check {
		_, err := os.Create(fileToWrite)
		if err != nil {
			log.Errorf("error when creating if %s file: %s", fileToWrite, err)
			return "", err
		}
		// defer file.Close()
	}

	return fileToWrite, nil
}
func setKeys(eMnemonics []EncryptedMnemonic) error {
	path, err := getEncryptedMnemonicsFilePath()
	if err != nil {
		log.Errorf("error when getting mnemonic file", err)
		return err
	}
	// defer file.Close()

	jsonData, err := json.Marshal(eMnemonics)
	if err != nil {
		log.Errorf("error when marshaling %v into mnemonic file: %s", eMnemonics, err)
		return err
	}

	// Create and write to file
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Error("Error when opening %s file: %s", path, err)
	}
	defer file.Close()
	_, err = file.Write(jsonData)
	if err != nil {
		log.Errorf("error when writing %s data to mnemonic file: %s", jsonData, err)
		return err
	}
	return nil
}
func GetKeys() ([]EncryptedMnemonic, error) {
	path, err := getEncryptedMnemonicsFilePath()
	if err != nil {
		log.Errorf("error when getEncryptedMnemonicsFile file: %s", err)
		return nil, err
	}
	file, err := os.Open(path)
	if err != nil {
		log.Error("Error when opening %s file: %s", path, err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		log.Errorf("error while reading %v file: %s", file, err)
		return nil, err
	}

	var mnemonics []EncryptedMnemonic
	if len(data) == 0 {
		return mnemonics, nil
	}
	err = json.Unmarshal(data, &mnemonics)
	if err != nil {
		log.Errorf("error while unmarshaling mnemonic file: %s", err)
		return nil, err
	}
	return mnemonics, nil
}

func GetKey(name string) (EncryptedMnemonic, error) {
	keys, err := GetKeys()
	if err != nil {
		log.Errorf("error when getting key with %s name: %s", name, err)
		return EncryptedMnemonic{}, err
	}
	for _, k := range keys {
		if k.Name == name {
			return k, nil
		}
	}
	return EncryptedMnemonic{}, fmt.Errorf("key with %s was not found", name)
}

func AddKey(eMnemonic EncryptedMnemonic) error {
	keys, err := GetKeys()
	if err != nil {
		log.Errorf("error while adding key: %s", err)
		return err
	}
	keys = append(keys, eMnemonic)
	err = setKeys(keys)
	if err != nil {
		log.Errorf("error while adding key: %s", err)
		return err
	}
	return nil
}
func GeneratePrivateP256KeyFromMnemonic(mnemonic string) (*ecdsa.PrivateKey, error) {
	// Generate seed from mnemonic
	seed := bip39.NewSeed(mnemonic, "")

	// Use the seed to generate an ECDSA private key
	// Todo: is this a safe approach? Need security investigation
	privateKey, err := seedToPrivateKey(seed)
	if err != nil {
		fmt.Println("Error creating private key from seed:", err)
		return nil, err
	}
	privateKey, err = seedToPrivateKey(seed)
	if err != nil {
		fmt.Println("Error creating private key from seed:", err)
		return nil, err
	}
	return privateKey, nil
}

func seedToPrivateKey(seed []byte) (*ecdsa.PrivateKey, error) {
	curve := elliptic.P256()
	privKey := new(ecdsa.PrivateKey)
	privKey.PublicKey.Curve = curve
	privKey.D = new(big.Int).SetBytes(seed[:32]) // Using the first 32 bytes of the seed
	privKey.PublicKey.X, privKey.PublicKey.Y = curve.ScalarBaseMult(seed[:32])
	return privKey, nil
}

// Encrypt mnemonic to chacha20 with
// Key should be 32 byte long
// nonce should be 24 byte long
func EncryptMnemonic(mnemonic string, key, nonce []byte) (encryptedMnemonic []byte, err error) {
	//Todo: nonce should be hardcoded maybe inside "types" package
	// Create a new ChaCha20 cipher
	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return nil, err
	}
	// Encrypt the mnemonic
	encryptedMnemonic = make([]byte, len(mnemonic))
	cipher.XORKeyStream(encryptedMnemonic, []byte(mnemonic))
	return encryptedMnemonic, nil
}

// Decrypt Mnemonic from chacha20
// keyHex := "your 256-bit key in hex"
// nonceHex := "your nonce in hex"
func DecryptMnemonic(encryptedMnemonicHex, keyHex, nonceHex string) []byte {
	key, err := hex.DecodeString(keyHex)
	if err != nil {
		log.Fatalf("error when decoding keyHex: %s", err)
	}
	nonce, err := hex.DecodeString(nonceHex)
	if err != nil {
		log.Fatalf("error when decoding nonceHex %s", err)
	}
	// Encrypted data (hexadecimal format)
	// encryptedHex := "your encrypted data in hex"
	encrypted, err := hex.DecodeString(encryptedMnemonicHex)
	if err != nil {
		log.Fatalf("error when decoding encryptedHex %s", err)
	}
	// Create a new ChaCha20 cipher for decryption
	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		log.Fatal(err)
	}
	// Decrypt the data
	decrypted := make([]byte, len(encrypted))
	cipher.XORKeyStream(decrypted, encrypted)
	return decrypted
}

func Set32BytePassword(passw string) ([]byte, error) {
	if len(passw) > 32 {
		return []byte(""), fmt.Errorf("password is to large")
	} else if len(passw) == 32 {
		return []byte(passw), nil
	}
	key := make([]byte, 32)
	i := 0
	for i < len(passw) {
		key[i] = passw[i]
		i++
	}
	for i < 32 {
		key[i] = 0
		i++
	}
	return key, nil
}
