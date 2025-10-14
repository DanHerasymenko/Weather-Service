package weather_service

import (
	"Weather-API-Application/internal/client"
)

// WeatherServiceFactory creates weather services with their dependencies
type WeatherServiceFactory struct {
	clientFactory *client.WeatherClientFactory
}

// NewWeatherServiceFactory creates a new weather service factory
func NewWeatherServiceFactory() *WeatherServiceFactory {
	return &WeatherServiceFactory{
		clientFactory: client.NewWeatherClientFactory(),
	}
}

// CreateWeatherService creates a weather service with all dependencies
func (f *WeatherServiceFactory) CreateWeatherService(weatherClient client.WeatherClient) WeatherService {
	return NewService(weatherClient)
}
