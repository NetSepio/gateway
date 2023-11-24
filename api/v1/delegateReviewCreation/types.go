package delegatereviewcreation

type DelegateReviewCreationRequest struct {
	MetaDataUri   string `json:"metaDataUri" binding:"required"`
	Category      string `json:"category" binding:"required"`
	DomainAddress string `json:"domainAddress" binding:"required,hostname_rfc1123|hostname_port"`
	SiteUrl       string `json:"siteUrl" binding:"required,http_url"`
	SiteType      string `json:"siteType" binding:"required"`
	SiteTag       string `json:"siteTag" binding:"required"`
	SiteSafety    string `json:"siteSafety" binding:"required"`
}

type DelegateReviewCreationPayload struct {
	TransactionHash    string `json:"transactionHash"`
	TransactionVersion int64  `json:"transactionVersion"`
}
