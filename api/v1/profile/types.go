package profile

type PatchProfileRequest struct {
	Name              string `json:"name,omitempty"`
	Country           string `json:"country,omitempty"`
	Discord           string `json:"discord,omitempty"`
	Twitter           string `json:"twitter,omitempty"`
	ProfilePictureUrl string `json:"profilePictureUrl,omitempty"`
}

type GetProfilePayload struct {
	UserId            string `json:"userId,omitempty"`
	Name              string `json:"name,omitempty"`
	WalletAddress     string `json:"walletAddress"`
	ProfilePictureUrl string `json:"profilePictureUrl,omitempty"`
	Country           string `json:"country,omitempty"`
	Discord           string `json:"discord,omitempty"`
	Twitter           string `json:"twitter,omitempty"`
  Email             *string `json:"email,omitempty"`
	Plan string `json:"plan,omitempty"`
}
