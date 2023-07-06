package bootstrap

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
	cron "github.com/robfig/cron/v3"
	"log"
	"main/crontask"
	"main/routes"
	"main/utils"
	"path"
	"time"
)

func SetupApp(root string) *fiber.App {
	envPath := path.Join(root, ".env")
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	utils.ConnectDatabase()
	// 	utils.ConnectRedis()
	utils.ConnectS3()
	utils.GetLocales(root)

	app := fiber.New(fiber.Config{
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
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

	routes.SetupRoutes(app)

	// Создаем экземпляр cron
	c := cron.New()
	crontask.Handler(c)

	// Запускаем cron
	c.Start()

	return app
}
