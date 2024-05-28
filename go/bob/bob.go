// Package bob Lackadaisical teenager stuff
package bob

import "strings"

// Hey Whatever man
func Hey(remark string) string {
	isQuestion := strings.HasSuffix(strings.TrimSpace(remark), "?")
	isShouting := strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark
	isSilence := strings.TrimSpace(remark) == ""
	switch {
	case isQuestion && !isShouting:
		return "Sure."
	case isShouting && !isQuestion:
		return "Whoa, chill out!"
	case isQuestion && isShouting:
		return "Calm down, I know what I'm doing!"
	case isSilence:
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
