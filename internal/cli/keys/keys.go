package keys

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/spf13/cobra"
)

// log is the logger instance for this package.
var log = logging.Log

// Generate returns a cobra.Command that generates RSA or ECDSA keys
// with given length and output directory. The command's flags allow
// specification of key type, key length, and output directory.
func Generate() *cobra.Command {
	log.Debugln("Adding `keys` command...")

	keysCmd := &cobra.Command{
		Use:   "keys",
		Short: "Generates RSA or ECDSA keys",
		Long:  "Generates RSA or ECDSA keys with given length and output directory",
		RunE: func(cmd *cobra.Command, args []string) error {
			keyType, _ := cmd.Flags().GetString("type")
			keyLength, _ := cmd.Flags().GetInt("length")
			outDir, _ := cmd.Flags().GetString("out")

			switch keyType {
			case "rsa":
				return generateRSA(keyLength, outDir)
			case "ecdsa":
				return generateECDSA(keyLength, outDir)
			default:
				log.Errorf("invalid user input: %v\n", keyType)
				return fmt.Errorf("invalid key type: %s", keyType)
			}
		},
	}

	keysCmd.Flags().StringP("type", "t", "rsa", "Type of keys to generate (rsa, ecdsa)")
	keysCmd.Flags().IntP("length", "l", 2048, "Length of keys to generate")
	keysCmd.Flags().StringP("out", "o", ".", "Output directory for keys")

	return keysCmd
}

// generateRSA generates an RSA key pair with the given number of bits and
// writes the keys to files in the specified output directory.
func generateRSA(bits int, outDir string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		log.Errorln("failed to generate RSA private key...")
		return err
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: privateKeyBytes})
	err = ioutil.WriteFile(filepath.Join(outDir, "private.pem"), privateKeyPEM, 0600)
	if err != nil {
		log.Errorf("failed writing RSA private key to file %v\n", filepath.Join(outDir, "private.pem"))
		return err
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Errorf("failed to marshal RSA public key\n")
		return err
	}

	publicKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: publicKeyBytes})
	return ioutil.WriteFile(filepath.Join(outDir, "public.pem"), publicKeyPEM, 0644)
}

// generateECDSA generates an ECDSA key pair with the given curve size and
// writes the keys to files in the specified output directory.
func generateECDSA(curveBits int, outDir string) error {
	var curve elliptic.Curve
	switch curveBits {
	case 224:
		curve = elliptic.P224()
	case 256:
		curve = elliptic.P256()
	case 384:
		curve = elliptic.P384()
	case 521:
		curve = elliptic.P521()
	default:
		log.Errorf("invalid user input `curve size`: %v\n", curveBits)
		return fmt.Errorf("invalid curve size: %d", curveBits)

	}

	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Errorln("failed to generate ECDSA private key")
		return err
	}

	privateKeyBytes, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		log.Errorln("failed to marshal ECDSA private key")
		return err
	}

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: privateKeyBytes})
	err = ioutil.WriteFile(filepath.Join(outDir, "private.pem"), privateKeyPEM, 0600)
	if err != nil {
		log.Errorf("failed writing ECDSA private key to file %v\n", filepath.Join(outDir, "private.pem"))
		return err
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Errorln("failed to marshal ECDSA public key")
		return err
	}

	publicKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "EC PUBLIC KEY", Bytes: publicKeyBytes})
	return ioutil.WriteFile(filepath.Join(outDir, "public.pem"), publicKeyPEM, 0644)
}
