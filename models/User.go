package models

import (
	"time"
)

type User struct {
	UserId            string         `gorm:"primary_key" json:"userId,omitempty"`
	Name              string         `json:"name,omitempty"`
	WalletAddress     *string        `json:"walletAddress,omitempty"`
	Discord           string         `json:"discord"`
	Twitter           string         `json:"twitter"`
	FlowIds           []FlowId       `gorm:"foreignkey:UserId" json:"-"`
	ProfilePictureUrl string         `json:"profilePictureUrl,omitempty"`
	Country           string         `json:"country,omitempty"`
	Feedbacks         []UserFeedback `gorm:"foreignkey:UserId" json:"userFeedbacks"`
	Email             *string        `json:"emailId,omitempty"`
	ChainName         string         `json:"chainName,omitempty"`
	AppleId           *string        `json:"appleId,omitempty"`
	Apple             *string        `json:"apple,omitempty"`
	Google            *string        `json:"google,omitempty"` // Google
	Telegram          string         `json:"telegram,omitempty"`
	Farcaster         *string        `json:"farcaster,omitempty"` // Farcaster ID
	ReferralCode      string         `gorm:"unique" json:"referralCode"`
	CreatedAt         time.Time      `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

type TStripePiType string

type UserStripePi struct {
	Id           string        `gorm:"primary_key" json:"id,omitempty"`
	UserId       string        `json:"userId,omitempty"`
	StripePiId   string        `json:"stripePiId,omitempty"`
	StripePiType TStripePiType `json:"stripePiType,omitempty"`
	CreatedAt    time.Time     `json:"createdAt,omitempty"`
}

var Erebrus111NFT TStripePiType = "Erebrus111NFT"
var ThreeMonthSubscription TStripePiType = "ThreeMonthSubscription"
var TrialSubscription TStripePiType = "TrialSubscription"

type EmailAuth struct {
	Id        string    `gorm:"primary_key" json:"id,omitempty"`
	Email     string    `json:"email,omitempty"`
	AuthCode  string    `json:"authCode,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type SchemaMigration struct {
	Version int64 `gorm:"column:version"`
	Dirty   bool  `gorm:"column:dirty"`
}
