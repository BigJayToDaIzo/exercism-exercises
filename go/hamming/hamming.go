// package providing API to HAMMING DISTANCE yo
package hamming

// package providing an API to display the errors we are all prone to mistake
// for debugging purposes, and not humiliation for sure
import "errors"

// Distance function takes two strings and returns the Hamming Distance between them
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strings must be of equal length")
	}
	distance := 0
	for i := range a {
		if b[i] != a[i] {
			distance++
		}
	}
	return distance, nil
}
