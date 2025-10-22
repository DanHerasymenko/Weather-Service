package handler

import (
	"Weather-API-Application/internal/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockWeatherService struct {
	mock.Mock
}

func (m *MockWeatherService) FetchWeatherForCity(city string) (*model.Weather, error, int) {
	args := m.Called(city)

	var weather *model.Weather
	if args.Get(0) != nil {
		weather = args.Get(0).(*model.Weather)
	}

	return weather, args.Error(1), args.Int(2)
}

func TestGetWeather(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		city           string
		mockSetup      func(*MockWeatherService)
		expectedStatus int
		expectedBody   string
		reason         string
	}{
		{
			name: "Success - valid city returns weather data",
			city: "Kyiv",
			mockSetup: func(m *MockWeatherService) {
				expectedWeather := &model.Weather{
					Temperature: 25.5,
					Humidity:    60.0,
					Description: "Sunny",
				}
				m.On("FetchWeatherForCity", "Kyiv").Return(expectedWeather, nil, http.StatusOK)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"temperature":25.5,"humidity":60,"description":"Sunny"}`,
			reason:         "Handler should return weather data when service succeeds",
		},
		{
			name: "Error - empty city parameter",
			city: "",
			mockSetup: func(m *MockWeatherService) {
				// Do not setup mock - func must not be called
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "City parameter is required",
			reason:         "Handler should validate input before calling service",
		},
		{
			name: "Error - service returns city not found",
			city: "UnknownCity",
			mockSetup: func(m *MockWeatherService) {
				m.On("FetchWeatherForCity", "UnknownCity").Return(nil, assert.AnError, http.StatusNotFound)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   "City not found",
			reason:         "Handler should handle service errors correctly",
		},
		{
			name: "Error - service returns internal server error",
			city: "InvalidCity",
			mockSetup: func(m *MockWeatherService) {
				m.On("FetchWeatherForCity", "InvalidCity").Return(nil, assert.AnError, http.StatusInternalServerError)
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   "Internal server error",
			reason:         "Handler should map service error codes correctly",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockWeatherService)
			tt.mockSetup(mockService)

			handler := NewWeatherHandler(mockService)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest("GET", "/weather?city="+tt.city, nil)

			handler.GetWeather(c)

			require.Equal(t, tt.expectedStatus, w.Code, tt.reason)

			if tt.expectedStatus == http.StatusOK {
				assert.JSONEq(t, tt.expectedBody, w.Body.String(), "Response body should match expected JSON")
			} else {
				assert.Contains(t, w.Body.String(), tt.expectedBody, "Error message should contain expected text")
			}

			mockService.AssertExpectations(t)
		})
	}
}
