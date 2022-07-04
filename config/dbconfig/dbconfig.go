package dbconfig

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	log "github.com/sirupsen/logrus"

	"github.com/NetSepio/gateway/config/netsepio"
	"github.com/NetSepio/gateway/models"
	"github.com/NetSepio/gateway/util/pkg/envutil"
	"github.com/NetSepio/gateway/util/pkg/logwrapper"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// Return singleton instance of db, initiates it before if it is not initiated already
func GetDb() *gorm.DB {
	if db != nil {
		return db
	}
	var (
		host     = envutil.MustGetEnv("DB_HOST")
		username = envutil.MustGetEnv("DB_USERNAME")
		password = envutil.MustGetEnv("DB_PASSWORD")
		dbname   = envutil.MustGetEnv("DB_NAME")
		port     = envutil.MustGetEnv("DB_PORT")
	)

	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%s",
		host, username, password, dbname, port)
	var err error
	db, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	if err = db.DB().Ping(); err != nil {
		log.Fatal("failed to ping database", err)
	}
	if err := db.AutoMigrate(&models.UserFeedback{}, &models.FlowId{}, &models.User{}, &models.Role{}).Error; err != nil {
		log.Fatal(err)
	}
	//Create user_roles table
	db.Exec(`create table if not exists user_roles (
			wallet_address text,
			role_id text,
			unique (wallet_address,role_id)
			)`)

	//Create flow id
	db.Exec(`
	DO $$ BEGIN
		CREATE TYPE flow_id_type AS ENUM (
			'AUTH',
			'ROLE');
	EXCEPTION
    	WHEN duplicate_object THEN null;
	END $$;`)

	voterRoleId, err := netsepio.GetRole(netsepio.VOTER_ROLE)
	if err != nil {
		logwrapper.Fatal(err)
	}

	voterEula := envutil.MustGetEnv("VOTER_EULA")

	rolesToBeAdded := []models.Role{
		{Name: "Voter Role", RoleId: hexutil.Encode(voterRoleId[:]), Eula: voterEula}}
	for _, role := range rolesToBeAdded {
		if err := db.Model(&models.Role{}).FirstOrCreate(&role).Error; err != nil {
			log.Fatal(err)
		}
	}
	return db
}
