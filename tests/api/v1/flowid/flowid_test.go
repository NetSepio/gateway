package flowid

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"netsepio-api/app"
	"netsepio-api/util/testingcommon"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

//TODO add test for testing when wallet address exist
func Test_GetFlowId(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Cleanup(testingcommon.ClearTables)

	app.Init()

	testWalletAddress := testingcommon.GenerateWallet().WalletAddress
	u, err := url.Parse("/api/v1.0/flowid")
	if err != nil {
		t.Fatal(err)
	}
	q := url.Values{}
	q.Set("walletAddress", testWalletAddress)
	u.RawQuery = q.Encode()
	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		t.Error(err)
	}
	app.GinApp.ServeHTTP(rr, req)
	assert.Equal(t, rr.Result().StatusCode, http.StatusOK)
}
