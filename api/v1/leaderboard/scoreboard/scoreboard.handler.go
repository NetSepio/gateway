package scoreboard

import (
	"strconv"

	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {

	scoreboard := r.Group("/scoreboard")

	scoreboard.GET("/get-top-list", getScoreBoardTopLeast)

}

func getScoreBoardTopLeast(c *gin.Context) {

	limitRequestBody := c.Request.FormValue("limit")
	scoreBoardType := c.Request.FormValue("score_board_type")

	limit, err := strconv.Atoi(limitRequestBody)
	if err != nil {
		httpo.NewErrorResponse(500, "Failed to convert string into int = getScoreBoardTopLeast").SendD(c)
	}

	data, err := GetTop10ScoreBoards(scoreBoardType, limit)
	if err != nil {
		httpo.NewErrorResponse(500, "Failed to get top 10 scoreboards = getScore")
	}

	httpo.NewSuccessResponseP(200, "Leaderboard fetched successfully", data).SendD(c)
}
