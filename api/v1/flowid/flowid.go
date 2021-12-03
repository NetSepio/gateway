package flowid

import (
	"net/http"
	"netsepio-api/db"
	"netsepio-api/models"
	"netsepio-api/util/pkg/flowid"

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
	var request GetFlowIdRequest
	err := c.BindJSON(&request)
	if err != nil {
		log.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	dbRes := db.Db.Model(&models.User{}).Where("wallet_address = ?", request.WalletAddress).First(&user)
	// If there is an error and that error is not of "record not found"
	if dbRes.Error != nil && dbRes.Error != gorm.ErrRecordNotFound {
		log.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}
	// If wallet address exist
	if dbRes.Error != gorm.ErrRecordNotFound {
		flowId, err := flowid.GenerateFlowId(request.WalletAddress, true, models.AUTH, 0)
		if err != nil {
			log.Error(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, GetFlowIdResponse{
			Message: "flow id generated",
			FlowId:  flowId,
		})
	} else {
		//If wallet address doesn't exist
		flowId, err := flowid.GenerateFlowId(request.WalletAddress, false, models.AUTH, 0)
		if err != nil {
			log.Error(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, GetFlowIdResponse{
			Message: "TODO eula",
			FlowId:  flowId,
		})
	}
}
