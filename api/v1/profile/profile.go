package profile

import (
	"net/http"

	jwtMiddleWare "github.com/TheLazarusNetwork/netsepio-engine/api/middleware/auth/jwt"
	"github.com/TheLazarusNetwork/netsepio-engine/config/dbconfig"
	"github.com/TheLazarusNetwork/netsepio-engine/models"
	"github.com/TheLazarusNetwork/netsepio-engine/util/pkg/httphelper"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/profile")
	{
		g.Use(jwtMiddleWare.JWT)
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
		Update(requestBody)
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

	payload := GetProfilePayload{
		user.Name, user.WalletAddress, user.ProfilePictureUrl, user.Country,
	}
	httphelper.SuccessResponse(c, "Profile fetched successfully", payload)
}
