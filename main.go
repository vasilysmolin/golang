package main

import (
    "gorm.io/gorm"
    "os"
    "gorm.io/driver/mysql"
	"github.com/gofiber/fiber/v2"
)

var DB *gorm.DB

var (
	dbHost       = os.Getenv("DB_HOST")
	dbUserName   = os.Getenv("DB_USERNAME")
	dbDatabase   = os.Getenv("DB_DATABASE")
	dbPassword  = os.Getenv("DB_PASSWORD")
)


type Image struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"type:varchar(300)" json:"title"`
}


func Index(c *fiber.Ctx) error {
	var images []Image
	DB.Find(&images)

	return c.JSON(images)
}


func main() {
	ConnectDatabase()

	app := fiber.New()

	api := app.Group("/api")
	book := api.Group("/books")

	book.Get("/", Index)

	app.Listen(":4090")
}


func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open(""+dbUserName+":"+dbPassword+"@tcp("+dbHost+":3306)/"+dbDatabase+""))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Image{})
	DB = db
}

