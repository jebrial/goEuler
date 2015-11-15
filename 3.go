//What is the largest prime factor of the number 600851475143 ?
package main

import "fmt"

func main() {
	testNumber, number := 13195, 600851475143
	fmt.Println("This should be 29: ", getLargestPrimeFactor(testNumber))
	fmt.Println("This should be a bit bigger: ", getLargestPrimeFactor(number))
}

func getLargestPrimeFactor(number int) int {
	for i := 2; i <= number/2; i += 2 {
		if number%i == 0 {
			return getLargestPrimeFactor(number / i)
		}
		if i == 2 {
			i = 1
		}
	}
	return number
}
