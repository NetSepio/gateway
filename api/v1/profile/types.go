package profile

type PatchProfileRequest struct {
	Name              string  `json:"name,omitempty"`
	Country           string  `json:"country,omitempty"`
	EmailId           string  `json:"emailId,omitempty"`
	Discord           string  `json:"discord,omitempty"`
	Twitter           string  `json:"twitter,omitempty"`
	Google            *string `json:"google,omitempty"`
	Apple             *string `json:"apple,omitempty"`
	Telegram          string  `json:"telegram,omitempty"`
	Farcaster         *string `json:"farcaster,omitempty"`
	ProfilePictureUrl string  `json:"profilePictureUrl,omitempty"`
}

type GetProfilePayload struct {
	UserId            string  `json:"userId,omitempty"`
	Name              string  `json:"name,omitempty"`
	WalletAddress     *string `json:"walletAddress,omitempty"`
	ProfilePictureUrl string  `json:"profilePictureUrl,omitempty"`
	Country           string  `json:"country,omitempty"`
	Discord           string  `json:"discord,omitempty"`
	Twitter           string  `json:"twitter,omitempty"`
	Email             *string `json:"email,omitempty"`
	Apple             *string `json:"apple,omitempty"`
	Telegram          string  `json:"telegram,omitempty"`
	Farcaster         *string `json:"farcaster,omitempty"`
	Google            *string `json:"google,omitempty"`
	ChainName         string `json:"chainName,omitempty"`
}

type UpdateUserRequest struct {
	Discord   string  `json:"discord"`   // Required
	Twitter   string  `json:"twitter"`   // Required for X (formerly Twitter)
	Google    *string `json:"google"`    // Required for Google
	Apple     *string `json:"apple"`     // Required for Apple
	Telegram  string  `json:"telegram"`  // Required for Telegram
	Farcaster *string `json:"farcaster"` // Required for Farcaster
}
