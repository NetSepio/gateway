package flowid

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"netsepio-gateway-v1.1/internal/database"
	"netsepio-gateway-v1.1/models"
)

func GenerateFlowId(walletAddress string, flowIdType models.FlowIdType, relatedRoleId string, userId string) (string, error) {
	db := database.GetDb()
	flowId := uuid.NewString()
	fmt.Printf("userId: %s\n", userId)
	var update bool = true
	if userId == "" {
		var fetchUser models.User
		lowerWalletAddress := strings.ToLower(walletAddress)
		findResult := db.Model(&models.User{}).Find(&fetchUser, &models.User{WalletAddress: &lowerWalletAddress})
		if err := findResult.Error; err != nil {
			err = fmt.Errorf("while finding user error occured, %s", err)
			logrus.Error(err)
			return "", err
		}

		rowsAffected := findResult.RowsAffected
		if rowsAffected == 0 {
			update = false
		} else {
			userId = fetchUser.UserId
		}
	}

	if update {
		// User exist so update
		association := db.Model(&models.User{
			UserId: userId,
		}).Association("FlowIds")
		if err := association.Error; err != nil {
			logrus.Error(err)
			return "", err
		}
		err := association.Append(&models.FlowId{FlowIdType: flowIdType, FlowId: flowId, RelatedRoleId: relatedRoleId, WalletAddress: walletAddress})
		if err != nil {
			return "", err
		}
	} else {
		// User doesn't exist so create
		userId := uuid.NewString()
		newUser := &models.User{
			UserId: userId,
			FlowIds: []models.FlowId{{
				FlowIdType: flowIdType, UserId: userId, FlowId: flowId, RelatedRoleId: relatedRoleId, WalletAddress: walletAddress,
			}},
		}
		if err := db.Create(newUser).Error; err != nil {
			return "", err
		}

	}

	return flowId, nil
}

func GenerateFlowIdSol(walletAddress string, flowIdType models.FlowIdType, relatedRoleId string, userId string) (string, error) {
	db := database.GetDb()
	flowId := uuid.NewString()
	fmt.Printf("userId: %s\n", userId)
	var update bool = true
	if userId == "" {
		var fetchUser models.User
		findResult := db.Model(&models.User{}).Find(&fetchUser, &models.User{WalletAddress: &walletAddress})
		if err := findResult.Error; err != nil {
			err = fmt.Errorf("while finding user error occured, %s", err)
			logrus.Error(err)
			return "", err
		}

		rowsAffected := findResult.RowsAffected
		if rowsAffected == 0 {
			update = false
		} else {
			userId = fetchUser.UserId
		}
	}

	if update {
		// User exist so update
		association := db.Model(&models.User{
			UserId: userId,
		}).Association("FlowIds")
		if err := association.Error; err != nil {
			logrus.Error(err)
			return "", err
		}
		err := association.Append(&models.FlowId{FlowIdType: flowIdType, FlowId: flowId, RelatedRoleId: relatedRoleId, WalletAddress: walletAddress})
		if err != nil {
			return "", err
		}
	} else {
		// User doesn't exist so create
		userId := uuid.NewString()
		newUser := &models.User{
			UserId: userId,
			FlowIds: []models.FlowId{{
				FlowIdType: flowIdType, UserId: userId, FlowId: flowId, RelatedRoleId: relatedRoleId, WalletAddress: walletAddress,
			}},
		}
		if err := db.Create(newUser).Error; err != nil {
			return "", err
		}

	}

	return flowId, nil
}
