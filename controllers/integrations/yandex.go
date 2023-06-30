package integrations

import (
 "fmt"
 "net/http"
 "os"
 "github.com/gofiber/fiber/v2"
)

func Weather(c *fiber.Ctx) error {

  apiKey := os.Getenv("YANDEX")
  url := fmt.Sprintf("https://api.weather.yandex.ru/v2/informers?lat=55.75396&lon=37.620393")

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
  }
  req.Header.Add("X-Yandex-API-Key", apiKey)

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
  }
  defer resp.Body.Close()


  return c.JSON(resp.Body)

}

