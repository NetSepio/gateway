package getreviews

import "time"

type GetReviewsQuery struct {
	Voter  string `form:"voter"`
	Domain string `form:"domain"`
	Page   *int   `form:"page" binding:"required,min=1"`
}
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
	Voter              string    `json:"voter"`
	Name               string    `json:"name"`
}

type GetReviewsPayload []GetReviewsItem
