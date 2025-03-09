package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Referral struct {
	Id           string    `json:"id" gorm:"type:uuid;primaryKey"`
	ReferrerId   string    `json:"referrerId" gorm:"type:uuid;not null"`       // User who referred
	RefereeId    string    `json:"refereeId" gorm:"type:uuid;not null;unique"` // User who was referred
	ReferralCode string    `json:"referralCode" gorm:"type:varchar(255);unique;not null"`
	CreatedAt    time.Time `json:"createdAt"`
}

type ReferralEarnings struct {
	Id           string    `json:"id" gorm:"type:uuid;primaryKey"`
	ReferrerId   string    `json:"referrerId" gorm:"type:uuid;not null"`
	RefereeId    string    `json:"refereeId" gorm:"type:uuid;not null"`
	AmountEarned float64   `json:"amountEarned" gorm:"not null"`
	CreatedAt    time.Time `json:"createdAt"`
}

// Auto-generate UUID before creating the record
func (r *Referral) BeforeCreate(tx *gorm.DB) error {
	r.Id = uuid.New().String()
	return nil
}

func (e *ReferralEarnings) BeforeCreate(tx *gorm.DB) error {
	e.Id = uuid.New().String()
	return nil
}
