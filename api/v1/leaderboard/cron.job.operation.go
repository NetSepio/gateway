package leaderboard

import (
	"fmt"
	"log"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/NetSepio/gateway/config/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CronJobLeaderBoardUpdate(column_name string, leaderboard Leaderboard) {
	// Database connection setup (replace with your actual connection details)
	db := dbconfig.GetDb()

	// // Check if the user_id exists in the LeaderBoard table
	// var leaderboard ScoreBoard

	var scoreBoard ScoreBoard

	data, err := GetActivityUnitXpByActivity(column_name)
	if err != nil {
		log.Printf("failed to get the ScoreBoard by ID: %v", err)
	}
	// leaderboard.UserId = user_id
	err = db.Debug().Where("user_id = ?", leaderboard.UserId).First(&scoreBoard).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {

			CreateScoreBoard(ScoreBoard{
				ID:        uuid.New().String(),
				Reviews:   leaderboard.Reviews,
				Domain:    leaderboard.Domain,
				UserId:    leaderboard.UserId,
				Nodes:     leaderboard.Nodes,
				DWifi:     leaderboard.DWifi,
				Discord:   leaderboard.Discord,
				Twitter:   leaderboard.Twitter,
				Telegram:  leaderboard.Telegram,
				CreatedAt: leaderboard.CreatedAt,
				UpdatedAt: leaderboard.UpdatedAt,
			})
			log.Println("New record inserted and " + column_name + " count initialized successfully!")
			return
		}
		log.Printf("failed to query the ScoreBoard: %v", err)
	} else {

		err = UpdateScoreBoard(leaderboard.ID, ScoreBoard{
			ID:        uuid.New().String(),
			Reviews:   leaderboard.Reviews,
			Domain:    leaderboard.Domain,
			UserId:    leaderboard.UserId,
			Nodes:     leaderboard.Nodes,
			DWifi:     leaderboard.DWifi,
			Discord:   leaderboard.Discord,
			Twitter:   leaderboard.Twitter,
			Telegram:  leaderboard.Telegram,
			CreatedAt: leaderboard.CreatedAt,
			UpdatedAt: leaderboard.UpdatedAt,
		}, column_name, data.XP)
		if err != nil {
			log.Printf("failed to update the Reviews count: %v", err)
			return
		}
		log.Println(column_name + " count incremented successfully!")
		return
	}
}
func CronForReviewUpdate() {
	db := dbconfig.GetDb()

	var voters []string
	db.Model(&models.Review{}).Select("voter").Find(&voters)

	// fmt.Println("len voters : ", len(voters))
	// fmt.Println("voters : ", voters)

	if len(voters) > 0 {

		for _, v := range voters {
			var userIds []string
			db := dbconfig.GetDb()
			// Select only UserId column from the Leaderboard table
			if err := db.Raw("SELECT user_id FROM users WHERE wallet_address = ?", v).Scan(&userIds).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					fmt.Println("This user does not exist in the user table: wallet address =", v)
				} else {
					log.Printf("Failed to get the Reviews: %v\n", err)
				}
			} else {
				if len(userIds) > 0 {
					for _, id := range userIds {
						go DynamicLeaderBoardUpdate(id, "reviews")
					}
				}
			}
		}

	}

}
