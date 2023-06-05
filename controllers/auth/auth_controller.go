package auth

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"main/models"
	"os"
	"time"
)

type ErrorResponse struct {
	FailedField string `json:"field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

var validate = validator.New()

func ValidateStruct(user models.User) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func Register(c *fiber.Ctx) error {

	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := ValidateStruct(*user)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	result := models.DB.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(result.Error)
	}
	token := createTokenJwt(user.ID)
	return c.JSON(token)
}

func createTokenJwt(id uint64) string {
	// Создаем токен
	token := jwt.New(jwt.SigningMethodHS256)

	// Устанавливаем параметры токена
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Генерируем секретный ключ
	key := []byte(os.Getenv("SECRET_KEY"))

	// Подписываем токен с помощью секретного ключа
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println("Ошибка при создании токена:", err)
		return ""
	}

	return tokenString

}

func refreshTokenJwt(tokenString string) string {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что тип токена - JWT
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token")
		}

		// Возвращаем ключ для проверки подписи токена
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		// Обработка ошибки
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		// Обработка ошибки
	}

	exp := time.Unix(int64(claims["exp"].(float64)), 0)

	// Если токен еще действителен, вернем его
	if time.Until(exp) > 30*time.Second {
		return tokenString
	}

	return createTokenJwt(claims["user_id"].(uint64))

}
