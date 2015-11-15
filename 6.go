// Find the difference between the sum of the squares of the first
// one hundred natural numbers and the square of the sum.
package main

import "fmt"

func main() {
	fmt.Println(squareOfSum(10) - sumOfSquare(10))
	fmt.Println(squareOfSum(100) - sumOfSquare(100))
}

func sumOfSquare(num int) int {
	result := 0
	for i := 1; i <= num; i++ {
		result += square(i)
	}
	return result
}

func squareOfSum(num int) int {
	result := 0
	for i := 1; i <= num; i++ {
		result += i
	}
	return square(result)
}

func square(num int) int {
	return num * num
}
