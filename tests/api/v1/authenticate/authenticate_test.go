package authenticate_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"netsepio-api/api/v1/authenticate"
	"netsepio-api/api/v1/flowid"
	"netsepio-api/app"
	testingcommmon "netsepio-api/util/testingcommon"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TODO: Write test to verify expiry
func Test_PostAuthenticate(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app.Init()

	testWallet := testingcommmon.GenerateWallet()
	flowId := callFlowIdApi(testWallet.WalletAddress, t)
	t.Cleanup(testingcommmon.ClearTables)

	router := app.GinApp

	url := fmt.Sprintf("/api/v1.0/authenticate")

	t.Run("Should return 200 with correct wallet address", func(t *testing.T) {
		signature := getSignature(flowId, testWallet.PrivateKey)
		body := authenticate.AuthenticateRequest{Signature: signature, FlowId: flowId}
		jsonBody, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		//Request with signature created from correct wallet address
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code, rr.Body.String())
	})
	t.Run("Should return 403 with different wallet address", func(t *testing.T) {
		// Different private key will result in different wallet address
		differentPrivatekey := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
		signature := getSignature(flowId, differentPrivatekey)
		body := authenticate.AuthenticateRequest{Signature: signature, FlowId: flowId}
		jsonBody, err := json.Marshal(body)
		if err != nil {
			t.Fatal(err)
		}
		newWalletAddress := testWallet.WalletAddress + "b"
		callFlowIdApi(newWalletAddress, t)

		rr := httptest.NewRecorder()

		//Request with signature stil created from different walletAddress
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
		if err != nil {
			t.Fatal(err)
		}
		router.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusForbidden, rr.Code, rr.Body.String())
	})

}

func callFlowIdApi(walletAddress string, t *testing.T) (flowidString string) {
	// Call flowid api
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
	assert.Equal(t, http.StatusOK, rr.Code, "Failed to call flowApi")
	var flowIdResponse flowid.GetFlowIdResponse
	decoder := json.NewDecoder(rr.Result().Body)
	err = decoder.Decode(&flowIdResponse)
	if err != nil {
		t.Fatal(err)
	}
	return flowIdResponse.FlowId
}

func getSignature(flowId string, walletAddress string) string {
	hexPrivateKey := walletAddress
	message := flowId
	newMsg := fmt.Sprintf("\x19Ethereum Signed Message:\n%v%v", len(message), message)

	privateKey, err := crypto.HexToECDSA(hexPrivateKey)
	if err != nil {
		log.Fatal("HexToECDSA failed ", err)
	}

	// keccak256 hash of the data
	dataBytes := []byte(newMsg)
	hashData := crypto.Keccak256Hash(dataBytes)

	signatureBytes, err := crypto.Sign(hashData.Bytes(), privateKey)
	if err != nil {
		log.Fatal("len", err)
	}

	signature := hexutil.Encode(signatureBytes)

	return signature
}
