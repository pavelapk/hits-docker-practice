package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"go-app/models"
	"os"
)

var db *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[DB] Error loading .env file")
	}
	username := os.Getenv("db_user")
	//password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	conn, err := gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
			dbHost, dbPort, username, dbName))
	if err != nil {
		panic(err)
	} else {
		db = conn
		log.Info("[DB] Connected on " + dbHost)
	}

	err = migrateSchema()
	if err != nil {
		log.Fatal("[DB] Error migrating schema: " + err.Error())
	} else {
		log.Info("[DB] Schema migrated successfully")
	}
}

func GetDB() *gorm.DB {
	return db
}

func migrateSchema() error {
	err := db.AutoMigrate(
		models.Note{},
	).Error

	return err
}
