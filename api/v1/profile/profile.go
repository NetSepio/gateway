package profile

import (
	"net/http"
	"netsepio-api/db"
	jwtMiddleWare "netsepio-api/middleware/auth/jwt"
	"netsepio-api/models"

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
	var requestBody PatchProfileRequest
	c.BindJSON(&requestBody)
	walletAddress := c.GetString("walletAddress")
	result := db.Db.Model(&models.User{}).
		Where("wallet_address = ?", walletAddress).
		Update(requestBody)
	if result.Error != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		c.String(http.StatusNotFound, "Record not found")
		return
	}
	c.Status(http.StatusOK)

}

func getProfile(c *gin.Context) {
	walletAddress := c.GetString("walletAddress")
	var user models.User
	var userRoles []models.UserRole
	err := db.Db.Model(&models.User{}).Select("name, profile_picture_url,country, wallet_address").Where("wallet_address = ?", walletAddress).First(&user).Error
	if err != nil {
		logrus.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	err = db.Db.Model(&user).Association("Roles").Find(&userRoles).Error
	userRolesIds := make([]int, 0, 1)
	for _, userRole := range userRoles {
		userRolesIds = append(userRolesIds, userRole.RoleId)
	}
	if err != nil {
		logrus.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	responseData := GetProfileResponse{
		user.Name, user.WalletAddress, user.ProfilePictureUrl, user.Country, userRolesIds,
	}
	c.JSON(http.StatusOK, responseData)
}
