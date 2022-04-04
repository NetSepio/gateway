package paseto

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/NetSepio/gateway/config"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/models/claims"
	"github.com/NetSepio/gateway/util/pkg/auth"
	"github.com/NetSepio/gateway/util/pkg/envutil"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/NetSepio/gateway/util/testingcommon"
	"github.com/vk-rv/pvx"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_PASETO(t *testing.T) {
	config.Init("../../../../.env")
	logwrapper.Init("../../../../logs")
	db := dbconfig.GetDb()
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	gin.SetMode(gin.TestMode)
	testWalletAddress := testingcommon.GenerateWallet().WalletAddress
	newUser := models.User{
		WalletAddress: testWalletAddress,
	}
	err := db.Model(&models.User{}).Create(&newUser).Error
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		db.Delete(&newUser)
	}()
	t.Run("Should return 200 with correct PASETO", func(t *testing.T) {
		newClaims := claims.New(testWalletAddress)
		token, err := auth.GenerateToken(newClaims, envutil.MustGetEnv("PASETO_PRIVATE_KEY"))
		if err != nil {
			t.Fatal(err)
		}
		statusCode := callApi(t, token)
		assert.Equal(t, http.StatusOK, statusCode)
	})

	t.Run("Should return 403 with incorret PASETO", func(t *testing.T) {
		newClaims := claims.New(testWalletAddress)
		token, err := auth.GenerateToken(newClaims, "this private key is valid key")
		if err != nil {
			t.Fatal(err)
		}
		statusCode := callApi(t, token)
		assert.Equal(t, http.StatusForbidden, statusCode)
	})

	t.Run("Should return 403 with expired PASETO", func(t *testing.T) {
		expiration := time.Now().Add(time.Second * 2)
		signedBy := envutil.MustGetEnv("SIGNED_BY")
		newClaims := claims.CustomClaims{
			testWalletAddress,
			signedBy,
			pvx.RegisteredClaims{
				Expiration: &expiration,
			},
		}
		time.Sleep(time.Second * 2)
		token, err := auth.GenerateToken(newClaims, envutil.MustGetEnv("PASETO_PRIVATE_KEY"))
		if err != nil {
			t.Fatal(err)
		}

		statusCode := callApi(t, token)
		assert.Equal(t, http.StatusForbidden, statusCode)
	})

}

func callApi(t *testing.T, token string) int {
	rr := httptest.NewRecorder()
	ginTestApp := gin.New()

	rq, err := http.NewRequest("POST", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	rq.Header.Add("Authorization", token)
	ginTestApp.Use(PASETO)
	ginTestApp.Use(successHander)
	ginTestApp.ServeHTTP(rr, rq)
	return rr.Result().StatusCode
}

func successHander(c *gin.Context) {
	c.Status(http.StatusOK)
}
