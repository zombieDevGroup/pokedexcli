package commands

import (
	"fmt"

	"github.com/zombiedevgroup/pokedexcli/internal/pokedex"
)

func CommandInspect(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: inspect <pokemon>")
	}
	pokemon, isFound := pokedex.FindPokemon(args[0])
	if isFound == false {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")

	for _, stat := range pokemon.Stats {
		fmt.Printf("\t- %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf("\t- %s\n", typeInfo.Type.Name)
	}

	return nil
}
