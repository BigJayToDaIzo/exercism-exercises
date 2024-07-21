package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	numToString := strconv.Itoa(n)
	powerMultiplier := len(numToString)
	sumDigitsToPower := 0
	for _, digit := range numToString {
		digitInt, _ := strconv.Atoi(string(digit))
		sumDigitsToPower += int(math.Pow(float64(digitInt), float64(powerMultiplier)))
	}
	return sumDigitsToPower == n
}
