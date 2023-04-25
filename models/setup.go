package models

import (
	"gorm.io/driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"os"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    var (
        dbHost     = os.Getenv("DB_HOST")
        dbUserName = os.Getenv("DB_USERNAME")
        dbDatabase = os.Getenv("DB_DATABASE")
        dbPassword = os.Getenv("DB_PASSWORD")
    )
	db, err := gorm.Open(mysql.Open("" + dbUserName + ":" + dbPassword + "@tcp(" + dbHost + ":3306)/" + dbDatabase + ""))
	if err != nil {
		panic(err)
	}

// 	db.AutoMigrate(&Setting{})
// 	db.Migrator().DropColumn(&Setting{}, "name_two")
	DB = db
}



