// diffSquares package provides API to diff the squares and what not
package diffsquares

// import math becauce that's how we do math
import "math"

// SquareOfSum returns the square of sum of n
func SquareOfSum(n int) int {
	sum := 0
	// 1 base iter because it's our iteration count not an array index
	for i := 1; i <= n; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

// SumOfSquares returns the sum of squares of n
func SumOfSquares(n int) int {
	sumSq := 0
	// 1 base iter because it's our power base and not an array index
	for i := 1; i <= n; i++ {
		sumSq += int(math.Pow(float64(i), 2))
	}
	return sumSq
}

// Difference returns the difference of SquareOfSum(n) and SumOfSquares(n)
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
