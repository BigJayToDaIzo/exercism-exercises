package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	cypher := ""
	for _, c := range plain {
		if c >= 'a' && c <= 'z' {
			cypher += string('a' + (c-'a'+rune(shiftKey))%26)
		} else if c >= 'A' && c <= 'Z' {
			cypher += string('A' + (c-'A'+rune(shiftKey))%26)
		} else {
			cypher += string(c)
		}
	}
	return cypher
}
