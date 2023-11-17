package profile

import (
	"net/http"
	"strings"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/TheLazarusNetwork/go-helpers/httpo"

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
	err := c.BindJSON(&requestBody)
	if err != nil {
		httpo.NewErrorResponse(http.StatusForbidden, "payload is invalid").SendD(c)
		return
	}

	profileUpdate := models.User{
		Name:              requestBody.Name,
		ProfilePictureUrl: requestBody.ProfilePictureUrl,
		Country:           requestBody.Country,
		Discord:           requestBody.Discord,
		Twitter:           requestBody.Twitter,
	}
	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES)
	result := db.Model(&models.User{}).
		Where("wallet_address = ?", strings.ToLower(walletAddress)).
		Updates(&profileUpdate)
	if result.Error != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)

		return
	}
	if result.RowsAffected == 0 {
		httpo.NewErrorResponse(http.StatusNotFound, "Record not found").SendD(c)

		return
	}
	httpo.NewSuccessResponse(200, "Profile successfully updated").SendD(c)

}

func getProfile(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES)
	var user models.User
	err := db.Model(&models.User{}).Select("name, profile_picture_url,country, wallet_address, discord, twitter").Where("wallet_address = ?", strings.ToLower(walletAddress)).First(&user).Error
	if err != nil {
		logrus.Error(err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		return
	}

	payload := GetProfilePayload{
		user.Name, user.WalletAddress, user.ProfilePictureUrl, user.Country, user.Discord, user.Twitter,
	}
	httpo.NewSuccessResponseP(200, "Profile fetched successfully", payload).SendD(c)
}
