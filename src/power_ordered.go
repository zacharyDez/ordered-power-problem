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

func addPowerTwo(dst []int, src []int, idx int) []int {
	dst = append(dst, int(math.Pow(float64(src[idx]), 2.0)))
	return dst
}

// GetPowerOrderedStartWithZero takes an ordered list as input and returns an ordered list of the elements to the power of 2.
// It is assumed that the list passed in ordered.
func GetPowerOrderedStartWithZero(l []int) []int {
	initLen := len(l)
	var res []int

	start := findAbsoluteSmallestIdx(l)
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

		if math.Abs(float64(l[right])) < math.Abs(float64(l[left])) {
			res = addPowerTwo(res, l, right)
			right++
		} else {
			res = addPowerTwo(res, l, left)
			left--
		}
	}
	return res
}

func findAbsoluteSmallestIdx(l []int) int {
	for i := 0; ; i++ {
		posIdx := binarySearch(l, i)
		negIdx := binarySearch(l, -i)
		if posIdx != -1 {
			return posIdx
		}

		if negIdx != -1 {
			return negIdx
		}
	}
}

func binarySearch(array []int, target int) int {
	startIndex := 0
	endIndex := len(array) - 1
	midIndex := len(array) / 2

	for startIndex <= endIndex {
		value := array[midIndex]

		if value == target {
			return midIndex
		}

		if value > target {
			endIndex = midIndex - 1
			midIndex = (startIndex + endIndex) / 2
			continue
		}

		startIndex = midIndex + 1
		midIndex = (startIndex + endIndex) / 2
	}

	return -1
}
