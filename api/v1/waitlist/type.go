package waitlist

type WaitListRequest struct {
	EmailId       string `binding:"required,email" json:"emailId"`
	WalletAddress string `json:"walletAddress,omitempty" binding:"omitempty,hexadecimal,startswith=0x"`
	Twitter       string `json:"twitter,omitempty" binding:"omitempty,startswith=@"`
}
