package main

import (
	"bufio"
	"fmt"
	"os"

	"weather"
)

func main() {
	apiKey := os.Getenv("WEATHER_API")

	request := ""

	for request != "exit" {
		fmt.Print("Enter in a location to get it's weather, or \"exit\" to exit")
		request = getLocation()
		service := weather.New(apiKey)
		currentWeather, err := service.GetWeather(request)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Current weather in %v is %v\n", request, currentWeather)
	}
}

func getLocation() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	return text
}
