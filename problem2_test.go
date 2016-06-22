// Package goEuler - Problem 2
// By considering the terms in the Fibonacci sequence whose values do not exceed four million,
// find the sum of the even-valued terms.
package goEuler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func sumOfEvenFibs() int {
	var i, j, sum int
	for i, j, sum = 1, 2, 0; j < 4000000; j, i = j+i, j {
		if j%2 == 0 {
			sum += j
		}
	}
	return sum
}

func TestSumOfEvenFibs(t *testing.T) {
	assert.Equal(t, 4613732, sumOfEvenFibs(), "Should equal 4613732")
}
