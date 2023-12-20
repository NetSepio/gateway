package models

type User struct {
	UserId            string         `gorm:"primary_key" json:"userId,omitempty"`
	Name              string         `json:"name,omitempty"`
	WalletAddress     string         `gorm:"unique" json:"walletAddress"`
	Discord           string         `json:"discord"`
	Twitter           string         `json:"twitter"`
	FlowIds           []FlowId       `gorm:"foreignkey:UserId" json:"-"`
	ProfilePictureUrl string         `json:"profilePictureUrl,omitempty"`
	Country           string         `json:"country,omitempty"`
	Feedbacks         []UserFeedback `gorm:"foreignkey:UserId" json:"userFeedbacks"`
	EmailId           string         `json:"emailId,omitempty"`
}
