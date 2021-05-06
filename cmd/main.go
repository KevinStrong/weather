package main

import (
	"fmt"
	"net/http"
	"os"
	"weather"
)

func main() {
	apiKey := os.Getenv("WEATHER_API")
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/forecast?zip=75080&appid=" + apiKey)
	if err != nil {
		fmt.Print(err)
		return
	}

	defer resp.Body.Close()
	toStruct, err := weather.ConvertWeatherOpenApiResponseToStruct(resp.Body)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("%v", toStruct)
}