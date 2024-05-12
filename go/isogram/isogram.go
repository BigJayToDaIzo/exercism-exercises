// Source: exercism/x/go/isogram
// package provides API to check isogramism
package isogram

// import strings becauce that's how we compare strings, or isograms
import "strings"

// IsIsogram returns a words isogramness as a bool
func IsIsogram(word string) bool {
	wordMap := make(map[rune]int)
	word = strings.ToLower(word)
	for _, char := range word {
		if char != '-' && char != ' ' {
			if wordMap[char] > 0 {
				return false
			}
			wordMap[char]++
		}
	}
	return true
}
