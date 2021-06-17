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
		request = getLocation()
		service := weather.New(apiKey)
		currentWeather, err := service.GetWeather(request)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%v\n", currentWeather)
	}
}

func getLocation() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	return text
}
