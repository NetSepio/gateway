package migrate

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserId            string `gorm:"primary_key;type:uuid"`
	Name              string
	WalletAddress     string
	DeviceType        string // web3.0,web2.0,mobile
	Discord           string
	Twitter           string
	FlowIds           []FlowId
	ProfilePictureUrl string
	Country           string
	Feedbacks         []UserFeedback
	Email             *string `gorm:"unique;index"`
	ChainName         string
	Apple             *string
	Google            *string
	Telegram          string
	Farcaster         *string
	Origin            *string   `gorm:"default:'web'" json:"origin"` // Origin of the user, e.g., web, mobile, etc.
	Metadata          *string   `gorm:"string" json:"metadata"`
	ReferralCode      string    `gorm:"unique" json:"referalCode"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`
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
type Leaderboard struct {
	ID           string `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Reviews      int
	Domain       int
	UserId       string `gorm:"type:uuid;not null"`
	Nodes        int
	DWifi        int
	Discord      int
	Twitter      int
	Telegram     int
	Subscription bool      `gorm:"default:false"`
	BetaTester   int       `gorm:"default:0" json:"beta_tester"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

type ScoreBoard struct {
	ID           string `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Reviews      int
	Domain       int
	UserId       string `gorm:"type:uuid;not null"`
	Nodes        int
	DWifi        int
	Discord      int
	Twitter      int
	Telegram     int
	Subscription bool      `gorm:"default:false"`
	BetaTester   int       `gorm:"default:0" json:"beta_tester"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

func (l *Leaderboard) BeforeCreate(tx *gorm.DB) (err error) {
	l.ID = uuid.New().String()
	return
}

type NftSubscription struct {
	ID              uint   `gorm:"primaryKey"`
	UserID          string `gorm:"index"`
	ContractAddress string
	ChainName       string
	Name            string
	Symbol          string
	TotalSupply     string
	Owner           string
	TokenURI        string
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

type DVPNNFTRecord struct {
	ID              uint `gorm:"primaryKey;autoIncrement"`
	Chain           string
	WalletAddress   string `gorm:"not null"`
	EmailID         string
	TransactionHash string
	CreatedAt       time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
}

type ActivityUnitXp struct {
	Activity string `gorm:"not null;unique"` // Name of the activity (e.g., Reviews, Domain, etc.)
	XP       int    `gorm:"not null"`
}

type Subscription struct {
	ID        uint      `gorm:"primary_key" json:"id,omitempty"`
	UserId    string    `json:"userId,omitempty"`
	Type      string    `json:"type,omitempty"`
	StartTime time.Time `json:"startTime,omitempty"`
	EndTime   time.Time `json:"endTime,omitempty"`
}

// SCOREBOARD = LEADERBOARD * ActivityUnitXp

type ReferralAccount struct {
	Id           string    `json:"id" gorm:"type:uuid;primaryKey"`
	ReferrerId   string    `json:"referrerId" gorm:"type:uuid;not null"` // User who referred
	ReferredId   string    `json:"referredId" gorm:"type:uuid;not null"` // User who was referred
	ReferralCode string    `json:"referralCode" gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time `json:"createdAt"`
}

type ReferralSubscription struct {
	Id           string    `json:"id" gorm:"type:uuid;primaryKey"`
	ReferrerId   string    `json:"referrerId" gorm:"type:uuid;not null"` // User who referred
	RefereeId    string    `json:"refereeId" gorm:"type:uuid;not null"`  // User who was referred
	ReferralCode string    `json:"referralCode" gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time `json:"createdAt"`
}

type ReferralEarnings struct {
	Id           string    `json:"id" gorm:"type:uuid;primaryKey"`
	ReferrerId   string    `json:"referrerId" gorm:"type:uuid;not null"`
	RefereeId    string    `json:"refereeId" gorm:"type:uuid;"`
	AmountEarned float64   `json:"amountEarned" gorm:"not null"`
	CreatedAt    time.Time `json:"createdAt"`
}

type ReferralDiscount struct {
	Id           string    `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserId       string    `json:"userId" gorm:"type:uuid;not null"` // The user receiving the discount
	ReferralCode string    `json:"referralCode" gorm:"type:varchar(255);unique;not null"`
	Discount     float64   `json:"discount" gorm:"type:decimal(10,2);not null"` // Discount amount or percentage
	Validity     time.Time `json:"validity" gorm:"not null"`                    // Expiration date of the discount
	CreatedAt    time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

type Node struct {
	//using for db operation
	PeerId           string  `json:"peerId" gorm:"primaryKey"`
	Name             string  `json:"name"`
	HttpPort         string  `json:"httpPort"`
	Host             string  `json:"host"` //domain
	PeerAddress      string  `json:"peerAddress"`
	Region           string  `json:"region"`
	Status           string  `json:"status"` // offline 1, online 2, maintainance 3,block 4
	DownloadSpeed    float64 `json:"downloadSpeed"`
	UploadSpeed      float64 `json:"uploadSpeed"`
	RegistrationTime int64   `json:"registrationTime"` //StartTimeStamp
	LastPing         int64   `json:"lastPing"`
	Chain            string  `json:"chainName"`
	WalletAddress    string  `json:"walletAddress"`
	Version          string  `json:"version"`
	CodeHash         string  `json:"codeHash"`
	SystemInfo       string  `json:"systemInfo" gorm:"type:jsonb"`
	IpInfo           string  `json:"ipinfo" gorm:"type:jsonb"`
	IpGeoData        string  `json:"ipGeoData" gorm:"type:jsonb"`
	NodeType         string  `json:"nodeType"`
	NodeConfig       string  `json:"nodeConfig"`
}

type Organisation struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(225);not null" json:"name"`
	IPAddress   string    `gorm:"type:varchar(225);not null" json:"ip_address" binding:"required"`
	APIKey      string    `gorm:"not null" json:"api_key"`
	Status      string    `gorm:"type:varchar(50);default:'inactive'" json:"status"` // e.g., active, inactive, suspended
	OrgMetaData string    `gorm:"type:varchar(50)" json:"org_meta_data"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type OrganisationApp struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	OrganisationId uuid.UUID `gorm:"type:uuid;not null" json:"organisation_id"`
	Name           string    `gorm:"type:varchar(225);not null" json:"name"`
	Description    string    `gorm:"type:text" json:"description"`
	MetaData       string    `gorm:"type:text" json:"meta_data_uri"`
	APIKey         string    `gorm:"not null" json:"api_key"`
	Status         string    `gorm:"type:varchar(50);default:'inactive'" json:"status"` // e.g., active, inactive, suspended
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Plan struct {
	ID            string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name          string         `gorm:"not null" json:"name"`
	Status        string         `gorm:"not null" json:"status"`           // e.g., active, inactive
	AllowedRegion []string       `gorm:"type:jsonb" json:"allowed_region"` // null or empty = all regions allowed
	MaxClients    int            `gorm:"not null" json:"max_clients"`
	Duration      int            `gorm:"not null" json:"duration"`    // in days
	PriceCents    int64          `gorm:"not null" json:"price_cents"` // stored in cents
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

type SubscriptionPlan struct {
	ID          string         `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	PlanID      string         `gorm:"type:uuid;not null" json:"plan_id"`                               // foreign key
	Plan        Plan           `gorm:"foreignKey:PlanID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"` // optional: preload plan
	DateCreated time.Time      `gorm:"not null" json:"date_created"`
	Status      string         `gorm:"not null" json:"status"` // e.g., active, expired
	AutoRenewal bool           `gorm:"default:false" json:"auto_renewal"`
	CreatedBy   string         `gorm:"type:uuid;not null" json:"created_by"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type SubscriptionRenewal struct {
	ID                 string           `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CreatedBy          string           `gorm:"type:uuid;not null" json:"created_by"`
	SubscriptionPlanID string           `gorm:"type:uuid;not null" json:"subscription_plan_id"` // 'type' field renamed for clarity
	SubscriptionPlan   SubscriptionPlan `gorm:"foreignKey:SubscriptionPlanID" json:"subscription_plan"`
	StartTime          time.Time        `gorm:"not null" json:"start_time"`
	EndTime            time.Time        `gorm:"not null" json:"end_time"`
	CreatedAt          time.Time        `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt          time.Time        `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt          gorm.DeletedAt   `gorm:"index" json:"-"`
}
type OrgSubscription struct {
	ID              uuid.UUID    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	OrganisationID  uuid.UUID    `gorm:"type:uuid;not null;index" json:"organisation_id"`
	Organisation    Organisation `gorm:"foreignKey:OrganisationID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"organisation"`
	StartTime       time.Time    `gorm:"not null" json:"start_time"`
	EndTime         *time.Time   `json:"end_time"`
	BillingCycle    string       `gorm:"type:varchar(20);not null" json:"billing_cycle"` // e.g., Monthly, Quarterly
	NextBillingDate time.Time    `gorm:"not null" json:"next_billing_date"`
	AmountDue       float64      `gorm:"type:numeric(10,2);not null" json:"amount_due"`
	Status          string       `gorm:"type:varchar(20);not null" json:"status"` // Active, Cancelled, Overdue
	LastPaymentDate *time.Time   `json:"last_payment_date"`
	CreatedAt       time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
}

type Agent struct {
	ID            string    `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time `json:"created_at" gorm:"autoCreaeTime"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	Name          string    `json:"name"`
	Clients       string    `json:"clients"`
	Status        string    `json:"status"`
	AvatarImg     string    `json:"avatar_img"`
	CoverImg      string    `json:"cover_img"`
	VoiceModel    string    `json:"voice_model"`
	Organization  string    `json:"organization"`
	WalletAddress string    `json:"wallet_address" gorm:"index"`
	ServerDomain  string    `json:"server_domain"`
	Domain        string    `json:"domain"`
	CharacterFile string    `json:"character_file"`
}

type UserActivity struct {
	Id        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserId    string    `json:"user_id"`
	Modules   string    `json:"modules"`
	Action    string    `json:"action"`
	Metadata  string    `json:"metadata"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

