package client

import (
	"Weather-API-Application/internal/config"
	"net/http"
)

// WeatherClientFactory creates weather clients based on configuration
type WeatherClientFactory struct{}

// NewWeatherClientFactory creates a new weather client factory
func NewWeatherClientFactory() *WeatherClientFactory {
	return &WeatherClientFactory{}
}

// CreateWeatherClient creates a weather client from configuration
func (f *WeatherClientFactory) CreateWeatherClient(cfg *config.Config) WeatherClient {
	return NewWeatherClient(cfg.WeatherApiKey)
}

// CreateWeatherClientWithHTTPClient creates a weather client with custom HTTP client (for testing)
func (f *WeatherClientFactory) CreateWeatherClientWithHTTPClient(cfg *config.Config, httpClient interface{}) WeatherClient {
	if client, ok := httpClient.(*http.Client); ok {
		return NewWeatherClientWithHTTPClient(cfg.WeatherApiKey, client)
	}
	return NewWeatherClient(cfg.WeatherApiKey)
}
