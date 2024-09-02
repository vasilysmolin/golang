# go-restapi-fiber

## Golang REST API menggunakan Fiber, Gorm, MySQL
Youtube: [Golang REST API menggunakan Fiber, Gorm, MySQL](https://youtu.be/X4USU6GRP2g)

Local setup:
- docker-compose up -d
- make start

Build prod:
- make build
- подготовить .env-go для переменных окружения

Add packet:
- go get -u github.com/spf13/cobra@latest
- import "github.com/spf13/cobra"

Console commands
- go run main.go goodbye (запускаются если есть переданный параметр)
- DB CurLocale RedisCon S3 (доступны глобально)
