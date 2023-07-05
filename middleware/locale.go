package middleware

import (
    "github.com/gofiber/fiber/v2"
    "main/utils"
)

func LocaleMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

         lang := c.Get("Accept-Language") // Получение значения заголовка Accept-Language

         // Определение поддерживаемых языков и выбор наиболее подходящего
         supportedLangs := utils.Langs
         i18n := "en"

         for _, str := range supportedLangs {
           if str == lang {
               i18n = lang
               break
           }
          }
         // Установка текущей локали в контексте Fiber
         c.Locals("locale", utils.CurLocale[i18n])

         return c.Next()
	}
}
