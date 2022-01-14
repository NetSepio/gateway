package roleid

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TheLazarusNetwork/marketplace-engine/app"
	"github.com/TheLazarusNetwork/marketplace-engine/util/testingcommon"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_GetRoleId(t *testing.T) {
	gin.SetMode(gin.TestMode)
	app.Init()
	testWallet := testingcommon.GenerateWallet()
	headers := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)
	t.Cleanup(testingcommon.ClearTables)
	url := "/api/v1.0/roleId/%v"
	t.Run("Get role EULA with flowId when roleId exist", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", fmt.Sprintf(url, 2), nil)
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
