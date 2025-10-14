package validate

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestIsValidCity - простий тест для валідації міста
func TestIsValidCity(t *testing.T) {
	tests := []struct {
		name     string
		city     string
		expected bool
		reason   string
	}{
		{
			name:     "Empty city should be invalid",
			city:     "",
			expected: false,
			reason:   "Empty string is not a valid city",
		},
		{
			name:     "Only spaces should be invalid",
			city:     "   ",
			expected: false,
			reason:   "Only spaces should be trimmed and result in empty string",
		},
		{
			name:     "Valid city name should be valid",
			city:     "Kyiv",
			expected: true,
			reason:   "Basic city name should work",
		},
		{
			name:     "City with spaces should be valid",
			city:     " New York ",
			expected: true,
			reason:   "Function should trim spaces and return true for non-empty result",
		},
		{
			name:     "City with special characters should be valid",
			city:     "São Paulo",
			expected: true,
			reason:   "Special characters in city names should be allowed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidCity(tt.city)
			require.Equal(t, tt.expected, result,
				"City validation failed for '%s'. Expected: %v, Got: %v. Reason: %s",
				tt.city, tt.expected, result, tt.reason)
		})
	}
}

// TestIsValidEmail - простий тест для валідації email
func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
		reason   string
	}{
		{
			name:     "Empty email should be invalid",
			email:    "",
			expected: false,
			reason:   "Empty string is not a valid email",
		},
		{
			name:     "Email without @ symbol should be invalid",
			email:    "email.com",
			expected: false,
			reason:   "Email must contain @ symbol",
		},
		{
			name:     "Email with multiple @ symbols should be invalid",
			email:    "user@@domain.com",
			expected: false,
			reason:   "Email should have exactly one @ symbol",
		},
		{
			name:     "Email without domain should be invalid",
			email:    "user@",
			expected: false,
			reason:   "Email must have a domain part",
		},
		{
			name:     "Email without local part should be invalid",
			email:    "@domain.com",
			expected: false,
			reason:   "Email must have a local part before @",
		},
		{
			name:     "Valid simple email should be valid",
			email:    "user@domain.com",
			expected: true,
			reason:   "Basic email format should work",
		},
		{
			name:     "Email with plus sign should be valid",
			email:    "user+tag@example.co.uk",
			expected: true,
			reason:   "Plus signs are allowed in email addresses",
		},
		{
			name:     "Email with dots should be valid",
			email:    "first.last@company.org",
			expected: true,
			reason:   "Dots are allowed in email addresses",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidEmail(tt.email)
			require.Equal(t, tt.expected, result,
				"Email validation failed for '%s'. Expected: %v, Got: %v. Reason: %s",
				tt.email, tt.expected, result, tt.reason)
		})
	}
}

// TestIsValidFrequency - покращений тест для валідації частоти без suite
func TestIsValidFrequency(t *testing.T) {
	// 🎯 Що ми покращили:
	// 1. Описові назви тестів
	// 2. Більше тестових випадків
	// 3. Поле reason для пояснення
	// 4. Кращі повідомлення про помилки

	tests := []struct {
		name      string
		frequency string
		expected  bool
		reason    string // Пояснення чому цей випадок важливий
	}{
		{
			name:      "Empty frequency should be invalid",
			frequency: "",
			expected:  false,
			reason:    "Empty string is not a valid frequency",
		},
		{
			name:      "Only spaces should be invalid",
			frequency: "   ",
			expected:  false,
			reason:    "Only spaces should be trimmed and result in empty string",
		},
		{
			name:      "Invalid frequency 'weekly' should be invalid",
			frequency: "weekly",
			expected:  false,
			reason:    "Only 'hourly' and 'daily' are supported frequencies",
		},
		{
			name:      "Invalid frequency 'monthly' should be invalid",
			frequency: "monthly",
			expected:  false,
			reason:    "Only 'hourly' and 'daily' are supported frequencies",
		},
		{
			name:      "Random text should be invalid",
			frequency: "random_text",
			expected:  false,
			reason:    "Any text other than hourly/daily should be invalid",
		},
		{
			name:      "Hourly with capital H should be valid",
			frequency: "Hourly",
			expected:  true,
			reason:    "Function should be case-insensitive",
		},
		{
			name:      "DAILY in uppercase should be valid",
			frequency: "DAILY",
			expected:  true,
			reason:    "Function should be case-insensitive",
		},
		{
			name:      "hourly in lowercase should be valid",
			frequency: "hourly",
			expected:  true,
			reason:    "Lowercase should work",
		},
		{
			name:      "daily in lowercase should be valid",
			frequency: "daily",
			expected:  true,
			reason:    "Lowercase should work",
		},
		{
			name:      "Hourly with spaces should be valid",
			frequency: " Hourly ",
			expected:  true,
			reason:    "Function should trim spaces",
		},
		{
			name:      "Daily with spaces should be valid",
			frequency: "  Daily  ",
			expected:  true,
			reason:    "Function should trim spaces",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidFrequency(tt.frequency)
			require.Equal(t, tt.expected, result,
				"Frequency validation failed for '%s'. Expected: %v, Got: %v. Reason: %s",
				tt.frequency, tt.expected, result, tt.reason)
		})
	}
}
