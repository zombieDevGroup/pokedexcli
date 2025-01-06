package pokeapi

import (
	"net/http"

	"github.com/zombiedevgroup/pokedexcli/internal/pokecache"
)

// Client is the PokeAPI client for making HTTP requests
type Client struct {
	baseURL    string
	httpClient *http.Client
	cache      *pokecache.Cache
}

// LocationArea represents a location area in the Pokemon world
type LocationArea struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Encounter represents an encounter in a location area
type PokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"pokemon"`
}

type Pokemon struct {
	Name           string `json:"name"`
	URL            string `json:"url"`
	BaseExperience int    `json:"base_experience"`
}
