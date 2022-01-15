package creatify

import (
	"os"

	"github.com/TheLazarusNetwork/marketplace-engine/generated/smartcontract/creatify"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var instance *creatify.Creatify

func GetInstance(client *ethclient.Client) *creatify.Creatify {
	if instance != nil {
		return instance
	}
	envContractAddress := os.Getenv("CREATIFY_CONTRACT_ADDRESS")
	if envContractAddress == "" {
		logwrapper.Fatalf("environment variable %v is required", "CREATIFY_CONTRACT_ADDRESS")
	}
	addr := common.HexToAddress(envContractAddress)
	var err error
	instance, err = creatify.NewCreatify(addr, client)
	if err != nil {
		logwrapper.Fatalf("failed to load creatify contract at address %v, error: %v", envContractAddress, err.Error())
	}
	return instance
}
