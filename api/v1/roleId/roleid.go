package roleid

import (
	"net/http"

	jwtMiddleWare "github.com/TheLazarusNetwork/marketplace-engine/api/middleware/auth/jwt"
	"github.com/TheLazarusNetwork/marketplace-engine/config/dbconfig"
	"github.com/TheLazarusNetwork/marketplace-engine/models"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/flowid"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/httphelper"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/roleId")
	{
		g.Use(jwtMiddleWare.JWT)
		g.GET(":roleId", getRoleId)
	}
}

func getRoleId(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	roleId, exist := c.Params.Get("roleId")
	if !exist {
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")

		return
	}
	var role models.Role
	err := db.Model(&models.Role{}).Where("role_id = ?", roleId).First(&role).Error
	if err == gorm.ErrRecordNotFound {
		httphelper.ErrResponse(c, http.StatusNotFound, err.Error())

	} else if err != nil {
		logwrapper.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")

	} else {
		flowId, err := flowid.GenerateFlowId(walletAddress, true, models.ROLE, roleId)
		if err != nil {
			logwrapper.Error(err)
			httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
			c.Status(http.StatusInternalServerError)
			return
		}

		payload := GetRoleIdPayload{
			role.Eula, flowId,
		}
		httphelper.SuccessResponse(c, "Flow id successfully generated", payload)

	}

}
