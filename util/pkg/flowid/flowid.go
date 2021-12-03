package flowid

import (
	"netsepio-api/db"
	"netsepio-api/models"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

func GenerateFlowId(walletAddress string, update bool) (string, error) {

	flowId := uuid.NewString()
	if update {
		// User exist so update
		err := db.Db.Model(&models.User{}).
			Where("wallet_address = ?", walletAddress).
			Update("flow_id", gorm.Expr("array_cat(flow_id,?)", pq.Array([]string{flowId}))).Error
		if err != nil {
			return "", err
		}
	} else {
		// User doesn't exist so create
		newUser := &models.User{
			WalletAddress: walletAddress,
			FlowId:        pq.StringArray([]string{flowId}),
		}
		if err := db.Db.Create(newUser).Error; err != nil {
			return "", err
		}
	}

	return flowId, nil
}
