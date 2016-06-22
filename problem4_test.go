// Package goEuler - Problem 4
// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

package goEuler

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func largestPalindromeProduct(num int) int {
	result := 0
	for i := num; i > num-(num/10); i-- {
		for j := num; j > (i - 10); j-- {
			testNumb := i * j
			if isPalindrome(testNumb) && testNumb > result {
				result = testNumb
			}
		}
	}
	return result
}

func isPalindrome(num int) bool {
	testStr := strconv.Itoa(num)
	strInRunes := []rune(testStr)
	for i, j := 0, len(strInRunes)-1; i < j; i, j = i+1, j-1 {
		strInRunes[i], strInRunes[j] = strInRunes[j], strInRunes[i]
	}
	return string(strInRunes) == testStr
}

func TestLargestPalindromeProduct(t *testing.T) {
	assert.Equal(t, 906609, largestPalindromeProduct(999), "Should equal 906609")
}
