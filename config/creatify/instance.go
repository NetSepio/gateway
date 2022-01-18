package creatify

import (
	"errors"
	"os"

	"github.com/TheLazarusNetwork/marketplace-engine/generated/smartcontract/creatify"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var instance *creatify.Creatify

var (
	errEnvVariableNotDefined = errors.New("environment variable CREATIFY_CONTRACT_ADDRESS is required")
)

func GetInstance(client *ethclient.Client) (*creatify.Creatify, error) {
	if instance != nil {
		return instance, nil
	}
	envContractAddress := os.Getenv("CREATIFY_CONTRACT_ADDRESS")
	if envContractAddress == "" {
		logwrapper.Errorf("environment variable %v is required", "CREATIFY_CONTRACT_ADDRESS")
		return nil, errEnvVariableNotDefined
	}
	addr := common.HexToAddress(envContractAddress)
	var err error
	instance, err = creatify.NewCreatify(addr, client)
	if err != nil {
		logwrapper.Errorf("failed to load creatify contract at address %v, error: %v", envContractAddress, err.Error())
		return nil, err
	}
	return instance, nil
}
