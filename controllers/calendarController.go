package controllers

import (
	"github.com/gofiber/fiber/v2"
    "main/models"
    "net/http"
)

const unknown = "unknown"

type Error struct {
  Code    string `json:"code"`
  Message string `json:"message"`
 }

func Calendar(c *fiber.Ctx) error {

         errorMap := map[string][]Error{
          "errors": []Error{
           Error{
            Code:    "0",
            Message: "Поле обязательное",
           },
          },
         }

        addressID := c.Query("addressID", unknown)
        if addressID == "" {
            addressID = unknown
        }
        if addressID == unknown {
           return c.Status(http.StatusUnprocessableEntity).JSON(errorMap)
        }

        view := c.Query("view", unknown)
        if view == "" {
            view = unknown
        }
        if view == unknown {
           return c.Status(http.StatusUnprocessableEntity).JSON(errorMap)
        }

        date := c.Query("date", unknown)
        if date == "" {
            view = unknown
        }
        if date == unknown {
           return c.Status(http.StatusUnprocessableEntity).JSON(errorMap)
        }

//         userID := c.Query("userID", unknown)
//         if userID == "" {
//             view = unknown
//         }
//         if userID == unknown {
//            return c.Status(http.StatusUnprocessableEntity).JSON(errorMap)
//         }

//         position := c.Query("position", "0")
//         countMaster := c.Query("countMaster", "15")

//         user, ok := c.Locals("authUser").(*models.User)
//         if !ok {
//             return c.Status(http.StatusUnprocessableEntity).JSON(ok)
//         }

      user := new(models.User)
      models.DB.Preload("Studio").Preload("Profile").First(&user, 152)

	  return c.JSON(user)
}


