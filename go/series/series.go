package series

func All(n int, s string) []string {
	if n > len(s) {
		return nil
	}
	var substrs []string
	for i := 0; i <= len(s)-n; i++ {
		substrs = append(substrs, s[i:i+n])
	}
	return substrs
}

func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		return ""
	}
	return s[:n]
}

func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
