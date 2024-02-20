package v2_5

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vinodhalaharvi/weather-api/weather/utils"
	"io"
	"log"
	"net/http"
)

type Client interface {
	GetCurrentWeather(*CurrentWeatherRequest) (*CurrentWeatherResponse, error)
}

type OpenWeatherClient25 struct {
}

func NewOpenWeatherClient25() *OpenWeatherClient25 {
	return &OpenWeatherClient25{}
}

func (owc *OpenWeatherClient25) GetCurrentWeather(request *CurrentWeatherRequest) (*CurrentWeatherResponse, error) {
	var weatherResponse *CurrentWeatherResponse
	format := "https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&units=metric&units=%s&lang=%s"
	url := fmt.Sprintf(
		format,
		request.Lat,
		request.Lon,
		request.APIKey,
		request.Units,
		request.Lang,
	)
	fmt.Printf("URL: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return weatherResponse, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return weatherResponse, err
	}
	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		return weatherResponse, err
	}
	weatherResponse.TemperatureAssessment = utils.InterpretWeather(weatherResponse.Main.Temp, request.Units)
	return weatherResponse, nil
}

func (owc *OpenWeatherClient25) CurrentWeatherRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		lat := c.Query("lat")
		lon := c.Query("lon")
		units := c.Query("units")
		apiKey := c.Query("appid")
		lang := c.Query("lang")

		if lat == "" || lon == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "lat and lon are required query parameters"})
		}
		weatherRequest := &CurrentWeatherRequest{
			Lat:    utils.ToFloat64(lat),
			Lon:    utils.ToFloat64(lon),
			Units:  units,
			Lang:   lang,
			APIKey: apiKey,
		}
		var weather, err = owc.GetCurrentWeather(weatherRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		indent, err := json.MarshalIndent(weather, "", "  ")
		if err != nil {
			return
		}
		c.String(http.StatusOK, string(indent))
	}
}
