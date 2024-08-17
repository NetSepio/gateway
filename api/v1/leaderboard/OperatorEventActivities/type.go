package OperatorEventActivities

import (
	"time"
)

type OperatorEventActivities struct { // mapped leaderBoard
	Reviews    int       // 1. Reviews
	Domain     int       // 2. Domain (all projects)
	UserId     string    // 3. User
	Nodes      int       // 4. Nodes
	DWifi      int       // 5. DWifi
	Discord    int       // 6. Discord
	Twitter    int       // 7. Twitter
	Telegram   int       // 8. Telegram
	CreatedAt  time.Time // 10. CreatedAt
	UpdatedAt  time.Time // 11. UpdatedAt
	BetaTest   int
	ErebrusNFT int
}

type NodeOperators struct {
	Id           string
	UserId       string
	OpeartorType string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
