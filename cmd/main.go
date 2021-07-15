package main

import (
	"fmt"
	"os"
	"weather"
)

func main() {
	// todo  pull all of this into a "run" method in weather
	service, err := weather.New(weather.WithConfigFromEnv())
	if err != nil {
		fmt.Printf("Please set environment variable WEATHER API key with our open weather api key value")
		os.Exit(1)
	}
	location := service.GetLocation(os.Args[1:])
	currentWeather, err := service.GetWeather(location)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Printf("Current weather in %v is %v\n", location, currentWeather)
}
