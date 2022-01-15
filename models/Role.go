package models

type Role struct {
	Name   string `gorm:"unique"`
	RoleId string `gorm:"primary_key"`
	Eula   string
}
