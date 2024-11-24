package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type WeatherResponse struct {
    Name    string `json:"name"`
    Main    struct {
        Temp float64 `json:"temp"`
    } `json:"main"`
    Weather []struct {
        Description string `json:"description"`
    } `json:"weather"`
}

func main() {
    apiKey := "API-KEY" // Replace with your OpenWeatherMap API key
    city := "Boston" // You can change this to any city

    url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

    response, err := http.Get(url)
    if err != nil {
        fmt.Println("Error fetching weather:", err)
        return
    }
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        fmt.Printf("Error: Received HTTP status %d\n", response.StatusCode)
        return
    }

    var weather WeatherResponse
    if err := json.NewDecoder(response.Body).Decode(&weather); err != nil {
        fmt.Println("Error parsing JSON:", err)
        return
    }

    fmt.Printf("Current weather in %s:\n", weather.Name)

    // Safely handle the Weather slice
    if len(weather.Weather) == 0 {
        fmt.Println("Weather description not available.")
    } else {
        fmt.Printf("Description: %s\n", weather.Weather[0].Description)
    }

    fmt.Printf("Temperature: %.2f°C\n", weather.Main.Temp)
}
