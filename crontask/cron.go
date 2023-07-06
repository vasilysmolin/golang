package crontask

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func Handler(c *cron.Cron) {
	// Добавляем задачи в cron
	_, err := c.AddFunc("* * * * *", func() {
		logrus.Info("Запуск крон задачи каждую минуту")
	})
	if err != nil {
		logrus.Fatal(err)
	}
}
