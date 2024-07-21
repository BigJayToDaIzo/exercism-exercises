package cipher

import (
	"strings"
	"unicode"
)

const alphabetSize = 26

var shiftMap = map[rune]int{
	'a': 0, 'b': 1, 'c': 2, 'd': 3, 'e': 4, 'f': 5, 'g': 6, 'h': 7, 'i': 8, 'j': 9, 'k': 10, 'l': 11, 'm': 12, 'n': 13, 'o': 14, 'p': 15, 'q': 16, 'r': 17, 's': 18, 't': 19, 'u': 20, 'v': 21, 'w': 22, 'x': 23, 'y': 24, 'z': 25,
}

type shift struct {
	shift int
}

type vigenere struct {
	key string
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance >= -25 && distance <= -1 {
		return shift{shift: distance + alphabetSize}
	} else if distance >= 1 && distance <= 25 {
		return shift{shift: distance}
	} else {
		return nil
	}
}

func (c shift) Encode(input string) string {
	encoded := ""
	for _, r := range stripInput(input) {
		if r >= 'a' && r <= 'z' {
			r += rune(c.shift)
			if r > 'z' {
				r -= alphabetSize
			}
		}
		encoded += string(r)
	}
	return encoded
}

func (c shift) Decode(input string) string {
	decoded := ""
	for _, r := range stripInput(input) {
		if r >= 'a' && r <= 'z' {
			r -= rune(c.shift)
			if r < 'a' {
				r += alphabetSize
			}
		}
		decoded += string(r)
	}
	return decoded
}

func NewVigenere(key string) Cipher {
	if !checkKey(key) {
		return nil
	}
	return vigenere{key: key}
}

func (v vigenere) Encode(input string) string {
	encoded := ""
	for i, r := range stripInput(input) {
		if r >= 'a' && r <= 'z' {
			shift := shiftMap[rune(v.key[i%len(v.key)])]
			r += rune(shift)
			if r > 'z' {
				r -= alphabetSize
			}
		}
		encoded += string(r)
	}
	return encoded
}

func (v vigenere) Decode(input string) string {
	decoded := ""
	for i, r := range stripInput(input) {
		if r >= 'a' && r <= 'z' {
			shift := shiftMap[rune(v.key[i%len(v.key)])]
			r -= rune(shift)
			if r < 'a' {
				r += alphabetSize
			}
		}
		decoded += string(r)
	}
	return decoded
}

func stripInput(input string) string {
	// strings.Map is beautiful!😍
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, input)
}

func checkKey(input string) bool {
	if strings.ReplaceAll(input, "a", "") == "" {
		return false
	}
	for _, r := range input {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}
