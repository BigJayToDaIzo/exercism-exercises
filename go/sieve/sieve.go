package sieve

const unmarked = "unmarked"
const notPrime = "not prime"

func Sieve(limit int) []int {
	m := make(map[int]string)
	primes := []int{}

	for i := 2; i <= limit; i++ {
		m[i] = unmarked
	}

	for i := 2; i <= limit; i++ {
		if m[i] == unmarked {
			primes = append(primes, i)
			for j := i * 2; j <= limit; j += i {
				m[j] = notPrime
			}
		}
	}
	return primes
}
