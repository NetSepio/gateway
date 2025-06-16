package registerDwifi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"netsepio-gateway-v1.1/internal/database"
	"netsepio-gateway-v1.1/models"
	"netsepio-gateway-v1.1/utils/httpo"
	"netsepio-gateway-v1.1/utils/logwrapper"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/registernode")
	{
		g.POST("", RegisterWifiNode)
	}
}

func RegisterWifiNode(c *gin.Context) {
	db := database.GetDB2()
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
