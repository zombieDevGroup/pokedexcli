package commands

import (
	"fmt"
)

func CommandExplore(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: explore <location>")
	}

	locationName := args[0]
	encounters, err := PokeClient.GetEncounters(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationName)
	fmt.Printf("Found Pokemon:\n")
	for _, encounter := range encounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
