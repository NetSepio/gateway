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
	DomainId           string `gorm:"primary_key" json:"domainId"`
	Domain             Domain `gorm:"foreignkey:DomainId" json:"-"`
	Admin              User   `gorm:"foreignkey:AdminWalletAddress" json:"-"`
	UpdatedBy          User   `gorm:"foreignkey:UpdatedByAddress" json:"-"`
	UpdatedByAddress   string `json:"updatedBy"`
	Name               string `json:"name"`
	Role               string `json:"role"`
	AdminWalletAddress string `gorm:"primary_key" json:"walletAddress"`
}
