package leaderboard

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/NetSepio/gateway/internal/database"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/utils/pkg/graphql"
)

func BetaGraphqlQueryForLeaderboard() ([]string, error) {
	// Database connection setup (replace with your actual connection details)
	db := database.GetDb()

	var userIds []string

	// Call Beta Test Query
	betaTestResponse, err := graphql.BetaTestQuery()
	if err != nil {
		log.Printf("Error performing Beta Test query: %v", err)
		return nil, err // Return the error if the query fails
	}

	fmt.Println("Beta Test Query Results:")
	for _, token := range betaTestResponse.Data.CurrentTokenDatasV2 {
		fmt.Printf("Token Name: %s, Description: %s\n", token.TokenName, token.Description)
		for _, ownership := range token.CurrentTokenOwnerships {
			fmt.Printf("Owner Address: %s, Last Transaction: %s\n", ownership.OwnerAddress, ownership.LastTransactionTimestamp)

			// Check if the user exists or create a new one if not
			var user models.User
			result := db.Where("wallet_address = ?", ownership.OwnerAddress).First(&user)
			if result.Error != nil {
				// If the user does not exist, create a new one
				if result.Error == gorm.ErrRecordNotFound {
					// Generate a new UUID for the user
					user = models.User{
						UserId:        uuid.New().String(), // UUID generation
						WalletAddress: &ownership.OwnerAddress,
					}
					if err := db.Create(&user).Error; err != nil {
						log.Printf("Error creating new user: %v", err)
						return userIds, err // Return if error occurs while creating user
					} else {
						userIds = append(userIds, user.UserId)
						// leaderboard.DynamicLeaderBoardUpdate(user.UserId, "reviews")

					}
					fmt.Printf("Created new user with UserID: %s\n", user.UserId)
				} else {
					// If another error occurred, log and return it
					log.Printf("Error querying user: %v", result.Error)
					return nil, result.Error
				}
			} else {
				// leaderboard.DynamicLeaderBoardUpdate(user.UserId, "reviews")
				userIds = append(userIds, user.UserId)
			}
		}
	}
	return userIds, nil
}
func ErebrusQueryForLeaderboard() ([]string, error) {
	// Database connection setup (replace with your actual connection details)
	db := database.GetDb()

	var userIds []string

	// Call Beta Test Query
	betaTestResponse, err := graphql.ErebrusQuery()
	if err != nil {
		log.Printf("Error performing Beta Test query: %v", err)
		return nil, err // Return the error if the query fails
	}

	fmt.Println("Beta Test Query Results:")
	for _, token := range betaTestResponse.Data.CurrentTokenDatasV2 {
		fmt.Printf("Token Name: %s, Description: %s\n", token.TokenName, token.Description)
		for _, ownership := range token.CurrentTokenOwnerships {
			fmt.Printf("Owner Address: %s, Last Transaction: %s\n", ownership.OwnerAddress, ownership.LastTransactionTimestamp)

			// Check if the user exists or create a new one if not
			var user models.User
			result := db.Where("wallet_address = ?", ownership.OwnerAddress).First(&user)
			if result.Error != nil {
				// If the user does not exist, create a new one
				if result.Error == gorm.ErrRecordNotFound {
					// Generate a new UUID for the user
					user = models.User{
						UserId:        uuid.New().String(), // UUID generation
						WalletAddress: &ownership.OwnerAddress,
					}
					if err := db.Create(&user).Error; err != nil {
						log.Printf("Error creating new user: %v", err)
						return userIds, err // Return if error occurs while creating user
					} else {
						userIds = append(userIds, user.UserId)
						// leaderboard.DynamicLeaderBoardUpdate(user.UserId, "reviews")

					}
					fmt.Printf("Created new user with UserID: %s\n", user.UserId)
				} else {
					// If another error occurred, log and return it
					log.Printf("Error querying user: %v", result.Error)
					return nil, result.Error
				}
			} else {
				// leaderboard.DynamicLeaderBoardUpdate(user.UserId, "reviews")
				userIds = append(userIds, user.UserId)
			}
		}
	}
	return userIds, nil
}
