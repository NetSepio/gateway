package profile

type PatchProfileRequest struct {
	Name              string `json:"name,omitempty"`
	Country           string `json:"country,omitempty"`
	ProfilePictureUrl string `json:"profilePictureUrl,omitempty"`
}
