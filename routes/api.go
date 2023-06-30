package routes

import (
    "github.com/gofiber/fiber/v2"
	"main/controllers"
	"main/controllers/auth"
	"main/controllers/integrations"
	"main/middleware"
)

func SetupRoutes(app *fiber.App) {

    api := app.Group("/api")

    authGrp := api.Group("/auth")
    yandexGrp := api.Group("/yandex")
    authGrp.Post("/register", auth.Register)
    authGrp.Post("/refresh", middleware.AuthMiddleware(), auth.Refresh)
    authGrp.Post("/login", auth.Login)

    vkGrp := authGrp.Group("/vk")
	vkGrp.Get("/", auth.RegisterVk)
	vkGrp.Get("/verify", auth.VerifyVk)

	usersGrp := api.Group("/users", middleware.AuthMiddleware())
	usersGrp.Get("/info", controllers.Info)


	yandexGrp.Get("/weather", integrations.Weather)
}

