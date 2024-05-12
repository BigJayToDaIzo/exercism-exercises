// Source: exercism/problem-specifications/exercises/luhn/canonical-data.json
// package provides API to check luhnism
package luhn

// Valid returns a number's luhnism as a bool
func Valid(id string) bool {
	if len(id) <= 1 || id[0] == ' ' {
		return false
	}
	// convert string to int array
	intArr := make([]int, 0)
	for _, char := range id {
		if char == ' ' {
			continue
		}
		if char >= '0' && char <= '9' {
			intArr = append(intArr, int(char-'0'))
		} else {
			return false
		}
	}
	// double every second digit
	for i := len(intArr) - 2; i >= 0; i -= 2 {
		intArr[i] *= 2
		if intArr[i] > 9 {
			intArr[i] -= 9
		}
	}
	return isValid(intArr)
}

func isValid(arr []int) bool {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum%10 == 0
}
