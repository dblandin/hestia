package commands

import (
	"encoding/json"
	"fmt"
	"github.com/codeclimate/hestia/internal/notifiers"
	"github.com/codeclimate/hestia/internal/secrets"
	"github.com/codeclimate/hestia/internal/types"
	"io/ioutil"
	"log"
	"net/http"
)

type Weather struct {
	User     string
	Input    types.Input
	Notifier notifiers.Notifier
}

type Forecast struct {
	City        string `json:"city"`
	Zip         string `json:"zip"`
	Description string `json:"description"`
}

type WeatherResponse struct {
	Main        string `json:"main"`
	Description string `json"description"`
}

type OpenWeatherMapResponse struct {
	Name    string            `json:"name"`
	Weather []WeatherResponse `json:"weather"`
}

func getWeather(zip string) (weather OpenWeatherMapResponse) {
	weather_url := fmt.Sprintf(
		"http://api.openweathermap.org/data/2.5/weather?zip=%s,us&appid=%s",
		zip,
		secrets.GetSecretValue("open_weather_api_key"),
	)

	resp, err := http.Get(weather_url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	weather = OpenWeatherMapResponse{}
	_ = json.Unmarshal(bodyBytes, &weather)

	return weather
}

func (c Weather) Run() {
	zip := c.Input.Args

	if len(zip) == 0 {
		zip = "10011"
	}

	weather := getWeather(zip)

	forecast := Forecast{
		City:        weather.Name,
		Zip:         zip,
		Description: weather.Weather[0].Description,
	}

	message := fmt.Sprintf("%s in %s (%s)", forecast.Description, forecast.City, forecast.Zip)

	c.Notifier.Log(message)
}
