package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/zombiedevgroup/pokedexcli/internal/pokeapi"
)

var currentLocationId = 1
var pokeClient *pokeapi.Client

type cliCommand struct {
	name     string
	usage    string
	callback func() error
}

func commandHelp(commands map[string]cliCommand) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.usage)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap() error {
	currentPage := (currentLocationId-1)/20 + 1
	fmt.Printf("Current Page: %d\n", currentPage)
	fmt.Printf("Current Location ID: %d\n", currentLocationId)

	for i := 0; i < 20; i++ {
		location, err := pokeClient.GetLocationArea(currentLocationId + i)
		if err != nil {
			return err
		}
		fmt.Printf("%d: %s\n", location.ID, location.Name)
	}
	currentLocationId += 20
	return nil
}

func commandMapb() error {
	currentPage := (currentLocationId - 1) / 20
	fmt.Printf("Current Page: %d\n", currentPage)
	fmt.Printf("Current Location ID: %d\n", currentLocationId)

	if currentPage == 1 {
		fmt.Println("You're on the first page!")
		currentLocationId = 1
		return nil
	}

	currentLocationId = ((currentLocationId - 1) / currentPage) - 19
	fmt.Printf("Current Location ID: %d\n", currentLocationId)
	fmt.Printf("Next Page: %d\n", currentPage+1)
	fmt.Printf("Previous Page: %d\n", currentPage-1)

	for i := 0; i < 20; i++ {
		location, err := pokeClient.GetLocationArea(currentLocationId + i)
		if err != nil {
			return err
		}
		fmt.Printf("%d: %s\n", currentLocationId+i, location.Name)
	}
	return nil
}

func main() {
	pokeClient = pokeapi.NewClient()
	commands := map[string]cliCommand{
		"exit": {
			name:     "exit",
			usage:    "exit: Exit the Pokedex",
			callback: commandExit,
		},
	}

	commands["help"] = cliCommand{
		name:  "help",
		usage: "help: Display a help message",
		callback: func() error {
			return commandHelp(commands)
		},
	}

	commands["map"] = cliCommand{
		name:  "map",
		usage: "map: Display a map",
		callback: func() error {
			return commandMap()
		},
	}

	commands["mapb"] = cliCommand{
		name:  "mapb",
		usage: "mapb: Display a map",
		callback: func() error {
			return commandMapb()
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		if command, ok := commands[commandName]; ok {
			if err := command.callback(); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("Unknown command: %s\n", commandName)
		}
	}
}
