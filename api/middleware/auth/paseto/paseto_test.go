package paseto

import (
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/NetSepio/gateway/api/types"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/models/claims"
	"github.com/NetSepio/gateway/util/pkg/auth"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/NetSepio/gateway/util/testingcommon"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/google/uuid"
	"github.com/vk-rv/pvx"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_PASETO(t *testing.T) {
	envconfig.InitEnvVars()
	logwrapper.Init()
	db := dbconfig.GetDb()
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	gin.SetMode(gin.TestMode)
	testWalletAddress := testingcommon.GenerateWallet().WalletAddress
	newUser := models.User{
		WalletAddress: strings.ToLower(testWalletAddress),
	}
	err := db.Model(&models.User{}).Create(&newUser).Error
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		db.Delete(&newUser)
	}()
	t.Run("Should return 200 with correct PASETO", func(t *testing.T) {
		newClaims := claims.NewWithWallet(uuid.NewString(), &testWalletAddress)
		pvKey, err := hex.DecodeString(envconfig.EnvVars.PASETO_PRIVATE_KEY[2:])
		if err != nil {
			panic(err)
		}
		token, err := auth.GenerateToken(newClaims, pvKey)
		if err != nil {
			t.Fatal(err)
		}
		rr := callApi(t, token)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	})

	t.Run("Should return 401 with incorret PASETO", func(t *testing.T) {
		newClaims := claims.NewWithWallet(uuid.NewString(), &testWalletAddress)
		pvKey, err := hex.DecodeString(envconfig.EnvVars.PASETO_PRIVATE_KEY[2:])
		if err != nil {
			panic(err)
		}
		pvKey[0] = 'a'
		pvKey[0] = 'f'
		token, err := auth.GenerateToken(newClaims, pvKey)
		if err != nil {
			t.Fatal(err)
		}
		rr := callApi(t, token)
		assert.Equal(t, http.StatusUnauthorized, rr.Result().StatusCode)
	})

	t.Run("Should return 401 and 4011 with expired PASETO", func(t *testing.T) {
		expiration := time.Now().Add(time.Second * 2)
		signedBy := envconfig.EnvVars.PASETO_SIGNED_BY
		walletAddrLower := strings.ToLower(testWalletAddress)
		newClaims := claims.CustomClaims{
			WalletAddress: &walletAddrLower,
			SignedBy:      signedBy,
			RegisteredClaims: pvx.RegisteredClaims{
				Expiration: &expiration,
			},
		}
		time.Sleep(time.Second * 2)
		pvKey, err := hex.DecodeString(envconfig.EnvVars.PASETO_PRIVATE_KEY[2:])
		if err != nil {
			panic(err)
		}
		token, err := auth.GenerateToken(newClaims, pvKey)

		rr := callApi(t, token)
		assert.Equal(t, http.StatusUnauthorized, rr.Result().StatusCode)
		var response types.ApiResponse
		body := rr.Body
		err = json.NewDecoder(body).Decode(&response)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, httpo.TokenExpired, response.StatusCode)
	})

}

func callApi(t *testing.T, token string) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	ginTestApp := gin.New()

	rq, err := http.NewRequest("POST", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	rq.Header.Add("Authorization", token)
	ginTestApp.Use(PASETO(false))
	ginTestApp.Use(successHander)
	ginTestApp.ServeHTTP(rr, rq)
	return rr
}

func successHander(c *gin.Context) {
	c.Status(http.StatusOK)
}
