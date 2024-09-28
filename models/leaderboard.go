package models

import (
	"time"

	"github.com/google/uuid"
)

type Leaderboard struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Reviews   int       `gorm:"not null" json:"reviews"`
	Domain    int       `gorm:"not null" json:"domain"`
	UserId    uuid.UUID `gorm:"type:uuid;not null" json:"userId"`
	Nodes     int       `gorm:"not null" json:"nodes"`
	DWifi     int       `gorm:"not null" json:"dwifi"`
	Discord   int       `gorm:"not null" json:"discord"`
	Twitter   int       `gorm:"not null" json:"twitter"`
	Telegram  int       `gorm:"not null" json:"telegram"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}

func (Leaderboard) TableName() string {
	return "leaderboards"
}

type ScoreBoard struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Reviews   int
	Domain    int
	UserId    string `gorm:"type:uuid;not null"`
	Nodes     int
	DWifi     int
	Discord   int
	Twitter   int
	Telegram  int
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
