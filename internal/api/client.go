package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pjgraczyk/pokedexcli/internal/cache"
)

const ApiBaseUrl = "https://pokeapi.co/api/v2"

func buildURL(baseUrl, objectType string, args ...string) string {
	url := fmt.Sprintf("%s/%s", baseUrl, objectType)
	if len(args) > 0 {
		url += "?" + strings.Join(args, "&")
	}
	return url
}

func FetchData(baseUrl, objectType string, args ...string) (Response, error) {
	var data Response
	url := buildURL(baseUrl, objectType, args...)
	res, err := http.Get(url)
	if err != nil {
		return Response{}, fmt.Errorf("Something went wrong! Err: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Response{}, fmt.Errorf("bad status: %s", res.Status)
	}

	json.NewDecoder(res.Body).Decode(&data)
	return data, nil
}

func GetLocationAreas(cache *cache.Cache, args ...string) (Response, error) {
	url := buildURL(ApiBaseUrl, "location-area", args...)
	if cached, ok := cache.Get(url); ok {
		var data Response
		json.Unmarshal(cached, &data)
		return data, nil
	}

	data, err := FetchData(ApiBaseUrl, "location-area", args...)
	if err != nil {
		return Response{}, err
	}

	jsonData, err := json.Marshal(data)
	if err == nil {
		cache.Add(url, jsonData)
	}

	return data, nil
}
