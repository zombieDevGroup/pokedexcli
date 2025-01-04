package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zombiedevgroup/pokedexcli/internal/commands"
	"github.com/zombiedevgroup/pokedexcli/internal/pokeapi"
)

var PaginatorInstance *commands.Paginator

func main() {
	pokeClient := pokeapi.NewClient()
	commands.PokeClient = pokeClient
	commands.PaginatorInstance = commands.NewPaginator(20) // 20 items per page

	cliCommands := map[string]commands.Command{
		"exit": {
			Name:     "exit",
			Usage:    "exit: Exit the Pokedex",
			Callback: commands.CommandExit,
		},
	}

	cliCommands["help"] = commands.Command{
		Name:     "help",
		Usage:    "help: Display a help message",
		Callback: func() error { return commands.CommandHelp(cliCommands) },
	}

	cliCommands["map"] = commands.Command{
		Name:     "map",
		Usage:    "map: Display next page of locations",
		Callback: commands.CommandMap,
	}

	cliCommands["mapb"] = commands.Command{
		Name:     "mapb",
		Usage:    "mapb: Display previous page of locations",
		Callback: commands.CommandMapb,
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := strings.Fields(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		if command, ok := cliCommands[commandName]; ok {
			if err := command.Callback(); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("Unknown command: %v\n", commandName)
			fmt.Printf("Use 'help' to see available commands\n")
		}
	}
}
