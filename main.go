package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gofiber/template/html"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"log"
	"fmt"
    "net/http"
    "strings"
    "time"
)

var DB *gorm.DB

type Film struct {
    Title    string
    IsViewed bool
}

var films = []Film{
    {
        Title:    "The Shawshank Redemption",
        IsViewed: true,
    },
    {
        Title:    "The Godfather",
        IsViewed: true,
    },
    {
        Title:    "The Godfather: Part II",
        IsViewed: false,
    },
}

func main() {

	ConnectDatabase()
	viewsEngine := html.New("./template", ".tmpl")
    app := fiber.New(fiber.Config{
        Views: viewsEngine,
        ReadTimeout:  3 * time.Second,
        WriteTimeout: 3 * time.Second,
    })
    app.Use(recover.New())

    app.Use(requestid.New())
    app.Use(logger.New(logger.Config{
        Format:     "${locals:requestid}: ${time} ${method} ${path} - ${status} - ${latency}\n",
        TimeFormat: "2023-01-01 15:04:05.000000",
    }))

    app.Use(limiter.New(limiter.Config{
        KeyGenerator: func(c *fiber.Ctx) string {
            return c.IP()
        },
        Max:        200,
        Expiration: 60 * time.Second,
    }))

    app.Get("/profile", func(c *fiber.Ctx) error {
        return c.Render("profile", fiber.Map{
            "name":  "John",
            "email": "john@doe.com",
        })
    })

    app.Get("/films", func(c *fiber.Ctx) error {
            return c.Render("film-list", films)
    })

	api := app.Group("/api")
	images := api.Group("/images")

	images.Get("/", AuthMiddleware() , Index)

	logrus.Fatal(app.Listen(":4090"))
}



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
   c.Locals("user", claims["UserID"])
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
    tokenString := c.Get("Authorization")[7:]
    token, _ := jwt.Parse(tokenString, nil)
    claims := token.Claims.(jwt.MapClaims)
    userID := claims["userID"]

	return c.JSON(userID)
}


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

	db.AutoMigrate(&Image{})
	DB = db
}

// func InsertSort(a []int) {
// 	for i := 1; i < len(a); i++ {
// 		j := i
// 		for ; j > 0 && a[i] < a[j-1]; j-- {
// 		}
// 		for ; i > j; i-- {
// 			a[i], a[i-1] = a[i-1], a[i]
// 		}
// 	}
// }


