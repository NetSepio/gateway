package getreviewerdetails

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"gorm.io/gorm"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/utils/httpo"
	"github.com/NetSepio/gateway/utils/load"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/reviewerdetails")
	{
		g.GET("", getProfile)
	}
}

func getProfile(c *gin.Context) {
	db := database.GetDb()
	var request GetReviewerDetailsQuery
	err := c.BindQuery(&request)
	if err != nil {
		httpo.NewErrorResponse(http.StatusForbidden, fmt.Sprintf("payload is invalid: %s", err)).SendD(c)
		return
	}
	var user models.User
	err = db.Model(&models.User{}).Select("name, profile_picture_url, wallet_address, discord, twitter").Where("wallet_address = ?", strings.ToLower(request.WalletAddress)).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			load.Logger.Error(err.Error())
			httpo.NewErrorResponse(http.StatusNotFound, "profile not found").SendD(c)
			return
		}
		load.Logger.Error(err.Error())
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		return
	}

	payload := GetReviewerDetailsPayload{
		Name:              user.Name,
		WalletAddress:     user.WalletAddress,
		ProfilePictureUrl: user.ProfilePictureUrl,
		Discord:           user.Discord,
		Twitter:           user.Twitter,
	}
	httpo.NewSuccessResponseP(200, "Profile fetched successfully", payload).SendD(c)
}
