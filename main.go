package main

import (
	"github.com/sirupsen/logrus"
	"main/bootstrap"
	"os"
)

func main() {
	app := bootstrap.SetupApp()
	logrus.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}
