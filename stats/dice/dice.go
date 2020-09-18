package dice

import "github.com/nathangreene3/math"

// ProbOfSum ...
func ProbOfSum(sum, sides, numDice int) float64 {
	var prob float64
	for i := 0; i < numDice; i++ {

	}

	return prob / float64(math.PowInt(sides, numDice))
}
