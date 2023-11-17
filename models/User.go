package models

type User struct {
	Name              string         `json:"name,omitempty"`
	WalletAddress     string         `gorm:"primary_key" json:"walletAddress"`
	Discord           string         `json:"discord"`
	Twitter           string         `json:"twitter"`
	FlowIds           []FlowId       `gorm:"foreignkey:WalletAddress" json:"-"`
	ProfilePictureUrl string         `json:"profilePictureUrl,omitempty"`
	Country           string         `json:"country,omitempty"`
	Feedbacks         []UserFeedback `gorm:"foreignkey:WalletAddress" json:"userFeedbacks"`
}
