package commands

import (
	"fmt"
	"strings"
)

func CommandMapb() error {
	start, end, ok := PaginatorInstance.PreviousPage()
	if !ok {
		fmt.Println("You're on the first page!")
		return nil
	}

	foundAny := false
	for i := start; i <= end; i++ {
		location, err := PokeClient.GetLocationArea(i)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				// If we found some locations but hit a gap, just skip this one
				continue
			}
			return err
		}
		foundAny = true
		fmt.Printf("%s\n", location.Name)
	}

	// If we didn't find any locations on this page, try going back another page
	if !foundAny && PaginatorInstance.currentPage > 1 {
		PaginatorInstance.currentPage++ // Undo the previous page change
		return CommandMapb()            // Recursively try the previous page
	}

	return nil
}
