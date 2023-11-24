package delegatereviewcreation

type DelegateReviewCreationRequest struct {
	MetaDataUri   string `json:"metaDataUri" binding:"required"`
	Category      string `json:"category" binding:"required"`
	DomainAddress string `json:"domainAddress" binding:"required"`
	SiteUrl       string `json:"siteUrl" binding:"required,http_url"`
	SiteType      string `json:"siteType" binding:"required"`
	SiteTag       string `json:"siteTag" binding:"required"`
	SiteSafety    string `json:"siteSafety" binding:"required"`
	SiteIpfsHash  string `json:"siteIpfsHash" binding:"required"`
}

type DelegateReviewCreationPayload struct {
	TransactionHash    string `json:"transactionHash"`
	TransactionVersion int64  `json:"transactionVersion"`
}
