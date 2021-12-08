package db

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"netsepio-api/models"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var Db *gorm.DB

func InitDB() {

	var (
		host     = os.Getenv("DB_HOST")
		username = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
		port     = os.Getenv("DB_PORT")
	)

	psqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%s",
		host, username, password, dbname, port)
	var err error
	Db, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	if err = Db.DB().Ping(); err != nil {
		log.Fatal("failed to ping database", err)
	}
	if err := Db.AutoMigrate(&models.FlowId{}, &models.User{}, &models.Role{}).Error; err != nil {
		log.Fatal(err)
	}
	//Create user_roles table
	Db.Exec(`create table if not exists user_roles (
			wallet_address text,
			role_id int,
			unique (wallet_address,role_id)
			)`)

	rolesToBeAdded := []models.Role{
		{Name: "Investor", RoleId: 1, Eula: "TODO Investor EULA"},
		{Name: "Manager", RoleId: 2, Eula: "TODO Manager EULA"}}
	for _, role := range rolesToBeAdded {
		if err := Db.Model(&models.Role{}).FirstOrCreate(&role).Error; err != nil {
			log.Fatal(err)
		}
	}

}
