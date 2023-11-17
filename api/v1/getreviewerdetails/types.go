package getreviewerdetails

type GetReviewerDetailsQuery struct {
	WalletAddress string `form:"walletAddress" binding:"required,startswith=0x,hexadecimal"`
}

type GetReviewerDetailsPayload struct {
	Name              string `json:"name,omitempty"`
	WalletAddress     string `json:"walletAddress"`
	ProfilePictureUrl string `json:"profilePictureUrl,omitempty"`
}
