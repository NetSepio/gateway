package testingcommon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/TheLazarusNetwork/marketplace-engine/db"
	"github.com/TheLazarusNetwork/marketplace-engine/models"
	"github.com/TheLazarusNetwork/marketplace-engine/models/claims"
	"github.com/TheLazarusNetwork/marketplace-engine/types"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/auth"

	"crypto/ecdsa"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func PrepareAndGetAuthHeader(t *testing.T, testWalletAddress string) string {
	gin.SetMode(gin.TestMode)
	db.InitDB()
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
		Roles: []models.UserRole{
			{
				WalletAddress: walletAddress, RoleId: 1,
			},
		},
	}
	err := db.Db.Model(&models.User{}).Create(&user).Error
	if err != nil {
		t.Fatal(err)
	}
}

func GenerateWallet() *TestWallet {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	privateKeyHex := hexutil.Encode(privateKeyBytes)
	testWallet := TestWallet{
		PrivateKey:    privateKeyHex[2:],
		WalletAddress: address,
	}
	return &testWallet
}

// Converts map created by json decoder to struct
// out should be pointer (&payload)
func ExtractPayload(response *types.ApiResponse, out interface{}) {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(response.Payload)
	json.NewDecoder(buf).Decode(out)
}

func ClearTables() {
	db.Db.Delete(&models.User{})
	db.Db.Delete(&models.FlowId{})
	db.Db.Delete(&models.UserRole{})
}
