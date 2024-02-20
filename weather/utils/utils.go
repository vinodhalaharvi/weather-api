package utils

import (
	"log"
	"strconv"
)

func ToFloat64(lat string) float64 {
	latFloat, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		log.Printf("Error parsing lat: %v", err)
	}
	return latFloat

}

func InterpretWeather(temp float64, units string) string {
	var tempAssessment string

	if units == "metric" {
		switch {
		case temp < 10: // Less than 10°C is considered cold
			tempAssessment = "cold"
		case temp > 25: // More than 25°C is considered hot
			tempAssessment = "hot"
		default:
			tempAssessment = "moderate"
		}
	} else if units == "imperial" {
		switch {
		case temp < 50: // Less than 50°F is considered cold
			tempAssessment = "cold"
		case temp > 77: // More than 77°F is considered hot
			tempAssessment = "hot"
		default:
			tempAssessment = "moderate"
		}
	} else {
		// Handle unexpected units by setting a default or logging an error
		tempAssessment = "unknown"
	}
	return tempAssessment
}
