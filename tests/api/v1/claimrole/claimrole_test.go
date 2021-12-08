package claimrole

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	claimrole "netsepio-api/api/v1/claimRole"
	roleid "netsepio-api/api/v1/roleId"
	"netsepio-api/app"
	testingcommmon "netsepio-api/util/testing"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func Test_PostClaimRole(t *testing.T) {
	app.Init()
	headers := testingcommmon.PrepareAndGetAuthHeader(t)
	t.Cleanup(testingcommmon.ClearTables)
	url := "/api/v1.0/claimrole"
	rr := httptest.NewRecorder()
	requestRoleRes := requestRole(t, headers)
	signature := getSignature(requestRoleRes.Message, os.Getenv("TEST_WALLET_PRIVATE_KEY"))
	reqBody := claimrole.ClaimRoleRequest{
		Signature: signature, FlowId: requestRoleRes.FlowId,
	}
	jsonBytes, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Authorization", headers)
	app.GinApp.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
}

func requestRole(t *testing.T, headers string) roleid.GetRoleIdResponse {
	url := "/api/v1.0/roleId/1"
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Authorization", headers)
	app.GinApp.ServeHTTP(rr, req)
	var res roleid.GetRoleIdResponse
	json.NewDecoder(rr.Result().Body).Decode(&res)
	return res
}
func getSignature(eula string, hexPrivateKey string) string {
	message := eula
	fmt.Println(message)

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

	lenOfString := len(signature)
	newLenOfString := lenOfString - 2
	newSignature := signature[:newLenOfString]
	// TODO: Fix end bytes
	newSignature = newSignature + "1b"
	return newSignature
}
