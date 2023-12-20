package models

import (
	"time"
)

type Report struct {
	ID            string    `gorm:"type:uuid;primary_key;"`
	Title         string    `gorm:"type:text;not null"`
	Description   string    `gorm:"type:text"`
	Document      string    `gorm:"type:text"`
	ProjectName   string    `gorm:"type:text"`
	ProjectDomain string    `gorm:"type:text"`
	CreatedBy     string    `gorm:"type:uuid"`
	CreatedAt     time.Time `gorm:"type:timestamp"`
	EndTime       time.Time `gorm:"type:timestamp"`
}

type ReportVote struct {
	ReportID  string    `gorm:"type:uuid;primaryKey;"`
	VoterID   string    `gorm:"type:uuid;primaryKey;"`
	VoteType  string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"type:timestamp with time zone"`
}

type ReportTag struct {
	ReportID string `gorm:"type:uuid;"`
	Tag      string `gorm:"type:text;"`
}

type ReportImage struct {
	ReportID string `gorm:"type:uuid;"`
	ImageURL string `gorm:"type:text;"`
}
