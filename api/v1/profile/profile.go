package profile

import (
	"net/http"
	"strings"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/httpo"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/profile")
	{
		g.Use(paseto.PASETO(false))
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
		Email:             &requestBody.EmailId,
		Country:           requestBody.Country,
		Discord:           requestBody.Discord,
		Twitter:           requestBody.Twitter,
		Google:            requestBody.Google,
		Apple:             requestBody.Apple,
		Telegram:          requestBody.Telegram,
		Farcaster:         requestBody.Farcaster,
	}
	userId := c.GetString(paseto.CTX_USER_ID)
	if userId == "" {
		httpo.NewErrorResponse(http.StatusForbidden, "User not found").SendD(c)
		return
	}

	result := db.Model(&models.User{}).
		Where("user_id = ?", userId).
		Updates(&profileUpdate)

	if result.Error != nil {
		// Check if the error is a PostgreSQL error
		if errMsg := result.Error.Error(); errMsg != "" {
			// Check for duplicate key violation (unique constraint)
			if strings.Contains(errMsg, "duplicate key value violates unique constraint") {
				httpo.NewErrorResponse(http.StatusConflict, "Email address already in use for another account").SendD(c)
				return
			}
		}
		// Handle other errors
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occurred").SendD(c)
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
	userId := c.GetString(paseto.CTX_USER_ID)
	var user models.User
	err := db.Model(&models.User{}).Select("user_id, name, profile_picture_url,country, wallet_address, discord, twitter, email, apple, telegram, farcaster, google, chain_name").Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		logrus.Error(err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		return
	}

	payload := GetProfilePayload{
		user.UserId, user.Name, user.WalletAddress, user.ProfilePictureUrl, user.Country, user.Discord, user.Twitter, user.Email, user.Apple, user.Telegram, user.Farcaster, user.Google, user.ChainName,
	}
	httpo.NewSuccessResponseP(200, "Profile fetched successfully", payload).SendD(c)
}
