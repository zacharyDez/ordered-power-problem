package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPowerOrderedStartWithLimits(t *testing.T) {
	l := []int{-10, -3, -2, 0, 1, 2, 3, 10}
	assert.EqualValues(t, []int{0, 1, 4, 4, 9, 9, 100, 100}, GetPowerOrderedStartWithLimits(l))
}
