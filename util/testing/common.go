package testingcommmon

import (
	"fmt"
	"netsepio-api/app"
	"netsepio-api/db"
	"netsepio-api/models"
	"netsepio-api/models/claims"
	"netsepio-api/util/pkg/auth"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func PrepareAndGetAuthHeader(t *testing.T) string {
	gin.SetMode(gin.TestMode)
	app.Init()
	testWalletAddress := os.Getenv("TEST_WALLET_ADDRESS")
	CreateTestUser(t, testWalletAddress)
	customClaims := claims.New(testWalletAddress)
	jwtPrivateKey := os.Getenv("JWT_PRIVATE_KEY")
	token, err := auth.GenerateToken(customClaims, jwtPrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	header := fmt.Sprintf("Bearer %v", token)

	return header
}

func CreateTestUser(t *testing.T, walletAddress string) {
	user := models.User{
		Name:              "Jack",
		ProfilePictureUrl: "https://revoticengineering.com/",
		WalletAddress:     walletAddress,
		Country:           "India",
		Roles:             pq.Int32Array([]int32{1}),
	}
	err := db.Db.Model(&models.User{}).Create(&user).Error
	if err != nil {
		t.Fatal(err)
	}
}

func ClearTables() {
	db.Db.Delete(&models.User{})
	db.Db.Delete(&models.FlowId{})
}
