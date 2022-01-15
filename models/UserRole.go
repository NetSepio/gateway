package models

type UserRole struct {
	WalletAddress string `gorm:"unique"`
	RoleId        string `gorm:"unique"`
}
