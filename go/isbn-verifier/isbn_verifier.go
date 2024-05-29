package isbn

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func IsValidISBN(isbn string) bool {
	// Remove hyphens
	refIsbn := strings.ReplaceAll(isbn, "-", "")
	if utf8.RuneCountInString(refIsbn) != 10 {
		return false
	}
	adder := 0
	for i, c := range refIsbn {
		// Check if last char is X or 0-9
		if i == 9 && c == 'X' || c == 'x' {
			// Do maths
			adder += 10
		} else if c >= '0' && c <= '9' {
			adder += int(c-'0') * (10 - i)
		} else {
			fmt.Printf("invalid char: %c\n", c)
			return false
		}
	}
	return adder%11 == 0
}
