package admin

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

func createAdmin(c *gin.Context) {
	db := dbconfig.GetDb()
	var request CreateAdminRequest
	err := c.BindJSON(&request)
	if err != nil {
		//TODO not override status or not set status again
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid %s", err)).SendD(c)
		return
	}

	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES)

	err = db.Model(&models.DomainAdmin{}).
		Where(&models.DomainAdmin{DomainId: request.DomainId, AdminWalletAddress: walletAddress}).
		First(&models.DomainAdmin{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httpo.NewErrorResponse(http.StatusNotFound, "domain not exist or user is not admin of the domain").SendD(c)
			return
		}

		logwrapper.Errorf("failed to get domain admin: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to update admin").SendD(c)
	}

	adminDetails := make([]models.DomainAdmin, len(request.Admins))

	i := 0
	for _, v := range request.Admins {
		adminDetails[i].DomainId = request.DomainId
		adminDetails[i].AdminWalletAddress = v.AdminWalletAddress
		adminDetails[i].UpdatedByAddress = walletAddress
		adminDetails[i].Name = v.AdminName
		adminDetails[i].Role = v.AdminRole
	}
	err = db.Create(&adminDetails).Error
	if err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			if pgError.Code == "23503" && pgError.ConstraintName == "fk_domain_admins_admin" {
				httpo.NewErrorResponse(http.StatusBadRequest, "admin address not found").SendD(c)
				return
			}

			if pgError.Code == "23505" && pgError.ConstraintName == "domain_admins_pkey" {
				httpo.NewErrorResponse(http.StatusBadRequest, "admin already exist").SendD(c)
				return
			}
		}
		logwrapper.Errorf("failed to get domain admin: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to update admin").SendD(c)
		return
	}

	httpo.NewSuccessResponse(200, "updated admins").SendD(c)
}
