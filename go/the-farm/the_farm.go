package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	numCows int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numCows, e.message)
}

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, numCows int) (float64, error) {
	fa, faErr := fc.FodderAmount(numCows)
	if faErr != nil {
		return 0, faErr
	}
	ff, ffErr := fc.FatteningFactor()
	if ffErr != nil {
		return 0, ffErr
	}
	return fa / float64(numCows) * ff, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, numCows int) (float64, error) {
	if numCows < 1 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, numCows)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numCows int) error {
	if numCows < 0 {
		return &InvalidCowsError{
			numCows: numCows,
			message: "there are no negative cows",
		}
	} else if numCows == 0 {
		return &InvalidCowsError{
			numCows: numCows,
			message: "no cows don't need food",
		}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
