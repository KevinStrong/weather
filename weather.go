package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	weatherapi "weather/gen"
)

// CurrentWeather is part of weather api response
type CurrentWeather struct {
	Temp float64
}

type Service struct {
	ApiKey  string
	baseUrl string
	client *http.Client
}

func (s *Service) GetWeather(location string) (CurrentWeather, error) {
	targetUrl := s.MakeURL(location)

	response, err := s.client.Get(targetUrl) //nolint:gosec

	if err != nil {
		return CurrentWeather{}, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return CurrentWeather{}, fmt.Errorf("got %d", response.StatusCode)
	}

	weather, err := ConvertWeatherOpenApiResponseToStruct(response.Body)
	if err != nil {
		return CurrentWeather{}, err
	}
	return weather, nil
}

func (s *Service) MakeURL(city string) string {
	return fmt.Sprintf("%s/weather?q=%s&appid=%s&units=imperial", s.baseUrl, url.QueryEscape(city), s.ApiKey)
}

type Option func(*Service)

func New(apiKey string, opts ...Option) *Service {
	service := &Service{
		baseUrl: "https://api.openweathermap.org/data/2.5",
		ApiKey: apiKey,
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
	for _, o := range opts {
		o(service)
	}
	return service
}

func WithBaseURL(baseUrl string) Option {
	return func(s *Service) {
		s.baseUrl = baseUrl
	}
}

func WithHttpClient(client *http.Client) Option {
	return func(s *Service) {
		s.client = client
	}
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
