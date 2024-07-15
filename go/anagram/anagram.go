package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	lowerSubject := strings.ToLower(subject)
	anagrams := []string{}
	for i, candidate := range candidates {
		lowerCandidate := strings.ToLower(candidate)
		if isAnagram(lowerSubject, lowerCandidate) {
			anagrams = append(anagrams, candidates[i])
		}
	}
	return anagrams
}

func isAnagram(subject, candidate string) bool {
	if len(subject) != len(candidate) {
		return false
	}
	if subject == candidate {
		return false
	}
	subjectMap := make(map[rune]int)
	for _, r := range subject {
		subjectMap[r]++
	}
	for _, r := range candidate {
		if subjectMap[r] == 0 {
			return false
		}
		subjectMap[r]--
	}
	return true
}
