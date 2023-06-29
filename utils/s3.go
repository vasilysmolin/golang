package utils

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/aws/credentials"
	"os"
)

var S3 *s3.S3

func ConnectS3() {
    sess := session.Must(session.NewSession(&aws.Config{
     Region:           aws.String("ru-central1"),
     Endpoint:         aws.String("http://storage.yandexcloud.net/"),
     DisableSSL:       aws.Bool(false),
     S3ForcePathStyle: aws.Bool(true),
     Credentials:      credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_ACCESS_SECRET"), ""),
    }))

    svc := s3.New(sess)
	S3 = svc
}
