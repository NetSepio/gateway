package org_subscription

import (
	"time"

	"github.com/google/uuid"
)

// Organisation model
type Organisation struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(225);not null" json:"name"`
	IPAddress   string    `gorm:"type:varchar(225);not null" json:"ip_address" binding:"required"`
	APIKey      string    `gorm:"not null" json:"api_key"`
	Status      string    `gorm:"type:varchar(50);default:'inactive'" json:"status"` // e.g., active, inactive, suspended
	OrgMetaData string    `gorm:"type:varchar(50)" json:"org_meta_data"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// CreateOrganisationInput defines input for POST
type CreateOrganisationInput struct {
	Name        string `json:"name" binding:"required"`
	IPAddress   string `json:"ip_address" binding:"required"`
	APIKey      string `json:"api_key" binding:"required"`
	OrgMetaData string `json:"org_meta_data"`
}

type OrganisationPaseto struct {
	OrganisationId string `json:"organisationId"`
	Token          string `json:"token"`
}
type OrgAppSubscriptionPayload struct {
	ID              uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	OrganisationID  uuid.UUID  `gorm:"type:uuid;not null;index" json:"organisation_id"`
	StartTime       time.Time  `gorm:"not null" json:"start_time"`
	EndTime         *time.Time `json:"end_time"`
	BillingCycle    string     `gorm:"type:varchar(20);not null" json:"billing_cycle"` // e.g., Monthly, Quarterly
	NextBillingDate time.Time  `gorm:"not null" json:"next_billing_date"`
	AmountDue       float64    `gorm:"type:numeric(10,2);not null" json:"amount_due"`
	Status          string     `gorm:"type:varchar(20);not null" json:"status"` // Active, Cancelled, Overdue
	LastPaymentDate *time.Time `json:"last_payment_date"`
	CreatedAt       time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
