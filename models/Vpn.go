package models

type Sotreus struct {
	Name          string `gorm:"primary_key" json:"name"`
	WalletAddress string `json:"walletAddress"`
	Region        string `json:"region"`
}

type Erebrus struct {
	UUID          string `gorm:"primary_key" json:"UUID"`
	Name          string `json:"name"`
	WalletAddress string `json:"walletAddress"`
	Region        string `json:"region"`
}
