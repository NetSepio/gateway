package profile

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/NetSepio/gateway/internal/api/handlers/profile/email"
	"github.com/NetSepio/gateway/internal/api/handlers/referral"
	useractivity "github.com/NetSepio/gateway/internal/api/handlers/userActivity"
	"github.com/NetSepio/gateway/internal/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/utils/actions"
	"github.com/NetSepio/gateway/utils/httpo"
	"github.com/NetSepio/gateway/utils/module"
	"gorm.io/gorm"

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
		r.GET("/origin", GetUserProfilesByOrigins)
		g.Group("/email")
		{
			g.POST("", email.SendOTP)
			g.POST("/verify", email.VerifyOTP)
		}

	}
}

func ApplyRoutesv11(r *gin.RouterGroup) {
	g := r.Group("/profile")
	{
		g.GET("/origin", GetUserProfilesByOrigins)
	}
}

func patchProfile(c *gin.Context) {
	db := database.GetDb()
	var requestBody PatchProfileRequest
	err := c.BindJSON(&requestBody)
	if err != nil {
		httpo.NewErrorResponse(http.StatusForbidden, "payload is invalid").SendD(c)
		return
	}

	var email *string
	if requestBody.EmailId == "" {
		email = nil // This will store NULL in the database
	} else {
		email = &requestBody.EmailId // Store the provided email
	}

	profileUpdate := models.User{
		Name:              requestBody.Name,
		ProfilePictureUrl: requestBody.ProfilePictureUrl,
		Email:             email,
		Country:           requestBody.Country,
		Discord:           requestBody.Discord,
		Twitter:           requestBody.Twitter,
		Google:            requestBody.Google,
		Apple:             requestBody.Apple,
		Telegram:          requestBody.Telegram,
		Farcaster:         requestBody.Farcaster,
		Metadata:          requestBody.Metadata,
		Origin:            requestBody.Origin,
	}

	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES)

	userId := c.GetString(paseto.CTX_USER_ID)

	if userId == "" || walletAddress == "" {
		httpo.NewErrorResponse(http.StatusForbidden, "User not found").SendD(c)
		return
	}

	if errr := func(google, walletAddress, userId string) error {
		var user models.User

		// Check if the user exists with the provided email and non-null wallet address
		err := db.Where("google = ? AND wallet_address != ? ", google, walletAddress).First(&user).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				// If google is already nil, delete the user
				if err := db.Where("google = ? AND (wallet_address IS NULL OR wallet_address = '')", google).Delete(&user).Error; err != nil {
					return fmt.Errorf("failed to delete user with wallet address %s: %v", *user.WalletAddress, err)
				} else {
					logrus.Info("User deleted successfully.")
					return nil
				}
			}
			return err
		}

		// If email exists, remove it from the user's account
		if user.Google != nil {
			user.Google = nil
			if err := db.Save(&user).Where("google = ? AND wallet_address != ? ", google, walletAddress).Error; err != nil {
				return fmt.Errorf("failed to remove email from user with wallet address %s: %v", *user.WalletAddress, err)
			}
			logrus.Info("Email removed successfully.")
		}

		return nil // Success
	}(*requestBody.Google, walletAddress, userId); errr != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, errr.Error()).SendD(c)
		return
	}

	result := db.Model(&models.User{}).
		Where("user_id = ?", userId).
		Updates(&profileUpdate).Debug()

	if result.Error != nil {
		// Check if the error is a PostgreSQL error
		if errMsg := result.Error.Error(); errMsg != "" {
			// Check for duplicate key violation (unique constraint)
			if strings.Contains(errMsg, "duplicate key value violates unique constraint") {
				go useractivity.Save(models.UserActivity{UserId: userId, Modules: module.Profile, Action: actions.Failed + " to " + actions.Updated, Metadata: " duplication"})
				httpo.NewErrorResponse(http.StatusConflict, "Email address already in use for another account").SendD(c)
				return
			}
			go useractivity.Save(models.UserActivity{UserId: userId, Modules: module.Profile, Action: actions.Failed + " to " + actions.Updated})
		}
		// Handle other errors
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occurred").SendD(c)
		return
	}

	go useractivity.Save(models.UserActivity{UserId: userId, Modules: module.Profile, Action: actions.Updated})

	httpo.NewSuccessResponse(200, "Profile successfully updated").SendD(c)
}

func getProfile(c *gin.Context) {
	db := database.GetDb()
	userId := c.GetString(paseto.CTX_USER_ID)
	var user models.User
	err := db.Model(&models.User{}).Select("user_id, name, profile_picture_url,country, wallet_address, discord, twitter, email, apple, telegram, farcaster, google, chain_name, referral_code, metadata").Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		logrus.Error(err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		return
	}

	if user.ReferralCode == "" {
		user.ReferralCode = referral.GenerateReferralCodeForUser(user)
		if user.ReferralCode == "" {
			logrus.Warnln("getProfile -> Failed in updating the referral code")
		}
	}

	payload := GetProfilePayload{
		user.UserId, user.Name, user.WalletAddress, user.ProfilePictureUrl, user.Country, user.Discord, user.Twitter, user.Email, user.Apple, user.Telegram, user.Farcaster, user.Google, user.ChainName, user.ReferralCode, user.Metadata,
	}
	httpo.NewSuccessResponseP(200, "Profile fetched successfully", payload).SendD(c)
}

// make one more get user profile by origin
func GetUserProfilesByOrigins(c *gin.Context) {
	db := database.GetDb()

	originsParam := c.Query("origins")
	if originsParam == "" {
		httpo.NewErrorResponse(http.StatusBadRequest, "Origins are required").SendD(c)
		return
	}

	var users []models.User
	err := db.Where("origin = ?", originsParam).Find(&users).Error
	if err != nil {
		logrus.Error(err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occurred").SendD(c)
		return
	}

	if len(users) == 0 {
		httpo.NewErrorResponse(http.StatusNotFound, "No profiles found").SendD(c)
		return
	} else {
		httpo.NewSuccessResponseP(200, "Profiles fetched successfully", users).SendD(c)
	}
}
