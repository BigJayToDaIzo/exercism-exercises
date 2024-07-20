package atbash

func Atbash(s string) string {
	// First we strip all white space and punct
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == ',' || s[i] == '.' {
			s = s[:i] + s[i+1:]
			// shift 5 here?
			i--
		}
	}
	// Then we translate in 5s and inject a space after each 5th

	return ""

}
