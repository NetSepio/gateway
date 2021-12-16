package flowid

import (
	"net/http"
	"netsepio-api/db"
	"netsepio-api/models"
	"netsepio-api/types"
	"netsepio-api/util/pkg/flowid"
	"netsepio-api/util/pkg/httphelper"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/flowid")
	{
		g.GET("", getFlowId)
	}
}

func getFlowId(c *gin.Context) {
	var user models.User
	walletAddress := c.Query("walletAddress")
	if walletAddress == "" {
		httphelper.ErrResponse(c, http.StatusBadRequest, "Wallet address (walletAddress) is required")
		return
	}
	dbRes := db.Db.Model(&models.User{}).Where("wallet_address = ?", walletAddress).First(&user)
	// If there is an error and that error is not of "record not found"
	if dbRes.Error != nil && dbRes.Error != gorm.ErrRecordNotFound {
		log.Error(dbRes.Error)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	// If wallet address exist
	if dbRes.Error != gorm.ErrRecordNotFound {
		flowId, err := flowid.GenerateFlowId(walletAddress, true, models.AUTH, 0)
		if err != nil {
			log.Error(err)
			httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")

			return
		}
		payload := GetFlowIdPayload{
			FlowId: flowId,
		}
		response := types.ApiResponse{
			Status:  http.StatusOK,
			Payload: payload,
		}
		c.JSON(http.StatusOK, response)
	} else {
		//If wallet address doesn't exist
		flowId, err := flowid.GenerateFlowId(walletAddress, false, models.AUTH, 0)
		if err != nil {
			log.Error(err)
			httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")

			return
		}
		payload := GetFlowIdPayload{
			FlowId: flowId,
			Eula:   "TODO eula",
		}
		httphelper.SuccessResponse(c, "Role successfully claimed", payload)
	}
}
