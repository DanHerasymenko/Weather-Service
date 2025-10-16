package weather_service

import (
	"Weather-API-Application/internal/client"
	"Weather-API-Application/internal/model"
	"fmt"
	"net/http"
	"strings"
)

type WeatherService interface {
	FetchWeatherForCity(city string) (*model.Weather, error, int)
}

type Service struct {
	weatherClient client.WeatherClient
}

func NewService(weatherClient client.WeatherClient) *Service {
	return &Service{
		weatherClient: weatherClient,
	}
}

func (s *Service) FetchWeatherForCity(city string) (*model.Weather, error, int) {

	weatherResp, err := s.weatherClient.GetCurrentWeather(city)
	if err != nil {
		if strings.Contains(err.Error(), "API key") {
			return nil, fmt.Errorf("weather API key is missing in config"), http.StatusInternalServerError
		}
		if strings.Contains(err.Error(), "status 404") {
			return nil, fmt.Errorf("city not found: %w", err), http.StatusNotFound
		}
		if strings.Contains(err.Error(), "failed to fetch") {
			return nil, fmt.Errorf("invalid request: %w", err), http.StatusBadRequest
		}
		if strings.Contains(err.Error(), "failed to decode") {
			return nil, fmt.Errorf("failed to decode weather data: %w", err), http.StatusInternalServerError
		}
		return nil, fmt.Errorf("failed to fetch weather data: %w", err), http.StatusInternalServerError
	}

	weather := &model.Weather{
		Temperature: weatherResp.Current.TempC,
		Humidity:    weatherResp.Current.Humidity,
		Description: weatherResp.Current.Condition.Text,
	}

	return weather, nil, http.StatusOK
}
