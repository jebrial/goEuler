// Package goEuler - Problem 3
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?
package goEuler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func largestPrimeFactorOf(num int) int {
	for i := 2; i < num/2; i += 2 {
		if (num % i) == 0 {
			return largestPrimeFactorOf(num / i)
		}
		if i == 2 {
			i = 1
		}
	}
	return num
}

func TestLargestPrimeFactorOf(t *testing.T) {
	assert.Equal(t, 6857, largestPrimeFactorOf(600851475143), "Should equal 6857")
}
