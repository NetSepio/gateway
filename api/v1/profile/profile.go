package profile

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/envutil"
	"github.com/NetSepio/gateway/util/pkg/httphelper"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/profile")
	{
		g.Use(paseto.PASETO)
		g.PATCH("", patchProfile)
		g.GET("", getProfile)
	}
}

func patchProfile(c *gin.Context) {
	db := dbconfig.GetDb()
	var requestBody PatchProfileRequest
	c.BindJSON(&requestBody)
	walletAddress := c.GetString("walletAddress")
	result := db.Model(&models.User{}).
		Where("wallet_address = ?", walletAddress).
		Updates(requestBody)
	if result.Error != nil {
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")

		return
	}
	if result.RowsAffected == 0 {
		httphelper.ErrResponse(c, http.StatusNotFound, "Record not found")

		return
	}
	httphelper.SuccessResponse(c, "Profile successfully updated", nil)

}

func getProfile(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	var user models.User
	err := db.Model(&models.User{}).Select("name, profile_picture_url,country, wallet_address").Where("wallet_address = ?", walletAddress).First(&user).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	roles, err := getRoles(user.WalletAddress)
	if err != nil {
		logwrapper.Errorf("Failed to fetch roles from graph api %s", err)
		return
	}
	payload := GetProfilePayload{
		user.Name, user.WalletAddress, user.ProfilePictureUrl, user.Country, roles,
	}
	httphelper.SuccessResponse(c, "Profile fetched successfully", payload)
}

type rolesResponse struct {
	Data struct {
		User struct {
			Roles []string `json:"roles"`
		} `json:"user"`
	} `json:"data"`

	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

var (
	errStatusCode = errors.New("status code is not 200")
)

func getRoles(walletAddress string) ([]string, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`
		{ 
			user(id:"%v"){
				roles
			}
		}
	`, strings.ToLower(walletAddress)),
	}

	jsonValue, _ := json.Marshal(jsonData)

	request, err := http.NewRequest("POST", envutil.MustGetEnv("GRAPH_API"), bytes.NewBuffer(jsonValue))
	if err != nil {
		return []string{}, err
	}
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)

	if err != nil {
		return []string{}, err
	}

	if response.StatusCode != 200 {
		return []string{}, errStatusCode
	}
	data, _ := ioutil.ReadAll(response.Body)
	var res rolesResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []string{}, err
	}

	if len(res.Errors) > 1 {
		return []string{}, errors.New(res.Errors[0].Message)
	}
	err = response.Body.Close()
	if err != nil {
		logwrapper.Warnf("failed to close body, error :%s", err)
	}
	return res.Data.User.Roles, nil
}
