package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var openWeatherMapAPIURL = "http://api.openweathermap.org/data/2.5/weather?q=Toronto&appid=8cff1b7c8f3c63704b9370b7545598bb"

type WeatherInfo struct {
	Main struct {
		Temperature float64 `json:"temp"`
	} `json:"main"`
}

// Update the function to accept the API URL as a parameter
func getTorontoWeather(apiURL string) (float64, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("OpenWeatherMap API request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	// Uncomment the line below to print the response body for debugging
	// fmt.Println(string(body))

	var weatherInfo WeatherInfo
	err = json.Unmarshal(body, &weatherInfo)
	if err != nil {
		return 0, err
	}
	return weatherInfo.Main.Temperature, nil
}

func torontoWeatherHandler(w http.ResponseWriter, r *http.Request) {
	temperature, err := getTorontoWeather(openWeatherMapAPIURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching Toronto weather: %s", err), http.StatusInternalServerError)
		return
	}

	// Convert temperature from Kelvin to Celsius
	temperatureCelsius := temperature - 273.15

	fmt.Fprintf(w, "Toronto temperature is %.2f°C", temperatureCelsius)
	resp := map[string]float64{"current_temperature_Toronto": temperatureCelsius}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api/torontoweather", torontoWeatherHandler)
	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
