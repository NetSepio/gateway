package profile

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"netsepio-api/api/v1/profile"
	"netsepio-api/app"
	testingcommmon "netsepio-api/util/testing"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func Test_PatchProfile(t *testing.T) {
	header := testingcommmon.PrepareAndGetAuthHeader(t)
	t.Cleanup(testingcommmon.ClearTables)

	url := "/api/v1.0/profile"

	// TODO: Write more tests
	t.Run("Update name", func(t *testing.T) {
		rr := httptest.NewRecorder()

		requestBody := profile.PatchProfileRequest{
			Name: "Yash",
		}
		jsonData, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
		req.Header.Add("Authorization", header)
		if err != nil {
			t.Fatal(err)
		}
		app.GinApp.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	})
}

func Test_GetProfile(t *testing.T) {
	gin.SetMode(gin.TestMode)

	header := testingcommmon.PrepareAndGetAuthHeader(t)
	t.Cleanup(testingcommmon.ClearTables)
	url := "/api/v1.0/profile"
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", header)
	if err != nil {
		t.Fatal(err)
	}
	app.GinApp.ServeHTTP(rr, req)
	var user profile.GetProfileResponse
	body := rr.Body
	json.NewDecoder(body).Decode(&user)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	assert.Equal(t, "Jack", user.Name)
	assert.Equal(t, "https://revoticengineering.com/", user.ProfilePictureUrl)
	assert.Equal(t, "India", user.Country)
	logrus.Debug(user)
}
