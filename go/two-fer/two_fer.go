// Package provides a function to pay it forward
package twofer

import "fmt"

// ShareWith returns a name content appropriate statement
func ShareWith(name string) string {
	// readability improvement
	// DRY principle application
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
