package testingcommon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/TheLazarusNetwork/netsepio-engine/api/types"
	"github.com/TheLazarusNetwork/netsepio-engine/config/dbconfig"
	"github.com/TheLazarusNetwork/netsepio-engine/models"
	"github.com/TheLazarusNetwork/netsepio-engine/models/claims"
	"github.com/TheLazarusNetwork/netsepio-engine/util/pkg/auth"
	"github.com/TheLazarusNetwork/netsepio-engine/util/pkg/envutil"

	"crypto/ecdsa"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func PrepareAndGetAuthHeader(t *testing.T, testWalletAddress string) string {
	gin.SetMode(gin.TestMode)
	CreateTestUser(t, testWalletAddress)
	customClaims := claims.New(testWalletAddress)
	jwtPrivateKey := envutil.MustGetEnv("JWT_PRIVATE_KEY")
	token, err := auth.GenerateToken(customClaims, jwtPrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	header := fmt.Sprintf("Bearer %v", token)

	return header
}

func CreateTestUser(t *testing.T, walletAddress string) {
	db := dbconfig.GetDb()
	user := models.User{
		Name:              "Jack",
		ProfilePictureUrl: "https://revoticengineering.com/",
		WalletAddress:     walletAddress,
		Country:           "India",
	}
	err := db.Model(&models.User{}).Create(&user).Error
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
