package controllers

import (
	"github.com/gofiber/fiber/v2"
    "main/models"
)

func Show(c *fiber.Ctx) error {

    user := new(models.User)
	return c.JSON(user)
}


