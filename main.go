package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/sirupsen/logrus"
	"github.com/joho/godotenv"
    "time"
    "log"
    "main/models"
    "main/middleware"
    "main/controllers"
)

func main() {

    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

	models.ConnectDatabase()
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

	api := app.Group("/api/crm")
	settings := api.Group("/settings")

	settings.Get("/online/info", middleware.AuthMiddleware() ,  controllers.Show)
	api.Get("/calendar", middleware.AuthMiddleware() ,  controllers.Calendar)
	logrus.Fatal(app.Listen(":4091"))
}








