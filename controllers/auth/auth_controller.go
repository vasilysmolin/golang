package auth

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"main/models"
	"os"
	"time"
	"math/rand"
// 	"strings"
	"context"
	"net/http"
	"github.com/go-vk-api/vk"
	"golang.org/x/crypto/bcrypt"
    "golang.org/x/oauth2"
	vkAuth "golang.org/x/oauth2/vk"
)

type ErrorResponse struct {
	FailedField string `json:"field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

var validate = validator.New()

func ValidateStruct(user User) []*ErrorResponse {
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

type User struct {
    ID  uint64  `json:"id"`
	Email string `json:"email" validate:"required,email,min=6,max=32"`
	Password string `json:"password" validate:"min=1,max=32"`
	Secret string `json:"secret" validate:"max=32"`
    UpdatedAt time.Time `json:"updated_at,omitempty"`
    CreatedAt time.Time `json:"created_at,omitempty"`
}

type Auth struct {
	Email string `json:"email" validate:"required,email,min=6,max=32"`
	Password string `json:"password" validate:"min=1,max=32"`
}

type JwtResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType string `json:"type_token"`
	ExpiredIn int64 `json:"expired_in"`
}

func Login(c *fiber.Ctx) error {
    // Получение переданных данных
    data := new(Auth)
    if err := c.BodyParser(data); err != nil {
    		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
    			"message": err.Error(),
    		})
    	}


    // Проверка наличия правильных данных и получение пользователя из базы данных
    user, err := getUserByEmailAndPassword(data.Email, data.Password)
    if err != nil {
        return c.Status(401).JSON(fiber.Map{
            "message": "Unauthorized",
        })
    }

    token := createTokenJwt(user.ID)
    return c.JSON(token)
}

// Получение пользователя из базы данных по электронной почте и паролю
func getUserByEmailAndPassword(email, password string) (*User, error) {
    // Получение пользователя из базы данных по адресу электронной почты
    user := new(User)
    models.DB.Where("email = ?", email).First(&user)

    // Проверка соответствия пароля пользователю
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return nil, err
    }

    return user, nil
}


func Register(c *fiber.Ctx) error {

	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := ValidateStruct(*user)
	if errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}

	password := []byte(user.Password)
    hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil {
        // обработка ошибки
    }

	user.Password = string(hash)
    user.Secret = randStr(10)
	result := models.DB.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(result.Error)
	}
	token := createTokenJwt(user.ID)
	return c.JSON(token)
}

func Refresh(c *fiber.Ctx) error {

	user, ok := c.Locals("authUser").(*models.User)
	if !ok {
            return c.Status(http.StatusUnprocessableEntity).JSON(ok)
        }
//     authHeader := c.Get("Authorization")
//     if authHeader == "" {
//         return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Missing Authorization Header"})
//     }
//
//     authHeaderParts := strings.Split(authHeader, " ")
//     if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
//         return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid Authorization Header"})
//     }
//
//     tokenString := authHeaderParts[1]
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

func AuthVk(c *fiber.Ctx) error {

    curUser, ok := c.Locals("authUser").(*models.User)
        if !ok {
            return c.Status(http.StatusUnprocessableEntity).JSON(ok)
        }

// 	conf := &oauth2.Config{
// 		ClientID:     os.Getenv("CLIENT_ID"),
// 		ClientSecret: os.Getenv("CLIENT_SECRET"),
// 		RedirectURL:  os.Getenv("REDIRECT_URL"),
// 		Scopes:       []string{"profile", "email"},
// 		Endpoint:     vkAuth.Endpoint,
// 	}
//
// 	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
//     fmt.Println("url:", url)
    return c.JSON(curUser)

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
    user.Secret = randStr(10)
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


func createTokenJwt(id uint64) JwtResponse {
	// Создаем токен
	token := jwt.New(jwt.SigningMethodHS256)
	exp := time.Now().Add(time.Hour * 24).Unix()

	// Устанавливаем параметры токена
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = id
	claims["exp"] = exp

	// Генерируем секретный ключ
	key := []byte(os.Getenv("SECRET_KEY"))

	// Подписываем токен с помощью секретного ключа
	tokenRes := new(JwtResponse)
	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println("Ошибка при создании токена:", err)
		return *tokenRes
	}

	// Создаем refresh токен
    tokenRefresh := jwt.New(jwt.SigningMethodHS256)
    expRefresh := time.Now().Add(time.Hour * 24 * 30).Unix()

    // Устанавливаем параметры токена
    claimsRefresh := tokenRefresh.Claims.(jwt.MapClaims)
    claimsRefresh["user_id"] = id
    claimsRefresh["exp"] = expRefresh

    tokenStringRefresh, errRefresh := tokenRefresh.SignedString(key)
    if errRefresh != nil {
        fmt.Println("Ошибка при создании токена:", errRefresh)
        return *tokenRes
    }

    tokenRes.AccessToken = tokenString
    tokenRes.RefreshToken = tokenStringRefresh
    tokenRes.TokenType = "bearer"
    tokenRes.ExpiredIn = exp
	return *tokenRes

}

func refreshTokenJwt(tokenString string, user *models.User) JwtResponse {

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
// 	fmt.Println("user:", user)
// 	sec := randStr(10)
// 	models.DB.(&user).Update()
// 	models.DB.(&user).Update(User{Secret: sec})

//	Если токен еще действителен, вернем его
	if time.Until(exp) > 30*time.Second {
	    jwtRes := new(JwtResponse)
	    jwtRes.AccessToken = tokenString
	    jwtRes.RefreshToken = tokenString
	    jwtRes.ExpiredIn = claims["exp"].(int64)
	    jwtRes.TokenType = "bearer"
		return *jwtRes
	}

	return createTokenJwt(claims["user_id"].(uint64))

}


var charset = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// n is the length of random string we want to generate
func randStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		// randomly select 1 character from given charset
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
