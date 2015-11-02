package main

import "fmt"

func main() {
	sum := 0

	for i, j := 1, 2; j < 4000000; j, i = j+i, j {
		if j%2 == 0 {
			sum += j
		}
	}

	fmt.Println("The sum is: ", sum)
}
