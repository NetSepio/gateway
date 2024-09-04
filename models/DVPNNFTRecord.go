package models

import (
	"time"
)

type DVPNNFTRecord struct {
	ID              uint   `gorm:"primaryKey;autoIncrement"`
	Chain           string `gorm:"not null"`
	WalletAddress   string `gorm:"not null"`
	EmailID         string
	TransactionHash string
	CreatedAt       time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
}
