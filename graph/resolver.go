package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	v25 "github.com/vinodhalaharvi/weather-api/weather/v2.5"
	v30 "github.com/vinodhalaharvi/weather-api/weather/v3.0"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	OpenWeatherClient30 *v30.OpenWeatherClient30
	OpenWeatherClient25 *v25.OpenWeatherClient25
}

func NewResolver() *Resolver {
	client25 := v25.NewOpenWeatherClient25()
	client30 := v30.NewOpenWeatherClient30()
	resolver := Resolver{
		OpenWeatherClient25: client25,
		OpenWeatherClient30: client30,
	}
	return &resolver
}

func (r *Resolver) strPtr(s string) *string {
	return &s
}
func (r *Resolver) intPtr(i int) *int {
	return &i
}
func (r *Resolver) floatPtr(f float64) *float64 {
	return &f
}
