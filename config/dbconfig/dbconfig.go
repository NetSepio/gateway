package dbconfig

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/NetSepio/gateway/config/envconfig"
	migrate "github.com/NetSepio/gateway/models/Migrate"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"

	"gorm.io/driver/postgres"
)

var db *gorm.DB

// Return singleton instance of db, initiates it before if it is not initiated already
func GetDb() *gorm.DB {
	if db != nil {
		return db
	}
	var (
		host     = envconfig.EnvVars.DB_HOST
		username = envconfig.EnvVars.DB_USERNAME
		password = envconfig.EnvVars.DB_PASSWORD
		dbname   = envconfig.EnvVars.DB_NAME
		port     = envconfig.EnvVars.DB_PORT
	)

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%d",
		host, username, password, dbname, port)

	var err error
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dns,
	}))
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal("failed to ping database", err)
	}
	if err = sqlDb.Ping(); err != nil {
		log.Fatal("failed to ping database", err)
	}

	return db
}

func Init() error {
	db := GetDb()
	if err := db.AutoMigrate(
		&migrate.User{},
		&migrate.Role{},
		&migrate.UserFeedback{},
		&migrate.FlowId{},
		&migrate.Report{},
		&migrate.ReportTag{},
		&migrate.ReportImage{},
		&migrate.ReportVote{},
		&migrate.Review{},
		&migrate.WaitList{},
		&migrate.Domain{},
		&migrate.DomainAdmin{},
		&migrate.DomainClaim{},
		&migrate.EmailAuth{},
		&migrate.SchemaMigration{},
		&migrate.SiteInsight{},
		&migrate.UserStripePi{},
		&migrate.Sotreus{},
		&migrate.Erebrus{},
		&migrate.Leaderboard{},
		&migrate.NftSubscription{},
		&migrate.DVPNNFTRecord{},
	); err != nil {
		log.Fatal(err)
	}

	// db.Exec(`ALTER TABLE leader_boards DROP COLUMN IF EXISTS users;`)
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	logwrapper.Log.Info("Congrats ! Automigration completed")

	return nil
}
