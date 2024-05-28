package grains

import "errors"

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("invalid input")
	}
	if number == 1 {
		return 1, nil
	}
	halfRiceCount, err := Square(number - 1)
	if err != nil {
		return 0, errors.New("invalid input")
	}
	return halfRiceCount * 2, nil
}

func Total() uint64 {
	totalGrains := uint64(0)
	for i := 1; i <= 64; i++ {
		num, err := Square(i)
		if err != nil {
			return 0
		}
		totalGrains += num
	}
	return totalGrains
}
