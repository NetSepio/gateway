package profile

import (
	"fmt"
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
	c.Status(http.StatusNotImplemented)
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
	err := db.Db.Model(&models.User{}).Select("name, profile_picture_url,country").Where("wallet_address = ?", walletAddress).First(&user).Error
	if err != nil {
		logrus.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	fmt.Println("user is", user)

	c.JSON(http.StatusOK, user)
}
