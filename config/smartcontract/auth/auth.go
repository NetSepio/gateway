package auth

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/envutil"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

var (
	auth       *bind.TransactOpts
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	chainID    *big.Int
	err        error
)

//Return error when it occurs
func GetAuth(client *ethclient.Client) (*bind.TransactOpts, error) {

	if auth == nil {
		// mnemonic := envutil.MustGetEnv("MNEMONIC")
		privateKeyHex := envutil.MustGetEnv("PRIVATE_KEY")
		privateKey, err := crypto.HexToECDSA(privateKeyHex)
		publicKey = &privateKey.PublicKey
		// privateKey, publicKey, _, err = ethwallet.HdWallet(mnemonic) // Verify: https://iancoleman.io/bip39/
		if err != nil {
			logwrapper.Errorf("error while getting private and public keu from mnemonic, error: %v", err.Error())
			return nil, err
		}
		chainID, err = client.NetworkID(context.Background())
		if err != nil {
			logwrapper.Errorf("failed to get network id, error: %v", err.Error())
			return nil, err
		}
		auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
		if err != nil {
			logwrapper.Errorf("failed to call NewKeyedTransactorWithChainID, error: %v", err.Error())
			return nil, err
		}
		auth.Value = big.NewInt(0)           // in wei
		auth.GasLimit = 0.0001 * params.GWei // in units
	}

	blcNo, err := client.BlockNumber(context.Background())
	if err != nil {
		logwrapper.Errorf("failed to get block number, error: %v", err.Error())
		return nil, err
	}
	nonce, err := client.NonceAt(context.Background(), crypto.PubkeyToAddress(*publicKey), big.NewInt(int64(blcNo)))
	// nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(*publicKey))
	if err != nil {
		logwrapper.Errorf("failed to get nonce, error: %v", err)
		return nil, err
	}
	logwrapper.Infof("nonce is %v", nonce)

	gasTipCap, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		logwrapper.Errorf("error while calling %v, error: %v", "client.SuggestGasTipCap", err.Error())
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasTipCap = gasTipCap
	auth.GasFeeCap = big.NewInt(0).Mul(gasTipCap, big.NewInt(6))
	return auth, nil
}
