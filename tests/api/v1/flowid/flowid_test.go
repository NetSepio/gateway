package flowid

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"netsepio-api/api/v1/flowid"
	"netsepio-api/app"
	testingcommmon "netsepio-api/util/testing"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

//TODO add test for testing when wallet address exist
func Test_GetFlowId(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Cleanup(testingcommmon.ClearTables)

	app.Init()
	var (
		walletAddress = os.Getenv("TEST_WALLET_ADDRESS")
	)
	url := "/api/v1.0/flowid"
	rr := httptest.NewRecorder()
	body := flowid.GetFlowIdRequest{
		WalletAddress: walletAddress,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Error(err)
	}
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Error(err)
	}
	app.GinApp.ServeHTTP(rr, req)
	assert.Equal(t, rr.Result().StatusCode, http.StatusOK)
}
