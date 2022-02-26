package profile

type PatchProfileRequest struct {
	Name              string `json:"name,omitempty"`
	Country           string `json:"country,omitempty"`
	ProfilePictureUrl string `json:"profilePictureUrl,omitempty"`
}

type GetProfilePayload struct {
	Name              string `json:"name,omitempty"`
	WalletAddress     string `json:"walletAddress"`
	ProfilePictureUrl string `json:"profilePictureUrl,omitempty"`
	Country           string `json:"country,omitempty"`
}
