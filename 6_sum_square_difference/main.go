// Project Euler Problem 6
// The sum of the squares of the first ten natural numbers is,
//
// 12 + 22 + ... + 102 = 385
//
// The square of the sum of the first ten natural numbers is,
//
// (1 + 2 + ... + 10)2 = 552 = 3025
//
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
//
// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Euler Project 6")
	sumSquare := 1
	squareSum := 1

	for i := 2; i <= 100; i++ {
		val := i * i
		sumSquare += val
		squareSum += i
	}

	squareSum = squareSum * squareSum
	fmt.Printf("difference between sum of squares and square of sum of 1000 = %d", squareSum-sumSquare)
}
