package leaderboard

import (
	"log"

	"github.com/NetSepio/gateway/config/dbconfig"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DynamicLeaderBoardUpdate(user_id, column_name string) {
	// Database connection setup (replace with your actual connection details)
	db := dbconfig.GetDb()

	// Check if the user_id exists in the LeaderBoard table
	var leaderboard Leaderboard
	err := db.Debug().Where("user_id = ?", user_id).First(&leaderboard).Error

	// If user_id does not exist, insert a new record with the initial review count
	newLeaderBoard := Leaderboard{
		UserId: user_id,
	}

	if err != nil {
		if err == gorm.ErrRecordNotFound {

			// Use reflection to dynamically set the specified column value
			switch column_name {
			case "reviews":
				newLeaderBoard.Reviews = 1
			case "domain":
				newLeaderBoard.Domain = 1
			case "nodes":
				newLeaderBoard.Nodes = 1
			case "d_wifi":
				newLeaderBoard.DWifi = 1
			case "discord":
				newLeaderBoard.Discord = 1
			case "twitter":
				newLeaderBoard.Twitter = 1
			case "telegram":
				newLeaderBoard.Telegram = 1
			default:
				log.Printf("Invalid column name")
			}
			newLeaderBoard.ID = uuid.New().String()
			// Initialize the specific column with 1 (assuming it's an integer field)
			err = db.Debug().Create(&newLeaderBoard).Error
			if err != nil {
				log.Fatal("failed to insert new record:", err)
			}
			log.Println("New record inserted and reviews count initialized successfully!")
			return
		}
		log.Printf("failed to query the LeaderBoard: %v", err)
	}

	// If user_id exists, increment the Reviews column by 1
	err = db.Debug().Model(&leaderboard).Update(column_name, gorm.Expr(column_name+" + ?", 1)).Error
	if err != nil {
		log.Printf("failed to update the Reviews count: %v", err)
	}

	log.Println("Reviews count incremented successfully!")
}
func InsertDataInScoreBoard() {
	

}
