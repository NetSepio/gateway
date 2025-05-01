package database

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"netsepio-gateway-v1.1/utils/load"
)

var DB *gorm.DB

type DBout struct {
	DB *gorm.DB
}

// SetDB sets the database connection
func SetDB(database *gorm.DB) {
	DB = database
}

type Logger struct {
	Log *zap.Logger
}

func GormLogger(l *zap.Logger) {
}

type ConfigWrapper struct {
	*load.Config
}

func (cfg ConfigWrapper) GetDB() (out DBout, err error) {

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
	}
	// Set the database connection

	out = DBout{DB: db}

	// Migrate the schema

	return
}

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
		log.Fatal("failed to connect database", err)
	}

	sqlDb, err := DB.DB()
	if err != nil {
		log.Fatal("failed to ping database", err)
	}
	if err = sqlDb.Ping(); err != nil {
		log.Fatal("failed to ping database", err)
	}

	return DB
}
