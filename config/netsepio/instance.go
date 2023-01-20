package netsepio

import (
	"errors"

	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/generated/smartcontract/netsepio"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var instance *netsepio.Netsepio

var (
	errEnvVariableNotDefined = errors.New("environment variable NETSEPIO_CONTRACT_ADDRESS is required")
)

func GetInstance(client *ethclient.Client) (*netsepio.Netsepio, error) {
	if instance != nil {
		return instance, nil
	}
	envContractAddress := envconfig.EnvVars.NETSEPIO_CONTRACT_ADDRESS
	if envContractAddress == "" {
		logwrapper.Errorf("environment variable %v is required", "NETSEPIO_CONTRACT_ADDRESS")
		return nil, errEnvVariableNotDefined
	}
	addr := common.HexToAddress(envContractAddress)
	var err error
	instance, err = netsepio.NewNetsepio(addr, client)
	if err != nil {
		logwrapper.Errorf("failed to load netsepio contract at address %v, error: %v", envContractAddress, err.Error())
		return nil, err
	}
	return instance, nil
}
