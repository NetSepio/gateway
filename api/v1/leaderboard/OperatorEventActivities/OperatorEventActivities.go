package OperatorEventActivities

import (
	"net/http"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/getleaderboard")
	{
		g.GET("", getLeaderboard)
	}
}

func getLeaderboard(c *gin.Context) {
	db := dbconfig.GetDb()

	var leaderboard []models.Leaderboard

	if err := db.Order("reviews desc").Find(&leaderboard).Error; err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occurred").SendD(c)
		logwrapper.Error("failed to get leaderboard", err)
		return
	}

	if len(leaderboard) == 0 {
		httpo.NewErrorResponse(404, "No leaderboard entries found").SendD(c)
		return
	}

	httpo.NewSuccessResponseP(200, "Leaderboard fetched successfully", leaderboard).SendD(c)
}
