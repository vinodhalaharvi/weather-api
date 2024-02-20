package v30

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/vinodhalaharvi/weather-api/weather/utils"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Client interface {
	GetCurrentWeather(*CurrentWeatherRequest) (*CurrentWeatherResponse, error)
	GetDaySummary(*DaySummaryRequest) (*DaySummaryResponse, error)
	GetTimeMachineWeather(*TimeMachineRequest) (*TimeMachineResponse, error)
}

type OpenWeatherClient30 struct {
}

func NewOpenWeatherClient30() *OpenWeatherClient30 {
	return &OpenWeatherClient30{}
}

func (owc *OpenWeatherClient30) GetDaySummary(request *DaySummaryRequest) (*DaySummaryResponse, error) {
	var weatherResponse *DaySummaryResponse
	format := "https://api.openweathermap.org/data/3.0/onecall/day_summary?lat=%f&lon=%f&date=%s&appid=%s&lang=%s&exclude=%s&units=%s"
	url := fmt.Sprintf(
		format,
		request.Lat,
		request.Lon,
		request.Date,
		request.APIKey,
		request.Lang,
		request.Exclude,
		request.Units,
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
	return weatherResponse, nil
}

func (owc *OpenWeatherClient30) GetTimeMachineWeather(request *TimeMachineRequest) (*TimeMachineResponse, error) {
	var weatherResponse *TimeMachineResponse
	format := "https://api.openweathermap.org/data/3.0/onecall/timemachine?lat=%f&lon=%f&dt=%d&appid=%s&lang=%s&units=%s"
	url := fmt.Sprintf(
		format,
		request.Lat,
		request.Lon,
		request.Dt,
		request.APIKey,
		request.Lang,
		request.Units,
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
	return weatherResponse, nil
}

func (owc *OpenWeatherClient30) GetCurrentWeather(request *CurrentWeatherRequest) (*CurrentWeatherResponse, error) {
	var weatherResponse *CurrentWeatherResponse
	format := "https://api.openweathermap.org/data/3.0/onecall?lat=%f&lon=%f&appid=%s&lang=%s&units=%s"
	url := fmt.Sprintf(
		format,
		request.Lat,
		request.Lon,
		request.APIKey,
		request.Lang,
		request.Units,
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
	weatherResponse.TemperatureAssessment = utils.InterpretWeather(weatherResponse.Current.Temp, request.Units)
	return weatherResponse, nil
}

func (owc *OpenWeatherClient30) CurrentWeatherRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		lat := c.Query("lat")
		lon := c.Query("lon")
		lang := c.Query("lang")
		units := c.Query("units")
		apiKey := c.Query("appid")
		if lat == "" || lon == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "lat and lon are required query parameters"})
		}
		weatherRequest := &CurrentWeatherRequest{
			Lat:    toFloat64(lat),
			Lon:    toFloat64(lon),
			Lang:   lang,
			Units:  units,
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

func (owc *OpenWeatherClient30) TimeMachineRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		lat := c.Query("lat")
		lon := c.Query("lon")
		dt := c.Query("dt")
		apiKey := c.Query("appid")
		lang := c.Query("lang")
		units := c.Query("units")
		if lat == "" || lon == "" || dt == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "lat, lon, and dt are required query parameters"})
		}
		weatherRequest := &TimeMachineRequest{
			Lat:    toFloat64(lat),
			Lon:    toFloat64(lon),
			APIKey: apiKey,
			Lang:   lang,
			Units:  units,
			Dt:     toInt(dt),
		}
		var weather, err = owc.GetTimeMachineWeather(weatherRequest)
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

func (owc *OpenWeatherClient30) DailySummaryRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		lat := c.Query("lat")
		lon := c.Query("lon")
		date := c.Query("date")
		units := c.Query("units")
		apiKey := c.Query("appid")
		exclude := c.Query("exclude")
		lang := c.Query("lang")

		if lat == "" || lon == "" || date == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "lat, lon, and date are required query parameters"})
		}
		weatherRequest := &DaySummaryRequest{
			Lat:     toFloat64(lat),
			Lon:     toFloat64(lon),
			APIKey:  apiKey,
			Units:   units,
			Lang:    lang,
			Date:    date,
			Exclude: exclude,
		}
		var weather, err = owc.GetDaySummary(weatherRequest)
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

func toInt(dt string) int {
	dtInt, err := strconv.Atoi(dt)
	if err != nil {
		log.Printf("Error parsing dt: %v", err)
	}
	return dtInt
}

func toFloat64(lat string) float64 {
	latFloat, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		log.Printf("Error parsing lat: %v", err)
	}
	return latFloat
}
