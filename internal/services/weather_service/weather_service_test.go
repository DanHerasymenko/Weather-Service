package weather_service

import (
	"Weather-API-Application/internal/client"
	"Weather-API-Application/internal/model"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchWeatherForCity(t *testing.T) {
	mockClient := new(client.MockWeatherClient)
	svc := NewService(mockClient)

	tests := []struct {
		name           string
		city           string
		mockSetup      func()
		expectedCode   int
		expectedError  bool
		expectedResult *model.Weather
	}{
		{
			name: "Missing API key error from client -> 500",
			city: "Kyiv",
			mockSetup: func() {
				mockClient.On("GetCurrentWeather", "Kyiv").Return(nil, errors.New("weather API key is missing"))
			},
			expectedCode:  http.StatusInternalServerError,
			expectedError: true,
		},
		{
			name: "City not found -> 404",
			city: "UnknownCity",
			mockSetup: func() {
				mockClient.On("GetCurrentWeather", "UnknownCity").Return(nil, errors.New("weather API returned status 404: 404 Not Found"))
			},
			expectedCode:  http.StatusNotFound,
			expectedError: true,
		},
		{
			name: "Valid response -> 200",
			city: "Kyiv",
			mockSetup: func() {
				resp := &model.WeatherAPIResponse{}
				resp.Current.TempC = 23.4
				resp.Current.Humidity = 55
				resp.Current.Condition.Text = "Cloudy"
				mockClient.On("GetCurrentWeather", "Kyiv").Return(resp, nil)
			},
			expectedCode:  http.StatusOK,
			expectedError: false,
			expectedResult: &model.Weather{
				Temperature: 23.4,
				Humidity:    55,
				Description: "Cloudy",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient.ExpectedCalls = nil
			mockClient.Calls = nil
			tt.mockSetup()

			result, err, code := svc.FetchWeatherForCity(tt.city)

			if tt.expectedError {
				require.Error(t, err)
				require.Nil(t, result)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedResult, result)
			}
			require.Equal(t, tt.expectedCode, code)
		})
	}
}
