package main

import (
	"fmt"
	"net/http"
	"os"
	"weather"
)

func main() {
	apiKey := os.Getenv("WEATHER_API")
	request := weather.Request{
		ZipCode: "75080",
	}
	requestUrl, err := weather.ConvertOurRequestStructToOpenApiRequest(request, apiKey)
	if err != nil {
		panic(err) // okay now can I panic here?
	}
	resp, err := http.Get(requestUrl)
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