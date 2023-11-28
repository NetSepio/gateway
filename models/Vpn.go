package models

type Sotreus struct {
	Name          string `json:"name,omitempty"`
	WalletAddress string `json:"walletAddress,omitempty"`
	Region        string `json:"region,omitempty"`
}
