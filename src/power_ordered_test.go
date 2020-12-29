package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPowerOrderedStartWithLimits(t *testing.T) {
	l := []int{-10, -3, -2, 0, 1, 2, 3, 10}
	assert.EqualValues(t, []int{0, 1, 4, 4, 9, 9, 100, 100}, GetPowerOrderedStartWithLimits(l))
}

func TestPowerOrderedStartWithZero(t *testing.T) {
	l := []int{-10, -3, -2, 0, 1, 2, 3, 10}
	assert.EqualValues(t, []int{0, 1, 4, 4, 9, 9, 100, 100}, GetPowerOrderedStartWithZero(l))
}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	sign := -1
	for j := n - 1; j >= 0; j-- {
		result[i] = j * sign
		i++
		sign = sign * -1
	}
	return result
}

func BenchmarkPowerOrderedStartWithZeroSingleZero(b *testing.B) {
	els := []int{0}
	GetPowerOrderedStartWithZero(els)
}

func BenchmarkPowerOrderedStartWithLimitsSingleZero(b *testing.B) {
	els := []int{0}
	GetPowerOrderedStartWithLimits(els)
}

func BenchmarkPowerOrderedStartWithZeroFiveEltsMillionTimesMillion(b *testing.B) {
	els := []int{100000000}
	GetPowerOrderedStartWithZero(els)
}

func BenchmarkPowerOrderedStartWithLimitsFiveEltsMillionTimesMillion(b *testing.B) {
	els := []int{100000000}
	GetPowerOrderedStartWithLimits(els)
}
func BenchmarkPowerOrderedStartWithZeroSimpleList(b *testing.B) {
	els := []int{-10, -3, -2, 0, 1, 2, 3, 10}
	GetPowerOrderedStartWithZero(els)
}

func BenchmarkPowerOrderedStartWithLimitsSimpleList(b *testing.B) {
	els := []int{-10, -3, -2, 0, 1, 2, 3, 10}
	GetPowerOrderedStartWithLimits(els)
}
