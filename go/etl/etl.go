package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	transMap := make(map[string]int)
	for k, v := range in {
		for _, s := range v {
			transMap[strings.ToLower(s)] = k
		}
	}
	return transMap
}
