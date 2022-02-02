package roleid

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TheLazarusNetwork/netsepio-engine/app"
	"github.com/TheLazarusNetwork/netsepio-engine/config/netsepio"
	"github.com/TheLazarusNetwork/netsepio-engine/util/testingcommon"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_GetRoleId(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app.Init("../../../../.env", "../../../../logs")
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	testWallet := testingcommon.GenerateWallet()
	headers := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)
	creatorRole, err := netsepio.GetRole(netsepio.VOTER_ROLE)
	if err != nil {
		t.Fatalf("failed to get role id for %v , error: %v", "CREATOR ROLE", err.Error())
	}

	url := "/api/v1.0/roleId/%v"
	t.Run("Get role EULA with flowId when roleId exist", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", fmt.Sprintf(url, hexutil.Encode(creatorRole[:])), nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("Authorization", headers)
		app.GinApp.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	})
	t.Run("Get not found message when roleId doesn't exist", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", fmt.Sprintf(url, 58), nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("Authorization", headers)
		app.GinApp.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusNotFound, rr.Result().StatusCode)
	})

}
