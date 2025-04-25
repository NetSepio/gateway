package cryptosign

import (
	"encoding/hex"
	"regexp"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"github.com/sirupsen/logrus"
)

func IsValidAptosAddress(address string) bool {
	// Remove optional "0x" prefix
	address = strings.TrimPrefix(address, "0x")

	// Check for exact length (64 hex characters = 32 bytes)
	if len(address) != 64 {
		logrus.Error("Invalid Aptos address length ❌")
		return false
	}

	// Check if it's a valid hex
	_, err := hex.DecodeString(address)

	if err != nil {
		logrus.Error("Invalid Aptos address ❌")
	} else {
		return err == nil
	}

	// Default return value in case no condition is met
	return false
}

// IsValidEVMAddress verifies Ethereum/Monad-style EVM addresses
func IsValidEVM_MONAD_Address(address string) bool {
	if !strings.HasPrefix(address, "0x") || len(address) != 42 {
		logrus.Error("Invalid EVM address: must start with '0x' and be 42 characters long ❌")
		return false
	}
	_, err := hex.DecodeString(address[2:])
	if err != nil {
		logrus.Error("Invalid EVM address: failed to decode hex string ❌")
	} else {
		return true
	}
	return false
}

// IsValidPeaqAddress verifies Peaq wallet addresses using base58 and length
func IsValidPeaqAddress(address string) bool {
	// Basic regex for base58-encoded SS58 addresses (Peaq/Polkadot style)
	if !regexp.MustCompile(`^[1-9A-HJ-NP-Za-km-z]{47,48}$`).MatchString(address) {
		logrus.Error("Invalid Peaq address: does not match base58 regex ❌")
		return false
	}

	decoded := base58.Decode(address)
	return len(decoded) >= 32 // Usually 35-36 bytes depending on prefix and checksum
}

// Common for Sui and Solana: base58, 32-byte decoded
func IsValidBase58_32ByteAddress(address string) bool {
	decoded := base58.Decode(address)
	return len(decoded) == 32
}
