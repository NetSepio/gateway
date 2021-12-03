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

}
