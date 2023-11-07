package feedback

import (
	"net/http"

	"github.com/NetSepio/gateway/api/middleware/auth/paseto"
	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"

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
	var newFeedback models.UserFeedback
	err := c.BindJSON(&newFeedback)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, "body is invalid").SendD(c)
		return
	}
	walletAddress := c.GetString("walletAddress")
	newFeedback.WalletAddress = walletAddress
	association := db.Model(&models.User{
		WalletAddress: walletAddress,
	}).Association("Feedbacks")

	if err = association.Error; err != nil {
		logwrapper.Errorf("failed to associate feedbacks with users, %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		return
	}
	err = association.Append(&newFeedback)
	if err != nil {
		logwrapper.Errorf("failed to add new feedback, %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
		return
	}
	httpo.NewSuccessResponse(200, "Feedback added").SendD(c)
}
