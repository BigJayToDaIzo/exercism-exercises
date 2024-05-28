package reverse

func Reverse(input string) string {
	rev := ""
	for _, r := range input {
		rev = string(r) + rev
	}
	return rev
}
