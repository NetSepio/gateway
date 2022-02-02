package delegateartifactcreation

type DelegateArtifactCreationRequest struct {
	Category      string `json:"category" binding:"required"`
	DomainAddress string `json:"domainAddress" binding:"required"`
	SiteUrl       string `json:"siteUrl" binding:"required"`
	SiteType      string `json:"siteType" binding:"required"`
	SiteTag       string `json:"siteTag" binding:"required"`
	SiteSafety    string `json:"siteSafety" binding:"required"`
	MetaDataUri   string `json:"metaDataUri" binding:"required"`
	Voter         string `json:"voter" binding:"required"`
}

type DelegateArtifactCreationPayload struct {
	TransactionHash string `json:"transactionHash"`
}
