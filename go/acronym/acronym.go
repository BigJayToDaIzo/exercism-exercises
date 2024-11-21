package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate creates an acronym from a string
func Abbreviate(s string) string {
	filteredStr := strings.ReplaceAll(s, "'s", "")
	f := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	words := strings.FieldsFunc(filteredStr, f)
	acronym := ""
	for _, word := range words {
		acronym += strings.ToUpper(string(word[0]))

	}
	return acronym
}
