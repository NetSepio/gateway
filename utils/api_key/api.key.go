package api_key

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateAPIKey() string {
	key := make([]byte, 32) // 256-bit key
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(key)
}
