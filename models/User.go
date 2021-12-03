package models

import "github.com/lib/pq"

type User struct {
	Name              string
	WalletAddress     string         `gorm:"unique"`
	FlowId            pq.StringArray `gorm:"type:text[]"`
	ProfilePictureUrl string
	Country           string
}
