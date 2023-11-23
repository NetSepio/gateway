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
	//Title, Headline, Description (supports RTF), name (example: netsepio.com), coverImageHash, logoHash, Category
}

type DomainAdmin struct {
	DomainId           string
	Domain             Domain `gorm:"foreignkey:DomainId"`
	Admin              User   `gorm:"foreignkey:AdminWalletAddress"`
	AdminWalletAddress string
}

// domains POST
