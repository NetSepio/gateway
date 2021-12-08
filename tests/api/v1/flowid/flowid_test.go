package flowid

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"netsepio-api/api/v1/flowid"
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
	url := "/api/v1.0/flowid"
	rr := httptest.NewRecorder()
	body := flowid.GetFlowIdRequest{
		WalletAddress: testWalletAddress,
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
