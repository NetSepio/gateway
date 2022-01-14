package smartcontract

import (
	"os"

	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func GetClient() *ethclient.Client {
	if client != nil {
		return client
	}
	nodeUrl := os.Getenv("NODE_URL")
	var err error
	client, err = ethclient.Dial(nodeUrl)
	if err != nil {
		logwrapper.Fatalf("failed to dial client at url %v, error: %v", nodeUrl, err.Error())
	}
	return client
}
