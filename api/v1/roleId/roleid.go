package roleid

import (
	"net/http"
	"netsepio-api/db"
	jwtMiddleWare "netsepio-api/middleware/auth/jwt"
	"netsepio-api/models"
	"netsepio-api/util/pkg/flowid"
	"netsepio-api/util/pkg/httphelper"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
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
	walletAddress := c.GetString("walletAddress")
	roleId, exist := c.Params.Get("roleId")
	if !exist {
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")

		return
	}
	roleIdInt, err := strconv.Atoi(roleId)
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")

		return
	}
	var role models.Role
	err = db.Db.Model(&models.Role{}).Where("role_id = ?", roleIdInt).First(&role).Error
	if err == gorm.ErrRecordNotFound {
		httphelper.ErrResponse(c, http.StatusNotFound, err.Error())

	} else if err != nil {
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")

	} else {
		flowId, err := flowid.GenerateFlowId(walletAddress, true, models.ROLE, roleIdInt)
		if err != nil {
			logrus.Error(err)
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
