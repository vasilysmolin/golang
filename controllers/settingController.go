package controllers

import (
	"github.com/gofiber/fiber/v2"
    "main/models"
)

func Show(c *fiber.Ctx) error {

// 	var settings models.AddressesOnlineSettings
    user := new(models.User)
//     models.DB.Preload("Profile.Address").First(&user, c.Locals("user"))
//     models.DB.Where("addressID = ?", user.Profile.Address.AddressID).First(&settings)

	return c.JSON(user)
}


