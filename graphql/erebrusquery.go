package graphql

// Function for APTOS Erebrus URL query
func ErebrusQuery() (*GraphQLResponse, error) {
	url := "https://indexer.mainnet.aptoslabs.com/v1/graphql"
	query := `
		query MyQuery {
			current_token_datas_v2(
				where: {collection_id: {_eq: "0x465ebc4eb9718e1555976796a4456fa1a2df8126b4e01ff5df7f9d14fb3eba19"}}
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

	return leaderboardQuery(url, query, nil)
}
