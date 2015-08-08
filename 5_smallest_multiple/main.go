// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package main

import "fmt"

func main() {
	isDivisible := true

	divisors := []int{11, 12, 13, 14, 15, 16, 17, 18, 19}

	for i := 20; ; i += 20 {
		isDivisible = true

		for j := 0; j < len(divisors); j++ {
			if i%divisors[j] != 0 {
				isDivisible = false
				break
			}
		}

		if isDivisible {
			fmt.Println(i)
			return
		}
	}
}
