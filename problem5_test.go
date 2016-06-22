// Package goEuler - Problem 5
// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package goEuler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func evenlyDivisibleUpTo(count, frm int) int {
	for i := frm; i < 10000000000; i++ {
		for j := count; j > 0; j-- {
			if i%j != 0 {
				break
			}
			if j == 1 {
				return i
			}
		}
	}
	return 0
}

func TestEvenlyDivisibleUpTo(t *testing.T) {
	assert.Equal(t, 232792560, evenlyDivisibleUpTo(20, 2520), "Should equal 232792560")
}
