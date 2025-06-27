package orgApp

import (
	"time"

	"github.com/google/uuid"
)

type OrganisationApp struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	OrganisationId uuid.UUID `gorm:"type:uuid;not null" json:"organisation_id"`
	Name           string    `gorm:"type:varchar(225);not null" json:"name"`
	Description    string    `gorm:"type:text" json:"description"`
	APIKey         string    `gorm:"not null" json:"api_key"`
	Status         string    `gorm:"type:varchar(50);default:'inactive'" json:"status"` // e.g., active, inactive, suspended
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type OrganisationAppPaseto struct {
	OrganisationId string `json:"organisationId"`
	Token          string `json:"token"`
}
