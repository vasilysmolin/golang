package routes

import (
	"github.com/gofiber/fiber/v2"
	"main/controllers"
	"main/controllers/auth"
	"main/controllers/integrations"
	"main/controllers/web"
	"main/middleware"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", middleware.LocaleMiddleware())
	webApp := app.Group("/", middleware.LocaleMiddleware())

	webApp.Get("/", web.Index).Name("index.main")

	authGrp := api.Group("/auth")
	yandexGrp := api.Group("/yandex")

	authGrp.Post("/register", auth.Register).Name("auth.register")
	authGrp.Post("/refresh", middleware.AuthMiddleware(), auth.Refresh).Name("auth.refresh")
	authGrp.Post("/login", auth.Login).Name("auth.login")

	vkGrp := authGrp.Group("/vk")
	vkGrp.Get("/", auth.RegisterVk).Name("auth.vk.login")
	vkGrp.Get("/verify", auth.VerifyVk).Name("auth.vk.verify")

	usersGrp := api.Group("/users", middleware.AuthMiddleware())
	usersGrp.Get("/info", controllers.Info).Name("users.info")

	yandexGrp.Get("/weather", integrations.Weather).Name("yandex.weather")
}
