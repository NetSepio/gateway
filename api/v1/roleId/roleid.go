package roleid

import (
	"net/http"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/flowid"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/roleId")
	{
		g.Use(paseto.PASETO)
		g.GET(":roleId", GetRoleId)
	}
}

func GetRoleId(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	roleId, exist := c.Params.Get("roleId")
	if !exist {
		httpo.NewErrorResponse(http.StatusBadRequest, "Param roleId is required").SendD(c)
		return
	}
	var role models.Role
	err := db.Model(&models.Role{}).Where("role_id = ?", roleId).First(&role).Error
	if err == gorm.ErrRecordNotFound {
		httpo.NewErrorResponse(http.StatusNotFound, err.Error()).SendD(c)

	} else if err != nil {
		logwrapper.Errorf("failed to fetch details for roleId: %v, error: %v", roleId, err.Error())
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
	} else {
		flowId, err := flowid.GenerateFlowId(walletAddress, models.ROLE, roleId)
		if err != nil {
			logwrapper.Errorf("failed to generate flow id, err: %v", err.Error())
			httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
			c.Status(http.StatusInternalServerError)
			return
		}

		payload := GetRoleIdPayload{
			role.Eula, flowId,
		}
		httpo.NewSuccessResponseP(200, "Flow id successfully generated", payload).SendD(c)

	}

}
