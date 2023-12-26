package models

import "time"

type SiteInsight struct {
	SiteURL   string    `gorm:"primary_key" json:"siteUrl"`
	Insight   string    `json:"insight"`
	CreatedAt time.Time `json:"createdAt"`
}
