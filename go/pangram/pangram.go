package pangram

import (
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	// iterate over string adding/iterating runes in a map
	m := make(map[string]bool)
	for _, r := range input {
		if unicode.IsLetter(r) {
			m[strings.ToLower(string(r))] = true
		}
	}
	// if map size is 26 return true
	return len(m) == 26
}
