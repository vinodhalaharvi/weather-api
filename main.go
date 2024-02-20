package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/vinodhalaharvi/weather-api/graph"
	v25 "github.com/vinodhalaharvi/weather-api/weather/v2.5"
	v30 "github.com/vinodhalaharvi/weather-api/weather/v3.0"
	"log"
)

func main() {
	r := gin.Default()
	weather25 := v25.NewOpenWeatherClient25()
	weather30 := v30.NewOpenWeatherClient30()

	// Initialize the GraphQL server
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver()}))

	// Set up the GraphQL playground handler
	r.GET("/graphql", func(c *gin.Context) {
		playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Writer, c.Request)
	})

	// Set up the GraphQL query handler
	r.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	// Define the endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/weather/v25/currentWeather", weather25.CurrentWeatherRoute())
	r.GET("/weather/v30/timeMachine", weather30.TimeMachineRoute())
	r.GET("/weather/v30/dailySummary", weather30.DailySummaryRoute())
	r.GET("/weather/v30/currentWeather", weather30.CurrentWeatherRoute())

	err := r.Run(":8080")
	if err != nil {
		log.Printf("Error: %v", err)
	}
}
