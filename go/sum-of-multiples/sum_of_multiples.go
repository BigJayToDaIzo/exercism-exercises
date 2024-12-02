package summultiples

import "slices"

func SumMultiples(limit int, divisors ...int) int {
	arr := []int{}
	for _, d := range divisors {
		if d == 0 {
			continue
		}
		for i := 1; i < limit; i++ {
			if i%d == 0 {
				arr = append(arr, i)
			}
		}
	}
	slices.Sort(arr)
	// Compact removes consecutive duplicates, which is why we sort first
	arr = slices.Compact(arr)
	sum := 0
	for _, a := range arr {
		sum += a
	}
	return sum
}
