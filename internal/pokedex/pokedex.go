package pokedex

import (
	"github.com/zombiedevgroup/pokedexcli/internal/pokeapi"
)

var (
	Pokedex = make(map[string]pokeapi.Pokemon)
)

func AddPokemon(pokemonName string, pokemon pokeapi.Pokemon) {
	Pokedex[pokemonName] = pokemon
	// fmt.Printf("Added %s to Pokedex\n", pokemonName)
}

func FindPokemon(pokemonName string) (pokeapi.Pokemon, bool) {
	pokemon, ok := Pokedex[pokemonName]
	return pokemon, ok
}
