package models

type DomainClaim struct {
	ID       string `gorm:"primary_key" json:"id"`
	DomainID string `gorm:"not null" json:"domain_id"`
	Txt      string `gorm:"not null;unique" json:"txt"`
	AdminId  string `gorm:"not null" json:"admin_id"`
}
