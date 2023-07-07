package main

import (
	"github.com/sirupsen/logrus"
	"main/bootstrap"
	"os"
)

func main() {
	bootstrap.SetupApp(".")
	logrus.Printf(os.Getenv("APP_NAME"))
}
