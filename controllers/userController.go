package controllers

import (
	"github.com/gofiber/fiber/v2"
    "main/models"
)

func Info(c *fiber.Ctx) error {

    user := new(models.User)
	return c.JSON(user)
}

func Store(c *fiber.Ctx) error {

    user := new(models.User)
	return c.JSON(user)
}


