package ethwallet

import (
	"crypto/ecdsa"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/tyler-smith/go-bip39"
)

func HdWallet(mnemonic string) (*ecdsa.PrivateKey, *ecdsa.PublicKey, *string, error) {
	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "")
	// Generate a new master node using the seed.
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, nil, nil, err
	}
	//fmt.Println("Master Public Key: ", masterKey.PublicKey())
	// This gives the path: m/44H
	acc44H, err := masterKey.Derive(hdkeychain.HardenedKeyStart + 44)
	if err != nil {
		return nil, nil, nil, err
	}
	// This gives the path: m/44H/60H
	acc44H60H, err := acc44H.Derive(hdkeychain.HardenedKeyStart + 60)
	if err != nil {
		return nil, nil, nil, err
	}
	// This gives the path: m/44H/60H/0H
	acc44H60H0H, err := acc44H60H.Derive(hdkeychain.HardenedKeyStart + 0)
	if err != nil {
		return nil, nil, nil, err
	}
	// This gives the path: m/44H/60H/0H/0
	acc44H60H0H0, err := acc44H60H0H.Derive(0)
	if err != nil {
		return nil, nil, nil, err
	}
	// This gives the path: m/44H/60H/0H/0/1
	acc44H60H0H00, err := acc44H60H0H0.Derive(1)
	if err != nil {
		return nil, nil, nil, err
	}
	btcecPrivKey, err := acc44H60H0H00.ECPrivKey()
	if err != nil {
		return nil, nil, nil, err
	}
	privateKey := btcecPrivKey.ToECDSA()
	publicKey := &privateKey.PublicKey // Starts with 0x04. Contains DER encoding of the public key (which is what Bitcoin and all its fork uses)
	path := "m/44H/60H/0H/0/1"
	return privateKey, publicKey, &path, nil
}
