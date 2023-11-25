package domain

type CreateDomainRequest struct {
	DomainName     string `json:"domainName" binding:"required"`
	Title          string `json:"title" binding:"required"`
	Headline       string `json:"headline"`
	Description    string `json:"description" binding:"required"`
	CoverImageHash string `json:"coverImageHash" binding:"required"`
	LogoHash       string `json:"logoHash" binding:"required"`
	Category       string `json:"category" binding:"required"`
	AdminName      string `json:"adminName" binding:"required,min=1"`
	AdminRole      string `json:"adminRole" binding:"required,min=1"`
}

type DeleteDomainQuery struct {
	DomainId string `form:"domainId" binding:"required"`
}

type CreateDomainResponse struct {
	TxtValue string `json:"txtValue"`
	DomainId string `json:"domainId"`
}

type VerifyDomainRequest struct {
	DomainId string `json:"domainId" binding:"required"`
}

type GetDomainsQuery struct {
	DomainId   string `form:"domainId"`
	DomainName string `form:"domainName"`
	Verified   *bool  `form:"verified"`
	Page       *int   `form:"page" binding:"required,min=1"`
}
