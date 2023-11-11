package getreviews

import "time"

type GetReviewsItem struct {
	MetaDataUri        string    `json:"metaDataUri"`
	Category           string    `json:"category"`
	DomainAddress      string    `json:"domainAddress"`
	SiteUrl            string    `json:"siteUrl"`
	SiteType           string    `json:"siteType"`
	SiteTag            string    `json:"siteTag"`
	SiteSafety         string    `json:"siteSafety"`
	SiteIpfsHash       string    `json:"siteIpfsHash"`
	TransactionHash    string    `json:"transactionHash"`
	TransactionVersion int64     `json:"transactionVersion"`
	CreatedAt          time.Time `json:"createdAt"`
}

type GetReviewsPayload []GetReviewsItem
