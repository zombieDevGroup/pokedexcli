package commands

import (
	"fmt"
	"strings"
)

func CommandMap() error {
	start, end := PaginatorInstance.NextPage()

	foundAny := false
	for i := start; i <= end; i++ {
		location, err := PokeClient.GetLocationArea(i)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				// If we haven't found any locations on this page, go back to previous page
				if !foundAny {
					PaginatorInstance.currentPage--
					fmt.Println("You've reached the end of all locations!")
					return nil
				}
				// If we found some locations but hit a gap, just skip this one
				continue
			}
			return err
		}
		foundAny = true
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}
