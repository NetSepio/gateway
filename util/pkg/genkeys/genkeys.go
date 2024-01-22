package genkeys

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
)

func GenerateWireGuardKeys() (privateKey, publicKey string, err error) {
	// Generate a new Ed25519 key pair
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return "", "", err
	}

	// Encode the private key in base64
	privKeyBase64 := base64.StdEncoding.EncodeToString(privKey.Seed())

	// Encode the public key in base64
	pubKeyBase64 := base64.StdEncoding.EncodeToString(pubKey)

	return privKeyBase64, pubKeyBase64, nil
}
