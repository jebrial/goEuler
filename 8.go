//Find the thirteen adjacent digits in the 1000-digit number that have the greatest product. What is the value of this product?

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fileOpen, _ := os.Open("8.dat")
	b1 := make([]byte, 1000)
	fileOpen.Read(b1)
	start := time.Now()
	intSlice := convertToIntSlice(b1)
	fmt.Println("Finished converting string in: ", time.Since(start))
	//test- should equal 5832
	getLargestProduct(intSlice, 4)
	getLargestProductV2(intSlice, 4)
	getLargestProductV3(intSlice, 4) //this one reports a spike in time to complete
	getLargestProductV3(intSlice, 4)
	//actual
	getLargestProduct(intSlice, 13)
	getLargestProductV2(intSlice, 13)
	getLargestProductV3(intSlice, 13)
}

func getLargestProduct(intSlice []int, increment int) {
	start := time.Now()
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
	fmt.Printf("The greatest product is: %d  After %s \n", top, time.Since(start))
}

func getLargestProductV2(intSlice []int, increment int) {
	start := time.Now()
	top, next := 1, 1
	for i := 0; i < len(intSlice)-increment; i++ {
		next = 1
		for j := i; j < increment+i; j++ {
			next *= intSlice[j]
		}
		if next > top {
			top = next
		}
	}
	fmt.Printf("The greatest product is: %d  After %s \n", top, time.Since(start))
}

func getLargestProductV3(intSlice []int, increment int) {
	start := time.Now()
	top := getLargestProductRecursive(intSlice, increment, 0, 0)
	fmt.Printf("The greatest product is: %d  After %s \n", top, time.Since(start))
}

func getLargestProductRecursive(intSlice []int, increment int, start int, prevTop int) int {
	temp := 1
	for i := start; i < increment; i++ {
		temp *= intSlice[i]
	}
	if temp > prevTop {
		prevTop = temp
	}
	if increment == len(intSlice) {
		return prevTop
	}
	return getLargestProductRecursive(intSlice, increment+1, start+1, prevTop)
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
