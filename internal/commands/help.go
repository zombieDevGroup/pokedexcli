package commands

import "fmt"

func CommandHelp(cliCommands map[string]Command) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, command := range cliCommands {
		fmt.Printf("%s: %s\n", command.Name, command.Usage)
	}
	return nil
}
