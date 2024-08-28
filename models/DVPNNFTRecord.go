package models

import (
	"time"
)

type DVPNNFTRecord struct {
	ID              uint `gorm:"primaryKey;autoIncrement"`
	WalletAddress   string
	TransactionHash string
	CreatedAt       time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
}
