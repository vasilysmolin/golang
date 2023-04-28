package controllers

import (
	"github.com/gofiber/fiber/v2"
    "main/models"
    "net/http"
    "strconv"
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

        positionReq := c.Query("position", "0")
        countMasterReq := c.Query("countMaster", "15")

        user, ok := c.Locals("authUser").(*models.User)
        if !ok {
            return c.Status(http.StatusUnprocessableEntity).JSON(ok)
        }

        // string to int
        position, err := strconv.Atoi(positionReq)
        if err != nil {
            panic(err)
            panic(user)
        }
        countMaster, err := strconv.Atoi(countMasterReq)
        if err != nil {
            panic(err)
        }
        skip :=  position * countMaster
        if skip < 0 {
            panic(skip)
        }

//         type Result struct {
//           UserID   int `json:"userID"`
//           AddressID   int `json:"addressID"`
//           ItemID   int `json:"itemID"`
//           Name string `json:"name"`
//           Avatar  string `json:"avatar"`
//           Sort  int `json:"sort"`
//         }
//
//         var result []Result
//         models.DB.Raw("select users.userID,name,avatar,MIN(schedule_address_user.sort) AS sort,schedule_address_user.itemID,addressID from `users` left join `schedule_address_user` on `users`.`userID` =`schedule_address_user`.`userID` where `users`.`userID` in (?) and`schedule_address_user`.`addressID` = ? and `users`.`deleted_at` is null group by `users`.`userID`,`schedule_address_user`.`itemID` order by `sort` asc", 36718, 6045).Scan(&result)

        type User struct {
//           UserID   int `json:"userID"`
          Name string `json:"name"`
        }

        var result User
        models.DB.Raw("select * from `users` where `userID` = ? and `users`.`deleted_at` is null limit 1", 36718).First(&user)

//
//        type MyData struct {
//            ID   int `gorm:"primary_key:column:id" json:"id"`
//        }
//
//        var result MyData
//        models.DB.Raw("SELECT id FROM plans WHERE id = ?", 1).First(&result)

	    return c.JSON(result)
}


