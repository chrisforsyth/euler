// Problem 10
// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
// Find the sum of all the primes below two million.

package main

import "fmt"

func main() {
	fmt.Println("Euler Project 10")
	prime := int64(0)
	c := primeGenerator()
	sum := int64(0)

	for n := 0; ; n++ {
		prime = <-c

		if prime >= 2000000 {
			break
		}

		sum += prime
	}

	fmt.Println(sum)
}

// This method is used to calculate the next Prime number in the sequence
// The iterator returns the next value on each call
func primeGenerator() chan int64 {
	c := make(chan int64)

	go func() {

		c <- 2

		for i := int64(3); ; i += 2 {
			if isPrime(i) {
				c <- i
			}
		}
	}()

	return c
}

// This method returns true if the number is prime and false otherwise
func isPrime(num int64) bool {
	if num == 2 { // If the number is two it is the one prime even number
		return true
	}

	if num%2 == 0 { // even numbers cannot be prime
		return false
	}

	// start at 3 and increment by 2 since primes > 2 aren't even
	for i := int64(3); i*i <= num; i += 2 {
		if num%i == 0 {
			return false
		}
	}

	return true
}
