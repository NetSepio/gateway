package roleid

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"netsepio-api/app"
	testingcommmon "netsepio-api/util/testing"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetRoleId(t *testing.T) {
	app.Init()
	headers := testingcommmon.PrepareAndGetAuthHeader(t)
	t.Cleanup(testingcommmon.ClearTables)
	url := "/api/v1.0/roleId/%v"
	t.Run("Get role EULA with flowId when roleId exist", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("GET", fmt.Sprintf(url, 1), nil)
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
