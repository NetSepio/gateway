package models

type WaitList struct {
	EmailId       string `gorm:"primary_key"`
	WalletAddress string 
	Twitter       string 
}
