package client

import (
	"Weather-API-Application/internal/model"
	"encoding/json"
	"fmt"
	"net/http"
)

// WeatherClient interface defines the contract for weather API client
type WeatherClient interface {
	GetCurrentWeather(city string) (*model.WeatherAPIResponse, error)
}

// weatherClient implements WeatherClient interface
type weatherClient struct {
	apiKey     string
	baseURL    string
	httpClient *http.Client
}

// NewWeatherClient creates a new weather API client
func NewWeatherClient(apiKey string) WeatherClient {
	return &weatherClient{
		apiKey:     apiKey,
		baseURL:    "https://api.weatherapi.com/v1/current.json",
		httpClient: &http.Client{},
	}
}

// NewWeatherClientWithHTTPClient creates a new weather API client with custom HTTP client (for testing)
func NewWeatherClientWithHTTPClient(apiKey string, httpClient *http.Client) WeatherClient {
	return &weatherClient{
		apiKey:     apiKey,
		baseURL:    "https://api.weatherapi.com/v1/current.json",
		httpClient: httpClient,
	}
}

// GetCurrentWeather fetches current weather data for the given city
func (c *weatherClient) GetCurrentWeather(city string) (*model.WeatherAPIResponse, error) {
	// Validate API key
	if c.apiKey == "" {
		return nil, fmt.Errorf("weather API key is missing")
	}

	// Build URL
	url := fmt.Sprintf("%s?key=%s&q=%s&aqi=no", c.baseURL, c.apiKey, city)

	// Make HTTP request
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather data: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weather API returned status %d: %s", resp.StatusCode, resp.Status)
	}

	// Decode response
	var weatherResp model.WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return nil, fmt.Errorf("failed to decode weather response: %w", err)
	}

	return &weatherResp, nil
}
