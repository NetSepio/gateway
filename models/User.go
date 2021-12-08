package models

type User struct {
	Name              string     `json:"name,omitempty"`
	WalletAddress     string     `gorm:"primary_key" json:"walletAddress"`
	FlowIds           []FlowId   `gorm:"foreignkey:WalletAddress" json:"-"`
	ProfilePictureUrl string     `json:"profilePictureUrl,omitempty"`
	Country           string     `json:"country,omitempty"`
	Roles             []UserRole `gorm:"foreignkey:WalletAddress;type:int[];not null" json:"roles"`
}
