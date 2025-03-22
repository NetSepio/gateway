package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReferralSubscription struct {
	Id           string    `json:"id" gorm:"type:uuid;primaryKey"`
	ReferrerId   string    `json:"referrerId" gorm:"type:uuid;not null"` // User who referred
	ReferredId   string    `json:"referredId" gorm:"type:uuid;not null"` // User who was referred
	ReferralCode string    `json:"referralCode" gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time `json:"createdAt"`
}

type ReferralAccount struct {
	Id           string    `json:"id" gorm:"type:uuid;primaryKey"`
	ReferrerId   string    `json:"referrerId" gorm:"type:uuid;not null"` // User who referred
	ReferredId   string    `json:"referredId" gorm:"type:uuid;not null"` // User who was referred
	ReferralCode string    `json:"referralCode" gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time `json:"createdAt"`
}

type ReferralEarnings struct {
	Id           string    `json:"id" gorm:"type:uuid;primaryKey"`
	ReferrerId   string    `json:"referrerId" gorm:"type:uuid;not null"`
	ReferredId   string    `json:"referredId" gorm:"type:uuid;not null"`
	AmountEarned float64   `json:"amountEarned" gorm:"not null"`
	CreatedAt    time.Time `json:"createdAt"`
}

type ReferralDiscount struct {
	Id           string    `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserId       string    `json:"userId" gorm:"type:uuid;not null"` // The user receiving the discount
	ReferralCode string    `json:"referralCode" gorm:"type:varchar(255);unique;not null"`
	Discount     float64   `json:"discount" gorm:"type:decimal(10,2);not null"` // Discount amount or percentage
	Validity     time.Time `json:"validity" gorm:"not null"`                    // Expiration date of the discount
	CreatedAt    time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

// Auto-generate UUID before creating the record
func (r *ReferralSubscription) BeforeCreate(tx *gorm.DB) error {
	r.Id = uuid.New().String()
	return nil
}

func (r *ReferralAccount) BeforeCreate(tx *gorm.DB) error {
	r.Id = uuid.New().String()
	return nil
}

func (e *ReferralEarnings) BeforeCreate(tx *gorm.DB) error {
	e.Id = uuid.New().String()
	return nil
}

func (e *ReferralDiscount) BeforeCreate(tx *gorm.DB) error {
	e.Id = uuid.New().String()
	return nil
}
