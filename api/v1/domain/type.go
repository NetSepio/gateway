package domain

type CreateDomainRequest struct {
	DomainName     string `json:"domainName" binding:"required"`
	Title          string `json:"title" binding:"required"`
	Headline       string `json:"headline" binding:"required"`
	Description    string `json:"description" binding:"required"`
	CoverImageHash string `json:"coverImageHash" binding:"required"`
	LogoHash       string `json:"logoHash" binding:"required"`
	Category       string `json:"category" binding:"required"`
}

type CreateDomainResponse struct {
	TxtValue string `json:"txtValue"`
	DomainId string `json:"domainId"`
}

type VerifyDomainRequest struct {
	DomainId string `json:"domainId" binding:"required"`
}

type GetDomainsQuery struct {
	DomainId string `form:"domainId"`
	Domain   string `form:"domain"`
	Verified *bool  `form:"verified"`
	Page     *int   `form:"page" binding:"required,min=1"`
}
