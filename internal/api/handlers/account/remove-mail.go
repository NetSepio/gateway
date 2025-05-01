package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"netsepio-gateway-v1.1/internal/api/middleware/auth/paseto"
	"netsepio-gateway-v1.1/internal/database"
	"netsepio-gateway-v1.1/models"
	"netsepio-gateway-v1.1/utils/httpo"
	"netsepio-gateway-v1.1/utils/logwrapper"
)

// router for removing mail from account
func removeMail(c *gin.Context) {
	db := database.GetDb()
	userId := c.GetString(paseto.CTX_USER_ID)

	err := db.Model(&models.User{}).Where("user_id = ?", userId).Update("email", nil).Error
	if err != nil {

		logwrapper.Errorf("failed to update user email: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}

	httpo.NewSuccessResponse(200, "email removed successfully").SendD(c)
}
