package database

import (
	"fmt"
	"strconv"
	"strings"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/NetSepio/gateway/utils/load"
)

var (
	DB  *gorm.DB
	DB2 *gorm.DB
)

type DBout struct {
	DB *gorm.DB
}

// SetDB sets the main database connection
func SetDB(database *gorm.DB) {
	DB = database
}

// SetDB2 sets the Erebrus database connection
func SetDB2(database *gorm.DB) {
	DB2 = database
}

type Logger struct {
	Log *zap.Logger
}

func GormLogger(l *zap.Logger) {
}

type ConfigWrapper struct {
	*load.Config
}

func (cfg ConfigWrapper) GetDB() (err error) {
	var b strings.Builder
	b.WriteString("host=")
	b.WriteString(cfg.DB_HOST)
	b.WriteString(" user=")
	b.WriteString(cfg.DB_USERNAME)
	b.WriteString(" password=")
	b.WriteString(cfg.DB_PASSWORD)
	b.WriteString(" dbname=")
	b.WriteString(cfg.DB_NAME)
	b.WriteString(" port=")
	b.WriteString(strconv.Itoa(cfg.DB_PORT))
	b.WriteString(" sslmode=disable TimeZone=UTC")
	dsn := b.String()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		SetDB(db)
	}

	return
}

// GetDb returns the main database connection
func GetDb() *gorm.DB {
	if DB != nil {
		return DB
	}
	var (
		host     = load.Cfg.DB_HOST
		username = load.Cfg.DB_USERNAME
		password = load.Cfg.DB_PASSWORD
		dbname   = load.Cfg.DB_NAME
		port     = load.Cfg.DB_PORT
	)

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%d",
		host, username, password, dbname, port)

	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dns,
	}))
	if err != nil {
		load.Logger.Error("failed to connect database", zap.Error(err))
	}

	sqlDb, err := DB.DB()
	if err != nil {
		load.Logger.Error("failed to get database instance", zap.Error(err))
	}
	if err = sqlDb.Ping(); err != nil {
		load.Logger.Error("failed to ping database", zap.Error(err))
	}

	return DB
}

// GetDB2 returns the Erebrus database connection
func GetDB2() *gorm.DB {
	if DB2 != nil {
		return DB2
	}
	var (
		host     = load.Cfg.DB_HOST
		username = load.Cfg.DB_USERNAME
		password = load.Cfg.DB_PASSWORD
		dbname   = `erebrus`
		port     = load.Cfg.DB_PORT
	)

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%d",
		host, username, password, dbname, port)

	var err error
	DB2, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dns,
	}))
	if err != nil {
		load.Logger.Error("failed to connect Erebrus database", zap.Error(err))
	}

	sqlDb, err := DB2.DB()
	if err != nil {
		load.Logger.Error("failed to get Erebrus database instance", zap.Error(err))
	}
	if err = sqlDb.Ping(); err != nil {
		load.Logger.Error("failed to ping Erebrus database", zap.Error(err))
	}

	return DB2
}
