package flowid

import (
	"netsepio-api/db"
	"netsepio-api/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GenerateFlowId(walletAddress string, update bool, flowIdType models.FlowIdType) (string, error) {
	flowId := uuid.NewString()
	if update {
		// User exist so update
		association := db.Db.Model(&models.User{
			WalletAddress: walletAddress,
		}).Association("FlowIds")
		if err := association.Error; err != nil {
			logrus.Error(err)
			return "", err
		}
		association.Append(&models.FlowId{FlowIdType: flowIdType, WalletAddress: walletAddress, FlowId: flowId})

	} else {
		// User doesn't exist so create
		newUser := &models.User{
			WalletAddress: walletAddress,
			FlowIds: []models.FlowId{{
				FlowIdType: flowIdType, WalletAddress: walletAddress, FlowId: flowId,
			},
			},
		}
		if err := db.Db.Create(newUser).Error; err != nil {
			return "", err
		}

	}

	return flowId, nil
}
