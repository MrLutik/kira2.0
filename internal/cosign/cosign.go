package cosign

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"

	"github.com/google/go-containerregistry/pkg/name"
	"github.com/mrlutik/kira2.0/internal/logging"
	"github.com/sigstore/cosign/pkg/cosign"
	"github.com/sigstore/sigstore/pkg/signature"
)

var (
	log = logging.Log

	ErrInvalidPublicKeyType = errors.New("public key is not of type *ecdsa.PublicKey")
	ErrDecodingPublicKey    = errors.New("failed to decode public key")
)

func VerifyImageSignature(ctx context.Context, imageRef, pubKey string) (bool, error) {
	log.Infof("Verifying image '%s' signature\n", imageRef)
	// Decode the PEM-encoded public key data
	block, _ := pem.Decode([]byte(pubKey))
	if block == nil {
		log.Errorln("Failed to decode public key")
		return false, ErrDecodingPublicKey
	}

	// Parse the public key from the decoded PEM block
	// x509.ParsePKIXPublicKey is used for parsing PKIX public keys, which include RSA, DSA, and ECDSA public keys
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Errorln("Failed to parse public key")
		return false, fmt.Errorf("failed to parse public key: %w", err)
	}

	// Get ecdsa.PublicKey by type assertion
	ecdsaPubKey, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		log.Errorln("Public key is not of type *ecdsa.PublicKey")
		return false, ErrInvalidPublicKeyType
	}

	// Get image reference
	ref, err := name.ParseReference(imageRef)
	if err != nil {
		log.Errorf("Failed to parse image reference: %s", err)
		return false, fmt.Errorf("failed to parse image reference: %w", err)
	}

	// Create a signature verifier for an ECDSA signature algorithm using a public key
	// and the SHA256 cryptographic hash function, and then setting the signature verifier as an option for verifying a signed image.
	verifier, err := signature.LoadECDSAVerifier(ecdsaPubKey, crypto.SHA256)
	if err != nil {
		log.Errorf("Failed to load ECDSA verifier: %s", err)
		return false, fmt.Errorf("failed to load ECDSA verifier: %w", err)
	}
	co := &cosign.CheckOpts{
		SigVerifier: verifier,
	}

	signatures, verified, err := cosign.VerifyImageSignatures(ctx, ref, co)
	if err != nil {
		log.Errorf("Failed verifying signatures: %s", err)
		return false, fmt.Errorf("failed verifying signatures: %w", err)
	}
	for _, sig := range signatures {
		base64Sig, err := sig.Base64Signature()
		if err != nil {
			log.Errorf("Failed to get Base64 Signature: %s", err)
			continue // Skip this signature and move to the next one
		}
		fmt.Fprintf(os.Stdout, "Signature: %s\n", base64Sig)

		payload, err := sig.Payload()
		if err != nil {
			log.Errorf("Failed to get Payload: %s", err)
			continue // Skip this signature and move to the next one
		}
		fmt.Fprintf(os.Stdout, "Payload: %s\n", string(payload))

		fmt.Fprintln(os.Stdout, "====")
	} // Maybe I will use it in future

	log.Infof("Verifying image '%s' signature was successfully\n", imageRef)
	return verified, nil
}
