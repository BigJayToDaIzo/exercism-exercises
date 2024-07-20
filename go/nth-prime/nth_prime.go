package prime

import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n must be greater than one")
	}
	for i := 2; ; i++ {
		if isPrimeSqrt(i) {
			n--
			if n == 0 {
				return i, nil
			}
		}
	}
}

func isPrimeSqrt(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}
