/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.

*/

package main

import (
	"fmt"
)

func main() {
	tripletCh := make(chan *Triplet)

	go constructTriplets(tripletCh)
	for {
		triplet := <-tripletCh
		if triplet.Equals(1000) {
			fmt.Println("The solution for Euler problem 9:")
			fmt.Print(triplet.PrintSum(), triplet.PrintProduct())
			break
		}
	}
	fmt.Println("Type any key to exit.")
	fmt.Scanln()
}

func constructTriplets(tripletCh chan *Triplet) {
	for m := 1; m < 1000/3; m++ {
		for n := m + 1; n < 1000/3; n++ {
			triplet := new(Triplet)
			triplet.compute(m, n)
			tripletCh <- triplet
		}
	}
}

type Triplet struct {
	A int
	B int
	C int
}

func (t *Triplet) Equals(numb int) bool {
	return (t.A + t.B + t.C) == numb
}

func (t *Triplet) PrintSum() string {
	return fmt.Sprintf("%d + %d + %d = %d\n", t.A, t.B, t.C, t.Sum())
}

func (t *Triplet) PrintProduct() string {
	return fmt.Sprintf("%d * %d * %d = %d\n", t.A, t.B, t.C, t.Product())
}

func (t *Triplet) Sum() int {
	return t.A + t.B + t.C
}

func (t *Triplet) Product() int {
	return t.A * t.B * t.C
}

func (t *Triplet) compute(m, n int) {
	t.A = (n * n) - (m * m)
	t.B = 2 * n * m
	t.C = (n * n) + (m * m)
}
