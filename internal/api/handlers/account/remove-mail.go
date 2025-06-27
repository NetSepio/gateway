package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/NetSepio/gateway/internal/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/utils/httpo"
	"github.com/NetSepio/gateway/utils/logwrapper"
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
