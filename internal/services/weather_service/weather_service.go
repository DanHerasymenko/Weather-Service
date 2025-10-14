package weather_service

import (
	"Weather-API-Application/internal/client"
	"Weather-API-Application/internal/model"
	"fmt"
	"net/http"
	"strings"
)

// WeatherService defines the interface for weather operations
type WeatherService interface {
	FetchWeatherForCity(city string) (*model.Weather, error, int)
}

// Service implements WeatherService interface
type Service struct {
	weatherClient client.WeatherClient
}

// NewService creates a new weather service
func NewService(weatherClient client.WeatherClient) *Service {
	return &Service{
		weatherClient: weatherClient,
	}
}

// FetchWeatherForCity retrieves the current weather data for the given city
// using the weather client to fetch data from external API.
//
// It performs the following steps:
//   - Uses weather client to fetch weather data
//   - Maps the API response to internal Weather model
//   - Returns appropriate HTTP status codes based on client response
func (s *Service) FetchWeatherForCity(city string) (*model.Weather, error, int) {
	// Fetch weather data from client
	weatherResp, err := s.weatherClient.GetCurrentWeather(city)
	if err != nil {
		// Map client errors to appropriate HTTP status codes
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

	// Map API response to internal model
	weather := &model.Weather{
		Temperature: weatherResp.Current.TempC,
		Humidity:    weatherResp.Current.Humidity,
		Description: weatherResp.Current.Condition.Text,
	}

	return weather, nil, http.StatusOK
}