package leaderboard

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"

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
	h := r.Group("/getScoreboard/top10")
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

		payload := models.User{UserId: user.UserId, Name: user.Name, WalletAddress: user.WalletAddress, ProfilePictureUrl: user.ProfilePictureUrl, Country: user.Country, Discord: user.Discord, Twitter: user.Twitter, EmailId: user.EmailId}

		UserScoreBoard.UserDetails = payload

		response = append(response, UserScoreBoard)

		// Sort the response by TotalScore in descending order
		sort.SliceStable(response, func(i, j int) bool {
			return response[i].TotalScore > response[j].TotalScore
		})

		// Take the top 10 entries
		if len(response) > 10 {
			response = response[:10]
		}
	}

	httpo.NewSuccessResponseP(200, "ScoreBoard fetched successfully", response).SendD(c)
}

func getAllUsersScoreBoard(c *gin.Context) {
	db := dbconfig.GetDb()
	var response []models.UserScoreBoard

	var users []models.User
	// Omit FlowIds and Feedbacks from the query
	if err := db.Omit("FlowIds", "Feedbacks").Find(&users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("no rows in the user table")
		}
		return
	}

	fmt.Println("len user_details : ", len(users))
	// fmt.Printf("%+v\n", users)

	for _, userDetail := range users {
		var scoreBoard ScoreBoard

		// Query to find the first ScoreBoard record by UserId
		if err := db.Where("user_id = ?", userDetail.UserId).First(&scoreBoard).Error; err != nil {
			// Handle error, e.g., record not found
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// No record found for the given UserId
				var data models.UserScoreBoard
				data.UserDetails = userDetail
				response = append(response, data)
			} else {
				logrus.Error(err)
			}
		} else {
			var data = models.UserScoreBoard{
				ID:        scoreBoard.ID,
				Reviews:   scoreBoard.Reviews,
				Domain:    scoreBoard.Domain,
				UserId:    scoreBoard.UserId,
				Nodes:     scoreBoard.Nodes,
				DWifi:     scoreBoard.DWifi,
				Discord:   scoreBoard.Discord,
				Twitter:   scoreBoard.Twitter,
				Telegram:  scoreBoard.Telegram,
				CreatedAt: scoreBoard.CreatedAt,
				UpdatedAt: scoreBoard.UpdatedAt,
			}
			data.UserDetails = userDetail
			response = append(response, data)

		}
	}

	httpo.NewSuccessResponseP(200, "ScoreBoard fetched successfully", response).SendD(c)
}
