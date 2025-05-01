package admin

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"netsepio-gateway-v1.1/internal/api/middleware/auth/paseto"
	"netsepio-gateway-v1.1/internal/database"
	"netsepio-gateway-v1.1/models"
	"netsepio-gateway-v1.1/utils/httpo"
	"netsepio-gateway-v1.1/utils/logwrapper"
)

func getAdmin(c *gin.Context) {
	db := database.GetDb()
	var request GetAdminQuery
	err := c.BindJSON(&request)
	if err != nil {
		//TODO not override status or not set status again
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid %s", err)).SendD(c)
		return
	}

	userId := c.GetString(paseto.CTX_USER_ID)

	err = db.Model(&models.DomainAdmin{}).
		Where(&models.DomainAdmin{DomainId: request.DomainId, AdminId: userId}).
		First(&models.DomainAdmin{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httpo.NewErrorResponse(http.StatusNotFound, "domain not exist or user is not admin of the domain").SendD(c)
			return
		}

		logwrapper.Errorf("failed to get domain admin: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to get admins").SendD(c)
	}

	var domainAdmins []models.DomainAdmin
	err = db.Where(&models.DomainAdmin{DomainId: request.DomainId}).Find(&domainAdmins).Error
	if err != nil {
		logwrapper.Errorf("failed to get domain admins: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to get admins").SendD(c)
	}

	httpo.NewSuccessResponseP(200, "fetch admins", domainAdmins).SendD(c)
}
