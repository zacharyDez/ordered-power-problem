package main

import (
	"math"
)

// GetPowerOrderedStartWithLimits takes an ordered list as input and returns an ordered list of the elements to the power of 2.
// It is assumed that the list passed in ordered.
func GetPowerOrderedStartWithLimits(l []int) []int {
	initLen := len(l)
	res := make([]int, initLen)

	for i := 1; i <= initLen; i++ {
		if math.Abs(float64(l[0])) > math.Abs(float64(l[len(l)-1])) {
			res[initLen-i] = int(math.Pow(float64(l[0]), 2.0))
			l = l[1:]
		} else {
			res[initLen-i] = int(math.Pow(float64(l[len(l)-1]), 2.0))
			l = l[:len(l)-1]
		}

	}

	return res
}
