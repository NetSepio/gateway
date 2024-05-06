package migrate

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserId            string `gorm:"primary_key;type:uuid"`
	Name              string
	WalletAddress     string `gorm:"unique"`
	Discord           string
	Twitter           string
	FlowIds           []FlowId
	ProfilePictureUrl string
	Country           string
	Feedbacks         []UserFeedback
	EmailId           string `gorm:"unique"`
}

// TODO: Make relations for field `relatedRoleId`
type FlowId struct {
	FlowIdType    FlowIdType
	UserId        string `gorm:"type:uuid"`
	FlowId        string `gorm:"primary_key"`
	RelatedRoleId string
	WalletAddress string
}

type FlowIdType string

func (fit *FlowIdType) Scan(value interface{}) error {
	*fit = FlowIdType([]byte(value.(string)))
	return nil
}

type UserFeedback struct {
	UserId    string `gorm:"primary_key"`
	Feedback  string `gorm:"primary_key"`
	Rating    int    `gorm:"primary_key"`
	CreatedAt time.Time
}
type Role struct {
	Name   string `gorm:"unique"`
	RoleId string `gorm:"primary_key"`
	Eula   string
}
type Report struct {
	ID                    string    `gorm:"type:uuid;primary_key;"`
	Title                 string    `gorm:"type:text;not null"`
	Description           string    `gorm:"type:text"`
	Document              string    `gorm:"type:text"`
	ProjectName           string    `gorm:"type:text"`
	ProjectDomain         string    `gorm:"type:text"`
	TransactionHash       *string   `gorm:"type:text"`
	TransactionVersion    *int64    `gorm:"type:text"`
	CreatedBy             string    `gorm:"type:uuid"`
	CreatedAt             time.Time `gorm:"type:timestamp"`
	EndTime               time.Time `gorm:"type:timestamp"`
	EndTransactionHash    *string   `gorm:"type:text"`
	EndTransactionVersion *int64    `gorm:"type:text"`
	MetaDataHash          *string   `gorm:"type:text"`
	EndMetaDataHash       *string   `gorm:"type:text"`
	Category              string    `gorm:"type:text"`
	UpVotes               int
	DownVotes             int
	NotSure               int
	TotalVotes            int
	Status                string
}
type ReportTag struct {
	ReportID string `gorm:"type:uuid;unique"`
	Tag      string `gorm:"type:text;unique"`
}

type ReportImage struct {
	ReportID string `gorm:"type:uuid;unique"`
	ImageURL string `gorm:"type:text;unique"`
}
type ReportVote struct {
	ReportID  string    `gorm:"type:uuid;primaryKey;"`
	VoterID   string    `gorm:"type:uuid;primaryKey;"`
	VoteType  string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"type:timestamp with time zone"`
}
type Review struct {
	Voter              string
	MetaDataUri        string
	Category           string
	DomainAddress      string
	SiteUrl            string
	SiteType           string
	SiteTag            string
	SiteSafety         string
	SiteIpfsHash       string
	TransactionHash    string
	TransactionVersion int64
	DeletedAt          gorm.DeletedAt
	CreatedAt          time.Time
	SiteRating         int
}
type WaitList struct {
	EmailId       string `gorm:"primary_key"`
	WalletAddress string
	Twitter       string
	Discord       string
}
type Domain struct {
	Id             string
	DomainName     string
	TxtValue       *string
	Verified       *bool `gorm:"not null;default:false"`
	CreatedAt      time.Time
	Title          string
	Headline       string
	Description    string
	CoverImageHash string
	LogoHash       string
	Category       string
	Blockchain     string
	CreatedBy      User
	UpdatedBy      User
	Claimable      bool
	CreatedById    string
	UpdatedById    string
}
type DomainAdmin struct {
	DomainId    string `gorm:"primary_key"`
	Domain      Domain
	Admin       User
	UpdatedBy   User
	UpdatedById string
	Name        string
	Role        string
	AdminId     string `gorm:"primary_key"`
}
type DomainClaim struct {
	ID       string `gorm:"primary_key;type:uuid"`
	DomainID string `gorm:"not null"`
	Txt      string `gorm:"not null;unique"`
	AdminId  string `gorm:"not null;type:uuid"`
}
type EmailAuth struct {
	Id        string `gorm:"primary_key"`
	Email     string `gorm:"unique"`
	AuthCode  string `gorm:"unique"`
	CreatedAt time.Time
}
type SchemaMigration struct {
	Version int64 `gorm:"primary_key;column:version"`
	Dirty   bool  `gorm:"column:dirty"`
}
type SiteInsight struct {
	SiteURL   string    `gorm:"primary_key"` //json:"siteUrl"
	Insight   string    //`json:"insight"`
	CreatedAt time.Time //`json:"createdAt"`
}
type TStripePiType string
type UserStripePi struct {
	Id           string `gorm:"primary_key;type:uuid"`
	UserId       string `gorm:"type:uuid"`
	StripePiId   string `gorm:"unique"`
	StripePiType TStripePiType
	CreatedAt    time.Time
}
type Sotreus struct {
	Name          string `gorm:"primary_key" json:"name"`
	WalletAddress string
	Region        string
}
type Erebrus struct {
	UUID          string `gorm:"primary_key" json:"UUID"`
	Name          string
	WalletAddress string
	Region        string
	CollectionId  string
}
