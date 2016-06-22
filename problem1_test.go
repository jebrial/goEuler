// Package goEuler - Problem 1
// Find the sum of all the multiples of 3 or 5 below 1000.
package goEuler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func sumOfMutilples() int {
	sum := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func TestSumOfMultiples(t *testing.T) {
	assert.Equal(t, 233168, sumOfMutilples(), "Should equal 233168")
}
