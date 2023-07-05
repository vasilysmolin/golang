package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"main/models"
	"main/utils"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Missing Authorization Header"})
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid Authorization Header"})
		}

		tokenString := authHeaderParts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			user := new(models.User)
			utils.DB.First(&user, claims["user_id"])
			c.Locals("authUser", user)
			return c.Next()
		}

		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid Authorization Token"})
	}
}
