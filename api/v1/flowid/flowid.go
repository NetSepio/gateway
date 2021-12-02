package flowid

import (
	"net/http"
	"netsepio-api/db"
	"netsepio-api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
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
	dbRes := db.Db.Where("wallet_address = ?", request.WalletAddress).First(&user)
	// If there is an error and that error is not of "record not found"
	if dbRes.Error != nil && dbRes.Error != gorm.ErrRecordNotFound {
		log.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	// If wallet address exist
	if dbRes.Error != gorm.ErrRecordNotFound {
		flowId, err := generateFlowId(request.WalletAddress, true)
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

		flowId, err := generateFlowId(request.WalletAddress, false)
		if err != nil {
			log.Error(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, GetFlowIdResponse{
			Message: "test eula",
			FlowId:  flowId,
		})
	}
}

func generateFlowId(walletAddress string, update bool) (string, error) {

	flowId := uuid.NewString()
	if update {
		// User exist so update
		err := db.Db.Model(&models.User{}).
			Where("wallet_address = ?", walletAddress).
			Update("flow_id", gorm.Expr("array_cat(flow_id,?)", pq.Array([]string{flowId}))).Error
		if err != nil {
			return "", err
		}
	} else {
		// User doesn't exist so create
		newUser := &models.User{
			WalletAddress: walletAddress,
			FlowId:        pq.StringArray([]string{flowId}),
		}
		if err := db.Db.Create(newUser).Error; err != nil {
			return "", err
		}
	}

	return flowId, nil
}
