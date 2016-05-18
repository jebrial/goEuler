/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	count, sum := 2000000, 2
	start := time.Now()
	for count > 0 {
		if count%2 == 0 {
			count--
			continue
		}
		if isPrime(count) {
			sum += count
		}
		count--
	}
	fmt.Println("\nThe sum of primes is: ", sum, "\n\n Completed after: ", time.Since(start))
}

func isPrime(c int) bool {
	if c < 3 {
		return false
	}
	for i := 3; i < c/2; i += 2 {
		if c%i == 0 {
			return false
		}
	}
	return true
}
