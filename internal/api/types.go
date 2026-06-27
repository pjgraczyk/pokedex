package api

type LocationAreaResponse struct {
	Count    int                  `json:"count"`
	Next     string               `json:"next"`
	Previous string               `json:"previous"`
	Results  []LocationAreaResult `json:"results"`
}

type LocationAreaResult struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
