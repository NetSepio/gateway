package models

import (
	"time"
)

type Report struct {
	ID                    string    `gorm:"type:uuid;primary_key;" json:"id"`
	Title                 string    `gorm:"type:text;not null" json:"title"`
	Description           string    `gorm:"type:text" json:"description"`
	Document              string    `gorm:"type:text" json:"document"`
	ProjectName           string    `gorm:"type:text" json:"projectName"`
	ProjectDomain         string    `gorm:"type:text" json:"projectDomain"`
	TransactionHash       *string   `gorm:"type:text" json:"transactionHash"`
	TransactionVersion    *int64    `gorm:"type:text" json:"transactionVersion"`
	CreatedBy             string    `gorm:"type:uuid" json:"createdBy"`
	CreatedAt             time.Time `gorm:"type:timestamp" json:"createdAt"`
	EndTime               time.Time `gorm:"type:timestamp" json:"endTime"`
	EndTransactionHash    *string   `gorm:"type:text" json:"endTransactionHash"`
	EndTransactionVersion *int64    `gorm:"type:text" json:"endTransactionVersion"`
	MetaDataHash          *string   `gorm:"type:text" json:"metaDataHash"`
	EndMetaDataHash       *string   `gorm:"type:text" json:"endMetaDataHash"`
	Category              string    `gorm:"type:text" json:"category"`
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
