package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	builder := ""
	lowerString := strings.ToLower(s)
	lowerWashed := ""
	rmap := map[rune]rune{
		'a': 'z', 'b': 'y', 'c': 'x', 'd': 'w', 'e': 'v',
		'f': 'u', 'g': 't', 'h': 's', 'i': 'r', 'j': 'q',
		'k': 'p', 'l': 'o', 'm': 'n', 'n': 'm', 'o': 'l',
		'p': 'k', 'q': 'j', 'r': 'i', 's': 'h', 't': 'g',
		'u': 'f', 'v': 'e', 'w': 'd', 'x': 'c', 'y': 'b',
		'z': 'a',
	}

	for i := 0; i < len(s); i++ {
		if lowerString[i] == ' ' || lowerString[i] == ',' || lowerString[i] == '.' {
			continue
		} else {
			lowerWashed += string(lowerString[i])
		}
	}

	for i := 0; i < len(lowerWashed); i++ {
		if i%5 == 0 && i != 0 {
			builder += " "
		}
		if unicode.IsDigit(rune(lowerWashed[i])) {
			builder += string(lowerWashed[i])
			continue
		}
		builder += string(rmap[rune(lowerWashed[i])])
	}
	return builder
}
