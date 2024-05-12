// package provites api to warp your mind with algebraic principles
package collatzconjecture

// import errors package for that sweet sweet New interface
import "errors"

// CollatzConjecture returns the number of steps required
// to reach 1 starting from n
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("invalid input")
	}
	steps := 0
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		steps++
	}
	return steps, nil
}
