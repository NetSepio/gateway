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
