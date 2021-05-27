package weather

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	weatherapi "weather/gen"
)

// CurrentWeather is part of weather api response
type CurrentWeather struct {
	Temp float64
}

type Request struct {
	ZipCode string
}

type Service struct {
	ApiKey string
	client *weatherapi.Client
}

func (s *Service) GetWeather(request Request) (*CurrentWeather, error) {
	zip := weatherapi.Zip(request.ZipCode)
	unit := weatherapi.CurrentWeatherDataParamsUnits("imperial")
	req := &weatherapi.CurrentWeatherDataParams{Zip: &zip, Units: &unit}
	response, err := s.client.CurrentWeatherData(context.Background(), req)
	if err != nil {
		return nil, err
	}
	weather := &weatherapi.N200{}
	err = json.NewDecoder(response.Body).Decode(weather)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	toStruct := CurrentWeather{
		Temp: *weather.Main.Temp,
	}
	return &toStruct, nil
}

func New(apiKey string) Service {
	// We can ignore the error here because the request func
	// we provide doesn't have a case of returning an error.
	newClient, _ := weatherapi.NewClient("",
		weatherapi.WithBaseURL("https://api.openweathermap.org/data/2.5/"),
		weatherapi.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			query := req.URL.Query()
			query.Add("apiKey", apiKey)
			req.URL.RawQuery = query.Encode()
			return nil
		}),
	)
	return Service{
		ApiKey: apiKey,
		client: newClient,
	}
}

func ConvertOurRequestStructToOpenApiRequest(request Request, apiKey string) (string, error) {
	if request.ZipCode == "" || apiKey == "" {
		return "", errors.New("please specify zipcode and ApiKey")
	}
	return "http://api.openweathermap.org/data/2.5/forecast?zip=" + request.ZipCode + "&appid=" + apiKey, nil
}

func ConvertWeatherOpenApiResponseToStruct(r io.Reader) (CurrentWeather, error) {
	weatherResponse := &weatherapi.N200{}
	err := json.NewDecoder(r).Decode(weatherResponse)
	if err != nil {
		return CurrentWeather{}, err
	}
	response := CurrentWeather{}
	response.Temp = *weatherResponse.Main.Temp
	return response, nil
}
