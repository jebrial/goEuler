package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("This number is the largest palindrome: ", findLargestPalindrome(999, 900))
}

func findLargestPalindrome(high, low int) int {
	var palindrome int
	for i := high; i > low; i-- {
		for j := high; j > low; j-- {
			result := i * j
			resultStr := strconv.Itoa(result)
			if resultStr == reverse(resultStr) && result > palindrome {
				palindrome = result
			}
		}
	}
	return palindrome
}

func reverse(s string) string {
	strInRunes := []rune(s)
	for i, j := 0, len(strInRunes)-1; i < j; i, j = i+1, j-1 {
		strInRunes[i], strInRunes[j] = strInRunes[j], strInRunes[i]
	}
	return string(strInRunes)
}
