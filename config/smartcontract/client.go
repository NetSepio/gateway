package smartcontract

import (
	"os"

	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func GetClient() (*ethclient.Client, error) {
	if client != nil {
		return client, nil
	}
	nodeUrl := os.Getenv("POLYGON_RPC")
	var err error
	client, err = ethclient.Dial(nodeUrl)
	if err != nil {
		logwrapper.Errorf("failed to dial client at url %v, error: %v", nodeUrl, err.Error())
		return nil, err
	}
	return client, nil
}
