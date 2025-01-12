package commands

import (
	"fmt"

	"github.com/zombiedevgroup/pokedexcli/internal/pokedex"
)

func CommandPokedex(args []string) error {
	if len(pokedex.Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty!")
		return nil
	}

	fmt.Println("Your Pokedex:")

	for _, pokemon := range pokedex.Pokedex {
		fmt.Printf("\t- %s\n", pokemon.Name)
	}
	return nil
}
