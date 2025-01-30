package profile

import (
	"fmt"
	"net/http"

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
		g.PATCH("/socials", updateUser)
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
	userId := c.GetString(paseto.CTX_USER_ID)
	result := db.Model(&models.User{}).
		Where("user_id = ?", userId).
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
	userId := c.GetString(paseto.CTX_USER_ID)
	var user models.User
	err := db.Model(&models.User{}).Select("user_id, name, profile_picture_url,country, wallet_address, discord, twitter, email_id").Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		logrus.Error(err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		return
	}

	payload := GetProfilePayload{
		user.UserId, user.Name, user.WalletAddress, user.ProfilePictureUrl, user.Country, user.Discord, user.Twitter, user.EmailId,
	}
	httpo.NewSuccessResponseP(200, "Profile fetched successfully", payload).SendD(c)
}

func updateUser(c *gin.Context) {
	var updateReq UpdateUserRequest
	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES)
	if len(walletAddress) == 0 {
		logrus.Errorf("Wallet address not found in context")
		httpo.NewErrorResponse(http.StatusBadRequest, "Failed to get wallet address by paseto").SendD(c)
		return

	}

	// Fetch the user by wallet address
	var user models.User
	db := dbconfig.GetDb()
	if err := db.Where("wallet_address = ?", walletAddress).First(&user).Error; err != nil {
		logrus.Errorf("User not found with wallet address: %s, error: %v", walletAddress, err)
		httpo.NewErrorResponse(http.StatusNotFound, "User not found").SendD(c)
		return
	}

	// Bind the incoming JSON to the update request struct
	if err := c.ShouldBindJSON(&updateReq); err != nil {
		logrus.Errorf("Invalid input, error: %v", err)
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("Invalid input: %v", err)).SendD(c)
		return
	}

	// Update fields using the helper function
	updateField(updateReq.Discord, &user.Discord)
	updateField(updateReq.X, &user.X)
	if updateReq.Google != nil {
		updateField(*updateReq.Google, user.Google)
	}
	if updateReq.AppleId != nil {
		updateField(*updateReq.AppleId, user.AppleId)
	}
	updateField(updateReq.Telegram, &user.Telegram)
	if updateReq.Farcaster != nil {
		updateField(*updateReq.Farcaster, user.Farcaster)
	}

	// Save the updated user
	if err := db.Save(&user).Error; err != nil {
		logrus.Errorf("Failed to update user, error: %v", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Failed to update user").SendD(c)
		return
	}

	// Log the success of the update
	logrus.Infof("User updated successfully with wallet address: %s", walletAddress)
	httpo.NewSuccessResponse(http.StatusOK, "User updated successfully").SendD(c)
}

// Helper function to update a field if its length is greater than 0
func updateField(value string, pointer *string) {
	if len(value) > 0 {
		*pointer = value
	}
}
