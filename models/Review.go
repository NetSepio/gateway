package models

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	Voter              string `json:"voter"`
	MetaDataUri        string `json:"metaDataUri"`
	Category           string `json:"category"`
	DomainAddress      string `json:"domainAddress"`
	SiteUrl            string `json:"siteUrl"`
	SiteType           string `json:"siteType"`
	SiteTag            string `json:"siteTag"`
	SiteSafety         string `json:"siteSafety"`
	SiteIpfsHash       string `json:"siteIpfsHash"`
	TransactionHash    string `json:"transactionHash"`
	TransactionVersion int64  `json:"transactionVersion"`
	DeletedAt          gorm.DeletedAt
	CreatedAt          time.Time `json:"createdAt"`
	SiteRating         int       `json:"siteRating"`
}
