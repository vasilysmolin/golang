package web

import (
	"github.com/gofiber/fiber/v2"
	"main/utils"
)

func Index(c *fiber.Ctx) error {
   locale := c.Locals("locale").(utils.Locale)
   hello := locale["hello"]
   goodbye := locale["goodbye"]
   return c.JSON(hello + ", " + goodbye)
}



