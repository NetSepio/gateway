package graphql

// Function for Beta Test URL query
func BetaTestQuery() (*GraphQLResponse, error) {
	url := "https://indexer-testnet.staging.gcp.aptosdev.com/v1/graphql"
	query := `
		query MyQuery($_lt: timestamp = "2024-01-16T00:00:00.000000") {
			current_token_datas_v2(
				where: {
					collection_id: {_eq: "0x212ee7ca88024f75e20c79dfee04898048fb9de15cb2da27d793151a6d58db25"}, 
					current_token_ownerships: {last_transaction_timestamp: {_lt: $_lt}}
				}
			) {
				token_name
				description
				current_token_ownerships {
					owner_address
					last_transaction_timestamp
				}
			}
		}
	`
	variables := map[string]interface{}{
		"_lt": "2024-01-16T00:00:00.000000",
	}

	return leaderboardQuery(url, query, variables)
}
