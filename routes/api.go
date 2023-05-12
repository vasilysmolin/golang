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
    usersGrp := api.Group("/users", middleware.AuthMiddleware())

	authGrp.Post("/register", auth.Register)

	usersGrp.Get("/info", controllers.Info)
}

