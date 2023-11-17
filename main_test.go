package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTorontoWeatherMocked(t *testing.T) {
	// Create a mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate OpenWeatherMap API response
		response := `{"main": {"temp": 290.15}}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer mockServer.Close()

	// Call the function to get Toronto weather, passing the mock server URL
	temperature, err := getTorontoWeather(mockServer.URL)
	if err != nil {
		t.Fatal(err)
	}

	// Assuming the mocked response has a temperature of 290.15 Kelvin
	expectedTemperature := 290.15

	if temperature != expectedTemperature {
		t.Errorf("getTorontoWeather returned unexpected temperature: got %.2f want %.2f", temperature, expectedTemperature)
	}
}
