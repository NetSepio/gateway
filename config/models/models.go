package models

import (
	"time"
)

type User struct {
	UserId        string  `gorm:"primary_key" json:"userId,omitempty"`
	Name          string  `json:"name,omitempty"`
	WalletAddress *string `json:"walletAddress,omitempty"`
	Discord       string  `json:"discord"`
	Twitter       string  `json:"twitter"`
	// FlowIds           []FlowId       `gorm:"foreignkey:UserId" json:"-"`
	ProfilePictureUrl string `json:"profilePictureUrl,omitempty"`
	Country           string `json:"country,omitempty"`
	// Feedbacks         []UserFeedback `gorm:"foreignkey:UserId" json:"userFeedbacks"`
	EmailId *string `json:"emailId,omitempty"`
}

type FlowId struct {
	FlowIdType    FlowIdType
	UserId        string `gorm:"type:uuid"`
	FlowId        string `gorm:"primary_key"`
	RelatedRoleId string
	WalletAddress string
}
type FlowIdType string

type ReportVote struct {
	ReportID  string    `gorm:"column:report_id"`
	VoterID   string    `gorm:"column:voter_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	VoteType  string    `gorm:"column:vote_type"`
}

type Domain struct {
	CreatedByID    string    `gorm:"column:created_by_id"`
	UpdatedByID    string    `gorm:"column:updated_by_id"`
	Claimable      bool      `gorm:"column:claimable"`
	Verified       bool      `gorm:"column:verified"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	ID             string    `gorm:"column:id"`
	DomainName     string    `gorm:"column:domain_name"`
	TxtValue       string    `gorm:"column:txt_value"`
	Title          string    `gorm:"column:title"`
	Headline       string    `gorm:"column:headline"`
	Description    string    `gorm:"column:description"`
	CoverImageHash string    `gorm:"column:cover_image_hash"`
	LogoHash       string    `gorm:"column:logo_hash"`
	Category       string    `gorm:"column:category"`
	Blockchain     string    `gorm:"column:blockchain"`
}

type SiteInsight struct {
	CreatedAt time.Time `gorm:"column:created_at"`
	SiteURL   string    `gorm:"column:site_url"`
	Insight   string    `gorm:"column:insight"`
}

type DomainClaim struct {
	ID       string `gorm:"column:id"`
	DomainID string `gorm:"column:domain_id"`
	Txt      string `gorm:"column:txt"`
}

type FlowID struct {
	UserID        string `gorm:"column:user_id"`
	FlowIDType    string `gorm:"column:flow_id_type"`
	FlowID        string `gorm:"column:flow_id"`
	RelatedRoleID string `gorm:"column:related_role_id"`
	WalletAddress string `gorm:"column:wallet_address"`
}

type UserFeedback struct {
	UserID    string    `gorm:"column:user_id"`
	Rating    int64     `gorm:"column:rating"`
	CreatedAt time.Time `gorm:"column:created_at"`
	Feedback  string    `gorm:"column:feedback"`
}

// Define other structs similarly for the remaining tables in the schema

type Report struct {
	ID                    string    `gorm:"column:id"`
	CreatedBy             string    `gorm:"column:created_by"`
	EndTime               time.Time `gorm:"column:end_time"`
	CreatedAt             time.Time `gorm:"column:created_at"`
	TransactionVersion    int64     `gorm:"column:transaction_version"`
	EndTransactionVersion int64     `gorm:"column:end_transaction_version"`
	MetaDataHash          string    `gorm:"column:meta_data_hash"`
	EndMetaDataHash       string    `gorm:"column:end_meta_data_hash"`
	Category              string    `gorm:"column:category"`
	UpVotes               int       `gorm:"column:up_votes"`
	DownVotes             int       `gorm:"column:down_votes"`
	NotSure               int       `gorm:"column:not_sure"`
	TotalVotes            int       `gorm:"column:total_votes"`
	Title                 string    `gorm:"column:title"`
	Description           string    `gorm:"column:description"`
	Document              string    `gorm:"column:document"`
	ProjectName           string    `gorm:"column:project_name"`
	ProjectDomain         string    `gorm:"column:project_domain"`
	Status                string    `gorm:"column:status"`
	TransactionHash       string    `gorm:"column:transaction_hash"`
	EndTransactionHash    string    `gorm:"column:end_transaction_hash"`
}

type Review struct {
	DeletedAt          time.Time `gorm:"column:deleted_at"`
	CreatedAt          time.Time `gorm:"column:created_at"`
	SiteRating         int       `gorm:"column:site_rating"`
	TransactionVersion int64     `gorm:"column:transaction_version"`
	Voter              string    `gorm:"column:voter"`
	MetaDataURI        string    `gorm:"column:meta_data_uri"`
	Category           string    `gorm:"column:category"`
	DomainAddress      string    `gorm:"column:domain_address"`
	SiteURL            string    `gorm:"column:site_url"`
	SiteType           string    `gorm:"column:site_type"`
	SiteTag            string    `gorm:"column:site_tag"`
	SiteSafety         string    `gorm:"column:site_safety"`
	SiteIPFSHash       string    `gorm:"column:site_ipfs_hash"`
	TransactionHash    string    `gorm:"column:transaction_hash"`
}

type EmailAuth struct {
	ID        string    `gorm:"column:id"`
	Email     string    `gorm:"column:email"`
	AuthCode  string    `gorm:"column:auth_code"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

type Role struct {
	Name   string `gorm:"column:name"`
	RoleID string `gorm:"column:role_id"`
	EULA   string `gorm:"column:eula"`
}

type Sotreus struct {
	Name          string `gorm:"column:name"`
	WalletAddress string `gorm:"column:wallet_address"`
	Region        string `gorm:"column:region"`
}

type UserRole struct {
	WalletAddress string `gorm:"column:wallet_address"`
	RoleID        string `gorm:"column:role_id"`
}

type WaitList struct {
	EmailID       string `gorm:"column:email_id"`
	WalletAddress string `gorm:"column:wallet_address"`
	Twitter       string `gorm:"column:twitter"`
	Discord       string `gorm:"column:discord"`
}

type ReportImage struct {
	ReportID string `gorm:"column:report_id"`
	ImageURL string `gorm:"column:image_url"`
}

type Erebrus struct {
	UUID          string `gorm:"column:uuid"`
	Name          string `gorm:"column:name"`
	WalletAddress string `gorm:"column:wallet_address"`
	Region        string `gorm:"column:region"`
	CollectionID  string `gorm:"column:collection_id"`
}

type ReportTag struct {
	ReportID string `gorm:"column:report_id"`
	Tag      string `gorm:"column:tag"`
}

type UserStripePI struct {
	ID           string    `gorm:"column:id"`
	UserID       string    `gorm:"column:user_id"`
	StripePIID   string    `gorm:"column:stripe_pi_id"`
	StripePIType string    `gorm:"column:stripe_pi_type"`
	CreatedAt    time.Time `gorm:"column:created_at"`
}

type DomainAdmin struct {
	AdminID     string `gorm:"column:admin_id"`
	DomainID    string `gorm:"column:domain_id"`
	Name        string `gorm:"column:name"`
	Role        string `gorm:"column:role"`
	UpdatedByID string `gorm:"column:updated_by_id"`
}

type SchemaMigration struct {
	Version int64 `gorm:"column:version"`
	Dirty   bool  `gorm:"column:dirty"`
}

// Define other structs similarly for the remaining tables in the schema
