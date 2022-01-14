package claimrole

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	claimrole "github.com/TheLazarusNetwork/marketplace-engine/api/v1/claimRole"
	roleid "github.com/TheLazarusNetwork/marketplace-engine/api/v1/roleId"
	"github.com/TheLazarusNetwork/marketplace-engine/app"
	"github.com/TheLazarusNetwork/marketplace-engine/types"
	"github.com/TheLazarusNetwork/marketplace-engine/util/testingcommon"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_PostClaimRole(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app.Init()
	testWallet := testingcommon.GenerateWallet()
	headers := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)
	t.Cleanup(testingcommon.ClearTables)
	url := "/api/v1.0/claimrole"
	rr := httptest.NewRecorder()
	requestRoleRes := requestRole(t, headers)
	signature := getSignature(requestRoleRes.Eula, requestRoleRes.FlowId, testWallet.PrivateKey)
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
	assert.Equal(t, http.StatusOK, rr.Result().StatusCode, rr.Body.String())
}

func requestRole(t *testing.T, headers string) roleid.GetRoleIdPayload {
	url := "/api/v1.0/roleId/2"
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Authorization", headers)
	app.GinApp.ServeHTTP(rr, req)
	var res types.ApiResponse
	json.NewDecoder(rr.Result().Body).Decode(&res)
	var getRoleIdPayload roleid.GetRoleIdPayload
	testingcommon.ExtractPayload(&res, &getRoleIdPayload)
	return getRoleIdPayload
}
func getSignature(eula string, flowId string, hexPrivateKey string) string {
	message := eula + flowId

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
