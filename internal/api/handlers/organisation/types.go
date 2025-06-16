package organisation

import (
	"time"

	"github.com/google/uuid"
)

// Organisation model
type Organisation struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(225);not null" json:"name"`
	IPAddress string    `gorm:"type:varchar(225);not null" json:"ip_address" binding:"required"`
	APIKey    string    `gorm:"not null" json:"api_key"`
	Status    string    `gorm:"type:varchar(50);default:'inactive'" json:"status"` // e.g., active, inactive, suspended
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// CreateOrganisationInput defines input for POST
type CreateOrganisationInput struct {
	Name      string `json:"name" binding:"required"`
	IPAddress string `json:"ip_address" binding:"required"`
	APIKey    string `json:"api_key" binding:"required"`
}

type OrganisationPaseto struct {
	OrganisationId string
	PasetoToken    string
}
