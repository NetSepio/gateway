package account

import (
	"net/http"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/httpo"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"

	"github.com/gin-gonic/gin"
)

// router for removing mail from account
func removeMail(c *gin.Context) {
	db := dbconfig.GetDb()
	userId := c.GetString(paseto.CTX_USER_ID)

	err := db.Model(&models.User{}).Where("user_id = ?", userId).Update("email_id", nil).Error
	if err != nil {
		logwrapper.Errorf("failed to update user email: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}

	httpo.NewSuccessResponse(200, "email removed successfully").SendD(c)
}
