package graphql

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Struct for GraphQL response parsing
type TokenData struct {
	TokenName              string `json:"token_name"`
	Description            string `json:"description"`
	CurrentTokenOwnerships []struct {
		OwnerAddress             string `json:"owner_address"`
		LastTransactionTimestamp string `json:"last_transaction_timestamp"`
	} `json:"current_token_ownerships"`
}

type GraphQLResponse struct {
	Data struct {
		CurrentTokenDatasV2 []TokenData `json:"current_token_datas_v2"`
	} `json:"data"`
}

// Function to perform GraphQL request
func leaderboardQuery(url string, query string, variables map[string]interface{}) (*GraphQLResponse, error) {
	payload := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result GraphQLResponse
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
