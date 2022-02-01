package claimrole

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/TheLazarusNetwork/marketplace-engine/api/types"
	claimrole "github.com/TheLazarusNetwork/marketplace-engine/api/v1/claimRole"
	roleid "github.com/TheLazarusNetwork/marketplace-engine/api/v1/roleId"
	"github.com/TheLazarusNetwork/marketplace-engine/app"
	"github.com/TheLazarusNetwork/marketplace-engine/config/creatify"
	"github.com/TheLazarusNetwork/marketplace-engine/config/smartcontract"
	"github.com/TheLazarusNetwork/marketplace-engine/config/smartcontract/auth"
	smartcontractcreatify "github.com/TheLazarusNetwork/marketplace-engine/generated/smartcontract/creatify"
	"github.com/TheLazarusNetwork/marketplace-engine/util/testingcommon"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_PostClaimRole(t *testing.T) {
	defer time.Sleep(4 * time.Second)
	app.Init("../../../../.env", "../../../../logs")
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	gin.SetMode(gin.TestMode)
	testWallet := testingcommon.GenerateWallet()
	headers := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)
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
	ok := assert.Equal(t, http.StatusOK, rr.Result().StatusCode, rr.Body.String())
	if !ok {
		t.FailNow()
	}
	client, err := smartcontract.GetClient()
	if err != nil {
		t.Fatal(err)
	}
	instance, err := creatify.GetInstance(client)
	if err != nil {
		t.Fatalf("failed to get instance for %v , error: %v", "CREATIFY", err.Error())
	}
	creatorRole, err := creatify.GetRole(creatify.CREATOR_ROLE)
	if err != nil {
		t.Fatalf("failed to get role id for %v , error: %v", "CREATOR ROLE", err.Error())
	}
	addr := common.HexToAddress(testWallet.WalletAddress)
	roleGrantedChannel := make(chan *smartcontractcreatify.CreatifyRoleGranted, 10)

	authBindOpts, err := auth.GetAuth(client)

	if err != nil {
		t.Fatalf("failed to get auth, error: %v", err.Error())
	}
	subs, err := instance.WatchRoleGranted(nil, roleGrantedChannel, [][32]byte{creatorRole}, []common.Address{addr}, []common.Address{authBindOpts.From})
	if err != nil {
		t.Fatalf("failed to listen to an event %v, error: %v", "RoleGranted", err.Error())
	}

	//Check if role trasaction is successfull
	hasRole, err := instance.HasRole(nil, creatorRole, addr)
	if err != nil {
		t.Fatalf("failed to call %v smart contract function HasRole , error: %v", "CREATIFY", err.Error())
	}
	success := false
	if !hasRole {
		go failAfter(t, &success, 10*time.Second, roleGrantedChannel)
		event := <-roleGrantedChannel
		subs.Unsubscribe()
		if event != nil && event.Account.String() != addr.String() {
			log.Fatal("user doesn't have role in blockchain")
		} else {
			success = true
		}
	}

}

func failAfter(t *testing.T, success *bool, duration time.Duration, ch chan *smartcontractcreatify.CreatifyRoleGranted) {
	time.Sleep(duration)
	if !*success {
		close(ch)
		t.Errorf("didn't got any response from %v after %v", "RoleGranted", duration)
	}
}
func requestRole(t *testing.T, headers string) roleid.GetRoleIdPayload {
	creatorRole, err := creatify.GetRole(creatify.CREATOR_ROLE)
	if err != nil {
		t.Fatalf("failed to get role id for %v , error: %v", "CREATOR ROLE", err.Error())
	}

	url := "/api/v1.0/roleId/" + hexutil.Encode(creatorRole[:])
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
