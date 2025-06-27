package admin

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/NetSepio/gateway/internal/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/utils/httpo"
	"github.com/NetSepio/gateway/utils/logwrapper"
)

func deleteAdmin(c *gin.Context) {
	db := database.GetDb()
	var request DeleteAdminRequest
	err := c.BindJSON(&request)
	if err != nil {
		//TODO not override status or not set status again
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid %s", err)).SendD(c)
		return
	}

	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES)

	err = db.Model(&models.DomainAdmin{}).
		Where(&models.DomainAdmin{DomainId: request.DomainId, AdminId: walletAddress}).
		First(&models.DomainAdmin{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httpo.NewErrorResponse(http.StatusNotFound, "domain not exist or user is not admin of the domain").SendD(c)
			return
		}

		logwrapper.Errorf("failed to get domain admin: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to update admin").SendD(c)
	}

	res := db.Delete(&models.DomainAdmin{DomainId: request.DomainId, AdminId: strings.ToLower(request.AdminWalletAddres)})
	if res.Error != nil {
		logwrapper.Errorf("failed to delete domain admin: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to remove admin").SendD(c)
		return
	}

	if res.RowsAffected == 0 {
		httpo.NewErrorResponse(http.StatusNotFound, "admin not exist for that domain").SendD(c)
		return
	}

	httpo.NewSuccessResponse(200, "admin removed").SendD(c)
}
