package utils

import (
    "github.com/go-redis/redis"
    "os"
)

var RedisCon *redis.Client

func ConnectRedis() {
     client := redis.NewClient(&redis.Options{
      Addr:     os.Getenv("REDIS_HOST"),
      Password: os.Getenv("REDIS_PASSWORD"),
      DB:       0,
     })

     // Проверяем соединение
//      pong, err := client.Ping().Result()
//      logrus.Info(pong)

     // Пример выполнения запросов к Redis
//      err = client.Set("key", "value", 0).Err()
//      if err != nil {
//       logrus.Info(err)
//      }
//
//      val, errors := client.Get("key").Result()
//      if errors != nil {
//       logrus.Info(errors)
//      }
//      logrus.Info("key", val)
     RedisCon = client
}
