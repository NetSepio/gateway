package roleid

import (
	"net/http"
	"netsepio-api/db"
	jwtMiddleWare "netsepio-api/middleware/auth/jwt"
	"netsepio-api/models"
	"netsepio-api/util/pkg/flowid"
	"strconv"

	"github.com/gin-gonic/gin"
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
		c.Status(http.StatusInternalServerError)
		return
	}
	roleIdInt, err := strconv.Atoi(roleId)
	if err != nil {
		logrus.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	var role models.Role
	err = db.Db.Model(&models.Role{}).Where("role_id = ?", roleIdInt).First(&role).Error
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		flowId, err := flowid.GenerateFlowId(walletAddress, true, models.ROLE, roleIdInt)
		if err != nil {
			logrus.Error(err)
			c.Status(http.StatusInternalServerError)
			return
		}

		response := GetRoleIdResponse{
			role.Eula, flowId,
		}
		c.JSON(200, response)
	}

}
