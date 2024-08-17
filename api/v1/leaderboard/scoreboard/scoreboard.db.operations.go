package scoreboard

import (
	"fmt"
	"reflect"

	leaderboard "github.com/NetSepio/gateway/api/v1/Leaderboard"
	"github.com/NetSepio/gateway/config/dbconfig"
	"gorm.io/gorm"
)

type Automation interface {
	Execute() error
}

func OperatorEventActivitiesCalculator(userid, everType string) {

}

/* func DynamicScoreBoardUpdate(user_id, column_name string) error {
	// Database connection setup (replace with your actual connection details)
	db := dbconfig.GetDb()

	// Check if the user_id exists in the OperatorEventActivities table
	var (
		scoreBoard     ScoreBoard
		updating_value int
	)
	err := db.Debug().Where("user_id = ?", user_id).First(&scoreBoard).Error

	// If user_id does not exist, insert a new record with the initial review count
	newScoreBoard := ScoreBoard{
		UserId: user_id,
	}

	if err != nil {
		if err == gorm.ErrRecordNotFound {

			// Use switch to dynamically set the specified column value
			switch column_name {
			case "reviews":
				newScoreBoard.Reviews = leaderboard.Reviews
			case "domain":
				newScoreBoard.Domain = leaderboard.Domain
			case "nodes":
				newScoreBoard.Nodes = leaderboard.Nodes
			case "d_wifi":
				newScoreBoard.DWifi = leaderboard.DWifi
			case "discord":
				newScoreBoard.Discord = leaderboard.Discord
			case "twitter":
				newScoreBoard.Twitter = leaderboard.Twitter
			case "telegram":
				newScoreBoard.Telegram = leaderboard.Telegram
			case "beta_test":
				newScoreBoard.BetaTest = leaderboard.BetaTest
			case "erebrus_NFT":
				newScoreBoard.ErebrusNFT = leaderboard.ErebrusNFT
			default:
				log.Printf("Invalid column name")
				return fmt.Errorf("invalid column name : %v", column_name)
			}

			// Initialize the specific column with 1 (assuming it's an integer field)
			// IF ITS NEW USER THEN MAKE A NEW ENTRY OF THAT USER IN NEW scoreboard TABLE

			err = db.Debug().Create(&newScoreBoard).Error
			if err != nil {
				log.Fatal("failed to insert new record:", err)
			}
			log.Println("New record inserted and reviews count initialized successfully!")
			return nil
		}
		log.Printf("failed to query the OperatorEventActivities: %v", err)
	} else {
		switch column_name {
		case "reviews":
			updating_value = leaderboard.Reviews
		case "domain":
			updating_value = leaderboard.Domain
		case "nodes":
			updating_value = leaderboard.Nodes
		case "d_wifi":
			updating_value = leaderboard.DWifi
		case "discord":
			updating_value = leaderboard.Discord
		case "twitter":
			updating_value = leaderboard.Twitter
		case "telegram":
			updating_value = leaderboard.Telegram
		case "beta_test":
			updating_value = leaderboard.BetaTest
		case "erebrus_NFT":
			updating_value = leaderboard.ErebrusNFT
		default:
			log.Printf("Invalid column name")
			return fmt.Errorf("invalid column name : %v", column_name)
		}
	}

	// If user_id exists, increment the Reviews column by 1
	err = db.Debug().Model(&scoreBoard).Update(column_name, gorm.Expr(column_name+" + ?", updating_value)).Error
	if err != nil {
		log.Printf("failed to update the Reviews count: %v", err)
	}

	log.Println("Reviews count incremented successfully!")
	return nil
} */

func DynamicScoreBoardUpdate(user_id, column_name string) error {
	db := dbconfig.GetDb()

	// Database connection setup (replace with your actual connection details)

	// Validate column name
	validColumns := map[string]int{
		"reviews":     leaderboard.Reviews,
		"domain":      leaderboard.Domain,
		"nodes":       leaderboard.Nodes,
		"d_wifi":      leaderboard.DWifi,
		"discord":     leaderboard.Discord,
		"twitter":     leaderboard.Twitter,
		"telegram":    leaderboard.Telegram,
		"beta_test":   leaderboard.BetaTest,
		"erebrus_nft": leaderboard.ErebrusNFT,
	}

	value, valid := validColumns[column_name]
	if !valid {
		return fmt.Errorf("invalid column name: %v", column_name)
	}

	// Try to update existing record
	result := db.Model(&ScoreBoard{}).
		Where("user_id = ?", user_id).
		Update(column_name, gorm.Expr(column_name+" + ?", value))

	if result.Error != nil {
		return fmt.Errorf("failed to update score: %v", result.Error)
	}

	// If no rows were affected, create a new record
	if result.RowsAffected == 0 {
		newScoreBoard := ScoreBoard{UserId: user_id}
		reflect.ValueOf(&newScoreBoard).Elem().FieldByName(column_name).SetInt(int64(value))
		if err := db.Create(&newScoreBoard).Error; err != nil {
			return fmt.Errorf("failed to create new score record: %v", err)
		}

		return nil
	}
	return nil
}

// Get top 10 ScoreBoards with highest reviews
func GetTop10ScoreBoards(scoreBoardType string, limits int) ([]ScoreBoard, error) {
	db := dbconfig.GetDb()
	var scoreBoards []ScoreBoard
	result := db.Order(scoreBoardType + " DESC").Limit(limits).Find(&scoreBoards)
	if result.Error != nil {
		return nil, result.Error
	}
	return scoreBoards, nil
}
