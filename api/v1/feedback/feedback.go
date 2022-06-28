package feedback

import (
	"net/http"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/httphelper"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/feedback")
	{
		g.Use(paseto.PASETO)
		g.POST("", createFeedback)
	}
}

func createFeedback(c *gin.Context) {
	db := dbconfig.GetDb()
	var requestBody PostFeedbackRequest
	c.BindJSON(&requestBody)
	walletAddress := c.GetString("walletAddress")

	result := db.Model(&models.User{}).Where("wallet_address = ?", walletAddress).
		Update("feedbacks", gorm.Expr("feedbacks || ?", pq.StringArray([]string{requestBody.Feedback})))

	if result.Error != nil {
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")

		return
	}
	if result.RowsAffected == 0 {
		httphelper.ErrResponse(c, http.StatusNotFound, "Record not found")

		return
	}
	httphelper.SuccessResponse(c, "Feedback added", nil)

}
