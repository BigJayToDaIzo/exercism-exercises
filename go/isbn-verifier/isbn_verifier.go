package isbn

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func IsValidISBN(isbn string) bool {
	// Remove hyphens
	refIsbn := strings.ReplaceAll(isbn, "-", "")
	fmt.Printf("isbn: %s refIsbn: %s\n", isbn, refIsbn)

	if utf8.RuneCountInString(refIsbn) != 10 {
		fmt.Printf("wrong length (%s): %d\n", refIsbn, utf8.RuneCountInString(isbn))
		return false
	}
	// Check if last char is X or 0-9
	adder := 0
	// Do maths
	if adder != 0 && adder%11 == 0 {
		fmt.Printf("%s is a good ISBN!\n", refIsbn)
	} else {
		fmt.Printf("%s is NOT a good ISBN!\n", refIsbn)
	}
	return adder%11 == 0
}
