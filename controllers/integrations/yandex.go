package integrations

import (
 "fmt"
 "net/http"
 "os"
 "io/ioutil"
 "github.com/gofiber/fiber/v2"
 "encoding/json"
)

var apiUrl string = "https://api.weather.yandex.ru/"

type weatherJson struct {
    Now int  `json:"now"`
    NowDt string  `json:"now_dt"`
    Fact struct {
        Temp int `json:"temp"`
        FeelsLike int `json:"feels_like"`
        TempWater int `json:"temp_water"`
        Condition string `json:"condition"`
        WindSpeed int `json:"wind_speed"`
        WindDir string `json:"wind_dir"`
        PressureMm int `json:"pressure_mm"`
        Humidity int `json:"humidity"`
        Season string `json:"season"`
    }
    Forecast struct {
        Date string `json:"date"`
    }
}

func Weather(c *fiber.Ctx) error {

  lat := c.Query("lat", "anonymous")
  lon := c.Query("lon", "anonymous")
  apiKey := os.Getenv("YANDEX")
  url := fmt.Sprintf(apiUrl + "v2/informers?" + "lat=" + lat + "&" + "lon=" + lon)

  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
  }
  req.Header.Add("X-Yandex-API-Key", apiKey)
  req.Header.Set("Accept", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  var myResp weatherJson

  err = json.Unmarshal(body, &myResp)
  if err != nil {
  }

  return c.JSON(myResp)

}

