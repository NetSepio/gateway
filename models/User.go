package models

import "github.com/lib/pq"

type User struct {
	Name              string
	WalletAddress     string   `gorm:"primary_key"`
	FlowIds           []FlowId `gorm:"foreignkey:WalletAddress"`
	ProfilePictureUrl string
	Country           string
	Roles             pq.Int32Array `gorm:"type:int[]"`
}
