package creatify

import (
	"github.com/TheLazarusNetwork/marketplace-engine/config/smartcontract"
	"github.com/TheLazarusNetwork/marketplace-engine/generated/smartcontract/creatify"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
)

var creatorRoleId [32]byte
var initiated = false

func GetCreatorRoleId() [32]byte {
	if initiated {
		return creatorRoleId
	}
	instance := creatify.GetInstance(smartcontract.GetClient())
	var err error
	creatorRoleId, err = instance.CREATORROLE(nil)
	if err != nil {
		logwrapper.Fatalf("Failed to get %v, error: %v", "CREATORROLE", err.Error())
	} else {
		initiated = true
	}
	return creatorRoleId
}
