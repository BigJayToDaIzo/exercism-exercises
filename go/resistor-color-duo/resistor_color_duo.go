package resistorcolorduo

import (
	"fmt"
	"strconv"
	"strings"
)

// Value should return the resistance value of a resistor with a given colors, ignoring any after second.
func Value(colors []string) int {
	ref := map[string]int{
		"black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
	}
	ohms := ""
	if strings.ToLower(colors[0]) != "black" {
		ohms += fmt.Sprint(ref[colors[0]])
	}
	ohms += fmt.Sprint(ref[colors[1]])
	s, err := strconv.Atoi(ohms)
	if err != nil {
		return 0
	}
	return s
}
