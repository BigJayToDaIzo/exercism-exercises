package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(input string, span int) (int64, error) {
	if span < 0 {
		return 0.0, errors.New("span is negative")
	}
	if span > len(input) {
		return 0.0, errors.New("span is greater than input length")
	}
	series := []string{}
	for i := 0; i < len(input)-span+1; i++ {
		series = append(series, input[i:i+span])
	}
	largestSeriesProd := 0
	for _, s := range series {
		seriesProd := 1
		for _, c := range s {
			i, e := strconv.Atoi(string(c))
			if e != nil {
				return 0.0, errors.New("input contains non-digit characters")
			}
			seriesProd *= i
		}
		if seriesProd > largestSeriesProd {
			largestSeriesProd = seriesProd
		}
	}
	return int64(largestSeriesProd), nil
}
