//Find the thirteen adjacent digits in the 1000-digit number that have the greatest product. What is the value of this product?

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fileOpen, _ := os.Open("8.dat")
	b1 := make([]byte, 1000)
	fileOpen.Read(b1)
	intSlice := convertToIntSlice(b1)

	//test should equal
	getLargestProduct(intSlice, 4)
	//actual
	getLargestProduct(intSlice, 13)
}

func getLargestProduct(intSlice []int, increment int) {
	i, current, top, next := increment-1, 0, 1, 1
	for {
		if current == 1000 {
			break
		}
		if current < i {
			next *= intSlice[current]
			current++
		}
		if current == i {
			if next > top {
				top = next
			}
			current, i, next = current-increment+1, i+1, 1
		}
	}
	fmt.Println("The greatest product is: ", top)
}

func convertToIntSlice(b []byte) []int {
	var intSlice []int
	for _, e := range b {
		temp, err := strconv.Atoi(string(e))
		if err != nil {
			fmt.Println(err)
		}
		intSlice = append(intSlice, temp)
	}
	return intSlice
}
