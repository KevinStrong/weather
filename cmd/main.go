package main

import (
	"fmt"
	"os"
	"weather"
)

func main() {
	// todo verify api key and prompt user if not here
	// move logic into weather package
	apiKey := os.Getenv("WEATHER_API")
	service := weather.New(apiKey)
	location := service.GetLocation(os.Args[1:])
	currentWeather, err := service.GetWeather(location)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Current weather in %v is %v\n", location, currentWeather)
}
