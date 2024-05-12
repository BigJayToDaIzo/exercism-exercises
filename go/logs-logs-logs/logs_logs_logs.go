package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		if char == '❗' {
			return "recommendation"
		} else if char == '🔍' {
			return "search"
		} else if char == '☀' {
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newString []byte
	for _, char := range log {
		if char == oldRune {
			newString = utf8.AppendRune(newString, newRune)
		} else {
			newString = utf8.AppendRune(newString, char)
		}
	}
	return string(newString)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return limit >= utf8.RuneCountInString(log)
}
