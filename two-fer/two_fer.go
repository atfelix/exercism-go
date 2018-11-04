package twofer

import "fmt"

func ShareWith(person string) string {
	if person != "" {
		return fmt.Sprintf("One for %v, one for me.", person)
	} else {
		return "One for you, one for me."
	}
}
