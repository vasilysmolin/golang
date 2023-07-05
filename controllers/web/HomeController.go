package web

import (
	"github.com/gofiber/fiber/v2"
	"main/utils"
	"github.com/sirupsen/logrus"
)

func Index(c *fiber.Ctx) error {
   locale := c.Locals("locale").(utils.Locale)
   hello := locale["hello"]
   goodbye := locale["goodbye"]
   return c.JSON(hello + ", " + goodbye)
}



