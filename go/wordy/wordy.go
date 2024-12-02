package wordy

import (
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	ops := map[string]func(int, int) int{
		"plus":       func(a, b int) int { return a + b },
		"minus":      func(a, b int) int { return a - b },
		"multiplied": func(a, b int) int { return a * b },
		"divided":    func(a, b int) int { return a / b },
	}
	// trim question mark
	question = question[:len(question)-1]
	chunks := strings.Fields(question)
	intsArr := []int{}
	opsArr := []func(int, int) int{}
	opsInARow := 0
	intsInARow := 0

	for _, chunk := range chunks {
		// check for valid op after found int
		// and it's not the last chunk of the question
		if intsInARow == 1 {
			if _, ok := ops[chunk]; !ok {
				return 0, false
			}
		}
		// is chunk an int?
		if i, err := strconv.Atoi(chunk); err == nil {
			intsArr = append(intsArr, i)
			intsInARow++
			opsInARow = 0
		} else if op, ok := ops[chunk]; ok {
			opsArr = append(opsArr, op)
			opsInARow++
			if opsInARow > 1 {
				return 0, false
			}
			intsInARow = 0
		}
	}
	if len(opsArr) == 0 && len(intsArr) == 1 {
		return intsArr[0], true
	}
	if len(opsArr)+1 != len(intsArr) {
		return 0, false
	}
	return calculate(intsArr, opsArr), true
}

func calculate(ints []int, ops []func(int, int) int) int {
	accumulator := ints[0]
	for i, op := range ops {
		accumulator = op(accumulator, ints[i+1])
	}
	return accumulator
}
