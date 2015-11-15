//What is the 10001st prime number?
package main

import "fmt"

func main() {
	count, number := 0, 2
	for count < 10001 {
		count, number = checkPrime(count, number)
	}
	fmt.Printf("The %d prime is: %d \n", count, number-1)
}

func checkPrime(c, n int) (int, int) {
	for i := 2; i <= n/2; i += 2 {
		if n%i == 0 {
			return c, n + 1
		}
		if i == 2 {
			i = 1
		}
	}
	return c + 1, n + 1
}
