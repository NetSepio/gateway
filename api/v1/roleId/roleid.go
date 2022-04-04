package roleid

import (
	"net/http"

	jwtMiddleWare "github.com/NetSepio/gateway/api/middleware/auth/jwt"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/flowid"
	"github.com/NetSepio/gateway/util/pkg/httphelper"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/roleId")
	{
		g.Use(jwtMiddleWare.JWT)
		g.GET(":roleId", GetRoleId)
	}
}

func GetRoleId(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	roleId, exist := c.Params.Get("roleId")
	if !exist {
		httphelper.ErrResponse(c, http.StatusBadRequest, "Param roleId is required")
		return
	}
	var role models.Role
	err := db.Model(&models.Role{}).Where("role_id = ?", roleId).First(&role).Error
	if err == gorm.ErrRecordNotFound {
		httphelper.ErrResponse(c, http.StatusNotFound, err.Error())

	} else if err != nil {
		logwrapper.Errorf("failed to fetch details for roleId: %v, error: %v", roleId, err.Error())
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
	} else {
		flowId, err := flowid.GenerateFlowId(walletAddress, models.ROLE, roleId)
		if err != nil {
			logwrapper.Errorf("failed to generate flow id, err: %v", err.Error())
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
