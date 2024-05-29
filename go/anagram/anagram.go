package anagram

import (
	"fmt"
	"strings"
	"unicode"
)

func Detect(subject string, candidates []string) []string {
	fmt.Printf("subject: %v\n", subject)
	expected := []string{}
	for _, candidate := range candidates {
		fmt.Printf("candidate: %v\n", strings.ToLower(candidate))
		if strings.EqualFold(subject, candidate) || len(subject) != len(candidate) {
			fmt.Printf("subject: %v & candidate: %v rejected for sameness or length mismatch\n", subject, candidate)
			continue
		}
		for i, rune := range subject {
			if !strings.ContainsRune(strings.ToLower(candidate), unicode.ToLower(rune)) {
				fmt.Printf("rune: %v not contained in %v!\n", string(rune), candidate)
				break
			} else if i == len(subject)-1 {
				fmt.Printf("!anagram!: %v\n", candidate)
				expected = append(expected, candidate)
			}
		}
	}
	return expected
}
