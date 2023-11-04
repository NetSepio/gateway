package models

import "time"

type Review struct {
	Voter              string    `json:"voter" binding:"required,hexadecimal"`
	MetaDataUri        string    `json:"metaDataUri" binding:"required"`
	Category           string    `json:"category" binding:"required"`
	DomainAddress      string    `json:"domainAddress" binding:"required"`
	SiteUrl            string    `json:"siteUrl" binding:"required"`
	SiteType           string    `json:"siteType" binding:"required"`
	SiteTag            string    `json:"siteTag" binding:"required"`
	SiteSafety         string    `json:"siteSafety" binding:"required"`
	SiteIpfsHash       string    `json:"siteIpfsHash" binding:"required"`
	TransactionHash    string    `json:"transactionHash" gorm:"primary_key"`
	TransactionVersion int64     `json:"transactionVersion"`
	CreatedAt          time.Time `json:"createdAt"`
}
