package flowid

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"netsepio-api/api/v1/flowid"
	"netsepio-api/app"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetFlowId(t *testing.T) {
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
