package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestWeatherEndpoint(t *testing.T) {
	// Initialize httpmock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Mock the OpenWeather API response
	httpmock.RegisterResponder("GET", "https://api.openweathermap.org/data/2.5/weather",
		httpmock.NewStringResponder(
			200,
			`{"main": {"temp": 20.0}, "weather": [{"main": "Clouds", "description": "overcast clouds"}]}`,
		))

	// Set up Gin and the route
	gin.SetMode(gin.TestMode)
	router := gin.New()
	owc := OpenWeatherClient{APIKey: "test_api_key"}
	router.GET("/weather", WeatherRoute(owc))

	// Perform a GET request with that router
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/weather?lat=35&lon=139", nil)
	router.ServeHTTP(w, req)

	// Assert the response
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "temperature")
}
