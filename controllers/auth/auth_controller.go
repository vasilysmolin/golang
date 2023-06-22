package auth

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"main/models"
	"os"
	"time"
	"context"
	"github.com/go-vk-api/vk"
    "golang.org/x/oauth2"
	vkAuth "golang.org/x/oauth2/vk"
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

func RegisterVk(c *fiber.Ctx) error {

	conf := &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		Scopes:       []string{"profile", "email"},
		Endpoint:     vkAuth.Endpoint,
	}

	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
    fmt.Println("url:", url)
    return c.JSON(url)

}

func VerifyVk(c *fiber.Ctx) error {

    conf := &oauth2.Config{
        ClientID:     os.Getenv("CLIENT_ID"),
        ClientSecret: os.Getenv("CLIENT_SECRET"),
        RedirectURL:  os.Getenv("REDIRECT_URL"),
        Scopes:       []string{"profile", "email"},
        Endpoint:     vkAuth.Endpoint,
    }

    code := c.Query("code", "anonymous")
    ctx := context.Background()
    token, err := conf.Exchange(ctx, code)
    if err != nil {
        fmt.Println("Ошибка при код на access токен:", err)
    }
    client, err := vk.NewClientWithOptions(vk.WithToken(token.AccessToken))
    if err != nil {
        fmt.Println("Ошибка при создаем клиента для получения данных из API VK:", err)
    }
    userVk := getCurrentUser(client)

    user := new(models.User)
    user.Name = userVk.FirstName
    user.Surname = userVk.LastName
    user.Email = "vasya.bal@mail.ru"
    user.Avatar = userVk.Photo
    result := models.DB.Create(&user)
    if result.Error != nil {
        return c.Status(fiber.StatusUnprocessableEntity).JSON(result.Error)
    }

    userSocials := new(models.UserSocials)
    userSocials.UserID = user.ID
    userSocials.AccessToken = token.AccessToken
    userSocials.SocialID = userVk.ID
    resultSoc := models.DB.Create(&userSocials)
    if resultSoc.Error != nil {
        return c.Status(fiber.StatusUnprocessableEntity).JSON(resultSoc.Error)
    }

//     res := models.DB.Preload("UserSocials").Find(&user)

    return c.JSON(user)

}

type UserVk struct {
	ID        uint64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email  string `json:"email"`
	Photo     string `json:"photo_400_orig"`
}

func getCurrentUser(api *vk.Client) UserVk  {
	var users []UserVk

	err := api.CallMethod("users.get", vk.RequestParams{
		"fields": "photo_400_orig,city,email",
	}, &users)
	fmt.Println("user:", err)
	return users[0]
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
