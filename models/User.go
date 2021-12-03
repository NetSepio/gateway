package models

type User struct {
	Name              string
	WalletAddress     string   `gorm:"primary_key"`
	FlowIds           []FlowId `gorm:"foreignkey:WalletAddress"`
	ProfilePictureUrl string
	Country           string
}
