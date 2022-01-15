package auth

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/ethwallet"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

/* Global Auth variable can be accessed for auth purposes in smart contract transactions */
var auth *bind.TransactOpts

func GetAuth(client *ethclient.Client) *bind.TransactOpts {

	if auth != nil {
		//Auth is already initialized, return that object
		return auth
	}

	mnemonic := os.Getenv("MNEMONIC")
	privateKey, publicKey, _, err := ethwallet.HdWallet(mnemonic) // Verify: https://iancoleman.io/bip39/
	if err != nil {
		fmt.Printf("Error: %+v", err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(*publicKey))
	if err != nil {
		log.Fatal(err)
	}

	gasLimit := uint64(200000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		logwrapper.Fatalf("failed to get network id, error: %v", err.Error())
	}
	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		logwrapper.Fatalf("failed to call NewKeyedTransactorWithChainID, error: %v", err.Error())
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = gasLimit   // in units
	auth.GasPrice = gasPrice
	return auth
}
