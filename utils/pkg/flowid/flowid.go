package flowid

import (
	"fmt"
	"strings"

	"github.com/NetSepio/gateway/internal/api/handlers/referral"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GenerateFlowId(walletAddress string, flowIdType models.FlowIdType, relatedRoleId string, userId, origin string) (string, error, bool) {
	var verify bool
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
			return "", err, verify
		}

		rowsAffected := findResult.RowsAffected
		if rowsAffected == 0 {
			update = false
		} else {
			userId = fetchUser.UserId
		}

		if fetchUser.Email == nil || fetchUser.Name == "" {
			verify = false
		} else {
			verify = true
		}
	}

	if update {
		// User exist so update
		association := db.Model(&models.User{
			UserId: userId,
		}).Association("FlowIds")
		if err := association.Error; err != nil {
			logrus.Error(err)
			return "", err, verify
		}
		err := association.Append(&models.FlowId{FlowIdType: flowIdType, FlowId: flowId, RelatedRoleId: relatedRoleId, WalletAddress: walletAddress})
		if err != nil {
			return "", err, verify
		}

		var fetchUser models.User
		lowerWalletAddress := strings.ToLower(walletAddress)
		findResult := db.Model(&models.User{}).Find(&fetchUser, &models.User{WalletAddress: &lowerWalletAddress})
		if err := findResult.Error; err != nil {
			err = fmt.Errorf("while finding user error occured, %s", err)
			logrus.Error(err)
			return "", err, verify
		}

		if fetchUser.Email == nil || fetchUser.Name == "" {
			verify = false
		} else {
			verify = true
		}

	} else {
		// User doesn't exist so create
		userId := uuid.NewString()
		newUser := &models.User{
			UserId: userId,
			FlowIds: []models.FlowId{{
				FlowIdType: flowIdType, UserId: userId, FlowId: flowId, RelatedRoleId: relatedRoleId, WalletAddress: walletAddress,
			}},
			Origin:       &origin, // Default origin, can be changed later
			ReferralCode: referral.GetReferalCode(),
		}
		if err := db.Create(newUser).Error; err != nil {
			return "", err, verify
		}

		var fetchUser models.User
		lowerWalletAddress := strings.ToLower(walletAddress)
		findResult := db.Model(&models.User{}).Find(&fetchUser, &models.User{WalletAddress: &lowerWalletAddress})
		if err := findResult.Error; err != nil {
			err = fmt.Errorf("while finding user error occured, %s", err)
			logrus.Error(err)
			return "", err, verify
		}

		if fetchUser.Email == nil || fetchUser.Name == "" {
			verify = false
		} else {
			verify = true
		}

	}

	return flowId, nil, verify
}

func GenerateFlowIdSol(walletAddress string, flowIdType models.FlowIdType, relatedRoleId string, userId, origin string) (string, error, bool) {
	db := database.GetDb()
	flowId := uuid.NewString()
	var verify bool
	fmt.Printf("userId: %s\n", userId)
	var update bool = true
	if userId == "" {
		var fetchUser models.User
		findResult := db.Model(&models.User{}).Find(&fetchUser, &models.User{WalletAddress: &walletAddress})
		if err := findResult.Error; err != nil {
			err = fmt.Errorf("while finding user error occured, %s", err)
			logrus.Error(err)
			return "", err, verify
		}

		rowsAffected := findResult.RowsAffected
		if rowsAffected == 0 {
			update = false
		} else {
			userId = fetchUser.UserId
		}

		if fetchUser.Email == nil || fetchUser.Name == "" {
			verify = false
		} else {
			verify = true
		}
	}

	if update {
		// User exist so update
		association := db.Model(&models.User{
			UserId: userId,
		}).Association("FlowIds")
		if err := association.Error; err != nil {
			logrus.Error(err)
			return "", err, verify
		}
		err := association.Append(&models.FlowId{FlowIdType: flowIdType, FlowId: flowId, RelatedRoleId: relatedRoleId, WalletAddress: walletAddress})
		if err != nil {
			return "", err, verify
		}
		var fetchUser models.User
		findResult := db.Model(&models.User{}).Find(&fetchUser, &models.User{WalletAddress: &walletAddress})
		if err := findResult.Error; err != nil {
			err = fmt.Errorf("while finding user error occured, %s", err)
			logrus.Error(err)
			return "", err, verify
		}
		if fetchUser.Email == nil || fetchUser.Name == "" {
			verify = false
		} else {
			verify = true
		}
	} else {
		// User doesn't exist so create
		userId := uuid.NewString()
		newUser := &models.User{
			UserId: userId,
			FlowIds: []models.FlowId{{
				FlowIdType: flowIdType, UserId: userId, FlowId: flowId, RelatedRoleId: relatedRoleId, WalletAddress: walletAddress,
			}},
			ReferralCode: referral.GetReferalCode(),
			Origin:       &origin, // Default origin, can be changed later
		}
		if err := db.Create(newUser).Error; err != nil {
			return "", err, verify
		}

		var fetchUser models.User
		findResult := db.Model(&models.User{}).Find(&fetchUser, &models.User{WalletAddress: &walletAddress})
		if err := findResult.Error; err != nil {
			err = fmt.Errorf("while finding user error occured, %s", err)
			logrus.Error(err)
			return "", err, verify
		}
		if fetchUser.Email == nil || fetchUser.Name == "" {
			verify = false
		} else {
			verify = true
		}

	}

	return flowId, nil, verify
}
