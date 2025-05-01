package aptos

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CoinGeckoResponse struct {
	Aptos struct {
		USD float64 `json:"usd"`
	} `json:"aptos"`
}

func GetCoinPrice() (float64, error) {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=aptos&vs_currencies=usd"
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("error fetching data: %w", err)
	}
	defer resp.Body.Close()

	var result CoinGeckoResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("error decoding JSON: %w", err)
	}

	return result.Aptos.USD, nil
}
