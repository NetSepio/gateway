package roleid

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/config/netsepio"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/NetSepio/gateway/util/testingcommon"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_GetRoleId(t *testing.T) {
	gin.SetMode(gin.TestMode)

	envconfig.InitEnvVars()
	logwrapper.Init()
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	testWallet := testingcommon.GenerateWallet()
	headers := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)
	voterRole, err := netsepio.GetRole(netsepio.VOTER_ROLE)
	if err != nil {
		t.Fatalf("failed to get role id for %v , error: %v", "VOTER ROLE", err.Error())
	}

	url := "/api/v1.0/roleId/%v"
	t.Run("Get role EULA with flowId when roleId exist", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", fmt.Sprintf(url, hexutil.Encode(voterRole[:])), nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("Authorization", headers)
		c, _ := gin.CreateTestContext(rr)
		c.Request = req
		// c.Params = gin.Params{{Key: "", Value: }
		c.Set("walletAddress", testWallet.WalletAddress)
		c.Params = gin.Params{{Key: "roleId", Value: hexutil.Encode(voterRole[:])}}
		GetRoleId(c)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	})
	t.Run("Get not found message when roleId doesn't exist", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", fmt.Sprintf(url, 58), nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("Authorization", headers)
		c, _ := gin.CreateTestContext(rr)
		c.Request = req
		c.Set("walletAddress", testWallet.WalletAddress)
		c.Params = gin.Params{{Key: "roleId", Value: "58"}}
		GetRoleId(c)
		assert.Equal(t, http.StatusNotFound, rr.Result().StatusCode)
	})

}
