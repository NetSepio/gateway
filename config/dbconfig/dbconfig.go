package dbconfig

import (
	"fmt"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/NetSepio/gateway/config/envconfig"
	migrate "github.com/NetSepio/gateway/models/Migrate"

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

func Migrate() error {
	db := GetDb()

	// db.Exec(`ALTER TABLE leader_boards DROP COLUMN IF EXISTS users;`)
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	for _, model := range []interface{}{
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
		&migrate.ScoreBoard{},
		&migrate.ActivityUnitXp{},
	} {

		if err := db.AutoMigrate(model); err != nil {
			logrus.Fatalf("failed to migrate %T: %v", model, err.Error())
			return err
		}
	}

	logrus.Info("Migrated all models")

	return nil
}
