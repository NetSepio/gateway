package flowid

import (
	"fmt"
	"strings"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GenerateFlowId(walletAddress string, flowIdType models.FlowIdType, relatedRoleId string) (string, error) {
	db := dbconfig.GetDb()
	flowId := uuid.NewString()
	var update bool = true

	var fetchUser models.User
	findResult := db.Model(&models.User{}).Find(&fetchUser, &models.User{WalletAddress: strings.ToLower(walletAddress)})
	if err := findResult.Error; err != nil {
		err = fmt.Errorf("while finding user error occured, %s", err)
		logrus.Error(err)
		return "", err
	}

	rowsAffected := findResult.RowsAffected
	if rowsAffected == 0 {
		update = false
	}

	if update {
		// User exist so update
		association := db.Model(&models.User{
			UserId: fetchUser.UserId,
		}).Association("FlowIds")
		if err := association.Error; err != nil {
			logrus.Error(err)
			return "", err
		}
		err := association.Append(&models.FlowId{FlowIdType: flowIdType, FlowId: flowId, RelatedRoleId: relatedRoleId})
		if err != nil {
			return "", err
		}
	} else {
		// User doesn't exist so create
		userId := uuid.NewString()
		newUser := &models.User{
			WalletAddress: strings.ToLower(walletAddress),
			UserId:        userId,
			FlowIds: []models.FlowId{{
				FlowIdType: flowIdType, UserId: userId, FlowId: flowId, RelatedRoleId: relatedRoleId,
			}},
		}
		if err := db.Create(newUser).Error; err != nil {
			return "", err
		}

	}

	return flowId, nil
}
