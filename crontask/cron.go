package crontask

import (
	"github.com/sirupsen/logrus"
	"github.com/robfig/cron/v3"
)

func Handler(c *cron.Cron) {
    // Добавляем задачи в cron
    c.AddFunc("* * * * *", func() {
        logrus.Info("Запуск крон задачи каждую минуту")
    })
}
