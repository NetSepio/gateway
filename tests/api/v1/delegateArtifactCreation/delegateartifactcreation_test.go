package delegateartifactcreation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	delegateartifactcreation "github.com/TheLazarusNetwork/marketplace-engine/api/v1/delegateArtifactCreation"
	"github.com/TheLazarusNetwork/marketplace-engine/app"
	"github.com/TheLazarusNetwork/marketplace-engine/util/testingcommon"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDelegateArtifactCreation(t *testing.T) {
	app.Init()
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	gin.SetMode(gin.TestMode)
	testWallet := testingcommon.GenerateWallet()
	createrWallet := testingcommon.GenerateWallet()
	headers := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)
	url := "/api/v1.0/delegateArtifactCreation"
	rr := httptest.NewRecorder()

	reqBody := delegateartifactcreation.DelegateArtifactCreationRequest{
		CreatorAddress: createrWallet.WalletAddress,
		MetaDataHash:   "QmSYRXWGGqVDAHKTwfnYQDR74d4bfwXxudFosbGA695AWS",
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
}
