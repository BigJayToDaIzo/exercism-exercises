// https://golang.org/doc/effective_go.html#commentary
// package provides an API to IsLeapYear megabrain
package leap

// IsLeapYear returns true if the year is a leap year
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		}
		return true
	}
	return false
}
