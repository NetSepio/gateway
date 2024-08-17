package OperatorEventActivities

import (
	"fmt"
	"log"

	"github.com/NetSepio/gateway/config/dbconfig"
	"gorm.io/gorm"
)

type Automation interface {
	Execute() error
}

func OperatorEventActivitiesCalculator(userid, everType string) {
}

func DynamicOperatorEventActivitiesUpdate(user_id, column_name string) error {
	// Database connection setup (replace with your actual connection details)
	db := dbconfig.GetDb()

	// Check if the user_id exists in the OperatorEventActivities table
	var operatorEventActivities OperatorEventActivities
	err := db.Debug().Where("user_id = ?", user_id).First(&operatorEventActivities).Error

	// If user_id does not exist, insert a new record with the initial review count
	newOperatorEventActivities := OperatorEventActivities{
		UserId: user_id,
	}

	if err != nil {
		if err == gorm.ErrRecordNotFound {

			// Use switch to dynamically set the specified column value
			switch column_name {
			case "reviews":
				newOperatorEventActivities.Reviews = 1
			case "domain":
				newOperatorEventActivities.Domain = 1
			case "nodes":
				newOperatorEventActivities.Nodes = 1
			case "d_wifi":
				newOperatorEventActivities.DWifi = 1
			case "discord":
				newOperatorEventActivities.Discord = 1
			case "twitter":
				newOperatorEventActivities.Twitter = 1
			case "telegram":
				newOperatorEventActivities.Telegram = 1
			case "beta_test":
				newOperatorEventActivities.BetaTest = 1
			case "erebrus_NFT":
				newOperatorEventActivities.ErebrusNFT = 1
			default:
				log.Printf("Invalid column name")
				return fmt.Errorf("invalid column name : %v", column_name)
			}
			// newOperatorEventActivities.ID = uuid.New().String()

			// Initialize the specific column with 1 (assuming it's an integer field)
			// IF ITS NEW USER THEN MAKE A NEW ENTRY OF THAT USER IN NEW OperatorEventActivities TABLE
			err = db.Debug().Create(&newOperatorEventActivities).Error
			if err != nil {
				log.Fatal("failed to insert new record:", err)
			}
			log.Println("New record inserted and reviews count initialized successfully!")
			return nil
		}
		log.Printf("failed to query the OperatorEventActivities: %v", err)
	}

	// If user_id exists, increment the Reviews column by 1
	err = db.Debug().Model(&operatorEventActivities).Update(column_name, gorm.Expr(column_name+" + ?", 1)).Error
	if err != nil {
		log.Printf("failed to update the Reviews count: %v", err)
	}

	log.Println("Reviews count incremented successfully!")
	return nil
}
