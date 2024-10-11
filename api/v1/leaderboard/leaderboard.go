package leaderboard

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/getleaderboard")
	{
		g.GET("", getLeaderboard)
	}
	h := r.Group("/getScoreboard")
	{
		h.GET("", getScoreBoard)
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
func getScoreBoard(c *gin.Context) {
	db := dbconfig.GetDb()

	var response []models.UserScoreBoard

	var scoreBoard []models.ScoreBoard

	if err := db.Order("reviews desc").Find(&scoreBoard).Error; err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occurred").SendD(c)
		logwrapper.Error("failed to get scoreBoard", err)
		return
	}

	if len(scoreBoard) == 0 {
		httpo.NewErrorResponse(404, "No ScoreBoard entries found").SendD(c)
		return
	}

	for _, data := range scoreBoard {

		total := data.Reviews + data.Domain + data.Nodes + data.DWifi + data.Discord + data.Twitter + data.Telegram
		UserScoreBoard := models.UserScoreBoard{
			ID:         data.ID,
			Reviews:    data.Reviews,
			Domain:     data.Domain,
			UserId:     data.UserId,
			Nodes:      data.Nodes,
			DWifi:      data.DWifi,
			Discord:    data.Discord,
			Twitter:    data.Twitter,
			Telegram:   data.Telegram,
			CreatedAt:  data.CreatedAt,
			UpdatedAt:  data.UpdatedAt,
			TotalScore: total,
		}

		var user models.User
		err := db.Model(&models.User{}).Select("user_id, name, profile_picture_url,country, wallet_address, discord, twitter, email_id").Where("user_id = ?", data.UserId).First(&user).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				response = append(response, UserScoreBoard)
				continue
			} else {
				logrus.Error(err)
				httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
				return
			}

		}

		payload := models.GetProfilePayload{UserId: user.UserId, Name: user.Name, WalletAddress: user.WalletAddress, ProfilePictureUrl: user.ProfilePictureUrl, Country: user.Country, Discord: user.Discord, Twitter: user.Twitter, Email: user.EmailId}

		UserScoreBoard.UserDetails = payload

		response = append(response, UserScoreBoard)
	}
	fmt.Printf("%+v\n", response)

	httpo.NewSuccessResponseP(200, "ScoreBoard fetched successfully", response).SendD(c)
}
func AutoCalculateScoreBoard() {

	// fmt.Println("STARTING AUTO CALCULATE SCOREBOARD AT ", time.Now())
	border := strings.Repeat("=", 50) // Creates a border line

	func() {
		fmt.Println(border)
		fmt.Println("🚀 STARTING AUTO CALCULATE SCOREBOARD")
		fmt.Println("📅 Date & Time:", time.Now().Format("02-Jan-2006 15:04:05 MST"))
		fmt.Println("🔄 Status: In Progress")
		fmt.Println(border)
	}()

	// CronForReviewUpdate()

	// var leaderboard ScoreBoard
	leaderboards, err := GetAllLeaderBoard()
	if err != nil {
		logrus.Error(err)
		return
	}

	fmt.Println("leaderboards len : ", len(leaderboards))

	for _, leaderboard := range leaderboards {
		CronJobLeaderBoardUpdate("reviews", leaderboard)
		CronJobLeaderBoardUpdate("domain", leaderboard)
		CronJobLeaderBoardUpdate("nodes", leaderboard)
		CronJobLeaderBoardUpdate("d_wifi", leaderboard)
		CronJobLeaderBoardUpdate("discord", leaderboard)
		CronJobLeaderBoardUpdate("twitter", leaderboard)
		CronJobLeaderBoardUpdate("telegram", leaderboard)
	}

	func() {
		// After the task completes, print the "Completed" status
		// border := strings.Repeat("=", 50) // Creates a border line
		fmt.Println(border)
		fmt.Println("✅ SCOREBOARD CALCULATION COMPLETED")
		fmt.Println("📅 Date & Time:", time.Now().Format("02-Jan-2006 15:04:05 MST"))
		fmt.Println("✔️ Status: Completed")
		fmt.Println(border)
	}()

}
