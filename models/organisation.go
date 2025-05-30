package models

import (
	"time"

	"github.com/google/uuid"
)

type Organisation struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(225);not null" json:"name"`
	IPAddress string    `gorm:"type:varchar(225);not null" json:"ip_address"`
	APIKey    string    `gorm:"not null" json:"api_key"`
	Status    string    `gorm:"type:varchar(50);default:'inactive'" json:"status"` // e.g., active, inactive, suspended
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
