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

func FetchData[T any](baseUrl, objectType string, args ...string) (T, error) {
	var data T
	url := buildURL(baseUrl, objectType, args...)
	res, err := http.Get(url)
	if err != nil {
		return data, fmt.Errorf("Something went wrong! Err: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return data, fmt.Errorf("bad status: %s", res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return data, fmt.Errorf("failed to decode JSON: %v", err)
	}
	return data, nil
}

func GetCachedData[T any](cache *cache.Cache, baseUrl, objectType string, args ...string) (T, error) {
	var data T
	url := buildURL(baseUrl, objectType, args...)

	if cached, ok := cache.Get(url); ok {
		err := json.Unmarshal(cached, &data)
		if err == nil {
			return data, nil
		}
	}

	data, err := FetchData[T](baseUrl, objectType, args...)
	if err != nil {
		return data, err
	}

	jsonData, err := json.Marshal(data)
	if err == nil {
		cache.Add(url, jsonData)
	}

	return data, nil
}
