package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	runeFrequency := make(Frequency)
	rf := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '\''
	}
	pf := func(c rune) bool {
		return unicode.IsSpace(c) || unicode.IsPunct(c) || unicode.IsSymbol(c)
	}
	trimmedPhrase := strings.TrimFunc(phrase, pf)
	runes := strings.FieldsFunc(trimmedPhrase, rf)
	for _, r := range runes {
		if len(r) == 0 {
			continue
		}
		r = strings.Trim(r, "'\"")
		r = strings.ToLower(r)
		runeFrequency[r]++
	}
	return runeFrequency
}
