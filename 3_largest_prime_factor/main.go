// Project Euler Problem 3
// The prime factors of 13195 are 5, 7, 13 and 29.
//
// What is the largest prime factor of the number 600851475143 ?

package main

import (
	"fmt"
	"math"
)

func main() {
	number := 600851475143
	topPrime := 1

	// The larges prime factor cannot be larger than the ceiling of the square root
	for i := 2; i <= int(math.Ceil((math.Sqrt(600851475143)))); i++ {
		if number%i == 0 { // If the current value i is a factor of number
			number = number / i // Divide the number by i
			topPrime = i        // Set topPrime to the current i value
			i = 1               // Reset i to 1 which will then be incremented
		}
	}

	fmt.Println(topPrime) // Print the current topPrime
}
