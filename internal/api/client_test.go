package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFetchDataSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := LocationAreaResponse{
			Count:    2,
			Next:     "https://example.com/next",
			Previous: "",
			Results: []LocationAreaResult{
				{Name: "location-1", Url: "https://example.com/1"},
				{Name: "location-2", Url: "https://example.com/2"},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	data, err := FetchData[LocationAreaResponse](server.URL, "location-area")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if data.Count != 2 {
		t.Errorf("Expected Count=2, got %d", data.Count)
	}

	if len(data.Results) != 2 {
		t.Errorf("Expected 2 results, got %d", len(data.Results))
	}

	if data.Results[0].Name != "location-1" {
		t.Errorf("Expected first location 'location-1', got %s", data.Results[0].Name)
	}
}

func TestFetchDataWithQueryParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		query := r.URL.Query().Get("limit")
		if query != "20" {
			t.Errorf("Expected limit=20, got %s", query)
		}

		response := LocationAreaResponse{
			Count: 1,
			Results: []LocationAreaResult{
				{Name: "test-location", Url: "https://example.com/test"},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	data, err := FetchData[LocationAreaResponse](server.URL, "location-area", "limit=20")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(data.Results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(data.Results))
	}
}

func TestFetchDataNetworkError(t *testing.T) {
	data, err := FetchData[LocationAreaResponse]("http://invalid-url-that-does-not-exist-12345.local", "location-area")

	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if data.Count != 0 {
		t.Errorf("Expected empty Response, got Count=%d", data.Count)
	}

	if !strings.Contains(err.Error(), "Something went wrong") {
		t.Errorf("Expected error message containing 'Something went wrong', got %s", err.Error())
	}
}

func TestFetchDataBadStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not Found")
	}))
	defer server.Close()

	data, err := FetchData[LocationAreaResponse](server.URL, "location-area")

	if err == nil {
		t.Fatal("Expected error for bad status code, got nil")
	}

	if data.Count != 0 {
		t.Errorf("Expected empty Response, got Count=%d", data.Count)
	}

	if !strings.Contains(err.Error(), "bad status") {
		t.Errorf("Expected error message containing 'bad status', got %s", err.Error())
	}
}
