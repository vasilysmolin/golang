package utils

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"main/models"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	var (
		dbHost     = os.Getenv("DB_HOST")
		dbUserName = os.Getenv("DB_USERNAME")
		dbDatabase = os.Getenv("DB_DATABASE")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbPort = os.Getenv("DB_PORT")
	)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // выводим лог в консоль
		logger.Config{
			LogLevel: logger.Info, // уровень отладки
		},
	)

	db, err := gorm.Open(mysql.Open(""+dbUserName+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbDatabase+""),
		&gorm.Config{
			Logger: newLogger,
			NamingStrategy: schema.NamingStrategy{
				NoLowerCase: false,
			},
		})
	if err != nil {
		logrus.Fatal(err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Profile{}, &models.UserSocials{}, &models.Image{})
	if err != nil {
		logrus.Fatal(err)
	}
	DB = db
}
