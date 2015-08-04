// Project Euler Problem 4
// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
//
// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	largestPalindrome := 0

	// Brute force it
	for i := 999; i >= 100; i-- { // Check all numbers i where 100 <= i <= 999
		for j := 999; j >= 100; j-- { // Check all numbers j where 100 <= j <= 999
			if isPalindome(i*j) && i*j > largestPalindrome { // If the number is a palindrome and it's greater then the last largest palindrome
				largestPalindrome = i * j
			}
		}
	}

	fmt.Println(largestPalindrome)
}

// Used to determine if a given number is a palindrome
// Takes in an int value, num, to test
// Returns true if num is a palindrome, false otherwise
func isPalindome(num int) bool {
	stringNum := strconv.FormatInt(int64(num), 10)

	n := 0
	rune := make([]rune, len(stringNum))
	for _, r := range stringNum { // For all the items in stringNum
		rune[n] = r // add the character to the slice
		n++
	}

	rune = rune[0:n]

	// Flip the rune order
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}

	// Convert the character array back to an int
	flippedInt, _ := strconv.Atoi(string(rune))

	return flippedInt == num
}
