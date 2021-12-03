package profile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"netsepio-api/api/v1/profile"
	"netsepio-api/app"
	"netsepio-api/db"
	"netsepio-api/models"
	"netsepio-api/models/claims"
	"netsepio-api/util/pkg/auth"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func prepare(t *testing.T) string {
	gin.SetMode(gin.TestMode)
	app.Init()
	testWalletAddress := os.Getenv("TEST_WALLET_ADDRESS")
	t.Cleanup(createTestUser(t, testWalletAddress))

	customClaims := claims.New(testWalletAddress)
	jwtPrivateKey := os.Getenv("JWT_PRIVATE_KEY")
	token, err := auth.GenerateToken(customClaims, jwtPrivateKey)
	if err != nil {
		t.Fatal(err)
	}
	header := fmt.Sprintf("Bearer %v", token)

	return header
}
func Test_PatchProfile(t *testing.T) {
	header := prepare(t)
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
	header := prepare(t)
	url := "/api/v1.0/profile"
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", header)
	if err != nil {
		t.Fatal(err)
	}
	app.GinApp.ServeHTTP(rr, req)
	var user models.User
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
func createTestUser(t *testing.T, walletAddress string) func() {
	user := models.User{
		Name:              "Jack",
		ProfilePictureUrl: "https://revoticengineering.com/",
		WalletAddress:     walletAddress,
		Country:           "India",
	}
	err := db.Db.Model(&models.User{}).Create(&user).Error
	if err != nil {
		t.Fatal(err)
	}

	return func() {
		db.Db.Delete(&user)
	}
}
