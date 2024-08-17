package OperatorEventActivities

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/NetSepio/gateway/config/dbconfig"
	"gorm.io/gorm"
)

const (
	testGraphQLAptos = "https://indexer-testnet.staging.gcp.aptosdev.com/v1/graphql"
	mainGraphQLAptos = "https://indexer.mainnet.aptoslabs.com/v1/graphql"
)

type GraphQLResponse struct {
	Data struct {
		CurrentTokenDatasV2 []struct {
			CurrentTokenOwnerships []struct {
				OwnerAddress string `json:"owner_address"`
			} `json:"current_token_ownerships"`
		} `json:"current_token_datas_v2"`
	} `json:"data"`
}

func fetchAddressesFromIndexer(url, query string) ([]string, error) {
	body := map[string]string{
		"query": query,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var graphQLResp GraphQLResponse
	if err := json.NewDecoder(resp.Body).Decode(&graphQLResp); err != nil {
		return nil, err
	}

	var addresses []string
	for _, data := range graphQLResp.Data.CurrentTokenDatasV2 {
		if len(data.CurrentTokenOwnerships) > 0 {
			addresses = append(addresses, data.CurrentTokenOwnerships[0].OwnerAddress)
		}
	}

	return addresses, nil
}

func UpdateLeaderboardFromIndexer() error {
	db := dbconfig.GetDb()

	betaTestQuery := `
    query MyQuery($_lt: timestamp = "2024-01-16T00:00:00.000000") {
        current_token_datas_v2(
            where: {collection_id: {_eq: "0x212ee7ca88024f75e20c79dfee04898048fb9de15cb2da27d793151a6d58db25"}, 
            current_token_ownerships: {last_transaction_timestamp: {_lt: $_lt}}}
        ) {
            current_token_ownerships {
                owner_address
            }
        }
    }
    `

	erebrusNFTQuery := `
    query MyQuery {
        current_token_datas_v2(
            where: {collection_id: {_eq: "0x465ebc4eb9718e1555976796a4456fa1a2df8126b4e01ff5df7f9d14fb3eba19"}}
        ) {
            current_token_ownerships {
                owner_address
            }
        }
    }
    `

	betaTestAddresses, err := fetchAddressesFromIndexer(testGraphQLAptos, betaTestQuery)
	if err != nil {
		return err
	}

	erebrusNFTAddresses, err := fetchAddressesFromIndexer(mainGraphQLAptos, erebrusNFTQuery)
	if err != nil {
		return err
	}

	return db.Transaction(func(tx *gorm.DB) error {
		for _, address := range betaTestAddresses {
			if err := updateOrCreateLeaderboard(tx, address, 1, 0); err != nil {
				return err
			}
		}

		for _, address := range erebrusNFTAddresses {
			if err := updateOrCreateLeaderboard(tx, address, 0, 1); err != nil {
				return err
			}
		}

		return nil
	})
}

func updateOrCreateLeaderboard(tx *gorm.DB, walletAddress string, betaTest, erebrusNFT int) error {
	var leaderboard OperatorEventActivities
	if err := tx.Where("wallet_address = ?", walletAddress).First(&leaderboard).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			newLeaderboard := OperatorEventActivities{
				// WalletAddress: walletAddress,
				BetaTest:   betaTest,
				ErebrusNFT: erebrusNFT,
			}
			return tx.Create(&newLeaderboard).Error
		}
		return err
	}

	updates := map[string]interface{}{
		"beta_test":   betaTest,
		"erebrus_nft": erebrusNFT,
	}
	return tx.Model(&leaderboard).Updates(updates).Error
}
