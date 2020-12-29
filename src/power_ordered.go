package main

import (
	"math"
)

func isAbsValueGreater(a int, b int) bool {
	if abs(a) > abs(b) {
		return true
	}

	return false
}

func abs(a int) float64 {
	return math.Abs(float64(a))
}

func powerTwo(a int) int {
	return int(math.Pow(float64(a), 2.0))
}

// GetPowerOrderedStartWithLimits takes an ordered list as input and returns an ordered list of the elements to the power of 2.
// It is assumed that the list passed in ordered.
func GetPowerOrderedStartWithLimits(l []int) []int {
	initLen := len(l)
	res := make([]int, initLen)

	for i := 1; i <= initLen; i++ {
		if isAbsValueGreater(l[0], l[len(l)-1]) {
			res[initLen-i] = powerTwo(l[0])
			l = l[1:]
		} else {
			res[initLen-i] = powerTwo(l[len(l)-1])
			l = l[:len(l)-1]
		}

	}

	return res
}

func addPowerTwo(dst []int, src []int, idx int) []int {
	dst = append(dst, powerTwo(src[idx]))
	return dst
}

// GetPowerOrderedStartWithZero takes an ordered list as input and returns an ordered list of the elements to the power of 2.
// It is assumed that the list passed in ordered.
func GetPowerOrderedStartWithZero(l []int) []int {
	initLen := len(l)
	var res []int

	start := GetMinAbsIdx(l)
	right := start + 1
	left := start - 1

	res = addPowerTwo(res, l, start)
	for i := 0; i < initLen-1; i++ {
		if left < 0 {
			res = addPowerTwo(res, l, right)
			right++
			continue
		}

		if right > initLen-1 {
			res = addPowerTwo(res, l, left)
			left--
			continue
		}

		if !isAbsValueGreater(l[right], l[left]) {
			res = addPowerTwo(res, l, right)
			right++
		} else {
			res = addPowerTwo(res, l, left)
			left--
		}
	}
	return res
}

// GetMinAbsIdx returns the index of the minimum absolute value of a list.
// The algorithm is derived from binarySearch algorithm and has a complexity of O(log(n)).
func GetMinAbsIdx(l []int) int {
	startIndex := 0
	endIndex := len(l) - 1
	midIndex := len(l) / 2

	minIdx := midIndex
	min := abs(l[midIndex])
	for startIndex <= endIndex {
		value := l[midIndex]
		if abs(value) < min {
			minIdx = midIndex
			min = abs(value)
		}

		if min == 0 {
			break
		}

		if value > 0 {
			endIndex = midIndex - 1
			midIndex = (startIndex + endIndex) / 2
			continue
		}

		startIndex = midIndex + 1
		midIndex = (startIndex + endIndex) / 2
	}

	return minIdx
}
