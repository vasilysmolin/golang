package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/sirupsen/logrus"
	"io"
	"main/helpers"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

type File struct {
	Name      string
	Extension string
	Size      int64
	MimeType  string
}

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

func SaveAvatarByPath(pathFile string) File {
	tmpName := CreateName()

	file, err := os.Create("storage/" + tmpName + ".jpeg")
	if err != nil {
		logrus.Fatal(err)
	}
	defer file.Close()

	// Получаем содержимое фото
	resp, err := http.Get(pathFile)
	if err != nil {
		logrus.Fatal(err)
	}
	defer resp.Body.Close()

	// Копируем содержимое ответа в файл
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		logrus.Fatal(err)
	}

	fileLocal, err := os.Open("storage/" + tmpName + ".jpeg")

	fileInfo, errOs := os.Stat("storage/" + tmpName + ".jpeg")
	if errOs != nil {
		logrus.Fatal(errOs)
	}

	re := regexp.MustCompile(`\?.*`)

	nameFile := CreateName()
	randPathFile := CreatePath(nameFile)
	fileStruct := File{
		Name:      nameFile,
		Extension: re.ReplaceAllString(filepath.Ext(pathFile), ""),
		MimeType:  getFileMimeType(pathFile),
		Size:      fileInfo.Size(),
	}

	_, err = S3.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("AWS_BUCKET")),
		Key:         aws.String(randPathFile + nameFile + fileStruct.Extension),
		Body:        fileLocal,
		ContentType: aws.String(fileStruct.MimeType),
		ACL:         aws.String("public-read"),
		Metadata: map[string]*string{
			"Cache-Control": aws.String("max-age=31536000"),
		},
	})
	if err != nil {
		logrus.Fatal(err)
	}

	err = os.Remove("storage/" + tmpName + ".jpeg")
	if err != nil {
		logrus.Fatal(err)
	}

	return fileStruct
}

func CreateName() string {
	return helpers.RandStr(80)
}

func CreatePath(name string) string {
	return name[:2] + "/" + name[2:4] + "/" + name[4:6] + "/" + name + "/"
}

func getFileMimeType(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return "image/jpeg"
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "image/jpeg"
	}
	return http.DetectContentType(buffer)
}
