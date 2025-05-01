package claim

type ClaimDomainRequest struct {
	DomainId string `json:"domainId" binding:"required"`
}

type ClaimDomainResponse struct {
	TxtValue string `json:"txtValue"`
}

type FinishClaimDomainRequest struct {
	DomainId string `json:"domainId" binding:"required"`
}
