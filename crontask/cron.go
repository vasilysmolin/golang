package crontask

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func Handler(c *cron.Cron) {
	// Добавляем задачи в cron
	c.AddFunc("* * * * *", func() {
		logrus.Info("Запуск крон задачи каждую минуту")
	})
}
