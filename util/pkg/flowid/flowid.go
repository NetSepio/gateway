package flowid

import (
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/jinzhu/gorm"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GenerateFlowId(walletAddress string, flowIdType models.FlowIdType, relatedRoleId string) (string, error) {
	db := dbconfig.GetDb()
	flowId := uuid.NewString()
	var update bool
	update = true
	if err := db.Model(&models.User{}).Find(&models.User{}).Error; err == gorm.ErrRecordNotFound {
		update = false
	}
	if update {
		// User exist so update
		association := db.Model(&models.User{
			WalletAddress: walletAddress,
		}).Association("FlowIds")
		if err := association.Error; err != nil {
			logrus.Error(err)
			return "", err
		}
		err := association.Append(&models.FlowId{FlowIdType: flowIdType, WalletAddress: walletAddress, FlowId: flowId, RelatedRoleId: relatedRoleId}).Error
		if err != nil {
			return "", err
		}
	} else {
		// User doesn't exist so create

		newUser := &models.User{
			WalletAddress: walletAddress,
			FlowIds: []models.FlowId{{
				FlowIdType: flowIdType, WalletAddress: walletAddress, FlowId: flowId, RelatedRoleId: relatedRoleId,
			}},
		}
		if err := db.Create(newUser).Error; err != nil {
			return "", err
		}

	}

	return flowId, nil
}
