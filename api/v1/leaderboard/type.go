package leaderboard

import (
	"time"
)

type Leaderboard struct {
	ID        string    `gorm:"type:uuid;primary_key"` // 9. Id (Primary Key)
	Reviews   int       `gorm:"not null"`              // 1. Reviews
	Domain    int       `gorm:"not null"`              // 2. Domain (all projects)
	UserId    string    `gorm:"type:uuid;not null"`    // 3. User
	Nodes     int       `gorm:"not null"`              // 4. Nodes
	DWifi     int       `gorm:"not null"`              // 5. DWifi
	Discord   int       `gorm:"not null"`              // 6. Discord
	Twitter   int       `gorm:"not null"`              // 7. Twitter
	Telegram  int       `gorm:"not null"`              // 8. Telegram
	CreatedAt time.Time `gorm:"autoCreateTime"`        // 10. CreatedAt
	UpdatedAt time.Time `gorm:"autoUpdateTime"`        // 11. UpdatedAt
}

type ActivityUnitXp struct {
	Activity string `gorm:"not null;unique"` // Name of the activity (e.g., Reviews, Domain, etc.)
	XP       int    `gorm:"not null"`
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
