package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"main/bootstrap"
)


func main() {
    app := bootstrap.SetupApp()
	logrus.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}
