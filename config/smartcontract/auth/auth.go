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
var Auth *bind.TransactOpts

func InitAuth(client *ethclient.Client) *bind.TransactOpts {

	if Auth != nil {
		//Auth is already initialized, return that object
		return Auth
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

	gasLimit := uint64(21000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		logwrapper.Fatalf("failed to get network id, error: %v", err.Error())
	}
	Auth, err = bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		logwrapper.Fatalf("failed to call NewKeyedTransactorWithChainID, error: %v", err.Error())
	}
	Auth.Nonce = big.NewInt(int64(nonce))
	Auth.Value = big.NewInt(0) // in wei
	Auth.GasLimit = gasLimit   // in units
	Auth.GasPrice = gasPrice
	return Auth
}
