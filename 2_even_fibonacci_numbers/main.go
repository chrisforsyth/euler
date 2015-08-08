// Project Euler Problem 2
// Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:
//
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//
// By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

package main

import "fmt"

func main() {
	sum := 0
	c := fibGenerator()

	for n := 0; ; n++ {
		val := <-c

		if val > 4000000 {
			break
		} // if the value is greater than four million, break out of loop

		if val%2 == 0 { // if the value is even, add it to sum
			sum += val
		}
	}

	fmt.Println(sum)
}

// This method is used to calculate the next Fibonacci number in the sequence
// The iterator returns the next value on each call
func fibGenerator() chan int {
	c := make(chan int)

	go func() {
		for i, j := 0, 1; ; i, j = i+j, i {
			c <- i // return i which is the sum of i + j
		}
	}()

	return c
}
