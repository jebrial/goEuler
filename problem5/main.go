//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
package main

import "fmt"

func main() {
	fmt.Println(smallestMutliple(10, 10))
	fmt.Println(smallestMutliple(2520, 20))
}

func smallestMutliple(number int, divisor int) int {
	for j := number; ; j++ {
		for i := divisor; i > 0; i-- {
			if j%i != 0 {
				break
			}
			if i == 1 {
				return j
			}
		}
	}
	return 0
}
