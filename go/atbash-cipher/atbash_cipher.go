package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	builder := ""
	lowerString := strings.ToLower(s)
	lowerWashed := ""
	rmap := map[rune]int{
		'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5,
		'f': 6, 'g': 7, 'h': 8, 'i': 9, 'j': 10,
		'k': 11, 'l': 12, 'm': 13, 'n': 14, 'o': 15,
		'p': 16, 'q': 17, 'r': 18, 's': 19, 't': 20,
		'u': 21, 'v': 22, 'w': 23, 'x': 24, 'y': 25,
		'z': 26,
	}
	revRmap := map[int]rune{
		1: 'z', 2: 'y', 3: 'x', 4: 'w', 5: 'v',
		6: 'u', 7: 't', 8: 's', 9: 'r', 10: 'q',
		11: 'p', 12: 'o', 13: 'n', 14: 'm', 15: 'l',
		16: 'k', 17: 'j', 18: 'i', 19: 'h', 20: 'g',
		21: 'f', 22: 'e', 23: 'd', 24: 'c', 25: 'b',
		26: 'a',
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
		val := rmap[rune(lowerWashed[i])]
		builder += string(revRmap[val])
	}
	return builder
}
