package models

import "time"

type Domain struct {
	Id               string    `json:"id" gorm:"primary_key"`
	DomainName       string    `json:"domainName"`
	TxtValue         *string   `json:"txtValue"`
	Verified         *bool     `json:"verified" gorm:"not null;default:false"`
	CreatedAt        time.Time `json:"createdAt"`
	Title            string    `json:"title"`
	Headline         string    `json:"headline"`
	Description      string    `json:"description"`
	CoverImageHash   string    `json:"coverImageHash"`
	LogoHash         string    `json:"logoHash"`
	Category         string    `json:"category"`
	Blockchain       string    `json:"blockchain"`
	CreatedBy        User      `gorm:"foreignkey:CreatedByAddress"`
	UpdatedBy        User      `gorm:"foreignkey:UpdatedByAddress"`
	CreatedByAddress string
	UpdatedByAddress string
}

type DomainAdmin struct {
	DomainId           string
	Domain             Domain `gorm:"foreignkey:DomainId"`
	Admin              User   `gorm:"foreignkey:AdminWalletAddress"`
	UpdatedBy          User   `gorm:"foreignkey:UpdatedByAddress"`
	UpdatedByAddress   string
	Name               string
	Role               string
	AdminWalletAddress string
}