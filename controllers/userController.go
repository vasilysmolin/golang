package controllers

import (
	"github.com/gofiber/fiber/v2"
    "main/models"
    "main/utils"
    "net/http"
    "net/url"
    "github.com/sirupsen/logrus"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/s3"
    "os"
)

func Info(c *fiber.Ctx) error {

    curUser, ok := c.Locals("authUser").(*models.User)

    file, err := os.Open("file2.jpeg")
    if err != nil {
     // Обработка ошибки
    }
    logrus.Info("file: %v", file)
    defer file.Close()

    utils.S3.PutObject(&s3.PutObjectInput{
     Bucket: aws.String(os.Getenv("AWS_BUCKET")),
     Key:    aws.String("00/80/1.jpeg"),
     Body:   file,
     ContentType: aws.String("image/jpeg"),
     ACL: aws.String("public-read"),
     Metadata: map[string]*string{
         "Cache-Control": aws.String("max-age=31536000"),
      },
    })

    req, _ := utils.S3.GetObjectRequest(&s3.GetObjectInput{
         Bucket: aws.String(os.Getenv("AWS_BUCKET")),
         Key:    aws.String("00/80/1.jpeg"),
        })
    urlStr, errS3 := req.Presign(1 * 60 * 60) // Время жизни URL в секундах
    if err != nil {
        logrus.Fatal("file: %v", errS3)
    }
     // Формирование URL без параметров GET
    parsedURL, errP := url.Parse(urlStr)
     if errP != nil {
      logrus.Fatal("errP: %v", errP)
     }
    parsedURL.RawQuery = ""
    curUser.Avatar = parsedURL.String()

    if !ok {
        return c.Status(http.StatusUnprocessableEntity).JSON(ok)
    }
	return c.JSON(curUser)
}



