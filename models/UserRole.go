package models

type UserRole struct {
	WalletAddress string `gorm:"unique"`
	RoleId        int    `gorm:"unique"`
}
