package commands

import "github.com/zombiedevgroup/pokedexcli/internal/pokeapi"

// Command represents a CLI command with its name, usage, and callback function
type Command struct {
	Name     string
	Usage    string
	Callback func() error
}

// Shared state for the commands package
var (
	PokeClient *pokeapi.Client
)
