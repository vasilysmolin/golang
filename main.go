package main

import (
	"github.com/sirupsen/logrus"
	"main/bootstrap"
	"os"
)

func main() {
	root := "."
	app := bootstrap.SetupApp(root)
	logrus.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}
