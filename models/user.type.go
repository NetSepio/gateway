package models

import (
	"time"

	"github.com/google/uuid"
)

type UserActivity struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserId    string    `json:"user_id"`
	Modules   string    `json:"modules"`
	Action    string    `json:"action"`
	Metadata  string    `json:"metadata"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
