package models

type NftSubscription struct {
	ID              uint   `gorm:"primaryKey"`
	UserID          string `gorm:"index"`
	ContractAddress string
	ChainName       string
	Name            string
	Symbol          string
	TotalSupply     string
	Owner           string
	TokenURI        string
}
