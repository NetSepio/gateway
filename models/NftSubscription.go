package models

import (
	"time"

	"gorm.io/gorm"
)

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
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
