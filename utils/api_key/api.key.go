package api_key

import (
	"crypto/rand"
	"encoding/hex"

	"crypto/sha256"
	"encoding/binary"

	"github.com/williepotgieter/keymaker"
)

func GenerateAPIKey() string {
	key := make([]byte, 32) // 256-bit key
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(key)
}

func GenerateAPIKey1(id string) string {

	key := keymaker.ApiKey{}
	key.Label = "skp"
	secret := make([]byte, 32) // 256-bit key
	_, err := rand.Read(secret)
	if err != nil {
		panic(err)
	}
	key.Secret = hex.EncodeToString(secret)
	// If keymaker.GenerateChecksum does not exist, use a simple checksum (e.g., SHA256)
	// or replace with the correct function from keymaker if available.
	// Example using SHA256:

	checksum := sha256.Sum256(secret)

	keys := make([]byte, 32) // 256-bit key
	_, err1 := rand.Read(checksum[:])
	if err1 != nil {
		panic(err1)
	}
	hex.EncodeToString(keys)

	// Convert the first 4 bytes of the checksum to uint32

	key.Checksum = binary.BigEndian.Uint32(checksum[0:4])

	return key.String() // or return the desired value
}
