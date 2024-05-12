// Package provides a fizzbuzz puzzle with a twist
package raindrops

import "fmt"

// Convert parses a number into raindrops on the tin roof
func Convert(number int) string {
	str := ""
	if number%3 == 0 {
		str += "Pling"
	}
	if number%5 == 0 {
		str += "Plang"
	}
	if number%7 == 0 {
		str += "Plong"
	}
	if str == "" {
		return fmt.Sprintf("%d", number)
	}
	return str
}
