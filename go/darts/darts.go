package darts

import "math"

func Score(x, y float64) int {
	distFromCenter := math.Sqrt(x*x + y*y)
	switch {
	case distFromCenter <= 1:
		return 10
	case distFromCenter <= 5:
		return 5
	case distFromCenter <= 10:
		return 1
	default:
		return 0
	}
}
