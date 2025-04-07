package service

import (
	"encoding/json"
	"fmt"
	"go-weather-api/utils"
	"net/http"
	"io"
)

func GetWeatherWithCache(city string) (map[string]interface{}, error) {
	cached, err := GetCache(city)
	if err == nil && cached != "" {
		var data map[string]interface{}
		_ = json.Unmarshal([]byte(cached), &data)
		return data, nil
	}

	// Call Visual Crossing
	apiKey := utils.GetEnv("VISUAL_CROSSING_API_KEY", "")
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&key=%s&contentType=json", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var data map[string]interface{}
	_ = json.Unmarshal(body, &data)

	// Cache result
	ttl := utils.GetEnv("CACHE_TTL", "43200")
	_ = SetCache(city, string(body), utils.Atoi(ttl))

	return data, nil
}
