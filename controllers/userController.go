package controllers

import (
	"github.com/gofiber/fiber/v2"
    "main/models"
    "net/http"
)

func Info(c *fiber.Ctx) error {

    curUser, ok := c.Locals("authUser").(*models.User)
    if !ok {
        return c.Status(http.StatusUnprocessableEntity).JSON(ok)
    }
	return c.JSON(curUser)
}



