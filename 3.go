package main

import "fmt"

func main() {
	number, testNumber := 600851475143, 13195
	fmt.Println("This should be 29: ", getLargestPrimeFactor(testNumber))
	fmt.Println("This should be big: ", getLargestPrimeFactor(number))
}

func getLargestPrimeFactor(number int) int {
	for i := 2; i < number/2; i++ {
		if number%i == 0 {
			return getLargestPrimeFactor(number / i)
		}
	}
	return number
}
