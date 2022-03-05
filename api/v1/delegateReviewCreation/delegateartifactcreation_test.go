package delegatereviewcreation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/TheLazarusNetwork/netsepio-engine/config"
	"github.com/TheLazarusNetwork/netsepio-engine/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/netsepio-engine/util/testingcommon"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDelegateReviewCreation(t *testing.T) {
	time.Sleep(4 * time.Second)
	config.Init("../../../.env")
	logwrapper.Init("../../../logs")
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	gin.SetMode(gin.TestMode)
	testWallet := testingcommon.GenerateWallet()
	voterWallet := testingcommon.GenerateWallet()
	headers := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)
	url := "/api/v1.0/delegateReviewCreation"
	rr := httptest.NewRecorder()

	reqBody := DelegateReviewCreationRequest{
		Voter:         voterWallet.WalletAddress,
		MetaDataUri:   "QmSYRXWGGqVDAHKTwfnYQDR74d4bfwXxudFosbGA695AWS",
		Category:      "Website",
		DomainAddress: "ommore.me",
		SiteUrl:       "todo.ommore.me",
		SiteType:      "Productivity app",
		SiteTag:       "react",
		SiteSafety:    "Very safe and smooth",
	}
	jsonBytes, _ := json.Marshal(reqBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Authorization", headers)
	c, _ := gin.CreateTestContext(rr)
	c.Request = req
	deletegateReviewCreation(c)
	ok := assert.Equal(t, http.StatusOK, rr.Result().StatusCode, rr.Body.String())
	if !ok {
		t.FailNow()
	}
}
