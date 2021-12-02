package models

import "github.com/lib/pq"

type User struct {
	WalletAddress string
	FlowId        pq.StringArray `gorm:"type:text[]"`
}
