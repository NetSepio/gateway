package leaderboard

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"sync"

	"gorm.io/gorm"
	"netsepio-gateway-v1.1/internal/database"
	"netsepio-gateway-v1.1/models"
	"netsepio-gateway-v1.1/utils/httpo"
	"netsepio-gateway-v1.1/utils/load"
	"netsepio-gateway-v1.1/utils/logwrapper"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/getleaderboard")
	{
		g.GET("", getLeaderboard)
	}
	h := r.Group("/getScoreboard/top")
	{
		h.GET("", getScoreBoard)
	}
	u := r.Group("/updateOldUsersLeaderBoard")
	{
		u.GET("", UpdateLeaderBoardForAllUsers)
	}
	a := r.Group("/BetaGraphqlQueryForLeaderboard")
	{
		a.GET("", BetaGraphqlQueryForLeaderboardHandler)
	}
	b := r.Group("/ErebrusQueryForLeaderboardHandler")
	{
		b.GET("", ErebrusQueryForLeaderboardHandler)
	}
}

func getLeaderboard(c *gin.Context) {
	db := database.GetDb()

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
	db := database.GetDb()

	var response []models.UserScoreBoard

	var scoreBoard []models.ScoreBoard

	limitParam := c.Query("limit")
	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occurred").SendD(c)
		logwrapper.Error("failed to get scoreBoard", err)
		return
	}

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
			ID:           data.ID,
			Reviews:      data.Reviews,
			Domain:       data.Domain,
			UserId:       data.UserId,
			Nodes:        data.Nodes,
			DWifi:        data.DWifi,
			Discord:      data.Discord,
			Twitter:      data.Twitter,
			Telegram:     data.Telegram,
			Subscription: data.Subscription,
			BetaTester:   data.BetaTester,
			CreatedAt:    data.CreatedAt,
			UpdatedAt:    data.UpdatedAt,
			TotalScore:   total,
		}

		var user models.User
		err := db.Model(&models.User{}).Select("user_id, name, profile_picture_url,country, wallet_address, discord, twitter, email, chain_name").Where("user_id = ?", data.UserId).First(&user).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				response = append(response, UserScoreBoard)
				continue
			} else {
				load.Logger.Error(err.Error())
				httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").SendD(c)
				return
			}

		}

		if user.ChainName == "" || user.ChainName == "null" {
			user.ChainName = "-"
		}

		payload := models.User{UserId: user.UserId, Name: user.Name, WalletAddress: user.WalletAddress, ProfilePictureUrl: user.ProfilePictureUrl, Country: user.Country, Discord: user.Discord, Twitter: user.Twitter, Email: user.Email, ChainName: user.ChainName}

		UserScoreBoard.UserDetails = payload

		response = append(response, UserScoreBoard)

		// Sort the response by TotalScore in descending order
		sort.SliceStable(response, func(i, j int) bool {
			return response[i].TotalScore > response[j].TotalScore
		})

		// Take the top 10 entries
		if len(response) >= limit {
			response = response[:limit]
		}
	}

	httpo.NewSuccessResponseP(200, "ScoreBoard fetched successfully length := "+strconv.Itoa(len(response)), response).SendD(c)
}

func getAllUsersScoreBoard(c *gin.Context) {
	db := database.GetDb()
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
				load.Logger.Error(err.Error())
			}
		} else {
			var data = models.UserScoreBoard{
				ID:           scoreBoard.ID,
				Reviews:      scoreBoard.Reviews,
				Domain:       scoreBoard.Domain,
				UserId:       scoreBoard.UserId,
				Nodes:        scoreBoard.Nodes,
				DWifi:        scoreBoard.DWifi,
				Discord:      scoreBoard.Discord,
				Twitter:      scoreBoard.Twitter,
				Telegram:     scoreBoard.Telegram,
				Subscription: scoreBoard.Subscription,
				BetaTester:   scoreBoard.BetaTester,
				CreatedAt:    scoreBoard.CreatedAt,
				UpdatedAt:    scoreBoard.UpdatedAt,
			}
			data.UserDetails = userDetail
			response = append(response, data)

		}
	}

	httpo.NewSuccessResponseP(200, "ScoreBoard fetched successfully", response).SendD(c)
}

func UpdateLeaderBoardForAllUsers(c *gin.Context) {
	var response []interface{}

	ReviewUpdateforOldUsers()

	httpo.NewSuccessResponseP(200, "Leaderboard updated successfully", response).SendD(c)
}

func BetaGraphqlQueryForLeaderboardHandler(c *gin.Context) {

	var payload interface{}

	userids, err := BetaGraphqlQueryForLeaderboard()
	if err != nil {
		httpo.NewErrorResponse(500, err.Error()).SendD(c)

	}

	if len(userids) != 0 {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Wait()
			for _, userId := range userids {
				DynamicLeaderBoardUpdate(userId, "beta_tester")
			}
		}()
		wg.Done()
	}

	httpo.NewSuccessResponseP(200, "request successfully will be done withing few min, BetaGraphqlQueryForLeaderboardHandler", payload).SendD(c)
}
func ErebrusQueryForLeaderboardHandler(c *gin.Context) {

	var payload interface{}

	userids, err := ErebrusQueryForLeaderboard()
	if err != nil {
		httpo.NewErrorResponse(500, err.Error()).SendD(c)

	}

	if len(userids) != 0 {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Wait()
			for _, userId := range userids {
				DynamicLeaderBoardUpdate(userId, "subscription")
			}
		}()
		wg.Done()
	}

	httpo.NewSuccessResponseP(200, "request successfully will be done withing few min, ErebrusQueryForLeaderboardHandler", payload).SendD(c)
}
