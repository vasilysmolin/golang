package integrations

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
)

var apiUrl string = "https://api.weather.yandex.ru/"

type weatherJson struct {
	Now   int    `json:"now"`
	NowDt string `json:"now_dt"`
	Fact  struct {
		Temp       int    `json:"temp"`
		FeelsLike  int    `json:"feels_like"`
		TempWater  int    `json:"temp_water"`
		Condition  string `json:"condition"`
		WindSpeed  int    `json:"wind_speed"`
		WindDir    string `json:"wind_dir"`
		PressureMm int    `json:"pressure_mm"`
		Humidity   int    `json:"humidity"`
		Season     string `json:"season"`
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
		logrus.Fatal(err)
	}
	req.Header.Add("X-Yandex-API-Key", apiKey)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var myResp weatherJson

	err = json.Unmarshal(body, &myResp)
	if err != nil {
		logrus.Fatal(err)
	}

	return c.JSON(myResp)

}
