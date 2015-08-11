// Problem 9
// A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
//
// a2 + b2 = c2
//
// For example, 32 + 42 = 9 + 16 = 25 = 52.
//
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

package main

import "fmt"

func main() {
	fmt.Println("Euler Project ")

	for c := 3; c <= 1000; c++ {
		for b := 2; b < c; b++ {
			for a := 1; a < b; a++ {
				cSquared := a*a + b*b
				if cSquared == c*c && a+b+c == 1000 { // if a^2 + b^2 = c^2 and a+b+c = 1000
					fmt.Println(a * b * c)
				}
			}
		}
	}
}
