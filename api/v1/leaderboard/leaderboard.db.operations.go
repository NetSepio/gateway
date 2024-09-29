package leaderboard

import (
	"fmt"
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
	leaderboard.UserId = user_id
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
				return
			}
			newLeaderBoard.ID = uuid.New().String()
			// Initialize the specific column with 1 (assuming it's an integer field)
			err = db.Debug().Create(&newLeaderBoard).Error
			if err != nil {
				log.Println("[ ERROR ] failed to insert new record:", err)
			} else {

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

	data, err := GetActivityUnitXpByActivity(column_name)
	if err != nil {
		log.Printf("failed to get the ScoreBoard by ID: %v", err)
	}
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

	log.Println("Reviews count incremented successfully!")
}

// Function to create a new ScoreBoard entry
func CreateScoreBoard(score ScoreBoard) error {
	db := dbconfig.GetDb()
	result := db.Create(&score)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Println("ScoreBoard record not found")
			return gorm.ErrRecordNotFound
		} else {
			log.Printf("failed to create a new ScoreBoard entry: %v", result.Error)
			return result.Error
		}
	}
	fmt.Println("New scoreboard created with ID:", score.ID)
	return nil
}

// Function to fetch all ScoreBoard records
func GetAllScoreBoards() ([]ScoreBoard, error) {
	var scoreboards []ScoreBoard
	db := dbconfig.GetDb()

	result := db.Find(&scoreboards)
	if result.Error != nil {
		return nil, result.Error
	}
	return scoreboards, nil
}

// Function to fetch a ScoreBoard by ID
func GetScoreBoardByID(id string) (*ScoreBoard, error) {
	var score ScoreBoard
	db := dbconfig.GetDb()

	result := db.First(&score, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &score, nil
}

// Function to update an existing ScoreBoard
func UpdateScoreBoard(id string, updatedScore ScoreBoard, column_name string, value int) error {
	var score ScoreBoard
	db := dbconfig.GetDb()

	result := db.First(&score, "user_id = ?", updatedScore.UserId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			switch column_name {
			case "reviews":
				score.Reviews = updatedScore.Reviews * value
			case "domain":
				score.Domain = updatedScore.Domain * value
			case "nodes":
				score.Nodes = updatedScore.Nodes * value
			case "d_wifi":
				score.DWifi = updatedScore.DWifi * value
			case "discord":
				score.Discord = updatedScore.Discord * value
			case "twitter":
				score.Twitter = updatedScore.Twitter * value
			case "telegram":
				score.Telegram = updatedScore.Telegram * value
			default:
				log.Printf("Invalid column name: %s", column_name)
				// return
			}

			score.ID = uuid.New().String()
			score.UserId = updatedScore.UserId
			// Initialize the specific column with 1 (assuming it's an integer field)
			err := db.Debug().Create(&score).Error
			if err != nil {
				log.Println("[ ERROR ] failed to insert new record:", err)
				return err
			}
			return nil
		} else {
			return result.Error
		}
	}

	var columnValue int

	switch column_name {
	case "reviews":
		columnValue = updatedScore.Reviews * value
	case "domain":
		columnValue = updatedScore.Domain * value
	case "nodes":
		columnValue = updatedScore.Nodes * value
	case "d_wifi":
		columnValue = updatedScore.DWifi * value
	case "discord":
		columnValue = updatedScore.Discord * value
	case "twitter":
		columnValue = updatedScore.Twitter * value
	case "telegram":
		columnValue = updatedScore.Telegram * value
	default:
		log.Printf("Invalid column name: %s", column_name)
		// return
	}

	// If user_id exists, increment the Reviews column by 1
	err := db.Debug().Model(&score).Update(column_name, columnValue).Error
	if err != nil {
		log.Printf("failed to update the ScoreBoard count: %v", err)
	}

	fmt.Println("ScoreBoard updated:", score.ID)
	return nil
}

// GetAllActivityUnitXp retrieves all records from the ActivityUnitXp table.
func GetAllActivityUnitXp() ([]ActivityUnitXp, error) {
	var activities []ActivityUnitXp
	db := dbconfig.GetDb()

	if err := db.Find(&activities).Error; err != nil {
		return nil, fmt.Errorf("error retrieving all activity unit xp records: %v", err)
	}
	return activities, nil
}

func GetActivityUnitXpByActivity(activity string) (*ActivityUnitXp, error) {
	var activityUnitXp ActivityUnitXp
	db := dbconfig.GetDb()
	if err := db.Where("activity = ?", activity).First(&activityUnitXp).Error; err != nil {
		return nil, fmt.Errorf("error retrieving activity unit xp: %v", err)
	}
	return &activityUnitXp, nil
}

func UpdateActivityUnitXp(activity string, xp int) error {
	db := dbconfig.GetDb()
	if err := db.Model(&ActivityUnitXp{}).Where("activity = ?", activity).Update("xp", xp).Error; err != nil {
		return fmt.Errorf("error updating activity unit xp: %v", err)
	}
	return nil
}

func DeleteActivityUnitXp(activity string) error {
	db := dbconfig.GetDb()
	if err := db.Where("activity = ?", activity).Delete(&ActivityUnitXp{}).Error; err != nil {
		return fmt.Errorf("error deleting activity unit xp: %v", err)
	}
	return nil
}
