package weather_service

import (
	"Weather-API-Application/internal/config"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFetchWeatherForCity(t *testing.T) {
	tests := []struct {
		name          string
		city          string
		apiKey        string
		expectedError bool
		expectedCode  int
		reason        string
	}{
		{
			name:          "Missing API key should return error",
			city:          "Kyiv",
			apiKey:        "",
			expectedError: true,
			expectedCode:  http.StatusInternalServerError,
			reason:        "Service should fail when API key is not configured",
		},
		{
			name:          "Empty city should return error",
			city:          "",
			apiKey:        "test-key",
			expectedError: true,
			expectedCode:  http.StatusBadRequest,
			reason:        "Service should handle empty city parameter",
		},
		// TODO: Додати більше кейсів коли зробимо моки
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 🎯 Arrange - підготовка
			cfg := &config.Config{
				WeatherApiKey: tt.apiKey,
			}
			service := NewService(cfg)

			// 🎯 Act - виконання
			weather, err, code := service.FetchWeatherForCity(tt.city)

			// 🎯 Assert - перевірка
			if tt.expectedError {
				require.Error(t, err, tt.reason)
				require.Nil(t, weather, "Weather should be nil when there's an error")
			} else {
				require.NoError(t, err, tt.reason)
				require.NotNil(t, weather, "Weather should not be nil when successful")
			}

			require.Equal(t, tt.expectedCode, code, "HTTP status code should match expected")
		})
	}
}
