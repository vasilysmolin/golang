package main

import (
	"github.com/vasilysmolin/fiber-rest-api/controllers/bookcontroller"
	"github.com/vasilysmolin/fiber-rest-api/controllers/imagecontroller"

	"github.com/gofiber/fiber/v2"
	"github.com/vasilysmolin/fiber-rest-api/models"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()

	api := app.Group("/api")
	book := api.Group("/books")
	images := api.Group("/images")

	book.Get("/", bookcontroller.Index)
	book.Get("/:id", bookcontroller.Show)
	book.Post("/", bookcontroller.Create)
	book.Put("/:id", bookcontroller.Update)
	book.Delete("/:id", bookcontroller.Delete)

//     images.Get("/", imagecontroller.Index)
//     images.Get("/:id", imagecontroller.Show)
//     images.Post("/", imagecontroller.Create)
//     images.Put("/:id", imagecontroller.Update)
//     images.Delete("/:id", imagecontroller.Delete)

	app.Listen(":4090")
}
