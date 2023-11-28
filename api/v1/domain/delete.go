package domain

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	errAdminNotFound = errors.New("domain admin not found")
)

func deleteDomain(c *gin.Context) {
	db := dbconfig.GetDb()
	var request DeleteDomainQuery
	err := c.BindQuery(&request)
	if err != nil {
		//TODO not override status or not set status again
		httpo.NewErrorResponse(http.StatusBadRequest, fmt.Sprintf("payload is invalid %s", err)).SendD(c)
		return
	}

	walletAddress := c.GetString(paseto.CTX_WALLET_ADDRES)

	err = db.Transaction(func(tx *gorm.DB) error {
		result := tx.Exec(`
		delete from domain_admins where domain_id=? and admin_wallet_address=?;
		`, request.DomainId, strings.ToLower(walletAddress))

		if err := result.Error; err != nil {
			return err
		}
		if result.RowsAffected == 0 {
			return errAdminNotFound
		}
		result = tx.Exec(`
		DELETE from domains 
		WHERE id = ?;
    	`, request.DomainId)

		if err := result.Error; err != nil {
			logwrapper.Errorf("failed to delete domain: %s", err)
			return err
		}

		if result.RowsAffected == 0 {
			return fmt.Errorf("no domain was deleted")
		}

		return nil
	})

	if err != nil {
		if errors.Is(err, errAdminNotFound) {
			httpo.NewErrorResponse(http.StatusNotFound, "domain not exist or user is not admin of the domain").SendD(c)
			return
		}
		logwrapper.Errorf("failed to delete domain records: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "failed to delete domain").SendD(c)
	}
	httpo.NewSuccessResponse(200, "domain deleted").SendD(c)

}
