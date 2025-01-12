package commands

import (
	"fmt"
	"math/rand"

	"github.com/zombiedevgroup/pokedexcli/internal/pokedex"
)

func CommandCatch(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: catch <pokemon>")
	}

	pokemonName := args[0]
	fmt.Println("Throwing a Pokeball at " + pokemonName + "...")

	pokemon, err := PokeClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	baseExp := pokemon.BaseExperience
	catchRandom := rand.Float64()
	catchChance := float64(baseExp) * catchRandom * 0.01
	catchRandomRoll := rand.Float64()

	if catchRandomRoll > catchChance {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	pokedex.AddPokemon(pokemonName, pokemon)

	fmt.Printf("%s caught!\n", pokemon.Name)
	fmt.Println("You may now inspect it with the inspect command.")
	return nil
}
