package jwt

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"netsepio-api/models/claims"
	"netsepio-api/util/pkg/auth"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_JWT(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("Should return 200 with correct JWT", func(t *testing.T) {
		testWalletAddress := os.Getenv("TEST_WALLET_ADDRESS")
		newClaims := claims.New(testWalletAddress)
		token, err := auth.GenerateToken(newClaims, os.Getenv("JWT_PRIVATE_KEY"))
		if err != nil {
			t.Fatal(err)
		}
		statusCode := callApi(t, token)
		assert.Equal(t, http.StatusOK, statusCode)
	})

	t.Run("Should return 403 with incorret JWT", func(t *testing.T) {
		newClaims := claims.New(os.Getenv("TEST_WALLET_ADDRESS"))
		token, err := auth.GenerateToken(newClaims, "this private key is valid key")
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

	header := fmt.Sprintf("Bearer %v", token)
	rq, err := http.NewRequest("POST", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	rq.Header.Add("Authorization", header)
	ginTestApp.Use(JWT)
	ginTestApp.Use(successHander)
	ginTestApp.ServeHTTP(rr, rq)
	return rr.Result().StatusCode
}

func successHander(c *gin.Context) {
	c.Status(http.StatusOK)
}
