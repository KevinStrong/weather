package main

import (
	"fmt"
	"os"
	"weather"
)

func main() {
	apiKey := os.Getenv("WEATHER_API")
	request := weather.Request{
		ZipCode: "75080",
	}
	service := weather.New(apiKey)
	currentWeather, err := service.GetWeather(request)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", currentWeather)
}