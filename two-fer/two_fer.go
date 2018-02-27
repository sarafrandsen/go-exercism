package twofer

import (
	"fmt"
)

// ShareWith prints a statement telling with whom something is shared with.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)

}
