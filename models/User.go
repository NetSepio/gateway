package models

import "github.com/lib/pq"

type User struct {
	Name              string         `json:"name,omitempty"`
	WalletAddress     string         `gorm:"primary_key" json:"walletAddress"`
	FlowIds           []FlowId       `gorm:"foreignkey:WalletAddress" json:"-"`
	ProfilePictureUrl string         `json:"profilePictureUrl,omitempty"`
	Country           string         `json:"country,omitempty"`
	Feedbacks         pq.StringArray `json:"-" gorm:"type:text[]"`
}
