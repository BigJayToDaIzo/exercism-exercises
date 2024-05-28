package romannumerals

import "errors"

// ToRomanNumeral converts an integer to a Roman numeral
// I forgot how tricksey this problem was 😁
// I'm sure there's a more elegant way to solve this
// but I'm happy with my solution
func ToRomanNumeral(input int) (string, error) {
	transNumeral := ""
	if input <= 0 || input >= 4000 {
		return transNumeral, errors.New("out of range")
	}
	// no subtraction case possible
	// so loop on addition case
	for input >= 1000 {
		transNumeral += "M"
		input -= 1000
	}
	for input >= 500 {
		// subtraction case
		if input >= 900 {
			transNumeral += "CM"
			input -= 900
			// addition case
		} else {
			transNumeral += "D"
			input -= 500
		}
	}
	for input >= 100 {
		// subtraction case
		if input >= 400 {
			transNumeral += "CD"
			input -= 400
			// addition case
		} else {
			transNumeral += "C"
			input -= 100
		}
	}
	for input >= 50 {
		// subtraction case
		if input > 90 {
			transNumeral += "XC"
			input -= 90
			// addition case
		} else {
			transNumeral += "L"
			input -= 50
		}
	}
	for input < 50 && input >= 10 {
		if input > 40 {
			transNumeral += "XL"
			input -= 40
		} else {
			transNumeral += "X"
			input -= 10
		}
	}
	if input == 9 {
		transNumeral += "IX"
		input -= 9
	}
	for input >= 5 {
		transNumeral += "V"
		input -= 5
	}
	if input == 4 {
		transNumeral += "IV"
		input -= 4
	}
	for input >= 1 {
		transNumeral += "I"
		input--
	}

	return transNumeral, nil
}
