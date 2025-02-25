package registerDwifi

import (
	"net/http"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/httpo"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/registernode")
	{
		g.POST("", RegisterWifiNode)
	}
}

func RegisterWifiNode(c *gin.Context) {
	db := dbconfig.GetDb()
	var wifiNode models.WifiNode
	if err := c.ShouldBindJSON(&wifiNode); err != nil {
		logwrapper.Errorf("failed to bind JSON: %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, err.Error()).SendD(c)
		return
	}

	// Save the WiFi node to the database
	if err := db.Create(&wifiNode).Error; err != nil {
		logwrapper.Errorf("failed to save node to DB: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, err.Error()).SendD(c)
		return
	}

	httpo.NewSuccessResponseP(http.StatusOK, "WiFi node registered successfully", wifiNode).SendD(c)
}
