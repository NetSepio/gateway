package rawtransaction

import (
	"context"
	"math/big"
	"strings"

	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/config/smartcontract"
	"github.com/NetSepio/gateway/util/pkg/ethwallet"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

func SendRawTransaction(contractABI string, method string, args ...interface{}) (*types.Transaction, error) {

	abiParsed, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		logwrapper.Errorf("failed to parse JSON abi, error %v", err)
		return nil, err
	}
	client, err := smartcontract.GetClient()
	if err != nil {
		return nil, err
	}
	mnemonic := envconfig.EnvVars.MNEMONIC
	privateKey, publicKey, _, err := ethwallet.HdWallet(mnemonic) // Verify: https://iancoleman.io/bip39/
	if err != nil {
		logwrapper.Errorf("failed to get private and public key from mnemonic, error %v", err.Error())
		return nil, err
	}

	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(*publicKey))
	if err != nil {
		logwrapper.Warnf("failed to get nonce")
		return nil, err
	}
	envContractAddress := envconfig.EnvVars.NETSEPIO_CONTRACT_ADDRESS

	toAddress := common.HexToAddress(envContractAddress)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		logwrapper.Errorf("failed to call client.NetworkID, error: %v", err.Error())
		return nil, err
	}

	bytesData, err := abiParsed.Pack(method, args...)
	if err != nil {
		logwrapper.Errorf("failed to pack trasaction of method %v, error: %v", method, err)
		return nil, err
	}

	logwrapper.Infof("nonce is %v", nonce)

	maxPriorityFeePerGas, err := client.SuggestGasTipCap(context.Background())
	if err != nil {
		logwrapper.Errorf("failed to suggestGasTipCap, error %v", err)
		return nil, err
	}
	config := &params.ChainConfig{
		ChainID: big.NewInt(80001),
	}
	bn, _ := client.BlockNumber(context.Background())

	bignumBn := big.NewInt(0).SetUint64(bn)
	blk, _ := client.BlockByNumber(context.Background(), bignumBn)
	baseFee := misc.CalcBaseFee(config, blk.Header())
	big2 := big.NewInt(2)
	mulRes := big.NewInt(0).Mul(baseFee, big2)
	maxFeePerGas := big.NewInt(0).Add(mulRes, maxPriorityFeePerGas)
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasFeeCap: maxFeePerGas,
		GasTipCap: maxPriorityFeePerGas,
		Gas:       1310000,
		To:        &toAddress,
		Data:      bytesData,
	})
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	if err != nil {
		logwrapper.Errorf("failed to sign trasaction %v, error: %v", tx, err.Error())
		return nil, err
	}

	err = client.SendTransaction(context.TODO(), signedTx)
	if err != nil {
		logwrapper.Error("failed to send trasaction, error: ", err)
		return nil, err
	}
	return signedTx, nil
}
