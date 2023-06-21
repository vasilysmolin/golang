package routes

import (
    "github.com/gofiber/fiber/v2"
	"main/controllers"
	"main/controllers/auth"
	"main/middleware"
)

func SetupRoutes(app *fiber.App) {

    api := app.Group("/api")

    authGrp := api.Group("/auth")
    authGrp.Post("/register", auth.Register)

    vkGrp := authGrp.Group("/vk")
	vkGrp.Get("/", auth.RegisterVk)
	vkGrp.Get("/verify", auth.VerifyVk)

	usersGrp := api.Group("/users", middleware.AuthMiddleware())
	usersGrp.Get("/info", controllers.Info)
}

