package account

import (
	"fmt"
	"net/http"

	useractivity "github.com/NetSepio/gateway/internal/api/handlers/userActivity"
	"github.com/NetSepio/gateway/internal/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/utils/actions"
	"github.com/NetSepio/gateway/utils/httpo"
	"github.com/NetSepio/gateway/utils/logwrapper"
	"github.com/NetSepio/gateway/utils/module"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// router for removing mail from account
func removeMail(c *gin.Context) {
	db := database.GetDb()
	userId := c.GetString(paseto.CTX_USER_ID)

	user := models.User{}

	if err := db.Where("user_id = ?", userId).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			httpo.NewErrorResponse(http.StatusNotFound, "User not found").SendD(c)
		} else {
			httpo.NewErrorResponse(http.StatusInternalServerError, fmt.Sprintf("Database error: %s", err)).SendD(c)
		}
		return
	}

	err := db.Model(&models.User{}).Where("user_id = ?", userId).Update("email", nil).Error
	if err != nil {

		logwrapper.Errorf("failed to update user email: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "internal server error").SendD(c)
		return
	}

	idenity := user.Name
	if idenity == "" && user.WalletAddress != nil {
		idenity = *user.WalletAddress
	}

	meta := "Email removed for " + idenity

	go useractivity.Save(models.UserActivity{UserId: userId, Modules: module.Account, Action: actions.Updated, Metadata: &meta})

	httpo.NewSuccessResponse(200, "email removed successfully").SendD(c)
}
