package client

import (
	"Weather-API-Application/internal/model"

	"github.com/stretchr/testify/mock"
)

// MockWeatherClient is a Testify mock implementing WeatherClient
type MockWeatherClient struct {
	mock.Mock
}

func (m *MockWeatherClient) GetCurrentWeather(city string) (*model.WeatherAPIResponse, error) {
	args := m.Called(city)

	var resp *model.WeatherAPIResponse
	if v := args.Get(0); v != nil {
		resp = v.(*model.WeatherAPIResponse)
	}
	return resp, args.Error(1)
}
