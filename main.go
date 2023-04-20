package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"fmt"
    "net/http"
    "strings"
)

var DB *gorm.DB

var (
	dbHost     = os.Getenv("DB_HOST")
	dbUserName = os.Getenv("DB_USERNAME")
	dbDatabase = os.Getenv("DB_DATABASE")
	dbPassword = os.Getenv("DB_PASSWORD")
)

func AuthMiddleware() func(*fiber.Ctx) error {
 return func(c *fiber.Ctx) error {
  authHeader := c.Get("Authorization")
  if authHeader == "" {
   return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Missing Authorization Header"})
  }

  authHeaderParts := strings.Split(authHeader, " ")
  if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
   return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid Authorization Header"})
  }

  tokenString := authHeaderParts[1]
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
   if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
    return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
   }

   return []byte(os.Getenv("SECRET_KEY")), nil
  })

  if err != nil {
   return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
  }

  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
   c.Locals("user", claims["user"])
   return c.Next()
  }

  return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid Authorization Token"})
 }
}

type Image struct {
	Id   int64  `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(300)" json:"title"`
}

func Index(c *fiber.Ctx) error {

	var images []Image
	DB.Find(&images)

	return c.JSON(images)
}

func main() {
	ConnectDatabase()
	a := []int{12, 8, 22, 11, 1, 3}
    InsertSort(a)
	app := fiber.New()

	api := app.Group("/api")
	book := api.Group("/books")

	book.Get("/", AuthMiddleware() , Index)

	app.Listen(":4090")
}


func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("" + dbUserName + ":" + dbPassword + "@tcp(" + dbHost + ":3306)/" + dbDatabase + ""))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Image{})
	DB = db
}

func InsertSort(a []int) {
	for i := 1; i < len(a); i++ {
		j := i
		for ; j > 0 && a[i] < a[j-1]; j-- {
		}
		for ; i > j; i-- {
			a[i], a[i-1] = a[i-1], a[i]
		}
	}
}

