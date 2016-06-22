// By considering the terms in the Fibonacci sequence whose values do not exceed four million,
// find the sum of the even-valued terms.
package main

import "fmt"

func main() {
	var i, j, sum uint64

	for i, j, sum = 1, 2, 0; j < 4000000; j, i = j+i, j {
		if j%2 == 0 {
			sum += j
		}
	}

	fmt.Println("The sum is: ", sum)
}
