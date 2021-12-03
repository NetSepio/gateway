package models

type Role struct {
	Name   string `gorm:"unique"`
	RoleId int    `gorm:"primary_key"`
	Eula   string
}
