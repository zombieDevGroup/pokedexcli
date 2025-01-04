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
