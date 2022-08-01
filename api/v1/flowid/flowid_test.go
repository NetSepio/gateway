package flowid

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/NetSepio/gateway/config"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/NetSepio/gateway/util/testingcommon"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

//TODO add test for testing when wallet address exist
func Test_GetFlowId(t *testing.T) {
	config.Init("../../../.env")
	logwrapper.Init("../../../logs")
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	gin.SetMode(gin.TestMode)
	testWalletAddress := testingcommon.GenerateWallet().WalletAddress
	u, err := url.Parse("/api/v1.0/flowid")
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Should fail if wallet address is not hexadecimal", func(t *testing.T) {
		q := url.Values{}
		q.Set("walletAddress", "invalidwalletaddr")
		u.RawQuery = q.Encode()
		rr := httptest.NewRecorder()

		req, err := http.NewRequest("GET", u.String(), nil)
		if err != nil {
			t.Error(err)
		}
		c, _ := gin.CreateTestContext(rr)
		c.Request = req
		GetFlowId(c)
		assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)
	})
	t.Run("Should be able to get flow id", func(t *testing.T) {

		q := url.Values{}
		q.Set("walletAddress", testWalletAddress)
		u.RawQuery = q.Encode()
		rr := httptest.NewRecorder()

		req, err := http.NewRequest("GET", u.String(), nil)
		if err != nil {
			t.Error(err)
		}
		c, _ := gin.CreateTestContext(rr)
		c.Request = req
		GetFlowId(c)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	})

}
