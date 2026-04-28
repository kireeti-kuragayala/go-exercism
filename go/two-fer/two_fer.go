package twofer

import "fmt"

func ShareWith(name string) string {
	// Sharing something with someone
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
