package testingcommon

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"log"
	"strings"
	"testing"

	"github.com/NetSepio/gateway/api/types"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/models/claims"
	"github.com/NetSepio/gateway/util/pkg/auth"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/google/uuid"
	"golang.org/x/crypto/nacl/sign"
	"golang.org/x/crypto/sha3"

	"github.com/gin-gonic/gin"
)

func PrepareAndGetAuthHeader(t *testing.T, testWalletAddress string) string {
	gin.SetMode(gin.TestMode)
	userId := uuid.NewString()
	CreateTestUser(t, testWalletAddress, userId)
	customClaims := claims.New(userId, testWalletAddress)

	pvKey, err := hex.DecodeString(envconfig.EnvVars.PASETO_PRIVATE_KEY[2:])
	if err != nil {
		t.Fatal(err)
	}
	token, err := auth.GenerateToken(customClaims, pvKey)
	if err != nil {
		t.Fatal(err)
	}

	header := token
	return header
}

func CreateTestUser(t *testing.T, walletAddress string, uuid string) {
	db := dbconfig.GetDb()
	user := models.User{
		Name:              "Jack",
		ProfilePictureUrl: "https://revoticengineering.com/",
		WalletAddress:     strings.ToLower(walletAddress),
		Country:           "India",
		UserId:            uuid,
	}
	err := db.Model(&models.User{}).Create(&user).Error
	if err != nil {
		t.Fatal(err)
	}
}

func GenerateWallet() *TestWallet {
	publicKey, privateKeyBytes, err := sign.GenerateKey(nil)
	if err != nil {
		log.Fatal(err)
	}
	sha3_i := sha3.New256()

	sha3_i.Write(publicKey[:])
	sha3_i.Write([]byte{0})
	hash := sha3_i.Sum(nil)
	addr := hexutil.Encode(hash)
	privateKeyHex := hexutil.Encode(privateKeyBytes[:])
	pubKeyHex := hexutil.Encode(publicKey[:])
	testWallet := TestWallet{
		PrivateKey:    privateKeyHex[2:],
		PubKey:        pubKeyHex,
		WalletAddress: addr,
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
