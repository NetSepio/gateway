package admin

type GetAdminQuery struct {
	DomainId string `json:"domainId" binding:"required"`
}

type DeleteAdminRequest struct {
	AdminWalletAddres string `json:"adminWalletAddress" binding:"required,min=5"`
	DomainId          string `json:"domainId" binding:"required,min=10"`
}

type AdminDetails struct {
	AdminName          string `json:"adminName" binding:"required,min=1"`
	AdminRole          string `json:"adminRole" binding:"required,min=1"`
	AdminWalletAddress string `json:"adminWalletAddress" binding:"required,min=1"`
}

type CreateAdminRequest struct {
	Admins   []AdminDetails `json:"admins" binding:"min=1,dive"`
	DomainId string         `json:"domainId" binding:"required,min=10"`
}
