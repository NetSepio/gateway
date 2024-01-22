package models

import "time"

type Domain struct {
	Id             string    `json:"id" gorm:"primary_key"`
	DomainName     string    `json:"domainName"`
	TxtValue       *string   `json:"txtValue"`
	Verified       *bool     `json:"verified" gorm:"not null;default:false"`
	CreatedAt      time.Time `json:"createdAt"`
	Title          string    `json:"title"`
	Headline       string    `json:"headline"`
	Description    string    `json:"description"`
	CoverImageHash string    `json:"coverImageHash"`
	LogoHash       string    `json:"logoHash"`
	Category       string    `json:"category"`
	Blockchain     string    `json:"blockchain"`
	CreatedBy      User      `gorm:"foreignkey:CreatedById"`
	UpdatedBy      User      `gorm:"foreignkey:UpdatedById"`
	Claimable      bool      `json:"claimable"`
	CreatedById    string
	UpdatedById    string
}

type DomainAdmin struct {
	DomainId    string `gorm:"primary_key" json:"domainId"`
	Domain      Domain `gorm:"foreignkey:DomainId" json:"-"`
	Admin       User   `gorm:"foreignkey:AdminId" json:"-"`
	UpdatedBy   User   `gorm:"foreignkey:UpdatedById" json:"-"`
	UpdatedById string `json:"updatedBy"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	AdminId     string `gorm:"primary_key" json:"walletAddress"`
}
