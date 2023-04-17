package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:1111@tcp(localhost:3306)/beautybox"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Book{})
	DB = db
}
