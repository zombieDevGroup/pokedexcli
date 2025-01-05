package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/zombiedevgroup/pokedexcli/internal/pokecache"
)

func NewClient() *Client {
	return &Client{
		baseURL:    "https://pokeapi.co/api/v2",
		httpClient: &http.Client{},
		cache:      pokecache.NewCache(5 * time.Minute),
	}
}

func (c *Client) GetLocationArea(id int) (LocationArea, error) {
	url := fmt.Sprintf("%s/location-area/%d", c.baseURL, id)

	// Check cache first
	if cached, ok := c.cache.Get(url); ok {
		var location LocationArea
		err := json.Unmarshal(cached, &location)
		if err != nil {
			return LocationArea{}, err
		}
		return location, nil
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return LocationArea{}, fmt.Errorf("location area %d not found", id)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	// Cache the raw response
	c.cache.Add(url, body)

	var location LocationArea
	err = json.Unmarshal(body, &location)
	if err != nil {
		return LocationArea{}, err
	}

	return location, nil
}

func (c *Client) GetEncounters(name string) ([]PokemonEncounter, error) {
	url := fmt.Sprintf("%s/location-area/%s", c.baseURL, name)

	// Check cache first
	if cached, ok := c.cache.Get(url); ok {
		var locationArea struct {
			PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
		}
		err := json.Unmarshal(cached, &locationArea)
		if err != nil {
			return []PokemonEncounter{}, err
		}
		return locationArea.PokemonEncounters, nil
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return []PokemonEncounter{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return []PokemonEncounter{}, fmt.Errorf("location area %s not found", name)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []PokemonEncounter{}, err
	}

	var locationArea struct {
		PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
	}
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return []PokemonEncounter{}, err
	}

	c.cache.Add(url, body)
	return locationArea.PokemonEncounters, nil
}
