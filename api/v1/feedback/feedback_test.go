package feedback

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NetSepio/gateway/config/envconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/NetSepio/gateway/util/testingcommon"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_PostFeedback(t *testing.T) {

	envconfig.InitEnvVars()
	logwrapper.Init()
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	testWallet := testingcommon.GenerateWallet()
	header := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)

	url := "/api/v1.0/feedback"

	t.Run("Should be able to add feedback", func(t *testing.T) {
		rr := httptest.NewRecorder()

		requestBody := models.UserFeedback{
			Feedback: "Very helpfull for avoiding spam and harmfull domains",
			Rating:   5,
		}
		jsonData, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		req.Header.Add("Authorization", header)
		if err != nil {
			t.Fatal(err)
		}
		c, _ := gin.CreateTestContext(rr)
		c.Request = req
		c.Set("walletAddress", testWallet.WalletAddress)

		createFeedback(c)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	})
}
