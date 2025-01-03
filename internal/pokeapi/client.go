package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

type LocationArea struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewClient() *Client {
	return &Client{
		baseURL:    "https://pokeapi.co/api/v2",
		httpClient: &http.Client{},
	}
}

func (c *Client) GetLocationArea(id int) (LocationArea, error) {
	url := fmt.Sprintf("%s/location-area/%d", c.baseURL, id)
	
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	var location LocationArea
	err = json.Unmarshal(body, &location)
	if err != nil {
		return LocationArea{}, err
	}

	return location, nil
}
