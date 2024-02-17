package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type WeatherClient interface {
	GetWeather(lat, lon string) (WeatherResponse, error)
}

type OpenWeatherClient struct {
	APIKey string
}

func (owc OpenWeatherClient) GetWeather(lat, lon string) (WeatherResponse, error) {
	var weatherResponse WeatherResponse
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=metric", lat, lon, owc.APIKey)
	resp, err := http.Get(url)
	if err != nil {
		return weatherResponse, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	weatherResponse.TemperatureAssessment = interpretWeather(body)
	if err != nil {
		return weatherResponse, err
	}

	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		return weatherResponse, err
	}

	return weatherResponse, nil
}

func WeatherRoute(owc WeatherClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		lat := c.Query("lat")
		lon := c.Query("lon")
		if lat == "" || lon == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "lat and lon are required query parameters"})
		}
		weather, err := owc.GetWeather(lat, lon)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		indent, err := json.MarshalIndent(weather, "", "  ")
		if err != nil {
			return
		}
		//c.JSON(http.StatusOK, gin.H{
		//	"temperature": weather.Main.Temp,
		//	"description": weather.Weather[0].Description,
		//})
		c.String(http.StatusOK, string(indent))
	}
}

type WeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		Id      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone              int    `json:"timezone"`
	Id                    int    `json:"id"`
	Name                  string `json:"name"`
	Cod                   int    `json:"cod"`
	TemperatureAssessment string
}

// WeatherData WeatherResponse Define a struct to unmarshal the OpenWeather API response
type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
}

func interpretWeather(data []byte) string {
	var weatherData WeatherData
	err := json.Unmarshal(data, &weatherData)
	if err != nil {
		fmt.Println("Error parsing data:", err)
		return ""
	}

	// Determine the weather condition
	//weatherCondition := "clear"
	//if len(weatherData.Weather) > 0 {
	//	weatherCondition = weatherData.Weather[0].Main
	//}

	// Assess temperature
	temp := weatherData.Main.Temp
	var tempAssessment string
	switch {
	case temp < 10:
		tempAssessment = "cold"
	case temp > 25:
		tempAssessment = "hot"
	// No need for a default case since we've already set tempAssessment to "moderate"
	default:
		tempAssessment = "moderate"
	}

	// Print the interpretation
	//fmt.Printf(
	//	"In %s, the weather is currently %s with %s conditions.\n",
	//	weatherData.Name,
	//	tempAssessment,
	//	weatherCondition,
	//)
	return tempAssessment
}
