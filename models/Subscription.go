package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subscription struct {
	ID        uint      `gorm:"primary_key" json:"id,omitempty"`
	UserId    string    `json:"userId,omitempty"`
	Type      string    `json:"type,omitempty"`
	StartTime time.Time `json:"startTime,omitempty"`
	EndTime   time.Time `json:"endTime,omitempty"`
}

type OrgAppClientSubscription struct {
	ID        uint      `gorm:"primary_key" json:"id,omitempty"`
	UserId    string    `json:"userId,omitempty"`
	Type      string    `json:"type,omitempty"`
	StartTime time.Time `json:"startTime,omitempty"`
	EndTime   time.Time `json:"endTime,omitempty"`
}

type Plan struct {
	ID            string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name          string         `gorm:"not null" json:"name"`
	Status        string         `gorm:"not null" json:"status"`           // e.g., active, inactive
	AllowedRegion []string       `gorm:"type:jsonb" json:"allowed_region"` // null or empty = all regions allowed
	MaxClients    int            `gorm:"not null" json:"max_clients"`
	Duration      int            `gorm:"not null" json:"duration"`    // in days
	PriceCents    int64          `gorm:"not null" json:"price_cents"` // stored in cents
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

type SubscriptionPlan struct {
	ID          string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	PlanID      string         `gorm:"type:uuid;not null" json:"plan_id"`                               // foreign key
	Plan        Plan           `gorm:"foreignKey:PlanID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // optional: preload plan
	DateCreated time.Time      `gorm:"not null" json:"date_created"`
	Status      string         `gorm:"not null" json:"status"` // e.g., active, expired
	AutoRenewal bool           `gorm:"default:false" json:"auto_renewal"`
	CreatedBy   string         `gorm:"type:uuid;not null" json:"created_by"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
type SubscriptionRenewal struct {
	ID                 string           `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedBy          string           `gorm:"type:uuid;not null" json:"created_by"`
	SubscriptionPlanID string           `gorm:"type:uuid;not null" json:"subscription_plan_id"` // 'type' field renamed for clarity
	SubscriptionPlan   SubscriptionPlan `gorm:"foreignKey:SubscriptionPlanID" json:"subscription_plan"`
	StartTime          time.Time        `gorm:"not null" json:"start_time"`
	EndTime            time.Time        `gorm:"not null" json:"end_time"`
	CreatedAt          time.Time        `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt          time.Time        `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt          gorm.DeletedAt   `gorm:"index" json:"-"`
}
type OrgSubscription struct {
	ID              uuid.UUID    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	OrganisationID  uuid.UUID    `gorm:"type:uuid;not null;index" json:"organisation_id"`
	Organisation    Organisation `gorm:"foreignKey:OrganisationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"organisation"`
	StartTime       time.Time    `gorm:"not null" json:"start_time"`
	EndTime         *time.Time   `json:"end_time"`
	BillingCycle    string       `gorm:"type:varchar(20);not null" json:"billing_cycle"` // e.g., Monthly, Quarterly
	NextBillingDate time.Time    `gorm:"not null" json:"next_billing_date"`
	AmountDue       float64      `gorm:"type:numeric(10,2);not null" json:"amount_due"`
	Status          string       `gorm:"type:varchar(20);not null" json:"status"` // Active, Cancelled, Overdue
	LastPaymentDate *time.Time   `json:"last_payment_date"`
	CreatedAt       time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
}
