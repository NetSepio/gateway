package roleid

import (
	"fmt"
	"net/http"
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

type Roles []Role
type Role struct {
	id   int
	name string
	eula string
}

func getRoleId(c *gin.Context) {
	roles := Roles{{1, "Manager", "TODO managerEula"}}
	walletAddress := c.GetString("walletAddress")

	roleId, exist := c.Params.Get("roleId")
	if !exist {
		c.Status(http.StatusInternalServerError)
		return
	}
	roleIdInt, err := strconv.Atoi(roleId)

	for _, v := range roles {
		if err != nil {
			logrus.Error(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		if v.id == roleIdInt {
			fmt.Println(v)
			flowId, err := flowid.GenerateFlowId(walletAddress, true, models.ROLE)

			if err != nil {
				logrus.Error(err)
				c.Status(http.StatusInternalServerError)
				return
			}

			response := GetRoleIdResponse{
				v.eula, flowId,
			}
			c.JSON(200, response)
			return
		}
	}
	c.String(http.StatusNotFound, fmt.Sprintf("Role with id %v is not found", roleIdInt))
}
